package handlers

import (
	"fmt"
	"net/http"
)

type SalutationStrct struct{}

func (receiver SalutationStrct) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello, Gopher\n")
}
