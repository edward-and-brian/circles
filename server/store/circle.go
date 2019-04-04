package store

import (
	"context"
	"fmt"

	"circles/server/types"
)

var (
	circlesTable = table("circles")

	allCirclesByChatIDSQL = `SELECT * FROM circles WHERE chat_id=$1 ORDER BY id ASC`

	createCircleSQL = `
	INSERT INTO circles (id, chat_id, name)
	VALUES (:id, :chat_id, :name)`

	updateCircleSQL = `UPDATE circles SET name=:name WHERE id=:id`
)

// AllCircles finds all Circle entries in the db
func (gs *GeneralStore) AllCircles(ctx context.Context) ([]*types.Circle, error) {
	var circles []*types.Circle
	if err := gs.findAllEntity(ctx, &circles, circlesTable); err != nil {
		return nil, fmt.Errorf("error with AllCircles: %v", err.Error())
	}

	return circles, nil
}

// AllCirclesByChatID finds all Circle entries for a given Chat in the db
func (gs *GeneralStore) AllCirclesByChatID(ctx context.Context, chid string) ([]*types.Circle, error) {
	var circles []*types.Circle
	if err := gs.sqlite.Select(&circles, allCirclesByChatIDSQL, chid); err != nil {
		return nil, fmt.Errorf("error with AllCirclesByChatID: %v", err.Error())
	}

	return circles, nil
}

// CreateCircle creates a Circle entry in the db
func (gs *GeneralStore) CreateCircle(ctx context.Context, circle *types.Circle) error {
	if err := gs.findEntity(ctx, &types.Chat{}, chatsTable, circle.ChatID); err != nil {
		return err
	}

	if err := gs.createEntity(ctx, circle, createCircleSQL); err != nil {
		return fmt.Errorf("error with CreateCircle: %v", err.Error())
	}

	return nil
}

// DeleteCircle deletes a Circle entry in the db
func (gs *GeneralStore) DeleteCircle(ctx context.Context, id string) error {
	if err := gs.deleteEntity(ctx, circlesTable, id); err != nil {
		return fmt.Errorf("error with DeleteCircle: %v", err.Error())
	}

	return nil
}

// FindCircle finds a Circle entry in the db
func (gs *GeneralStore) FindCircle(ctx context.Context, id string) (*types.Circle, error) {
	circle := &types.Circle{}
	if err := gs.findEntity(ctx, circle, circlesTable, id); err != nil {
		return nil, fmt.Errorf("error with FindCircle: %v", err.Error())
	}

	return circle, nil
}

// UpdateCircle updates a Circle entry in the db
func (gs *GeneralStore) UpdateCircle(ctx context.Context, circle *types.Circle) error {
	if err := gs.updateEntity(ctx, circle, updateCircleSQL); err != nil {
		return fmt.Errorf("error with UpdateCircle: %v", err.Error())
	}

	return nil
}
