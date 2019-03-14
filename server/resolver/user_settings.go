package resolver

import (
	"time"

	"circles/server/types"

	graphql "github.com/graph-gophers/graphql-go"
)

type UserSettingsResolver struct {
	us *types.UserSettings
}

func (r *UserSettingsResolver) ID() graphql.ID {
	return graphql.ID(r.us.ID)
}

func (r *UserSettingsResolver) UserID() graphql.ID {
	return graphql.ID(r.us.UserID)
}

func (r *UserSettingsResolver) CreatedAt() (graphql.Time, error) {
	t, err := time.Parse(time.RFC3339, r.us.CreatedAt)
	return graphql.Time{Time: t}, err
}
