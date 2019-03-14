package handler

import (
	"net/http"
)

// GraphiQL struct used to fullfil GraphiQL requests
type GraphiQL struct{}

func (h *GraphiQL) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "handler/graphiql.html")
}
