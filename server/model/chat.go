package model

import (
	"context"
	"time"

	"circles/server/types"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/rs/xid"
)

// AddUserToChat adds a User to an existing Chat
func AddUserToChat(ctx context.Context, gs generalStore, uid, chid string) (*ChatModel, error) {
	var chat *types.Chat
	id := xid.New().String()

	if err := gs.AddUserToChat(ctx, id, uid, chid); err != nil {
		return nil, err

	} else if chat, err = gs.FindChat(ctx, chid); err != nil {
		return nil, err
	}

	return &ChatModel{chat, gs}, nil
}

// CreateChat creates a new Chat with the given data and returns it as a ChatModel
func CreateChat(ctx context.Context, gs generalStore, chat *types.Chat) (*ChatModel, error) {
	chat.ID = xid.New().String()

	if err := gs.CreateChat(ctx, chat); err != nil {
		return nil, err

	} else if chat, err = gs.FindChat(ctx, chat.ID); err != nil {
		return nil, err
	}

	return &ChatModel{chat, gs}, nil
}

// DeleteChat deletes the Chat specified by ID and returns it as a ChatModel
func DeleteChat(ctx context.Context, gs generalStore, id string) (*ChatModel, error) {
	Chat, err := gs.FindChat(ctx, id)
	if err != nil {
		return nil, err

	} else if err = gs.DeleteChat(ctx, id); err != nil {
		return nil, err
	}

	return &ChatModel{Chat, gs}, nil
}

// FindChat retrieves the Chat specified by id
func FindChat(ctx context.Context, gs generalStore, id string) (*ChatModel, error) {
	chat, err := gs.FindChat(ctx, id)
	if err != nil {
		return nil, err
	}

	return &ChatModel{chat, gs}, nil
}

// UpdateChat updates the Chat specified by ID with the given data and returns it as a ChatModel
func UpdateChat(ctx context.Context, gs generalStore, chat *types.Chat) (*ChatModel, error) {
	if err := gs.UpdateChat(ctx, chat); err != nil {
		return nil, err

	} else if chat, err = gs.FindChat(ctx, chat.ID); err != nil {
		return nil, err
	}

	return &ChatModel{chat, gs}, nil
}

// UserChats retrieves all Chats for a User
func UserChats(ctx context.Context, gs generalStore, uid string) ([]*ChatModel, error) {
	chats, err := gs.AllChatsByUserID(ctx, uid)
	if err != nil {
		return nil, err
	}

	var chatModels []*ChatModel
	for _, ch := range chats {
		chatModels = append(chatModels, &ChatModel{ch, gs})
	}

	return chatModels, nil
}

// ChatModel is the resolvable struct for the Chat struct
type ChatModel struct {
	*types.Chat
	store generalStore
}

// ID field resolver
func (r *ChatModel) ID() graphql.ID {
	return graphql.ID(r.Chat.ID)
}

// Name field resolver
func (r *ChatModel) Name() string {
	return r.Chat.Name
}

// CreatedAt field resolver
func (r *ChatModel) CreatedAt() (graphql.Time, error) {
	t, err := time.Parse(time.RFC3339, r.Chat.CreatedAt)
	return graphql.Time{Time: t}, err
}

// Circles field resolver
func (r *ChatModel) Circles(ctx context.Context) ([]*CircleModel, error) {
	return ChatCircles(ctx, r.store, r.Chat.ID)
}
