package resolver

import (
	"circles/server/store"
	"circles/server/types"
	"context"
	"fmt"

	"github.com/rs/xid"
)

// CreateUser ...
func (r *Resolver) CreateUser(ctx context.Context, args *struct{ Input *types.User }) (*UserResolver, error) {
	var (
		err error
		db  *store.SqlxStore

		user = args.Input
	)

	if v := ctx.Value(DBkey); v != nil {
		db = v.(*store.SqlxStore)
	} else {
		return nil, fmt.Errorf("ctx.db is nil")
	}

	user.ID = xid.New().String()
	if err = db.CreateUser(user); err != nil {
		return nil, err

	} else if user, err = db.FindUser(&user.ID); err != nil {
		return nil, err
	}

	return &UserResolver{user, ctx}, nil
}

// UpdateUser ...
func (r *Resolver) UpdateUser(ctx context.Context, args *struct{ Input *types.User }) (*UserResolver, error) {
	var (
		err error
		db  *store.SqlxStore

		user = args.Input
	)

	if v := ctx.Value(DBkey); v != nil {
		db = v.(*store.SqlxStore)
	} else {
		return nil, fmt.Errorf("ctx.db is nil")
	}

	if err = db.UpdateUser(user); err != nil {
		return nil, err

	} else if user, err = db.FindUser(&user.ID); err != nil {
		return nil, err
	}

	return &UserResolver{user, ctx}, nil
}

// DeleteUser                                                                                                                                                                                                                        leteUser ...
func (r *Resolver) DeleteUser(ctx context.Context, args *IDArgs) (*UserResolver, error) {
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

	} else if err = db.DeleteUser(&args.ID); err != nil {
		return nil, err
	}

	return &UserResolver{user, ctx}, nil
}
