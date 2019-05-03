package model

import (
	"context"
	"fmt"
	"strings"
	"time"

	"circles/server/types"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/rs/xid"
)

// AddUserToChat adds a User to an existing Chat
func AddUserToChat(ctx context.Context, gs generalStore, uid, chid string) (*ChatModel, error) {
	var chat *types.Chat

	if err := gs.CreateMembership(ctx, uid, chid); err != nil {
		return nil, err

	} else if chat, err = gs.FindChat(ctx, chid); err != nil {
		return nil, err
	}

	return &ChatModel{chat, gs}, nil
}

// AllChats retrieves all users and returns them as ChatModels
func AllChats(ctx context.Context, gs generalStore) ([]*ChatModel, error) {
	chats, err := gs.AllChats(ctx)
	if err != nil {
		return nil, err
	}

	var chatModels []*ChatModel
	for _, chat := range chats {
		chatModels = append(chatModels, &ChatModel{chat, gs})
	}

	return chatModels, nil
}

// CreateChat creates a new Chat with the given data and returns it as a ChatModel
func CreateChat(ctx context.Context, gs generalStore, chat *types.Chat, userIDs []string) (*ChatModel, error) {
	chat.ID = xid.New().String()
	chat.CreatedAt = time.Now().Format(time.RFC3339)

	gs.BeginTransaction(ctx)
	defer gs.EndTransaction(ctx)

	if err := gs.CreateChat(ctx, chat, userIDs); err != nil {
		return nil, err

	} else if chat, err = gs.FindChat(ctx, chat.ID); err != nil {
		return nil, err
	}

	if _, err := gs.FindChat(ctx, chat.ID); err != nil {
		return nil, err
	}

	for _, uid := range userIDs {
		if err := gs.CreateMembership(ctx, uid, chat.ID); err != nil {
			return nil, err
		}
	}

	return &ChatModel{chat, gs}, nil
}

// DeleteChat deletes the Chat specified by ID and returns it as a ChatModel
func DeleteChat(ctx context.Context, gs generalStore, id string) (*ChatModel, error) {
	chat, err := gs.FindChat(ctx, id)
	if err != nil {
		return nil, err

	} else if err := gs.DeleteChat(ctx, id); err != nil {
		return nil, err
	}

	return &ChatModel{chat, gs}, nil
}

// FindChat retrieves the Chat specified by id
func FindChat(ctx context.Context, gs generalStore, id string) (*ChatModel, error) {
	chat, err := gs.FindChat(ctx, id)
	if err != nil {
		return nil, err
	}

	return &ChatModel{chat, gs}, nil
}

// UpdateChatInput ...
type UpdateChatInput struct {
	ID   string
	Name *string
}

// UpdateChat updates the chat specified by ID with the given data and returns it as a ChatModel
func UpdateChat(ctx context.Context, gs generalStore, input *UpdateChatInput) (*ChatModel, error) {
	chat, err := gs.FindChat(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	if input.Name != nil {
		chat.Name = *input.Name
	}

	if err = gs.UpdateChat(ctx, chat); err != nil {
		return nil, err
	}

	return &ChatModel{chat, gs}, nil
}

// UserChats retrieves all Chats for a User
func UserChats(ctx context.Context, gs generalStore, uid string) ([]*ChatModel, error) {
	chats, err := gs.AllChatsByUserMembership(ctx, uid)
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
func (r *ChatModel) Name(ctx context.Context) (string, error) {
	if r.Chat.Name == "" {
		return chatUsersAsName(ctx, r.store, r.Chat.ID)
	}

	return r.Chat.Name, nil
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

// Users field resolver
func (r *ChatModel) Users(ctx context.Context) ([]*UserModel, error) {
	return ChatUsers(ctx, r.store, r.Chat.ID)
}

func chatUsersAsName(ctx context.Context, store generalStore, chid string) (string, error) {
	var (
		val   interface{}
		uid   string
		users []*UserModel
		ok    bool
		err   error
	)

	if val = ctx.Value("uid"); val == nil {
		return "", fmt.Errorf("could not retrieve uid from context in ChatModel.Name")
	} else if uid, ok = val.(string); !ok {
		return "", fmt.Errorf("could not convert val to string in ChatModel.Name")
	}

	if users, err = ChatUsers(ctx, store, chid); err != nil {
		return "", err
	}

	chatUserNames := []string{}
	for _, user := range users {
		if user.User.ID != uid {
			chatUserNames = append(chatUserNames, user.User.Name)
		}
	}

	return strings.Join(chatUserNames, ", "), nil
}
