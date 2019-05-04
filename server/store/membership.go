package store

import (
	"context"
	"fmt"
	"time"

	"circles/server/types"

	"github.com/rs/xid"
)

var (
	membershipTable = table("memberships")

	createMembershipSQL = `
	INSERT INTO memberships (id, user_id, chat_id, created_at)
	VALUES (:id, :user_id, :chat_id, :created_at)`
)

// CreateMembership adds an entry to the memberships table with the given user_id and chat_id
func (gs *GeneralStore) CreateMembership(ctx context.Context, uid, chid string) error {
	if err := gs.findEntity(ctx, &types.Chat{}, chatsTable, chid); err != nil {
		return err

	} else if err := gs.findEntity(ctx, &types.User{}, usersTable, uid); err != nil {
		return err
	}

	membershipsEntry := map[string]interface{}{
		"id":         xid.New().String(),
		"user_id":    uid,
		"chat_id":    chid,
		"created_at": time.Now().Format(time.RFC3339),
	}

	if err := gs.createEntity(ctx, membershipsEntry, createMembershipSQL); err != nil {
		return fmt.Errorf("CreateMembership error: %v", err.Error())
	}

	return nil
}

// DeleteMembershipsByChatID deletes any membership entries with the given chat_id
func (gs *GeneralStore) DeleteMembershipsByChatID(ctx context.Context, chid string) error {
	deleteMembershipsByChatIDSQL := `DELETE FROM memberships WHERE chat_id=$1`

	if err := gs.exec(ctx, deleteMembershipsByChatIDSQL, chid); err != nil {
		return fmt.Errorf("DeleteMembershipsByChatID error: %v", err.Error())
	}

	return nil
}

// DeleteMembershipsByUserID deletes any membership entries with the given user_id
func (gs *GeneralStore) DeleteMembershipsByUserID(ctx context.Context, uid string) error {
	deleteMembershipsByUserIDSQL := `DELETE FROM memberships WHERE chat_id=$1`

	if err := gs.exec(ctx, deleteMembershipsByUserIDSQL, uid); err != nil {
		return fmt.Errorf("DeleteMembershipsByUserID error: %v", err.Error())
	}

	return nil
}
