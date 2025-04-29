package main

import (
	"fmt"
	"net/http"

	"github.com/olahol/melody"
)

func main() {
	// Create an instance of Melody
	m := melody.New() // creates a new melody instance to manage WebSocket connections

	// handler for the chat window
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.tmpl")
	})

	// ws handler
	// provides a WebSocket endpoint for clients to connect to
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		err := m.HandleRequest(w, r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// broadcast message
	// handles incoming messages from clients
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		// broadcast the message to all connected clients
		err := m.Broadcast(msg)
		if err != nil {
			fmt.Printf("Error broadcasting message: %v\n", err)
			return
		}

	})

	// start the server
	fmt.Println("Server started on :8800")
	err := http.ListenAndServe(":8800", nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}
}
