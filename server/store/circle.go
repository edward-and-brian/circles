package store

import (
	"context"
	"fmt"

	"circles/server/types"

	"github.com/rs/xid"
)

// AllCirclesByChatID finds all Circle entries for a given Chat in the db
func (gs *GeneralStore) AllCirclesByChatID(ctx context.Context, chid string) ([]*types.Circle, error) {
	var (
		circles   []*types.Circle
		circleSQL = `SELECT * FROM circles ORDER BY id ASC`
	)

	if err := gs.OpenSQLite(); err != nil {
		return nil, err

	} else if err = gs.sqlite.Select(&circles, circleSQL); err != nil {
		return nil, err
	}

	return circles, nil
}

// CreateCircle creates a Circle entry in the db
func (gs *GeneralStore) CreateCircle(ctx context.Context, circle *types.Circle) error {
	circle.ID = xid.New().String()

	circleSQL := `
	INSERT INTO circles (id, chat_id, name)
	VALUES (:id, :chat_id, :name)`

	if err := gs.OpenSQLite(); err != nil {
		return err

	} else if r, err := gs.sqlite.NamedExec(circleSQL, circle); err != nil {
		return err

	} else if count, err := r.RowsAffected(); err != nil {
		return err

	} else if count == 0 {
		return fmt.Errorf("CreateCircle error: circle was not created")
	}

	return nil
}

// DeleteCircle deletes a Circle entry in the db
func (gs *GeneralStore) DeleteCircle(ctx context.Context, id string) error {
	circleSQL := `DELETE FROM circles WHERE id=$id`

	if err := gs.OpenSQLite(); err != nil {
		return err

	} else if count, err := gs.sqlite.MustExec(circleSQL, id).RowsAffected(); err != nil {
		return err

	} else if count == 0 {
		return fmt.Errorf("DeleteCircle error: circle was not deleted")
	}
	return nil
}

// FindCircle finds a Circle entry in the db
func (gs *GeneralStore) FindCircle(ctx context.Context, id string) (*types.Circle, error) {
	var (
		circle    = &types.Circle{}
		circleSQL = `SELECT * FROM circles WHERE id=$1`
	)

	if err := gs.OpenSQLite(); err != nil {
		return nil, err

	} else if err = gs.sqlite.Get(circle, circleSQL, id); err != nil {
		return nil, err
	}

	return circle, nil
}

// UpdateCircle updates a Circle entry in the db
func (gs *GeneralStore) UpdateCircle(ctx context.Context, circle *types.Circle) error {
	circleSQL := `UPDATE circles SET name=:name WHERE id=:id`

	if err := gs.OpenSQLite(); err != nil {
		return err

	} else if r, err := gs.sqlite.NamedExec(circleSQL, circle); err != nil {
		return err

	} else if count, err := r.RowsAffected(); err != nil {
		return err

	} else if count == 0 {
		return fmt.Errorf("UpdateCircle error: circle was not updated")
	}

	return nil
}
