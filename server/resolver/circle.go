package resolver

import (
	"context"
	"time"

	"circles/server/types"

	graphql "github.com/graph-gophers/graphql-go"
)

type CircleResolver struct {
	ci *types.Circle
}

func (r *CircleResolver) ID() graphql.ID {
	return graphql.ID(r.ci.ID)
}

func (r *CircleResolver) ChatID() graphql.ID {
	return graphql.ID(r.ci.ChatID)
}

func (r *CircleResolver) Name() string {
	return r.ci.Name
}

func (r *CircleResolver) CreatedAt() (graphql.Time, error) {
	t, err := time.Parse(time.RFC3339, r.ci.CreatedAt)
	return graphql.Time{Time: t}, err
}

func (r *CircleResolver) Messages(ctx context.Context) []*MessageResolver {
	// db := ctx.Value("db").(*store.SqlxStore)
	// db.getCircleMessages(r.ci.ID)

	l := make([]*MessageResolver, len(r.ci.Messages))
	for i := range l {
		l[i] = &MessageResolver{
			m: r.ci.Messages[i],
		}
	}
	return l
}
