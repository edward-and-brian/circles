package store

import (
	"log"

	"github.com/jmoiron/sqlx"
	// sqlite necessary
	_ "github.com/mattn/go-sqlite3"
)

// SqliteStore defines an implementation of model store interfaces
type SqliteStore struct {
	*sqlx.DB
}

// Open opens a new session with SQLite
func (ss *SqliteStore) Open() (err error) {
	ss.DB, err = sqlx.Open("sqlite3", "circles.db")

	if err != nil {
		log.Fatal(err)

	} else if err = ss.DB.Ping(); err != nil {
		log.Fatal(err)
	}

	return nil
}
