package main

import (
	"log"
	"net/http"
	"time"

	"circles/server/handler"
	"circles/server/resolver"
	"circles/server/schema"
	"circles/server/store"

	graphql "github.com/graph-gophers/graphql-go"
)

func main() {
	resolver := &resolver.Resolver{Store: &store.SqliteStore{}}
	schema := graphql.MustParseSchema(schema.GetRootSchema(), resolver)

	mux := http.NewServeMux()
	mux.Handle("/", &handler.GraphiQL{})
	mux.Handle("/query", &handler.GraphQL{Schema: schema})

	s := &http.Server{
		Addr:              ":3000",
		Handler:           mux,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       90 * time.Second,
		MaxHeaderBytes:    http.DefaultMaxHeaderBytes,
	}

	log.Fatal(s.ListenAndServe())
}
