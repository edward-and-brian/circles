package resolver

import (
	"circles/server/model"
	"circles/server/store"
	"circles/server/types"
	"context"
)

// CreateUser creates a new User with the given data
func (r *Resolver) CreateUser(ctx context.Context, args *struct{ Input *types.User }) (*model.UserModel, error) {
	return model.CreateUser(ctx, &store.GeneralStore{}, args.Input)
}

// DeleteUser deletes the user specified by ID
func (r *Resolver) DeleteUser(ctx context.Context, args *IDArgs) (*model.UserModel, error) {
	return model.DeleteUser(ctx, &store.GeneralStore{}, args.ID)
}

// UpdateUser updates the user specified by ID with the given data
func (r *Resolver) UpdateUser(ctx context.Context, args *struct{ Input *model.UpdateUserInput }) (*model.UserModel, error) {
	return model.UpdateUser(ctx, &store.GeneralStore{}, args.Input)
}

// User retrieves the user specified by id
func (r *Resolver) User(ctx context.Context, args *IDArgs) (*model.UserModel, error) {
	return model.FindUser(ctx, &store.GeneralStore{}, args.ID)
}

// Users retrieves all users
func (r *Resolver) Users(ctx context.Context) ([]*model.UserModel, error) {
	return model.AllUsers(ctx, &store.GeneralStore{})
}
