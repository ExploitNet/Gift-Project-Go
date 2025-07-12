package authService

import (
	"context"
	"errors"
	"gift-buyer/internal/config"
	"gift-buyer/internal/infrastructure/logsWriter"
	"gift-buyer/internal/service/authService/apiChecker"
	"gift-buyer/internal/service/authService/authInterfaces"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gotd/td/telegram"
	"github.com/gotd/td/tg"
)

type AuthManagerImpl struct {
	api             *tg.Client
	botApi          *tg.Client
	mu              sync.RWMutex
	sessionManager  authInterfaces.SessionManager
	apiChecker      authInterfaces.ApiChecker
	cfg             *config.TgSettings
	reconnect       chan struct{}
	stopCh          chan struct{}
	wg              sync.WaitGroup
	monitor         authInterfaces.GiftMonitorAndAuthController
	infoLogsWriter  logsWriter.LogsWriter
	errorLogsWriter logsWriter.LogsWriter
}

func NewAuthManager(sessionManager authInterfaces.SessionManager, apiChecker authInterfaces.ApiChecker, cfg *config.TgSettings, infoLogsWriter logsWriter.LogsWriter, errorLogsWriter logsWriter.LogsWriter) *AuthManagerImpl {
	return &AuthManagerImpl{
		sessionManager:  sessionManager,
		apiChecker:      apiChecker,
		cfg:             cfg,
		reconnect:       make(chan struct{}, 1),
		stopCh:          make(chan struct{}),
		infoLogsWriter:  infoLogsWriter,
		errorLogsWriter: errorLogsWriter,
	}
}

func (f *AuthManagerImpl) InitClient(ctx context.Context) (*tg.Client, error) {
	if f.cfg == nil {
		return nil, errors.New("configuration is nil")
	}

	if f.sessionManager == nil {
		return nil, errors.New("session manager is nil")
	}

	client := telegram.NewClient(f.cfg.AppId, f.cfg.ApiHash, telegram.Options{
		SessionStorage: &telegram.FileSessionStorage{
			Path: "session.json",
		},
	})

	api, err := f.sessionManager.InitUserAPI(client, ctx)
	if err != nil {
		return nil, err
	}
	f.mu.Lock()
	f.api = api
	f.mu.Unlock()
	return api, nil
}

func (f *AuthManagerImpl) RunApiChecker(ctx context.Context) {
	if f.apiChecker == nil {
		f.logError("API checker is nil, skipping")
		return
	}

	f.logInfo("Starting API monitoring")

	f.wg.Add(1)
	go func() {
		defer f.wg.Done()
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				f.logInfo("API monitoring stopped due to context cancellation")
				return
			case <-f.stopCh:
				f.logInfo("API monitoring stopped due to stop signal")
				return
			case <-ticker.C:
				if err := f.apiChecker.Run(ctx); err != nil {
					f.logErrorf("API check failed: %v", err)
					if f.isCriticalError(err) {
						f.logErrorf("Critical API error detected, triggering reconnect: %v", err)
						select {
						case f.reconnect <- struct{}{}:
							f.stopCh <- struct{}{}
							f.logInfo("Reconnect signal sent")
						default:
							f.logError("Reconnect channel is full")
						}
					}
				} else {
					f.logInfo("API check successful")
				}
			}
		}
	}()

	f.wg.Add(1)
	go func() {
		defer f.wg.Done()
		f.handleReconnectSignals(ctx)
	}()
}

func (f *AuthManagerImpl) isCriticalError(err error) bool {
	if err == nil {
		return false
	}

	errStr := strings.ToLower(err.Error())

	return strings.Contains(errStr, "auth_key_unregistered") ||
		strings.Contains(errStr, "connection_not_inited") ||
		strings.Contains(errStr, "session_revoked")
}

func (f *AuthManagerImpl) handleReconnectSignals(ctx context.Context) {
	for {
		select {
		case <-f.reconnect:
			f.logInfo("Processing reconnect signal")

			if f.monitor != nil {
				f.logInfo("Pausing gift monitoring during reconnection")
				f.monitor.Pause()
			}

			if _, err := f.Reconnect(ctx); err != nil {
				f.logErrorf("Reconnect failed: %v", err)
				if strings.Contains(err.Error(), "timeout") || strings.Contains(err.Error(), "deadline") {
					f.logError("Reconnection timeout, exiting program")
					os.Exit(1)
				}
			} else {
				if f.monitor != nil {
					f.logInfo("Resuming gift monitoring after reconnection")
					f.monitor.Resume()
				}
				<-f.stopCh
				f.logInfo("Reconnection completed successfully")
			}
		case <-f.stopCh:
			f.logInfo("Stopping reconnect handler")
			return
		case <-ctx.Done():
			f.logInfo("Context cancelled, stopping reconnect handler")
			return
		}
	}
}

func (f *AuthManagerImpl) InitBotClient(ctx context.Context) (*tg.Client, error) {
	if f.sessionManager == nil {
		return nil, errors.New("session manager is nil")
	}

	botApi, err := f.sessionManager.InitBotAPI(ctx)
	if err != nil {
		return nil, err
	}
	f.mu.Lock()
	f.botApi = botApi
	f.mu.Unlock()
	return botApi, nil
}

func (f *AuthManagerImpl) GetBotApi() *tg.Client {
	f.mu.RLock()
	defer f.mu.RUnlock()
	return f.botApi
}

func (f *AuthManagerImpl) GetApi() *tg.Client {
	f.mu.RLock()
	defer f.mu.RUnlock()
	return f.api
}

func (f *AuthManagerImpl) Reconnect(ctx context.Context) (*tg.Client, error) {
	f.logInfo("Starting reconnection process")

	tgc, err := f.InitClient(ctx)
	if err != nil {
		return nil, err
	}

	if f.apiChecker != nil {
		newApiChecker := apiChecker.NewApiChecker(tgc, time.NewTicker(2*time.Second))
		f.SetApiChecker(newApiChecker)
		f.logInfo("API checker updated with new client")
	}

	f.logInfo("Reconnection successful")
	return tgc, nil
}

func (f *AuthManagerImpl) Stop() {
	f.logInfo("Stopping AuthManager...")

	close(f.stopCh)

	if f.apiChecker != nil {
		f.apiChecker.Stop()
	}

	close(f.reconnect)

	f.wg.Wait()

	f.logInfo("AuthManager stopped")
}

func (f *AuthManagerImpl) SetApiChecker(apiChecker authInterfaces.ApiChecker) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.apiChecker = apiChecker
}

func (f *AuthManagerImpl) SetMonitor(monitor authInterfaces.GiftMonitorAndAuthController) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.monitor = monitor
	f.logInfo("Gift monitor set for auth manager")
}
