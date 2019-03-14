package resolver

import (
	"time"

	"circles/server/types"

	graphql "github.com/graph-gophers/graphql-go"
)

type ChatResolver struct {
	ch *types.Chat
}

func (r *ChatResolver) ID() graphql.ID {
	return graphql.ID(r.ch.ID)
}

func (r *ChatResolver) Name() string {
	return r.ch.Name
}

func (r *ChatResolver) CreatedAt() (graphql.Time, error) {
	t, err := time.Parse(time.RFC3339, r.ch.CreatedAt)
	return graphql.Time{Time: t}, err
}

func (r *ChatResolver) Circles() []*CircleResolver {
	l := make([]*CircleResolver, len(r.ch.Circles))
	for i := range l {
		l[i] = &CircleResolver{
			ci: r.ch.Circles[i],
		}
	}
	return l
}
