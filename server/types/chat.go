package types

// Chat type
type Chat struct {
	ID        string
	Name      string
	CreatedAt string `db:"created_at"`
	// Avatar
	Circles []*Circle
}
