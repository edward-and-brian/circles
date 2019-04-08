package store

import (
	"context"
	"fmt"

	"circles/server/types"

	"github.com/rs/xid"
)

var (
	membershipTable = table("memberships")

	createMembershipSQL = `
	INSERT INTO memberships (id, user_id, chat_id)
	VALUES (:id, :user_id, :chat_id)`
)

// CreateMembership adds an entry to the memberships table with the given user_id and chat_id
func (gs *GeneralStore) CreateMembership(ctx context.Context, uid, chid string) error {
	if err := gs.findEntity(ctx, &types.Chat{}, chatsTable, chid); err != nil {
		return err

	} else if err := gs.findEntity(ctx, &types.User{}, usersTable, uid); err != nil {
		return err
	}

	membershipsEntry := map[string]interface{}{
		"id":      xid.New().String(),
		"user_id": uid,
		"chat_id": chid,
	}

	if err := gs.createEntity(ctx, membershipsEntry, createMembershipSQL); err != nil {
		return fmt.Errorf("error with CreateMembership: %v", err.Error())
	}

	return nil
}
