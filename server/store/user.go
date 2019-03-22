package store

import (
	"context"
	"fmt"

	"circles/server/types"
)

// AllUsers finds all User entries in the db
func (gs *GeneralStore) AllUsers(ctx context.Context) ([]*types.User, error) {
	var (
		users   []*types.User
		userSQL = `SELECT * FROM users ORDER BY id ASC`
	)

	if err := gs.OpenSQLite(ctx); err != nil {
		return nil, err

	} else if err = gs.sqlite.Select(&users, userSQL); err != nil {
		return nil, err
	}

	return users, nil
}

// CreateUser creates a User entry in the db
func (gs *GeneralStore) CreateUser(ctx context.Context, user *types.User) error {
	userSQL := `
	INSERT INTO users (id, name, phone_number, display_name) 
	VALUES (:id, :name, :phone_number, :display_name)`

	if err := gs.OpenSQLite(ctx); err != nil {
		return err

	} else if r, err := gs.sqlite.NamedExec(userSQL, user); err != nil {
		return err

	} else if count, err := r.RowsAffected(); err != nil {
		return err

	} else if count == 0 {
		return fmt.Errorf("CreateUser error: user was not created")
	}

	return nil
}

// DeleteUser deletes a User entry in the db
func (gs *GeneralStore) DeleteUser(ctx context.Context, id string) error {
	userSQL := `
	DELETE FROM users
	WHERE id=$id`

	if err := gs.OpenSQLite(ctx); err != nil {
		return err

	} else if count, err := gs.sqlite.MustExec(userSQL, id).RowsAffected(); err != nil {
		return err

	} else if count == 0 {
		return fmt.Errorf("DeleteUser error: user was not deleted")
	}
	return nil
}

// FindUser finds a User entry in the db
func (gs *GeneralStore) FindUser(ctx context.Context, id string) (*types.User, error) {
	var (
		user    = &types.User{}
		userSQL = `SELECT * FROM users WHERE id=$1`
	)

	if err := gs.OpenSQLite(ctx); err != nil {
		return nil, err

	} else if err = gs.sqlite.Get(user, userSQL, id); err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUser updates a User entry in the db
func (gs *GeneralStore) UpdateUser(ctx context.Context, user *types.User) error {
	userSQL := `
	UPDATE users 
	SET name=:name, 
		phone_number=:phone_number,
		display_name=:display_name 
	WHERE id=:id`

	if err := gs.OpenSQLite(ctx); err != nil {
		return err

	} else if r, err := gs.sqlite.NamedExec(userSQL, user); err != nil {
		return err

	} else if count, err := r.RowsAffected(); err != nil {
		return err

	} else if count == 0 {
		return fmt.Errorf("UpdateUser error: user was not updated")
	}

	return nil
}
