package resolver

import (
	"time"

	"circles/server/types"

	graphql "github.com/graph-gophers/graphql-go"
)

type MessageResolver struct {
	m *types.Message
}

func (r *MessageResolver) ID() graphql.ID {
	return graphql.ID(r.m.ID)
}

func (r *MessageResolver) ChatID() graphql.ID {
	return graphql.ID(r.m.ChatID)
}

func (r *MessageResolver) CircleID() graphql.ID {
	return graphql.ID(r.m.CircleID)
}

func (r *MessageResolver) SenderID() graphql.ID {
	return graphql.ID(r.m.SenderID)
}

func (r *MessageResolver) Content() string {
	return r.m.Content
}

func (r *MessageResolver) CreatedAt() (graphql.Time, error) {
	t, err := time.Parse(time.RFC3339, r.m.CreatedAt)
	return graphql.Time{Time: t}, err
}
