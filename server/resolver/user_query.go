package resolver

import (
	"circles/server/types"
	"context"
)

// User retrieves the user specified by id
func (r *Resolver) User(ctx context.Context, args *IDArgs) (*UserResolver, error) {
	var user *types.User

	if err := r.Store.Open(); err != nil {
		return nil, err

	} else if user, err = r.Store.FindUser(&args.ID); err != nil {
		return nil, err
	}

	return &UserResolver{user, r.Store}, nil
}

// Users retrieves all users
func (r *Resolver) Users(ctx context.Context) ([]*UserResolver, error) {
	var (
		users         []*types.User
		usersResolver []*UserResolver
	)

	if err := r.Store.Open(); err != nil {
		return nil, err

	} else if users, err = r.Store.AllUsers(); err != nil {
		return nil, err
	}

	for _, u := range users {
		usersResolver = append(usersResolver, &UserResolver{u, r.Store})
	}

	return usersResolver, nil
}
