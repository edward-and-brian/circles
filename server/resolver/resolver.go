package resolver

import "circles/server/types"

type store interface {
	Open() error

	AllUsers() ([]*types.User, error)
	CreateUser(u *types.User) error
	DeleteUser(id *string) error
	FindUser(id *string) (*types.User, error)
	UpdateUser(user *types.User) error

	AllUserChats(id *string) ([]*types.Chat, error)
	CreateChat(ch *types.Chat) error
	DeleteChat(id *string) error
	FindChat(id *string) (*types.Chat, error)
	UpdateChat(chat *types.Chat) error
}

// Resolver is the entry point for all top-level operations.
type Resolver struct {
	Store store
}
