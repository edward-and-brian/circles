package store

import (
	"context"
	"fmt"

	"circles/server/types"
)

var (
	messagesTable = table("messages")

	allMessagesByCircleIDSQL = `SELECT * FROM messages WHERE circle_id=$1 ORDER BY datetime(created_at) ASC`

	createMessageSQL = `
	INSERT INTO messages (id, circle_id, sender_id, content)
	VALUES (:id, :circle_id, :sender_id, :content)`

	updateMessageSQL = `UPDATE messages SET content=:content WHERE id=:id`
)

// AllMessages finds all Message entries in the db
func (gs *GeneralStore) AllMessages(ctx context.Context) ([]*types.Message, error) {
	var messages []*types.Message
	if err := gs.findAllEntity(ctx, &messages, messagesTable); err != nil {
		return nil, fmt.Errorf("error with AllMessages: %v", err.Error())
	}

	return messages, nil
}

// AllMessagesByCircleID finds all Message entries for a given Circle in the db
func (gs *GeneralStore) AllMessagesByCircleID(ctx context.Context, ciid string) ([]*types.Message, error) {
	var messages []*types.Message
	if err := gs.sqlite.Select(&messages, allMessagesByCircleIDSQL, ciid); err != nil {
		return nil, fmt.Errorf("error with AllMessagesByCircleID: %v", err.Error())
	}

	return messages, nil
}

// CreateMessage creates a Message entry in the db
func (gs *GeneralStore) CreateMessage(ctx context.Context, message *types.Message) error {
	if err := gs.findEntity(ctx, &types.Circle{}, circlesTable, message.CircleID); err != nil {
		return err

	} else if err := gs.findEntity(ctx, &types.User{}, usersTable, message.SenderID); err != nil {
		return err
	}

	if err := gs.createEntity(ctx, message, createMessageSQL); err != nil {
		return fmt.Errorf("error with CreateMessage: %v", err.Error())
	}

	return nil
}

// DeleteMessage deletes a Message entry in the db
func (gs *GeneralStore) DeleteMessage(ctx context.Context, id string) error {
	if err := gs.deleteEntity(ctx, messagesTable, id); err != nil {
		return fmt.Errorf("error with DeleteMessage: %v", err.Error())
	}

	return nil
}

// FindMessage finds a Message entry in the db
func (gs *GeneralStore) FindMessage(ctx context.Context, id string) (*types.Message, error) {
	message := &types.Message{}
	if err := gs.findEntity(ctx, message, messagesTable, id); err != nil {
		return nil, fmt.Errorf("error with FindMessage: %v", err.Error())
	}

	return message, nil
}

// UpdateMessage updates a Message entry in the db
func (gs *GeneralStore) UpdateMessage(ctx context.Context, message *types.Message) error {
	if err := gs.updateEntity(ctx, message, updateMessageSQL); err != nil {
		return fmt.Errorf("error with UpdateMessage: %v", err.Error())
	}

	return nil
}
