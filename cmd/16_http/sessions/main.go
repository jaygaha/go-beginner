package main

import (
	"fmt"
	"net/http"

	"github.com/jaygaha/go-beginner/tree/main/cmd/16_http/sessions/handlers"
)

/*
Sessions:
- Session is a collection of key-value pairs
- Session is a way to store data between requests


Cookies:
- Cookie is a small piece of data that is stored on the client side
- Cookie is a way to store data between requests

Sessions and Cookies:
 - useful for storing data between requests
 - authentication
 - shopping cart
 - user preferences
 - language preference

 Packages:
 - github.com/gorilla/sessions
 - github.com/gorilla/securecookie
 - github.com/gorilla/mux
*/

func main() {
	//  register a handlers
	http.HandleFunc("/", handlers.HomeHandler)

	// cookies
	http.HandleFunc("/cookie/set", handlers.SetCookie)
	http.HandleFunc("/cookie/get", handlers.GetCookie)
	http.HandleFunc("/cookie/delete", handlers.DeleteCookie)

	// sessions
	http.HandleFunc("/auth/login", handlers.Login)
	http.HandleFunc("/user/profile", handlers.Profile)
	http.HandleFunc("/auth/logout", handlers.Logout)

	// start the server
	fmt.Println("Server starting at port 8800...")

	// http.ListenAndServe(":8080", nil)
	err := http.ListenAndServe(":8800", nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
