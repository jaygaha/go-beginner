package handlers

import (
	"fmt"
	"net/http"
)

type WelcomeStrct struct{}

func (receiver WelcomeStrct) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Welcome to the Goland!\n")
}
