package resolver

// Resolver is the entry point for all top-level operations.
type Resolver struct{}

// GetRootResolver returns a new empty resolver struct
func GetRootResolver() *Resolver {
	return &Resolver{}
}

// IDArgs is used as an argument for multipe resolver functions
type IDArgs struct {
	ID string
}
