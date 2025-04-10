package main

import (
	"net/http"

	"github.com/jaygaha/go-beginner/tree/main/cmd/16_http/middleware/handlers"
	"github.com/jaygaha/go-beginner/tree/main/cmd/16_http/middleware/middleware"
)

/*
Middlware in Go

a middleware is a function that takes a http.handler as an input and processes the request before forwarding it to
the the given handler. This is especially handy, if you have multiple handlers that share some commen functionality.

- function wrapper for http handler
- executed before/after handler
- can modify request/response
- can terminate request
- can call next handler in chain
- can be chained
- can be reused
- can be used for logging, authentication, authorization, etc.
- can be used for profiling, monitoring, etc.
- can be used for caching, compression, etc.
*/

// we can create a list of middleware functions, which will be applied in reverse order
var middlewareList = []func(http.Handler) http.Handler{
	middleware.BasicAuthMiddleware,
	middleware.LogBasicAuth,
}

var middlewareBasicReqList = []func(http.Handler) http.Handler{
	middleware.BasicAuthRequestMiddleware,
	middleware.LogAuth,
}

var middlewareAuthorizationReqList = []func(http.Handler) http.Handler{
	middleware.AuthorizationHeaderMiddleware,
	middleware.LogAuth,
}

func main() {
	// Logging
	// Handle: it is a function that takes a http.Handler and returns a http.Handler
	// Log the request and manupulate by adding uuid to response header
	http.Handle("/", middleware.UUID(middleware.Log(handlers.WelcomeStrct{})))
	http.Handle("/salutation", middleware.UUID(middleware.Log(handlers.SalutationStrct{})))

	// Basic Auth
	authHandler := http.HandlerFunc(handlers.BasicAuthHandler)
	for _, m := range middlewareList {
		authHandler = http.HandlerFunc(m(authHandler).ServeHTTP)
	}

	http.Handle("/basic", authHandler)

	// BasicAuth in request
	basicAuthHandler := http.HandlerFunc(handlers.BasicAuthHandler)
	for _, m := range middlewareBasicReqList {
		basicAuthHandler = http.HandlerFunc(m(basicAuthHandler).ServeHTTP)
	}

	http.Handle("/basic-request", basicAuthHandler)

	// Using Authorization header
	AuthorizationHandler := http.HandlerFunc(handlers.AuthHandler)
	for _, m := range middlewareAuthorizationReqList {
		AuthorizationHandler = http.HandlerFunc(m(AuthorizationHandler).ServeHTTP)
	}

	http.Handle("/authorization", AuthorizationHandler)

	// Using panic
	// useful for handling errors or panic gracefully
	http.HandleFunc("/panic", middleware.RecoveryMiddleware(handlers.PanicHandler))

	http.ListenAndServe(":8080", nil)
}
