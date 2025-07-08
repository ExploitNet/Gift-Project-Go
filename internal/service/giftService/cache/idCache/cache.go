package idCache

import (
	"errors"
	"sync"

	"github.com/gotd/td/tg"
)

type idCacheImpl struct {
	users    map[string]*tg.User
	channels map[string]*tg.Channel
	mu       sync.RWMutex
}

func NewIDCache() *idCacheImpl {
	return &idCacheImpl{
		users:    make(map[string]*tg.User),
		channels: make(map[string]*tg.Channel),
	}
}

func (c *idCacheImpl) SetUser(user *tg.User) {
	if user == nil {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.users[user.Username] = user
}

func (c *idCacheImpl) GetUser(username string) (*tg.User, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	user, ok := c.users[username]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (c *idCacheImpl) SetChannel(channel *tg.Channel) {
	if channel == nil {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.channels[channel.Username] = channel
}

func (c *idCacheImpl) GetChannel(channelName string) (*tg.Channel, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	channel, ok := c.channels[channelName]
	if !ok {
		return nil, errors.New("channel not found")
	}
	return channel, nil
}
