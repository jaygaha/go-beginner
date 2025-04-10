package middleware

import (
	"net/http"
)

// Custom middleware to check if the request has a valid username and password
func BasicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		username := request.FormValue("username")
		password := request.FormValue("password")

		// Check if the username and password are correct
		// If not, return a 401 Unauthorized response
		// else call next handler in chain
		// you can check against a database or any other source of truth
		if username != "admin" || password != "admin" {
			writer.WriteHeader(http.StatusUnauthorized)
			writer.Write([]byte("Unauthorized"))

			return
		}

		// Call the next handler in the chain if the username and password are correct
		next.ServeHTTP(writer, request)
	})
}

// Using the BasicAuth() method to check if the request has a valid username and password
func BasicAuthRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// If the username and password are sent via header, you can use the BasicAuth() method
		username, password, ok := request.BasicAuth()
		if !ok || username != "admin" || password != "admin" {
			writer.WriteHeader(http.StatusUnauthorized)
			writer.Write([]byte("Unauthorized"))

			return
		}

		// Call the next handler in the chain if the username and password are correct
		next.ServeHTTP(writer, request)
	})
}

// Using authorization header to check if the request has a valid username and password
func AuthorizationHeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// If the username and password are sent via header, you can use the BasicAuth() method
		authorizationHeader := request.Header.Get("Authorization")
		if authorizationHeader != "Basic secret-token" {
			writer.WriteHeader(http.StatusUnauthorized)
			writer.Write([]byte("Unauthorized"))

			return
		}

		// Call the next handler in the chain if the username and password are correct
		next.ServeHTTP(writer, request)
	})
}
