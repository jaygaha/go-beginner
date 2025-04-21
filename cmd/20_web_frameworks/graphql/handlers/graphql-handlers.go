package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
	graphqlSchema "github.com/jaygaha/go-beginner/cmd/20_web_frameworks/graphql/graphql"
)

func GraphQLPostHandler(w http.ResponseWriter, r *http.Request) {
	// parse request from client
	var reqBody struct {
		Query string `json:"query"`
	}

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// execute GraphQL query
	result := graphql.Do(graphql.Params{
		Schema:        graphqlSchema.Schema,
		RequestString: reqBody.Query,
	})

	// send response back to client
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
