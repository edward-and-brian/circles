package model

import (
	"context"
	"time"

	"circles/server/types"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/rs/xid"
)

// AllCircles retrieves all circles and returns them as CircleModels
func AllCircles(ctx context.Context, gs generalStore) ([]*CircleModel, error) {
	circles, err := gs.AllCircles(ctx)
	if err != nil {
		return nil, err
	}

	var circleModels []*CircleModel
	for _, circle := range circles {
		circleModels = append(circleModels, &CircleModel{circle, gs})
	}

	return circleModels, nil
}

// CreateCircle creates a new Circle with the given data and returns it as a CircleModel
func CreateCircle(ctx context.Context, gs generalStore, circle *types.Circle) (*CircleModel, error) {
	circle.ID = xid.New().String()

	if err := gs.CreateCircle(ctx, circle); err != nil {
		return nil, err

	} else if circle, err = gs.FindCircle(ctx, circle.ID); err != nil {
		return nil, err
	}

	return &CircleModel{circle, gs}, nil
}

// DeleteCircle deletes the Circle specified by ID and returns it as a CircleModel
func DeleteCircle(ctx context.Context, gs generalStore, id string) (*CircleModel, error) {
	circle, err := gs.FindCircle(ctx, id)
	if err != nil {
		return nil, err

	} else if err = gs.DeleteCircle(ctx, id); err != nil {
		return nil, err
	}

	return &CircleModel{circle, gs}, nil
}

// FindCircle retrieves the Circle specified by id
func FindCircle(ctx context.Context, gs generalStore, id string) (*CircleModel, error) {
	circle, err := gs.FindCircle(ctx, id)
	if err != nil {
		return nil, err
	}

	return &CircleModel{circle, gs}, nil
}

// UpdateCircleInput ...
type UpdateCircleInput struct {
	ID   string
	Name *string
}

// UpdateCircle updates the circle specified by ID with the given data and returns it as a CircleModel
func UpdateCircle(ctx context.Context, gs generalStore, input *UpdateCircleInput) (*CircleModel, error) {
	circle, err := gs.FindCircle(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	if input.Name != nil {
		circle.Name = *input.Name
	}

	if err = gs.UpdateCircle(ctx, circle); err != nil {
		return nil, err
	}

	return &CircleModel{circle, gs}, nil
}

// ChatCircles retrieves all Circles for a Chat
func ChatCircles(ctx context.Context, gs generalStore, chid string) ([]*CircleModel, error) {
	circles, err := gs.AllCirclesByChatID(ctx, chid)
	if err != nil {
		return nil, err
	}

	var circleModels []*CircleModel
	for _, ch := range circles {
		circleModels = append(circleModels, &CircleModel{ch, gs})
	}

	return circleModels, nil
}

// CircleModel is the resolvable struct for the Circle struct
type CircleModel struct {
	*types.Circle
	store generalStore
}

// ID field resolver
func (c *CircleModel) ID() graphql.ID {
	return graphql.ID(c.Circle.ID)
}

// ChatID field resolver
func (c *CircleModel) ChatID() graphql.ID {
	return graphql.ID(c.Circle.ChatID)
}

// Name field resolver
func (c *CircleModel) Name() string {
	return c.Circle.Name
}

// CreatedAt field resolver
func (c *CircleModel) CreatedAt() (graphql.Time, error) {
	t, err := time.Parse(time.RFC3339, c.Circle.CreatedAt)
	return graphql.Time{Time: t}, err
}

// Messages field resolver
func (c *CircleModel) Messages(ctx context.Context) ([]*MessageModel, error) {
	return CircleMessages(ctx, c.store, c.Circle.ID)
}
