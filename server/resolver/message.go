package resolver

import (
	"circles/server/model"
	"circles/server/types"
	"context"
)

// CreateMessage creates a new Message with the given data
func (r *Resolver) CreateMessage(ctx context.Context, args *struct{ Message *types.Message }) (*model.MessageModel, error) {
	return model.CreateMessage(ctx, r.store, args.Message)
}

// DeleteMessage deletes the user specified by ID
func (r *Resolver) DeleteMessage(ctx context.Context, args *IDArgs) (*model.MessageModel, error) {
	return model.DeleteMessage(ctx, r.store, args.ID)
}

// Message retrieves the user specified by id
func (r *Resolver) Message(ctx context.Context, args *IDArgs) (*model.MessageModel, error) {
	return model.FindMessage(ctx, r.store, args.ID)
}

// Messages retrieves all users
func (r *Resolver) Messages(ctx context.Context) ([]*model.MessageModel, error) {
	return model.AllMessages(ctx, r.store)
}

// MessageDatePartitionsByCircle retrieves the DatePartitions for the messages of a given Circle.
func (r *Resolver) MessageDatePartitionsByCircle(ctx context.Context, args *IDArgs) ([]*model.DatePartitionModel, error) {
	return model.CircleMessageDatePartitions(ctx, r.store, args.ID)
}

// UpdateMessage updates the user specified by ID with the given data
func (r *Resolver) UpdateMessage(ctx context.Context, args *struct{ Message *model.UpdateMessageInput }) (*model.MessageModel, error) {
	return model.UpdateMessage(ctx, r.store, args.Message)
}
