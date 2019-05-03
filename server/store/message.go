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
	INSERT INTO messages (id, circle_id, sender_id, content, created_at)
	VALUES (:id, :circle_id, :sender_id, :content, :created_at)`

	updateMessageSQL = `UPDATE messages SET content=:content WHERE id=:id`
)

// AllMessages finds all Message entries in the db
func (gs *GeneralStore) AllMessages(ctx context.Context) ([]*types.Message, error) {
	var messages []*types.Message
	if err := gs.findAllEntity(ctx, &messages, messagesTable); err != nil {
		return nil, fmt.Errorf("AllMessages error: %v", err.Error())
	}

	return messages, nil
}

// AllMessagesByCircleID finds all Message entries for a given Circle in the db
func (gs *GeneralStore) AllMessagesByCircleID(ctx context.Context, ciid string) ([]*types.Message, error) {
	var messages []*types.Message
	if err := gs.getSet(ctx, &messages, allMessagesByCircleIDSQL, ciid); err != nil {
		return nil, fmt.Errorf("AllMessagesByCircleID error: %v", err.Error())
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
		return fmt.Errorf("CreateMessage error: %v", err.Error())
	}

	return nil
}

// DeleteMessage deletes a Message entry in the db
func (gs *GeneralStore) DeleteMessage(ctx context.Context, id string) error {
	if err := gs.deleteEntity(ctx, messagesTable, id); err != nil {
		return fmt.Errorf("DeleteMessage error: %v", err.Error())
	}

	return nil
}

// DeleteMessagesByCircleID deletes any messages entries with the given circle_id
func (gs *GeneralStore) DeleteMessagesByCircleID(ctx context.Context, ciid string) error {
	deleteMessagesByCircleIDSQL := `DELETE FROM messages WHERE circle_id=$1`

	if err := gs.exec(ctx, deleteMessagesByCircleIDSQL, ciid); err != nil {
		return fmt.Errorf("DeleteMessagesByCircleID error: %v", err)
	}

	return nil
}

// FindMessage finds a Message entry in the db
func (gs *GeneralStore) FindMessage(ctx context.Context, id string) (*types.Message, error) {
	message := &types.Message{}
	if err := gs.findEntity(ctx, message, messagesTable, id); err != nil {
		return nil, fmt.Errorf("FindMessage error: %v", err.Error())
	}

	return message, nil
}

// UpdateMessage updates a Message entry in the db
func (gs *GeneralStore) UpdateMessage(ctx context.Context, message *types.Message) error {
	if err := gs.updateEntity(ctx, message, updateMessageSQL); err != nil {
		return fmt.Errorf("UpdateMessage error: %v", err.Error())
	}

	return nil
}
