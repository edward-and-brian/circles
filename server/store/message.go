package store

import (
	"context"
	"fmt"

	"circles/server/types"
)

// AllMessagesByCircleID finds all Message entries for a given Circle in the db
func (gs *GeneralStore) AllMessagesByCircleID(ctx context.Context, ciid string) ([]*types.Message, error) {
	var (
		messages   []*types.Message
		messageSQL = `SELECT * FROM messages ORDER BY id ASC`
	)

	if err := gs.OpenSQLite(); err != nil {
		return nil, err

	} else if err = gs.sqlite.Select(&messages, messageSQL); err != nil {
		return nil, err
	}

	return messages, nil
}

// CreateMessage creates a Message entry in the db
func (gs *GeneralStore) CreateMessage(ctx context.Context, u *types.Message) error {
	messageSQL := `
	INSERT INTO messages (id, circle_id, sender_id, content)
	VALUES (:id, :circle_id, :sender_id, :content)`

	if err := gs.OpenSQLite(); err != nil {
		return err

	} else if r, err := gs.sqlite.NamedExec(messageSQL, u); err != nil {
		return err

	} else if count, err := r.RowsAffected(); err != nil {
		return err

	} else if count == 0 {
		return fmt.Errorf("CreateMessage error: message was not created")
	}

	return nil
}

// DeleteMessage deletes a Message entry in the db
func (gs *GeneralStore) DeleteMessage(ctx context.Context, id string) error {
	messageSQL := `DELETE FROM messages WHERE id=$id`

	if err := gs.OpenSQLite(); err != nil {
		return err

	} else if count, err := gs.sqlite.MustExec(messageSQL, id).RowsAffected(); err != nil {
		return err

	} else if count == 0 {
		return fmt.Errorf("DeleteMessage error: message was not deleted")
	}
	return nil
}

// FindMessage finds a Message entry in the db
func (gs *GeneralStore) FindMessage(ctx context.Context, id string) (*types.Message, error) {
	var (
		message    = &types.Message{}
		messageSQL = `SELECT * FROM messages WHERE id=$1`
	)

	if err := gs.OpenSQLite(); err != nil {
		return nil, err

	} else if err = gs.sqlite.Get(message, messageSQL, id); err != nil {
		return nil, err
	}

	return message, nil
}

// UpdateMessage updates a Message entry in the db
func (gs *GeneralStore) UpdateMessage(ctx context.Context, message *types.Message) error {
	messageSQL := `UPDATE messages SET content=:content WHERE id=:id`

	if err := gs.OpenSQLite(); err != nil {
		return err

	} else if r, err := gs.sqlite.NamedExec(messageSQL, message); err != nil {
		return err

	} else if count, err := r.RowsAffected(); err != nil {
		return err

	} else if count == 0 {
		return fmt.Errorf("UpdateMessage error: message was not updated")
	}

	return nil
}
