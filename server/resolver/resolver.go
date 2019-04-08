package resolver

import "circles/server/store"

// Resolver is the entry point for all top-level operations.
type Resolver struct {
	store *store.GeneralStore
}

// GetRootResolver returns a new empty resolver struct
func GetRootResolver() *Resolver {
	store := &store.GeneralStore{}
	if err := store.OpenSQLite(); err != nil {
		panic(err)
	}
	return &Resolver{store}
}

// IDArgs is used as an argument for multipe resolver functions
type IDArgs struct {
	ID string
}
