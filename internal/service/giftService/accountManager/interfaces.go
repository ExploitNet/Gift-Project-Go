package accountManager

import "github.com/gotd/td/tg"

// UserCache defines the interface for caching user information.
// It provides persistent storage for user data to avoid redundant API calls
// and maintain state across application restarts.
type UserCache interface {
	// SetUser stores a user in the cache with the specified ID as key.
	//
	// Parameters:
	//   - id: unique identifier for the user
	//   - user: the user object to cache
	SetUser(user *tg.User)

	// GetUser retrieves a cached user by their ID.
	//
	// Parameters:
	//   - id: unique identifier of the user to retrieve
	//
	// Returns:
	//   - *tg.User: the cached user object, nil if not found
	//   - error: retrieval error (currently always nil)
	GetUser(id string) (*tg.User, error)
}

type ChannelCache interface {
	// SetChannel stores a channel in the cache with the specified ID as key.
	//
	// Parameters:
	//   - id: unique identifier for the channel
	//   - channel: the channel object to cache
	SetChannel(channel *tg.Channel)

	// GetChannel retrieves a cached channel by its ID.
	//
	// Parameters:
	//   - id: unique identifier of the channel to retrieve
	//
	// Returns:
	//   - *tg.Channel: the cached channel object, nil if not found
	//   - error: retrieval error (currently always nil)
	GetChannel(id string) (*tg.Channel, error)
}
