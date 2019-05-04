package store

import (
	"context"
	"fmt"

	"circles/server/types"
)

var (
	circlesTable = table("circles")

	allCirclesByChatIDSQL = `SELECT * FROM circles WHERE chat_id=$1 ORDER BY last_message_at DESC`

	createCircleSQL = `
	INSERT INTO circles (id, chat_id, name, last_message_content, last_message_at, created_at)
	VALUES (:id, :chat_id, :name, :last_message_content, :last_message_at, :created_at)`

	updateCircleSQL = `
	UPDATE circles 
	SET name=:name,
		last_message_content=:last_message_content,
		last_message_at=:last_message_at
	WHERE id=:id`
)

// AllCircles finds all Circle entries in the db
func (gs *GeneralStore) AllCircles(ctx context.Context) ([]*types.Circle, error) {
	var circles []*types.Circle
	if err := gs.findAllEntity(ctx, &circles, circlesTable); err != nil {
		return nil, fmt.Errorf("AllCircles error: %v", err.Error())
	}

	return circles, nil
}

// AllCirclesByChatID finds all Circle entries for a given Chat in the db
func (gs *GeneralStore) AllCirclesByChatID(ctx context.Context, chid string) ([]*types.Circle, error) {
	var circles []*types.Circle
	if err := gs.getSet(ctx, &circles, allCirclesByChatIDSQL, chid); err != nil {
		return nil, fmt.Errorf("AllCirclesByChatID error: %v", err.Error())
	}

	return circles, nil
}

// CreateCircle creates a Circle entry in the db
func (gs *GeneralStore) CreateCircle(ctx context.Context, circle *types.Circle) error {
	if err := gs.findEntity(ctx, &types.Chat{}, chatsTable, circle.ChatID); err != nil {
		return err
	}

	if err := gs.createEntity(ctx, circle, createCircleSQL); err != nil {
		return fmt.Errorf("CreateCircle error: %v", err.Error())
	}

	return nil
}

// DeleteCircle deletes a Circle entry in the db
func (gs *GeneralStore) DeleteCircle(ctx context.Context, id string) error {
	if err := gs.deleteEntity(ctx, circlesTable, id); err != nil {
		return fmt.Errorf("DeleteCircle error: %v", err.Error())
	}

	return nil
}

// FindCircle finds a Circle entry in the db
func (gs *GeneralStore) FindCircle(ctx context.Context, id string) (*types.Circle, error) {
	circle := &types.Circle{}
	if err := gs.findEntity(ctx, circle, circlesTable, id); err != nil {
		return nil, fmt.Errorf("FindCircle error: %v", err.Error())
	}

	return circle, nil
}

// UpdateCircle updates a Circle entry in the db
func (gs *GeneralStore) UpdateCircle(ctx context.Context, circle *types.Circle) error {
	if err := gs.updateEntity(ctx, circle, updateCircleSQL); err != nil {
		return fmt.Errorf("UpdateCircle error: %v", err.Error())
	}

	return nil
}
