package main

import (
	"context"
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
	db, err := store.OpenDB()
	if err != nil {
		log.Fatalf("Unable to connect to db: %s \n", err)
	}

	ctx := context.WithValue(context.Background(), resolver.DBkey, db)
	schema := graphql.MustParseSchema(schema.GetRootSchema(), resolver.GetRootResolver(), graphql.UseFieldResolvers())

	mux := http.NewServeMux()
	mux.Handle("/", &handler.GraphiQL{})
	mux.Handle("/query", handler.AddContext(ctx, &handler.GraphQL{Schema: schema}))

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
