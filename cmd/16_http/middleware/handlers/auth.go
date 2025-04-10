package handlers

import (
	"fmt"
	"net/http"
)

type BasicAuthStrct struct{}

// takes a handlerFunc as an input and returns a handlerFunc
func BasicAuthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Authenticated Gopher from Basic Auth!\n")
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Authenticated Gopher from Auth!\n")
}
