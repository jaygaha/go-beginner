package middleware

import (
	"log"
	"net/http"
)

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// log.Printf("Request: %s %s\n", request.Method, request.URL.Path)
		log.Printf("Request path: %s, request id: %s\n", request.URL.Path, request.Context().Value("X-Request-Id"))

		// call next handler in chain
		next.ServeHTTP(writer, request)
	})
}

func LogBasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("User %s hit the endpoint %s\n", request.FormValue("username"), request.URL.Path)

		// call next handler in chain
		next.ServeHTTP(writer, request)
	})
}

func LogAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("Request: %s %s", request.Method, request.URL.Path)

		// call next handler in chain
		next.ServeHTTP(writer, request)
	})
}
