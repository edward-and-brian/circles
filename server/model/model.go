package model

import (
	"circles/server/types"
	"context"
)

type generalStore interface {
	chatStore
	circleStore
	messageStore
	userStore
}

// chatStore is an interface for database interaction by the ChatModel
type chatStore interface {
	AddUserToChat(context.Context, string, string, string) error
	AllChatsByUserID(context.Context, string) ([]*types.Chat, error)
	CreateChat(context.Context, *types.Chat) error
	DeleteChat(context.Context, string) error
	FindChat(context.Context, string) (*types.Chat, error)
	UpdateChat(context.Context, *types.Chat) error
}

// circleStore is an interface for database interaction by the CircleModel
type circleStore interface {
	AllCirclesByChatID(context.Context, string) ([]*types.Circle, error)
	CreateCircle(context.Context, *types.Circle) error
	DeleteCircle(context.Context, string) error
	FindCircle(context.Context, string) (*types.Circle, error)
	UpdateCircle(context.Context, *types.Circle) error
}

// MessageStore is an interface for database interaction by the MessageModel
type messageStore interface {
	AllMessagesByCircleID(context.Context, string) ([]*types.Message, error)
	CreateMessage(context.Context, *types.Message) error
	DeleteMessage(context.Context, string) error
	FindMessage(context.Context, string) (*types.Message, error)
	UpdateMessage(context.Context, *types.Message) error
}

// userStore is an interface for database interaction by the UserModel
type userStore interface {
	AllUsers(context.Context) ([]*types.User, error)
	CreateUser(context.Context, *types.User) error
	DeleteUser(context.Context, string) error
	FindUser(context.Context, string) (*types.User, error)
	UpdateUser(context.Context, *types.User) error
}
