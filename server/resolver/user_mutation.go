package resolver

import (
	"circles/server/types"
	"context"

	"github.com/rs/xid"
)

// CreateUser creates a new User with the given data
func (r *Resolver) CreateUser(ctx context.Context, args *struct{ Input *types.User }) (*UserResolver, error) {
	user := args.Input
	user.ID = xid.New().String()

	if err := r.Store.Open(); err != nil {
		return nil, err

	} else if err = r.Store.CreateUser(user); err != nil {
		return nil, err

	} else if user, err = r.Store.FindUser(&user.ID); err != nil {
		return nil, err
	}

	return &UserResolver{user, r.Store}, nil
}

// UpdateUser updates the user specified by ID with the given data
func (r *Resolver) UpdateUser(ctx context.Context, args *struct{ Input *types.User }) (*UserResolver, error) {
	user := args.Input

	if err := r.Store.Open(); err != nil {
		return nil, err

	} else if err = r.Store.UpdateUser(user); err != nil {
		return nil, err

	} else if user, err = r.Store.FindUser(&user.ID); err != nil {
		return nil, err
	}

	return &UserResolver{user, r.Store}, nil
}

// DeleteUser deletes the user specified by ID                                                                                                                                                                                                      leteUser ...
func (r *Resolver) DeleteUser(ctx context.Context, args *IDArgs) (*UserResolver, error) {
	var user *types.User

	if err := r.Store.Open(); err != nil {
		return nil, err

	} else if user, err = r.Store.FindUser(&args.ID); err != nil {
		return nil, err

	} else if err = r.Store.DeleteUser(&args.ID); err != nil {
		return nil, err
	}

	return &UserResolver{user, r.Store}, nil
}
