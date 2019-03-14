package store

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	// sqlite necessary
	_ "github.com/mattn/go-sqlite3"
)

// SqlxStore defines an implementation of model store interfaces
type SqlxStore struct {
	*sqlx.DB
}

// OpenDB opens a new session with PSQL
func OpenDB() (*SqlxStore, error) {
	log.Println("Database is connecting... ")
	db, err := sqlx.Open("sqlite3", "circles.db")
	if err != nil {
		panic(err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Println("Retry database connection in 5 seconds... ")
		time.Sleep(time.Duration(5) * time.Second)
		return OpenDB()
	}

	log.Println("Database is connected ")
	return &SqlxStore{db}, nil
}
