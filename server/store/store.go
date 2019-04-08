package store

import (
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	// sqlite necessary
	_ "github.com/mattn/go-sqlite3"
)

type table string
type key string

// GeneralStore is a struct used to allow dependency injection of different stores
type GeneralStore struct {
	sqlite    *sqlx.DB
	sqliteTxs map[string]*sqlx.Tx
}

// OpenSQLite opens a new session with SQLite
func (gs *GeneralStore) OpenSQLite() (err error) {
	gs.sqlite, err = sqlx.Connect("sqlite3", "circles.db")
	gs.sqliteTxs = map[string]*sqlx.Tx{}
	return
}

// BeginTransaction opens a unique connection to SQLite for a transaction
func (gs *GeneralStore) BeginTransaction(ctx context.Context) error {
	if rid, err := gs.getRequestID(ctx); err != nil {
		return err

	} else if tx, err := gs.sqlite.BeginTxx(ctx, nil); err != nil {
		return err

	} else {
		gs.sqliteTxs[rid] = tx
	}

	return nil
}

// EndTransaction closes a transaction
func (gs *GeneralStore) EndTransaction(ctx context.Context) error {
	if tx, err := gs.getRequestTransaction(ctx); err != nil {
		return err

	} else if tx != nil {
		err := tx.Commit()
		tx = nil
		return err
	}

	return nil
}

func (gs *GeneralStore) getRequestID(ctx context.Context) (string, error) {
	if val := ctx.Value("request_id"); val == nil {
		return "", fmt.Errorf("could not retrieve request_id from context in GeneralStore.getRequestID")

	} else if rid, ok := val.(string); ok {
		return rid, nil
	}

	return "", fmt.Errorf("could not use request_id as type string")
}

func (gs *GeneralStore) getRequestTransaction(ctx context.Context) (*sqlx.Tx, error) {
	rid, err := gs.getRequestID(ctx)
	if err != nil {
		return nil, err
	}

	return gs.sqliteTxs[rid], nil
}

// findAllOfEntity finds all entries in the db for the given entity
func (gs *GeneralStore) findAllEntity(ctx context.Context, entity interface{}, entityTable table) error {
	findAllEntitySQL := fmt.Sprintf(`SELECT * FROM %v ORDER BY id ASC`, entityTable)

	if tx, err := gs.getRequestTransaction(ctx); err != nil {
		return err

	} else if tx != nil {
		if err := tx.Select(entity, findAllEntitySQL); err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
		return nil
	}

	return gs.sqlite.Select(entity, findAllEntitySQL)
}

// CreateEntity creates a Entity entry in the db
func (gs *GeneralStore) createEntity(ctx context.Context, entity interface{}, createEntitySQL string) error {
	if tx, err := gs.getRequestTransaction(ctx); err != nil {
		return err

	} else if tx != nil {
		if r, err := tx.NamedExec(createEntitySQL, entity); err != nil {
			tx.Rollback()
			return err
		} else if count, err := r.RowsAffected(); err != nil {
			tx.Rollback()
			return err
		} else if count == 0 {
			tx.Rollback()
			return err
		}
		return nil
	}

	if r, err := gs.sqlite.NamedExec(createEntitySQL, entity); err != nil {
		return err
	} else if count, err := r.RowsAffected(); err != nil {
		return err
	} else if count == 0 {
		return fmt.Errorf("CreateEntity error: message was not created")
	}

	return nil
}

func (gs *GeneralStore) deleteEntity(ctx context.Context, entityTable table, entityID string) error {
	deleteEntitySQL := fmt.Sprintf(`DELETE FROM %v WHERE id=$id`, entityTable)

	if tx, err := gs.getRequestTransaction(ctx); err != nil {
		return err

	} else if tx != nil {
		if count, err := tx.MustExec(deleteEntitySQL, entityID).RowsAffected(); err != nil {
			tx.Rollback()
			return err
		} else if count == 0 {
			tx.Rollback()
			return err
		}
		return nil
	}

	if count, err := gs.sqlite.MustExec(deleteEntitySQL, entityID).RowsAffected(); err != nil {
		return err
	} else if count == 0 {
		return fmt.Errorf("DeleteEntity error: entity was not deleted")
	}

	return nil
}

func (gs *GeneralStore) findEntity(ctx context.Context, entity interface{}, entityTable table, id string) error {
	findEntitySQL := fmt.Sprintf(`SELECT * FROM %v WHERE id=$1`, entityTable)

	if tx, err := gs.getRequestTransaction(ctx); err != nil {
		return err

	} else if tx != nil {
		if err := tx.Get(entity, findEntitySQL, id); err != nil {
			tx.Rollback()
			return err
		}
		return nil
	}

	return gs.sqlite.Get(entity, findEntitySQL, id)
}

func (gs *GeneralStore) updateEntity(ctx context.Context, entity interface{}, updateEntitySQL string) error {
	if tx, err := gs.getRequestTransaction(ctx); err != nil {
		return err

	} else if tx != nil {
		if r, err := tx.NamedExec(updateEntitySQL, entity); err != nil {
			tx.Rollback()
			return err
		} else if count, err := r.RowsAffected(); err != nil {
			tx.Rollback()
			return err
		} else if count == 0 {
			tx.Rollback()
			return err
		}
		return nil
	}

	if r, err := gs.sqlite.NamedExec(updateEntitySQL, entity); err != nil {
		return err
	} else if count, err := r.RowsAffected(); err != nil {
		return err
	} else if count == 0 {
		return fmt.Errorf("UpdateEntity error: message was not updated")
	}

	return nil
}
