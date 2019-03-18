package resolver

import (
	"context"
	"time"

	"circles/server/types"

	graphql "github.com/graph-gophers/graphql-go"
)

// UserResolver is the resolvable struct for the User struct
type UserResolver struct {
	u     *types.User
	store store
}

// ID field resolver
func (r *UserResolver) ID() graphql.ID {
	return graphql.ID(r.u.ID)
}

// Name field resolver
func (r *UserResolver) Name() string {
	return r.u.Name
}

// PhoneNumber field resolver
func (r *UserResolver) PhoneNumber() string {
	return r.u.PhoneNumber
}

// DisplayName field resolver
func (r *UserResolver) DisplayName() string {
	return r.u.DisplayName
}

// CreatedAt field resolver
func (r *UserResolver) CreatedAt() (graphql.Time, error) {
	t, err := time.Parse(time.RFC3339, r.u.CreatedAt)
	return graphql.Time{Time: t}, err
}

// Chats field resolver
func (r *UserResolver) Chats(ctx context.Context) ([]*ChatResolver, error) {
	return nil, nil
}
