package store

import (
	"context"
	"fmt"

	"circles/server/types"

	"github.com/rs/xid"
)

// AddUserToChat adds an entry to the user_chats table with the given user_id and chat_id
func (gs *GeneralStore) AddUserToChat(ctx context.Context, uid, chid string) error {
	userChatEntry := map[string]interface{}{
		"id":      xid.New().String(),
		"user_id": uid,
		"chat_id": chid,
	}

	chatSQL := `
	INSERT INTO user_chats (id, user_id, chat_id)
	VALUES (:id, :user_id, :chat_id)`

	if err := gs.OpenSQLite(); err != nil {
		return err

	} else if r, err := gs.sqlite.NamedExec(chatSQL, userChatEntry); err != nil {
		return err

	} else if count, err := r.RowsAffected(); err != nil {
		return err

	} else if count == 0 {
		return fmt.Errorf("CreateChat error: chat was not created")
	}

	return nil
}

// AllChatsByUserID finds all Chat entries for a given User in the db
func (gs *GeneralStore) AllChatsByUserID(ctx context.Context, uid string) ([]*types.Chat, error) {
	var (
		chats   []*types.Chat
		chatSQL = `
		SELECT * 
		FROM chats
		WHERE id IN (
			SELECT chat_id 
			FROM user_chats
			WHERE user_id = $1
		)
		ORDER BY created_at ASC`
	)

	if err := gs.OpenSQLite(); err != nil {
		return nil, err

	} else if err = gs.sqlite.Select(&chats, chatSQL, uid); err != nil {
		return nil, err
	}

	return chats, nil
}

// CreateChat creates a Chat entry in the db
func (gs *GeneralStore) CreateChat(ctx context.Context, chat *types.Chat) error {
	chat.ID = xid.New().String()
	chatSQL := `
	INSERT INTO chats (id, name)
	VALUES (:id, :name)`

	if err := gs.OpenSQLite(); err != nil {
		return err

	} else if r, err := gs.sqlite.NamedExec(chatSQL, chat); err != nil {
		return err

	} else if count, err := r.RowsAffected(); err != nil {
		return err

	} else if count == 0 {
		return fmt.Errorf("CreateChat error: chat was not created")
	}

	return nil
}

// DeleteChat deletes a Chat entry in the db
func (gs *GeneralStore) DeleteChat(ctx context.Context, id string) error {
	chatSQL := `DELETE FROM chats WHERE id=$id`

	if err := gs.OpenSQLite(); err != nil {
		return err

	} else if count, err := gs.sqlite.MustExec(chatSQL, id).RowsAffected(); err != nil {
		return err

	} else if count == 0 {
		return fmt.Errorf("DeleteChat error: chat was not deleted")
	}
	return nil
}

// FindChat finds a Chat entry in the db
func (gs *GeneralStore) FindChat(ctx context.Context, id string) (*types.Chat, error) {
	var (
		chat    = &types.Chat{}
		chatSQL = `SELECT * FROM chats WHERE id=$1`
	)

	if err := gs.OpenSQLite(); err != nil {
		return nil, err

	} else if err = gs.sqlite.Get(chat, chatSQL, id); err != nil {
		return nil, err
	}

	return chat, nil
}

// UpdateChat updates a Chat entry in the db
func (gs *GeneralStore) UpdateChat(ctx context.Context, chat *types.Chat) error {
	chatSQL := `UPDATE chats SET name=:name WHERE id=:id`

	if err := gs.OpenSQLite(); err != nil {
		return err

	} else if r, err := gs.sqlite.NamedExec(chatSQL, chat); err != nil {
		return err

	} else if count, err := r.RowsAffected(); err != nil {
		return err

	} else if count == 0 {
		return fmt.Errorf("UpdateChat error: chat was not updated")
	}

	return nil
}
