package resolver

import (
	"context"
	"time"

	"circles/server/types"

	graphql "github.com/graph-gophers/graphql-go"
)

// UserResolver is the resolvable struct for the User struct
type UserResolver struct {
	u   *types.User
	ctx context.Context
}

// ID ...
func (r *UserResolver) ID() graphql.ID {
	return graphql.ID(r.u.ID)
}

// Name ...
func (r *UserResolver) Name() string {
	return r.u.Name
}

// PhoneNumber ...
func (r *UserResolver) PhoneNumber() string {
	return r.u.PhoneNumber
}

// DisplayName ...
func (r *UserResolver) DisplayName() string {
	return r.u.DisplayName
}

// CreatedAt ...
func (r *UserResolver) CreatedAt() (graphql.Time, error) {
	t, err := time.Parse(time.RFC3339, r.u.CreatedAt)
	return graphql.Time{Time: t}, err
}

// Chats searches the db for a list of User chats
func (r *UserResolver) Chats() ([]*ChatResolver, error) {
	var (
		// err   error
		chats []*ChatResolver

		// db = r.ctx.Value("db").(*store.SqlxStore)
	)

	// if chats, err = db.FindUserChats(r.u.ID); err != nil {
	// 	return nil, err
	// }

	return chats, nil
}
