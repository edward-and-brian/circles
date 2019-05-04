package types

// Chat type
type Chat struct {
	ID             string
	Name           string
	LastCircleName string `db:"last_circle_name"`
	LastMessageAt  string `db:"last_message_at"`
	CreatedAt      string `db:"created_at"`
	// Avatar
	Circles []*Circle
}
