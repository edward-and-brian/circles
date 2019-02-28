package types

// Message type
type Message struct {
	ID        string
	ChatID    string `db:"chat_id"`
	CircleID  string `db:"circle_id"`
	SenderID  string `db:"sender_id"`
	Content   string
	CreatedAt string `db:"created_at"`
}
