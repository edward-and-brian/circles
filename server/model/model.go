package model

import (
	"circles/server/types"
	"context"
)

type generalStore interface {
	BeginTransaction(ctx context.Context) error
	EndTransaction(ctx context.Context) error
	chatStore
	circleStore
	messageStore
	userStore
}

// chatStore is an interface for database interaction by the ChatModel
type chatStore interface {
	AllChats(ctx context.Context) ([]*types.Chat, error)
	AllChatsByUserMembership(ctx context.Context, uid string) ([]*types.Chat, error)
	CreateChat(ctx context.Context, chat *types.Chat, uids []string) error
	CreateMembership(ctx context.Context, uid string, chid string) error
	DeleteChat(ctx context.Context, id string) error
	FindChat(ctx context.Context, id string) (*types.Chat, error)
	UpdateChat(ctx context.Context, chat *types.Chat) error
}

// circleStore is an interface for database interaction by the CircleModel
type circleStore interface {
	AllCircles(ctx context.Context) ([]*types.Circle, error)
	AllCirclesByChatID(ctx context.Context, chid string) ([]*types.Circle, error)
	CreateCircle(ctx context.Context, circle *types.Circle) error
	DeleteCircle(ctx context.Context, id string) error
	FindCircle(ctx context.Context, id string) (*types.Circle, error)
	UpdateCircle(ctx context.Context, circle *types.Circle) error
}

// MessageStore is an interface for database interaction by the MessageModel
type messageStore interface {
	AllMessages(ctx context.Context) ([]*types.Message, error)
	AllMessagesByCircleID(ctx context.Context, ciid string) ([]*types.Message, error)
	CreateMessage(ctx context.Context, message *types.Message) error
	DeleteMessage(ctx context.Context, id string) error
	FindMessage(ctx context.Context, id string) (*types.Message, error)
	UpdateMessage(ctx context.Context, message *types.Message) error
}

// userStore is an interface for database interaction by the UserModel
type userStore interface {
	AllUsers(ctx context.Context) ([]*types.User, error)
	AllUsersByChatMembership(ctx context.Context, chid string) ([]*types.User, error)
	CreateUser(ctx context.Context, user *types.User) error
	DeleteUser(ctx context.Context, id string) error
	FindUser(ctx context.Context, id string) (*types.User, error)
	LoginUser(ctx context.Context, phoneNumber string, displayName string) (*types.User, error)
	UpdateUser(ctx context.Context, user *types.User) error
}
