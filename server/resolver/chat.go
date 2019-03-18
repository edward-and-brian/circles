package resolver

import (
	"context"
	"time"

	"circles/server/types"

	graphql "github.com/graph-gophers/graphql-go"
)

// ChatResolver is the resolvable struct for the Chat struct
type ChatResolver struct {
	ch    *types.Chat
	store store
}

// ID field resolver
func (r *ChatResolver) ID() graphql.ID {
	return graphql.ID(r.ch.ID)
}

// Name field resolver
func (r *ChatResolver) Name() string {
	return r.ch.Name
}

// CreatedAt field resolver
func (r *ChatResolver) CreatedAt() (graphql.Time, error) {
	t, err := time.Parse(time.RFC3339, r.ch.CreatedAt)
	return graphql.Time{Time: t}, err
}

// Circles field resolver
func (r *ChatResolver) Circles(ctx context.Context) []*CircleResolver {
	return []*CircleResolver{}
}
