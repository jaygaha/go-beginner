package main

import (
	"fmt"
	"net/http"

	"github.com/jaygaha/go_beginner/cmd/20_web_frameworks/rest/routes"
)

func main() {
	// define routes
	http.HandleFunc("/bags", routes.RegisterRoutes)

	// start server
	fmt.Println("Server started on :8800")
	if err := http.ListenAndServe(":8800", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
