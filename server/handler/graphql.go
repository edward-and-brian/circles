package handler

import (
	"circles/server/types"
	"context"
	"encoding/json"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/rs/xid"
)

// GraphQL struct used to fullfil GraphQL requests
type GraphQL struct {
	Schema *graphql.Schema
}

func (h *GraphQL) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var params struct {
		Query         string                 `json:"query"`
		OperationName string                 `json:"operationName"`
		Variables     map[string]interface{} `json:"variables"`
	}

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rid := r.Header.Get("X-Request-ID")
	if rid == "" {
		rid = xid.New().String()
		r.Header.Set("X-Request-ID", rid)
	}

	ctx := context.WithValue(context.Background(), types.Key("request_id"), rid)
	response := h.Schema.Exec(ctx, params.Query, params.OperationName, params.Variables)
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}
