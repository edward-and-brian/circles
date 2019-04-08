package resolver

import (
	"context"

	"circles/server/model"
	"circles/server/types"
)

// Circle retrieves the chat specified by id
func (r *Resolver) Circle(ctx context.Context, args *IDArgs) (*model.CircleModel, error) {
	return model.FindCircle(ctx, r.store, args.ID)
}

// Circles retrieves all users
func (r *Resolver) Circles(ctx context.Context) ([]*model.CircleModel, error) {
	return model.AllCircles(ctx, r.store)
}

// CreateCircle creates a new Circle with the given data
func (r *Resolver) CreateCircle(ctx context.Context, args *struct{ Circle *types.Circle }) (*model.CircleModel, error) {
	return model.CreateCircle(ctx, r.store, args.Circle)
}

// DeleteCircle deletes the Circle specified by ID
func (r *Resolver) DeleteCircle(ctx context.Context, args *IDArgs) (*model.CircleModel, error) {
	return model.DeleteCircle(ctx, r.store, args.ID)
}

// UpdateCircle updates the Circle specified by ID with the given data
func (r *Resolver) UpdateCircle(ctx context.Context, args *struct{ Circle *model.UpdateCircleInput }) (*model.CircleModel, error) {
	return model.UpdateCircle(ctx, r.store, args.Circle)
}
