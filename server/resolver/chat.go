package resolver

import (
	"context"

	"circles/server/model"
	"circles/server/store"
	"circles/server/types"
)

// ChatAddUser retrieves the chat specified by id
func (r *Resolver) ChatAddUser(ctx context.Context, args *struct{ UserID, ChatID string }) (*model.ChatModel, error) {
	return model.AddUserToChat(ctx, &store.GeneralStore{}, args.UserID, args.ChatID)
}

// Chat retrieves the chat specified by id
func (r *Resolver) Chat(ctx context.Context, args *IDArgs) (*model.ChatModel, error) {
	return model.FindChat(ctx, &store.GeneralStore{}, args.ID)
}

// CreateChat creates a new Chat with the given data
func (r *Resolver) CreateChat(ctx context.Context, args *struct{ Input *types.Chat }) (*model.ChatModel, error) {
	return model.CreateChat(ctx, &store.GeneralStore{}, args.Input)
}

// DeleteChat deletes the Chat specified by ID
func (r *Resolver) DeleteChat(ctx context.Context, args *IDArgs) (*model.ChatModel, error) {
	return model.DeleteChat(ctx, &store.GeneralStore{}, args.ID)
}

// UpdateChat updates the Chat specified by ID with the given data
func (r *Resolver) UpdateChat(ctx context.Context, args *struct{ Input *model.UpdateChatInput }) (*model.ChatModel, error) {
	return model.UpdateChat(ctx, &store.GeneralStore{}, args.Input)
}
