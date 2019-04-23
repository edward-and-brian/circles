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
		return nil, fmt.Errorf("error with AllUsers: %v", err.Error())
	}

	return users, nil
}

// AllUsersByChatMembership finds all User entries for a given Chat in the db
func (gs *GeneralStore) AllUsersByChatMembership(ctx context.Context, chid string) ([]*types.User, error) {
	var users []*types.User
	if err := gs.sqlite.Select(&users, allUsersByChatIDSQL, chid); err != nil {
		return nil, fmt.Errorf("error with AllUsersByChatMembership: %v", err.Error())
	}

	return users, nil
}

// CreateUser creates a User entry in the db
func (gs *GeneralStore) CreateUser(ctx context.Context, user *types.User) error {
	if err := gs.createEntity(ctx, user, createUserSQL); err != nil {
		return fmt.Errorf("error with CreateUser: %v", err.Error())
	}

	return nil
}

// DeleteUser deletes a User entry in the db
func (gs *GeneralStore) DeleteUser(ctx context.Context, id string) error {
	if err := gs.deleteEntity(ctx, usersTable, id); err != nil {
		return fmt.Errorf("error with DeleteUser: %v", err.Error())
	}

	return nil
}

// FindUser finds a User entry in the db
func (gs *GeneralStore) FindUser(ctx context.Context, id string) (*types.User, error) {
	user := &types.User{}

	if err := gs.findEntity(ctx, user, usersTable, id); err != nil {
		return nil, fmt.Errorf("error with FindUser: %v", err.Error())
	}

	return user, nil
}

// LoginUser finds a User entry in the db (using their login information)
func (gs *GeneralStore) LoginUser(ctx context.Context, phoneNumber string, displayName string) (*types.User, error) {
	user := &types.User{}

	loginUserSQL := `SELECT * FROM users WHERE phone_number=$1 AND display_name=$2`

	if err := gs.sqlite.Get(user, loginUserSQL, phoneNumber, displayName); err != nil {
		return nil, fmt.Errorf("error with LoginUser: %v", err.Error())
	}

	return user, nil
}

// UpdateUser updates a User entry in the db
func (gs *GeneralStore) UpdateUser(ctx context.Context, user *types.User) error {
	if err := gs.updateEntity(ctx, user, udpateUserSQL); err != nil {
		return fmt.Errorf("error with UpdateUser: %v", err.Error())
	}

	return nil
}
