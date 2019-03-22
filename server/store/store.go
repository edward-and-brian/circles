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
func (gs *GeneralStore) OpenSQLite() (err error) {
	gs.sqlite, err = sqlx.Connect("sqlite3", "circles.db")
	return err
}

// // SQLiteUpdateTableEntry is the generalized update function for sqlite entries
// func (gs *GeneralStore) SQLiteUpdateTableEntry(ctx context.Context, table, id string, updateFields map[string]interface{}) error {
// 	setFields := []string{}
// 	for key, value := range updateFields {
// 		setFields = append(setFields, fmt.Sprintf("%v:%v", key, value))
// 	}

// 	updateSQL := fmt.Sprintf("UPDATE %v SET %v WHERE id=%v", table, setFields, id)

// 	if err := gs.OpenSQLite(); err != nil {
// 		return err

// 	} else if count, err := gs.sqlite.MustExec(updateSQL).RowsAffected(); err != nil {
// 		return err

// 	} else if count == 0 {
// 		return fmt.Errorf("UpdateUser error: user was not updated")
// 	}

// 	return nil
// }
