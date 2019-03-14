package types

// Circle type
type Circle struct {
	ID        string
	ChatID    string
	Name      string
	CreatedAt string `db:"created_at"`
	// Avatar
	Messages []*Message
}
