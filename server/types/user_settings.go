package types

// UserSettings type
type UserSettings struct {
	ID        string
	UserID    string `db:"user_id"`
	CreatedAt string `db:"created_at"`
}
