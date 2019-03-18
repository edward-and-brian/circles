package resolver

import (
	"time"

	"circles/server/types"

	graphql "github.com/graph-gophers/graphql-go"
)

// MessageResolver ...
type MessageResolver struct {
	m *types.Message
}

// ID field resolver
func (r *MessageResolver) ID() graphql.ID {
	return graphql.ID(r.m.ID)
}

// ChatID field resolver
func (r *MessageResolver) ChatID() graphql.ID {
	return graphql.ID(r.m.ChatID)
}

// CircleID field resolver
func (r *MessageResolver) CircleID() graphql.ID {
	return graphql.ID(r.m.CircleID)
}

// SenderID field resolver
func (r *MessageResolver) SenderID() graphql.ID {
	return graphql.ID(r.m.SenderID)
}

// Content field resolver
func (r *MessageResolver) Content() string {
	return r.m.Content
}

// CreatedAt field resolver
func (r *MessageResolver) CreatedAt() (graphql.Time, error) {
	t, err := time.Parse(time.RFC3339, r.m.CreatedAt)
	return graphql.Time{Time: t}, err
}
