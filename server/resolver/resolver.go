package resolver

// Resolver is the entry point for all top-level operations.
type Resolver struct{}

// GetRootResolver returns a new Resolver
func GetRootResolver() *Resolver {
	return &Resolver{}
}
