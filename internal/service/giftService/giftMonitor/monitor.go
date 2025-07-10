// Package giftMonitor provides gift monitoring functionality for the gift buying system.
// It continuously monitors for new gifts, validates them against criteria,
// and triggers notifications when eligible gifts are discovered.
package giftMonitor

import (
	"context"
	"fmt"
	"gift-buyer/internal/infrastructure/logsWriter"
	"gift-buyer/internal/service/giftService/giftInterfaces"
	"gift-buyer/internal/service/giftService/giftServiceHelpers"
	"gift-buyer/internal/service/giftService/giftTypes"
	"gift-buyer/pkg/errors"

	"sync"
	"time"

	"github.com/gotd/td/tg"
)

// giftMonitorImpl implements the GiftMonitor interface for monitoring new gifts.
// It periodically checks for new gifts, validates them against criteria,
// and manages caching to avoid duplicate processing.
type giftMonitorImpl struct {
	// cache stores processed gifts to avoid duplicate notifications
	cache giftInterfaces.GiftCache

	// manager handles communication with Telegram API for gift retrieval
	manager giftInterfaces.Giftmanager

	// validator evaluates gifts against purchase criteria
	validator giftInterfaces.GiftValidator

	// notification sends alerts about new eligible gifts
	notification giftInterfaces.NotificationService

	// logsWriter is used to write logs to a file
	errorLogsWriter logsWriter.LogsWriter
	infoLogsWriter  logsWriter.LogsWriter

	// ticker controls the monitoring interval
	ticker *time.Ticker

	firstRun bool

	// paused indicates if monitoring is currently paused
	paused bool

	// mu protects the paused field from concurrent access
	mu sync.RWMutex
}

// NewGiftMonitor creates a new GiftMonitor instance with the specified dependencies.
// The monitor will check for new gifts at the specified interval and process
// them through the validation and notification pipeline.
//
// Parameters:
//   - cache: gift cache for tracking processed gifts
//   - manager: gift manager for retrieving available gifts
//   - validator: gift validator for eligibility checking
//   - notification: notification service for sending alerts
//   - tickTime: interval between gift checks
//
// Returns:
//   - giftInterfaces.GiftMonitor: configured gift monitor instance
func NewGiftMonitor(
	cache giftInterfaces.GiftCache,
	manager giftInterfaces.Giftmanager,
	validator giftInterfaces.GiftValidator,
	notification giftInterfaces.NotificationService,
	tickTime time.Duration,
	errorLogsWriter logsWriter.LogsWriter,
	infoLogsWriter logsWriter.LogsWriter,
) *giftMonitorImpl {
	return &giftMonitorImpl{
		cache:           cache,
		manager:         manager,
		validator:       validator,
		notification:    notification,
		ticker:          time.NewTicker(tickTime),
		firstRun:        true,
		errorLogsWriter: errorLogsWriter,
		infoLogsWriter:  infoLogsWriter,
	}
}

// Start begins the gift monitoring process and returns newly discovered eligible gifts.
// It runs continuously until the context is cancelled, checking for new gifts
// at the configured interval. When eligible gifts are found, it sends notifications
// and returns the gifts for purchase processing.
//
// The monitoring process:
//  1. Waits for the next tick or context cancellation
//  2. Checks for new gifts via the gift manager
//  3. Validates new gifts against criteria
//  4. Sends notifications for eligible gifts
//  5. Returns eligible gifts for purchase
//
// Parameters:
//   - ctx: context for cancellation and timeout control
//
// Returns:
//   - map[*tg.StarGift]int64: map of eligible gifts to their purchase quantities
//   - error: monitoring error, API communication error, or context cancellation
func (gm *giftMonitorImpl) Start(ctx context.Context) (map[*tg.StarGift]*giftTypes.GiftRequire, error) {
	resultCh := make(chan map[*tg.StarGift]*giftTypes.GiftRequire, 10)
	errCh := make(chan error, 10)

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-gm.ticker.C:
			if gm.IsPaused() {
				continue
			}

			go func() {
				newGifts, err := gm.checkForNewGifts(ctx)
				if err != nil {
					errCh <- err
					return
				}
				if len(newGifts) == 0 {
					giftServiceHelpers.LogInfo(gm.infoLogsWriter, "no new gifts found")
					return
				}
				resultCh <- newGifts
			}()
		case newGifts := <-resultCh:
			return newGifts, nil
		case err := <-errCh:
			if !gm.IsPaused() {
				gm.notification.SendErrorNotification(ctx, err)
				giftServiceHelpers.LogError(gm.errorLogsWriter, err.Error())
			}
			continue
		}
	}
}

// checkForNewGifts retrieves current gifts and identifies new eligible ones.
// It compares the current gift list against the cache to find new gifts,
// validates them against criteria, and updates the cache.
//
// Parameters:
//   - ctx: context for API request cancellation
//
// Returns:
//   - map[*tg.StarGift]int64: map of new eligible gifts to purchase quantities
//   - error: API communication error or validation error
func (gm *giftMonitorImpl) checkForNewGifts(ctx context.Context) (map[*tg.StarGift]*giftTypes.GiftRequire, error) {
	currentGifts, err := gm.manager.GetAvailableGifts(ctx)
	if err != nil {
		return nil, err
	}

	newValidGifts := make(map[*tg.StarGift]*giftTypes.GiftRequire, len(currentGifts))

	for _, gift := range currentGifts {
		if gm.cache.HasGift(gift.ID) {
			continue
		}
		if giftRequire, ok := gm.validator.IsEligible(gift); ok {
			giftServiceHelpers.LogInfo(gm.infoLogsWriter, fmt.Sprintf("gift id %d is valid", gift.ID))
			newValidGifts[gift] = giftRequire
		}

		gm.cache.SetGift(gift.ID, gift)
	}

	if gm.firstRun {
		gm.firstRun = false
		return nil, errors.Wrap(errors.New("first run"), "touch grass")
	}

	return newValidGifts, nil
}

// Pause pauses the gift monitoring process.
// It stops the monitoring goroutine and prevents new gifts from being discovered.
func (gm *giftMonitorImpl) Pause() {
	gm.mu.Lock()
	defer gm.mu.Unlock()
	if !gm.paused {
		gm.paused = true
		giftServiceHelpers.LogInfo(gm.infoLogsWriter, "Gift monitoring paused")
	}
}

// Resume resumes the gift monitoring process.
// It starts the monitoring goroutine and allows new gifts to be discovered.
func (gm *giftMonitorImpl) Resume() {
	gm.mu.Lock()
	defer gm.mu.Unlock()
	if gm.paused {
		gm.paused = false
		giftServiceHelpers.LogInfo(gm.infoLogsWriter, "Gift monitoring resumed")
	}
}

// IsPaused returns the status of the gift monitoring process.
//
// Returns:
//   - bool: true if the monitoring is paused, false if active
func (gm *giftMonitorImpl) IsPaused() bool {
	gm.mu.RLock()
	defer gm.mu.RUnlock()
	return gm.paused
}
