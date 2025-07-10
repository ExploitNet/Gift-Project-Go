// Package accountManager manages accounts and their validation.s
package accountManager

import (
	"context"
	"fmt"
	"gift-buyer/internal/service/giftService/accountManager/implement"
	"gift-buyer/pkg/errors"
	"gift-buyer/pkg/logger"
	"strings"

	"github.com/gotd/td/tg"
)

type accountManagerImpl struct {
	api                     *tg.Client
	mainName                string
	validator               implement.ValidatorClient
	usernames, channelNames []string
	userCache               UserCache
	channelCache            ChannelCache
}

func NewAccountManager(api *tg.Client, usernames, channelNames []string, userCache UserCache, channelCache ChannelCache, mainName string) *accountManagerImpl {
	return &accountManagerImpl{
		api:          api,
		usernames:    usernames,
		channelNames: channelNames,
		userCache:    userCache,
		channelCache: channelCache,
		mainName:     mainName,
		validator:    implement.CreateValidatorService(),
	}
}

func (am *accountManagerImpl) SetIds(ctx context.Context) error {
	if am.api == nil {
		return errors.New("API client is nil")
	}

	if len(am.usernames) > 0 {
		if err := am.loadUsersToCache(ctx); err != nil {
			return errors.Wrap(err, "failed to load users to cache")
		}
	}

	if len(am.channelNames) > 0 {
		if err := am.loadChannelsToCache(ctx); err != nil {
			return errors.Wrap(err, "failed to load channels to cache")
		}
	}

	return nil
}

func (am *accountManagerImpl) loadUsersToCache(ctx context.Context) error {
	if am.api == nil {
		return errors.New("API client is nil")
	}

	for _, username := range am.usernames {
		withoutTag := strings.TrimPrefix(username, "@")

		res, err := am.api.ContactsResolveUsername(ctx, &tg.ContactsResolveUsernameRequest{
			Username: withoutTag,
		})
		if err != nil {
			return errors.Wrap(err, "failed to resolve username")
		}
		for _, user := range res.Users {
			if u, ok := user.(*tg.User); ok {
				am.userCache.SetUser(u)
			}
		}
	}

	return nil
}

func (am *accountManagerImpl) loadChannelsToCache(ctx context.Context) error {
	cachedCount := 0
	notFoundChannels := []string{}

	for _, channelName := range am.channelNames {
		withoutTag := strings.TrimPrefix(channelName, "@")

		channel, err := am.loadSingleChannel(ctx, withoutTag)
		if err != nil {
			logger.GlobalLogger.Errorf("failed to load channel %s: %v", channelName, err)
			notFoundChannels = append(notFoundChannels, channelName)
			continue
		}

		am.channelCache.SetChannel(channel)
		cachedCount++
	}

	if len(notFoundChannels) > 0 {
		logger.GlobalLogger.Warnf("Channels not found or inaccessible: %v", notFoundChannels)
	}

	return nil
}

func (am *accountManagerImpl) loadSingleChannel(ctx context.Context, channelName string) (*tg.Channel, error) {
	if am.api == nil {
		return nil, errors.New("API client is nil")
	}
	res, err := am.api.ContactsResolveUsername(ctx, &tg.ContactsResolveUsernameRequest{
		Username: channelName,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to resolve username")
	}
	for _, channel := range res.Chats {
		if c, ok := channel.(*tg.Channel); ok {
			am.channelCache.SetChannel(c)
			return c, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("channel %s not found in response", channelName))
}

// CheckSubscription verifies whether a user is currently subscribed based on their tag.
// Returns true if the subscription is active; false otherwise.
// This method is typically used to determine access eligibility in the application.
func (am *accountManagerImpl) CheckSubscription(usertag string) bool {
	res, err := am.validator.CheckSubscription(context.Background(), &implement.Request{Data: usertag})
	if err != nil {
		return false
	}

	return res.Status
}

// ValidateSubscription performs a local check to determine if a user has an active subscription.
// It takes a user tag as input and returns true if the subscription is considered valid.
// If the check fails or an error occurs, it returns false and logs the reason internally.
// This function is used to control access to specific features or content.
func (am *accountManagerImpl) ValidateSubscription(usertag string) bool {
	res, err := am.validator.ValidateSubscription(context.Background(), &implement.Request{Data: usertag})
	if err != nil {
		logger.GlobalLogger.Errorf("failed to validate subscription: %v", err)
		return false
	}
	logger.GlobalLogger.Infof("subscription validated: %v", res.Status)
	return res.Status
}
