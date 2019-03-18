package store

import (
	"github.com/jmoiron/sqlx"
	// sqlite necessary
	_ "github.com/mattn/go-sqlite3"
)

// GeneralStore is a struct used to allow dependency injection of different stores
type GeneralStore struct {
	sqlite *sqlx.DB
}

// OpenSQLite opens a new session with SQLite
func (gs *GeneralStore) OpenSQLite() error {
	if gs.sqlite == nil {
		if db, err := sqlx.Open("sqlite3", "circles.db"); err != nil {
			return err

		} else if err = db.Ping(); err != nil {
			return err

		} else {
			gs.sqlite = db
		}
	}

	return nil
}
