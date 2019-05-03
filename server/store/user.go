package store

import (
	"context"
	"fmt"

	"circles/server/types"
)

var (
	usersTable = table("users")

	allUsersByChatIDSQL = `
	SELECT * 
	FROM users
	WHERE id IN (
		SELECT user_id 
		FROM memberships
		WHERE chat_id = $1
	)
	ORDER BY datetime(created_at) ASC`

	createUserSQL = `
	INSERT INTO users (id, name, phone_number, display_name, created_at) 
	VALUES (:id, :name, :phone_number, :display_name, :created_at)`

	loginUserSQL = `SELECT * FROM users WHERE phone_number=$1 AND display_name=$2`

	udpateUserSQL = `
	UPDATE users 
	SET name=:name, 
		phone_number=:phone_number,
		display_name=:display_name 
	WHERE id=:id`
)

// AllUsers finds all User entries in the db
func (gs *GeneralStore) AllUsers(ctx context.Context) ([]*types.User, error) {
	var users []*types.User
	if err := gs.findAllEntity(ctx, &users, usersTable); err != nil {
		return nil, fmt.Errorf("AllUsers error: %v", err.Error())
	}

	return users, nil
}

// AllUsersByChatMembership finds all User entries for a given Chat in the db
func (gs *GeneralStore) AllUsersByChatMembership(ctx context.Context, chid string) ([]*types.User, error) {
	var users []*types.User
	if err := gs.getSet(ctx, &users, allUsersByChatIDSQL, chid); err != nil {
		return nil, fmt.Errorf("AllUsersByChatMembership error: %v", err.Error())
	}

	return users, nil
}

// CreateUser creates a User entry in the db
func (gs *GeneralStore) CreateUser(ctx context.Context, user *types.User) error {
	if err := gs.createEntity(ctx, user, createUserSQL); err != nil {
		return fmt.Errorf("CreateUser error: %v", err.Error())
	}

	return nil
}

// DeleteUser deletes a User entry in the db
func (gs *GeneralStore) DeleteUser(ctx context.Context, id string) error {
	if err := gs.deleteEntity(ctx, usersTable, id); err != nil {
		return fmt.Errorf("DeleteUser error: %v", err.Error())
	}

	return nil
}

// FindUser finds a User entry in the db
func (gs *GeneralStore) FindUser(ctx context.Context, id string) (*types.User, error) {
	user := &types.User{}
	if err := gs.findEntity(ctx, user, usersTable, id); err != nil {
		return nil, fmt.Errorf("FindUser error: %v", err.Error())
	}

	return user, nil
}

// LoginUser finds a User entry in the db (using their login information)
func (gs *GeneralStore) LoginUser(ctx context.Context, phoneNumber string, displayName string) (*types.User, error) {
	user := &types.User{}
	if err := gs.get(ctx, user, loginUserSQL, phoneNumber, displayName); err != nil {
		return nil, fmt.Errorf("LoginUser error: %v", err.Error())
	}

	return user, nil
}

// UpdateUser updates a User entry in the db
func (gs *GeneralStore) UpdateUser(ctx context.Context, user *types.User) error {
	if err := gs.updateEntity(ctx, user, udpateUserSQL); err != nil {
		return fmt.Errorf("UpdateUser error: %v", err.Error())
	}

	return nil
}
