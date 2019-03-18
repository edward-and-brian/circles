package resolver

import (
	"context"

	"circles/server/model"
	"circles/server/store"
	"circles/server/types"
)

// Circle retrieves the chat specified by id
func (r *Resolver) Circle(ctx context.Context, args *IDArgs) (*model.CircleModel, error) {
	return model.FindCircle(ctx, &store.GeneralStore{}, args.ID)
}

// CreateCircle creates a new Circle with the given data
func (r *Resolver) CreateCircle(ctx context.Context, args *struct{ Input *types.Circle }) (*model.CircleModel, error) {
	return model.CreateCircle(ctx, &store.GeneralStore{}, args.Input)
}

// DeleteCircle deletes the Circle specified by ID
func (r *Resolver) DeleteCircle(ctx context.Context, args *IDArgs) (*model.CircleModel, error) {
	return model.DeleteCircle(ctx, &store.GeneralStore{}, args.ID)
}

// UpdateCircle updates the Circle specified by ID with the given data
func (r *Resolver) UpdateCircle(ctx context.Context, args *struct{ Input *types.Circle }) (*model.CircleModel, error) {
	return model.UpdateCircle(ctx, &store.GeneralStore{}, args.Input)
}
