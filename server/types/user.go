package types

// User type
type User struct {
	ID          string
	Name        string
	PhoneNumber string `db:"phone_number"`
	DisplayName string `db:"display_name"`
	CreatedAt   string `db:"created_at"`
	// Avatar
	Chats []*Chat
}
