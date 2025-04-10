package middleware

import (
	"log"
	"net/http"
	"runtime/debug"
)

func RecoveryMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// recover from panic
		// recover(): it serves as Go in-built try catch
		// allows a program to manage behavior of a paninking goroutine
		// avoiding program crash

		// Defer a function to catch any panics
		defer func() {
			if err := recover(); err != nil {
				// Log the panic and stack trace
				msg := "Caught panic: %v, Stack trace: %s"
				log.Printf(msg, err, string(debug.Stack()))

				// Return a generic error message to the client
				er := http.StatusInternalServerError
				http.Error(writer, http.StatusText(er), er)
			}
		}()

		// Call the next handler in the chain
		next(writer, request)
	})
}
