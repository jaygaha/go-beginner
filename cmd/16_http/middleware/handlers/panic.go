package handlers

import "net/http"

func PanicHandler(writer http.ResponseWriter, request *http.Request) {
	var tmp *int
	*tmp += 1 // simulate a panic (500)
}
