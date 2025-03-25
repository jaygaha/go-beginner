package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
)

type key string

const CusKeyServerAddr key = "CUSTOM_SERVER_ADDR" // to act as the key for the HTTP server address

func main() {
	// simple http server
	// StartServer()

	// multiplexing requests server
	StartMultiplexingServer()

	// Running multiple servers at the same time
	// StartMultipleServers()
}

// StartServer starts a simple http server
func StartServer() {
	fmt.Println("Starting server...")

	// register handlers (routes)
	http.HandleFunc("/", welcomeHandler)
	// uri with query params example: /hello?name=ジェイ
	http.HandleFunc("/hello", helloHandler)

	err := http.ListenAndServe(":8010", nil) // start server on port 8010

	if errors.Is(err, http.ErrServerClosed) { // server closed
		fmt.Printf("Server closed\n")
	} else if err != nil { // other error
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}

// StartMultiplexingServer starts a simple http server with multiplexing
func StartMultiplexingServer() {
	fmt.Println("Starting multiplexing server on 8010...")

	// create a new mux
	mux := http.NewServeMux() // it creates a new mux and returns a pointer to it
	// register handlers (routes)
	mux.HandleFunc("/", welcomeHandler)
	mux.HandleFunc("POST /hello", helloPostHandler)
	// uri with path params example: /hello/ジェイ
	mux.HandleFunc("GET /hello/{name}", helloGetHandler)
	err := http.ListenAndServe(":8010", mux) // start server on port 8010

	if errors.Is(err, http.ErrServerClosed) { // server closed
		fmt.Printf("Server closed\n")
	} else if err != nil { // other error
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}

// StartMultipleServers starts multiple servers at the same time
func StartMultipleServers() {
	fmt.Println("Starting multiple server at once on 8010...")

	// create a new mux
	mux := http.NewServeMux() // it creates a new mux and returns a pointer to it
	// register handlers (routes)
	mux.HandleFunc("/", welcomeMultipleHandler)
	ctx, cancelCtx := context.WithCancel(context.Background())

	// start server 1
	serverOne := &http.Server{
		Addr:    ":8010",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, CusKeyServerAddr, l.Addr().String())
			return ctx
		},
	}

	// start server 2
	serverTwo := &http.Server{
		Addr:    ":8011",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, CusKeyServerAddr, l.Addr().String())
			return ctx
		},
	}

	// start server 1 in a goroutine
	go func() {
		err := serverOne.ListenAndServe()

		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("Server 1 closed\n")
		} else if err != nil {
			fmt.Printf("Error starting server 1: %s\n", err)
		}
		cancelCtx()
	}()

	// start server 2 in a goroutine
	go func() {
		err := serverTwo.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("Server 2 closed\n")
		} else if err != nil {
			fmt.Printf("Error starting server 2: %s\n", err)
		}
		cancelCtx()
	}()

	<-ctx.Done() // wait for the context to be canceled
}

/*
Handlers are functions that take a response writer and a request as arguments
and return nothing.
Response writer is used to write the response to the client.
Request is used to get the request data (body, headers, url).
*/
func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to go-phers world!") // fprintf writes to the response writer
}

// handler with query params
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// r.URL.Query() returns a map of query params
	// r.URL.Query().Get("name") returns the value of the "name" query param
	// r.URL.Query().has() returns true if the key exists
	fmt.Fprintf(w, "Hello, %s!", r.URL.Query().Get("name"))
}

// handler with post request
func helloPostHandler(w http.ResponseWriter, r *http.Request) {
	// get name param from request body; if empty, return validation error
	name := r.PostFormValue("name")
	if name == "" {
		http.Error(w, "Name is required", http.StatusUnprocessableEntity)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Hello, %s!", name)
}

// handler with path params
func helloGetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Hello, %s!", name)
}

// handler with multiple servers
func welcomeMultipleHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // get context from request

	serverAddr := ctx.Value(CusKeyServerAddr).(string) // get server address from context
	fmt.Fprintf(w, "Welcome to go-phers world! from server %s", serverAddr)
}
