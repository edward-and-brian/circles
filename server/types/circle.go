package types

// Circle type
type Circle struct {
	ID        string
	ChatID    string `db:"chat_id"`
	Name      string
	CreatedAt string `db:"created_at"`
	// Avatar
	Messages []*Message
}
