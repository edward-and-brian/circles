package model

import (
	"context"
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
	chat.LastMessageAt = chat.CreatedAt
	chat.LastCircleName = "General"
	generalCircle := &types.Circle{
		ChatID: chat.ID,
		Name:   "General",
	}

	gs.BeginTransaction(ctx)
	defer gs.EndTransaction(ctx)

	if err := gs.CreateChat(ctx, chat, userIDs); err != nil {
		return nil, err

	} else if _, err = CreateCircle(ctx, gs, generalCircle); err != nil {
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
	gs.BeginTransaction(ctx)
	defer gs.EndTransaction(ctx)

	chat, err := gs.FindChat(ctx, id)
	if err != nil {
		return nil, err

	} else if err := gs.DeleteChat(ctx, id); err != nil {
		return nil, err

	} else if err := gs.DeleteMembershipsByChatID(ctx, id); err != nil {
		return nil, err

	} else if circles, err := gs.AllCirclesByChatID(ctx, id); err != nil {
		return nil, err

	} else {
		for _, circle := range circles {
			if err = gs.DeleteCircle(ctx, circle.ID); err != nil {
				return nil, err

			} else if err := gs.DeleteMessagesByCircleID(ctx, circle.ID); err != nil {
				return nil, err
			}
		}
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
	ID             string
	Name           *string
	LastCircleName *string
	LastMessageAt  *string
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

	if input.LastCircleName != nil {
		chat.LastCircleName = *input.LastCircleName
	}

	if input.LastMessageAt != nil {
		chat.LastMessageAt = *input.LastMessageAt
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
		if ch.Name == "" {
			if ch.Name, err = chatUsersAsName(ctx, gs, ch.ID, uid); err != nil {
				return nil, err
			}
		}
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
	return r.Chat.Name, nil
}

// LastCircleName field resolver
func (r *ChatModel) LastCircleName() string {
	return r.Chat.LastCircleName
}

// LastMessageAt field resolver
func (r *ChatModel) LastMessageAt() (graphql.Time, error) {
	t, err := time.Parse(time.RFC3339, r.Chat.LastMessageAt)
	return graphql.Time{Time: t}, err
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

func chatUsersAsName(ctx context.Context, store generalStore, chid string, uid string) (string, error) {
	users, err := store.AllUsersByChatMembership(ctx, chid)
	if err != nil {
		return "", err
	}

	chatUserNames := []string{}
	for _, user := range users {
		if user.ID != uid {
			chatUserNames = append(chatUserNames, user.Name)
		}
	}

	return strings.Join(chatUserNames, ", "), nil
}
