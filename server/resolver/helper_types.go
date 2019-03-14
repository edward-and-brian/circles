package resolver

type key string

// DBkey used to access db from context
var DBkey key = "db"

// IDArgs provides a struct for graphql query arguments
type IDArgs struct {
	ID string
}
