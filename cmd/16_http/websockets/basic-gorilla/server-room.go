package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type room struct {
	clients map[*client]bool // clients is a map of clients in the room.
	forward chan []byte      // forward is a channel that holds incoming messages that should be forwarded to the other clients.
	join    chan *client     // join is a channel for clients wishing to join the room.
	leave   chan *client     // leave is a channel for clients wishing to leave the room.
}

// newRoom makes a new room.
func newRoom() *room {
	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
	}
}

// run runs our room. It listens for various events and takes action on them.
func (r *room) run() {
	for {
		select {
		// <-r.join: describes the join channel, which is a channel that receives a client when a client joins the room.
		case client := <-r.join:
			// joining
			r.clients[client] = true
		// <-r.leave: denotes the leave channel, which is a channel that receives a client when a client leaves the room.
		case client := <-r.leave:
			// leaving
			delete(r.clients, client)
			close(client.receive)
		// <-r.forward: describes the forward channel, which is a channel that receives a message when a message is sent to the room.
		case msg := <-r.forward:
			// forward message to all clients
			for client := range r.clients {
				client.receive <- msg
			}
		}
	}
}

const (
	socketBufferSize  = 1024 // 1KB which means that we can send 1024 bytes at a time.
	messageBufferSize = 256  // 256 bytes which means that we can send 256 messages at a time.
)

// upgrader is used to upgrade the HTTP connection to a websocket connection. http -> ws protocol
// ReadBufferSize and WriteBufferSize are the size of the buffer used to read and write messages.
// CheckOrigin is a function that is used to check if the origin of the request is allowed to connect to the websocket.

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  socketBufferSize,
	WriteBufferSize: socketBufferSize,
	CheckOrigin:     func(r *http.Request) bool { return true }, // Allow all origins
}

// handler
func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// upgrade HTTP connection to a websocket connection.
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("Upgrade fails: ", err)
		return
	}

	// create a new client
	client := &client{
		socket:  socket,
		receive: make(chan []byte, messageBufferSize),
		room:    r,
	}
	// register the client in the room
	r.join <- client

	defer func() { r.leave <- client }()

	// start the client's writer
	go client.write()

	// start the client's reader
	client.read()
}
