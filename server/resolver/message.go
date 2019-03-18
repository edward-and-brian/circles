package resolver

import (
	"circles/server/model"
	"circles/server/store"
	"circles/server/types"
	"context"
)

// CreateMessage creates a new Message with the given data
func (r *Resolver) CreateMessage(ctx context.Context, args *struct{ Input *types.Message }) (*model.MessageModel, error) {
	return model.CreateMessage(ctx, &store.GeneralStore{}, args.Input)
}

// DeleteMessage deletes the user specified by ID
func (r *Resolver) DeleteMessage(ctx context.Context, args *IDArgs) (*model.MessageModel, error) {
	return model.DeleteMessage(ctx, &store.GeneralStore{}, args.ID)
}

// Message retrieves the user specified by id
func (r *Resolver) Message(ctx context.Context, args *IDArgs) (*model.MessageModel, error) {
	return model.FindMessage(ctx, &store.GeneralStore{}, args.ID)
}

// UpdateMessage updates the user specified by ID with the given data
func (r *Resolver) UpdateMessage(ctx context.Context, args *struct{ Input *types.Message }) (*model.MessageModel, error) {
	return model.UpdateMessage(ctx, &store.GeneralStore{}, args.Input)
}
