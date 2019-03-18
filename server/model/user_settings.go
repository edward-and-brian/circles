package model

import (
	"time"

	"circles/server/types"

	graphql "github.com/graph-gophers/graphql-go"
)

// UserSettingsStore is an interface for database interaction by the UserSettingsModel
type UserSettingsStore interface{}

// UserSettingsModel is the resolvable struct for the UserSettings struct
type UserSettingsModel struct {
	us    *types.UserSettings
	store UserSettingsStore
}

// ID field resolver
func (r *UserSettingsModel) ID() graphql.ID {
	return graphql.ID(r.us.ID)
}

// UserID field resolver
func (r *UserSettingsModel) UserID() graphql.ID {
	return graphql.ID(r.us.UserID)
}

// CreatedAt field resolver
func (r *UserSettingsModel) CreatedAt() (graphql.Time, error) {
	t, err := time.Parse(time.RFC3339, r.us.CreatedAt)
	return graphql.Time{Time: t}, err
}
