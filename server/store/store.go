package store

import (
	"circles/server/types"
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	// sqlite necessary
	_ "github.com/mattn/go-sqlite3"
)

type table string

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
	if rid, err := gs.getRequestID(ctx); err != nil {
		return err

	} else if tx := gs.sqliteTxs[rid]; tx != nil {
		err := tx.Commit()
		gs.sqliteTxs[rid] = nil
		return err
	}

	return nil
}

func (gs *GeneralStore) exec(ctx context.Context, query string, args ...interface{}) error {
	if rid, err := gs.getRequestID(ctx); err != nil {
		return err

	} else if tx := gs.sqliteTxs[rid]; tx != nil {
		if _, err := tx.Exec(query, args...); err != nil {
			tx.Rollback()
			gs.sqliteTxs[rid] = nil
			return err
		}
		return nil
	}

	_, err := gs.sqlite.Exec(query, args...)
	return err
}

func (gs *GeneralStore) namedExec(ctx context.Context, query string, arg interface{}) error {
	if rid, err := gs.getRequestID(ctx); err != nil {
		return err

	} else if tx := gs.sqliteTxs[rid]; tx != nil {
		if _, err := tx.NamedExec(query, arg); err != nil {
			tx.Rollback()
			gs.sqliteTxs[rid] = nil
			return err
		}

		return nil
	}

	_, err := gs.sqlite.NamedExec(query, arg)
	return err
}

func (gs *GeneralStore) get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	if rid, err := gs.getRequestID(ctx); err != nil {
		return err

	} else if tx := gs.sqliteTxs[rid]; tx != nil {
		if err := tx.Get(dest, query, args...); err != nil {
			tx.Rollback()
			gs.sqliteTxs[rid] = nil
			return err
		}
		return nil
	}

	return gs.sqlite.Get(dest, query, args...)
}

func (gs *GeneralStore) getRequestID(ctx context.Context) (string, error) {
	if val := ctx.Value(types.Key("request_id")); val == nil {
		return "", fmt.Errorf("could not retrieve request_id from context in GeneralStore.getRequestID")

	} else if rid, ok := val.(string); ok {
		return rid, nil
	}

	return "", fmt.Errorf("could not use request_id as type string")
}

func (gs *GeneralStore) getSet(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	if rid, err := gs.getRequestID(ctx); err != nil {
		return err

	} else if tx := gs.sqliteTxs[rid]; tx != nil {
		if err := tx.Select(dest, query, args...); err != nil {
			tx.Rollback()
			gs.sqliteTxs[rid] = nil
			log.Fatal(err)
		}
		return nil
	}

	return gs.sqlite.Select(dest, query, args...)
}
