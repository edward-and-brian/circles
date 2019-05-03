package types

// Circle type
type Circle struct {
	ID                 string
	ChatID             string `db:"chat_id"`
	Name               string
	LastMessageContent string `db:"last_message_content"`
	LastMessageAt      string `db:"last_message_at"`
	CreatedAt          string `db:"created_at"`
	// Avatar
	Messages []*Message
}
