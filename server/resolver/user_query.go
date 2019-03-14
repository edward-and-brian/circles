package resolver

import (
	"circles/server/store"
	"circles/server/types"
	"context"
	"fmt"
)

// User retrieves the user specified by id
func (r *Resolver) User(ctx context.Context, args *IDArgs) (*UserResolver, error) {
	var (
		err  error
		db   *store.SqlxStore
		user *types.User
	)

	if v := ctx.Value(DBkey); v != nil {
		db = v.(*store.SqlxStore)
	} else {
		return nil, fmt.Errorf("ctx.db is nil")
	}

	if user, err = db.FindUser(&args.ID); err != nil {
		return nil, err
	}

	return &UserResolver{user, ctx}, err
}

// Users retrieves all users
func (r *Resolver) Users(ctx context.Context) ([]*UserResolver, error) {
	var (
		err           error
		db            *store.SqlxStore
		users         []*types.User
		usersResolver []*UserResolver
	)

	if v := ctx.Value(DBkey); v != nil {
		db = v.(*store.SqlxStore)
	} else {
		return nil, fmt.Errorf("ctx.db is nil")
	}

	if users, err = db.AllUsers(); err != nil {
		return nil, err
	}

	for _, u := range users {
		usersResolver = append(usersResolver, &UserResolver{u, ctx})
	}

	return usersResolver, err
}
