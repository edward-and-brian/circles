package resolver

import (
	"context"

	"circles/server/model"
	"circles/server/types"
)

// ChatAddUser retrieves the chat specified by id
func (r *Resolver) ChatAddUser(ctx context.Context, args *struct{ UserID, ChatID string }) (*model.ChatModel, error) {
	return model.AddUserToChat(ctx, r.store, args.UserID, args.ChatID)
}

// Chat retrieves the chat specified by id
func (r *Resolver) Chat(ctx context.Context, args *IDArgs) (*model.ChatModel, error) {
	return model.FindChat(ctx, r.store, args.ID)
}

// Chats retrieves all users
func (r *Resolver) Chats(ctx context.Context) ([]*model.ChatModel, error) {
	return model.AllChats(ctx, r.store)
}

// CreateChatInput ...
type CreateChatInput struct {
	Chat    *types.Chat
	UserIDs []string
}

// CreateChat creates a new Chat with the given data
func (r *Resolver) CreateChat(ctx context.Context, args *CreateChatInput) (*model.ChatModel, error) {
	return model.CreateChat(ctx, r.store, args.Chat, args.UserIDs)
}

// DeleteChat deletes the Chat specified by ID
func (r *Resolver) DeleteChat(ctx context.Context, args *IDArgs) (*model.ChatModel, error) {
	return model.DeleteChat(ctx, r.store, args.ID)
}

// UpdateChat updates the Chat specified by ID with the given data
func (r *Resolver) UpdateChat(ctx context.Context, args *struct{ Chat *model.UpdateChatInput }) (*model.ChatModel, error) {
	return model.UpdateChat(ctx, r.store, args.Chat)
}
