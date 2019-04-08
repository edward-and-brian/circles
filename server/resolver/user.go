package resolver

import (
	"circles/server/model"
	"circles/server/types"
	"context"
)

// CreateUser creates a new User with the given data
func (r *Resolver) CreateUser(ctx context.Context, args *struct{ User *types.User }) (*model.UserModel, error) {
	return model.CreateUser(ctx, r.store, args.User)
}

// DeleteUser deletes the user specified by ID
func (r *Resolver) DeleteUser(ctx context.Context, args *IDArgs) (*model.UserModel, error) {
	return model.DeleteUser(ctx, r.store, args.ID)
}

// UpdateUser updates the user specified by ID with the given data
func (r *Resolver) UpdateUser(ctx context.Context, args *struct{ User *model.UpdateUserInput }) (*model.UserModel, error) {
	return model.UpdateUser(ctx, r.store, args.User)
}

// User retrieves the user specified by id
func (r *Resolver) User(ctx context.Context, args *IDArgs) (*model.UserModel, error) {
	return model.FindUser(ctx, r.store, args.ID)
}

// Users retrieves all users
func (r *Resolver) Users(ctx context.Context) ([]*model.UserModel, error) {
	return model.AllUsers(ctx, r.store)
}
