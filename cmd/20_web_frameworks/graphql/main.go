package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jaygaha/go-beginner/cmd/20_web_frameworks/graphql/handlers"
)

func main() {
	router := mux.NewRouter()

	// register routes
	router.HandleFunc("/books", handlers.GraphQLPostHandler).Methods("POST")

	// GraphQL Playground (optional, for testing)
	// testing queries in the browser:
	router.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "playground.html")
	})

	// start server
	log.Println("Server started on :8800")
	log.Fatal(http.ListenAndServe(":8800", router))
}
