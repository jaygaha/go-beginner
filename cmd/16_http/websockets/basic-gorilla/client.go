package main

import "github.com/gorilla/websocket"

/*
	client
	- client is a program that connects to a server and sends requests to it.
	- client is a program that connects to a server and receives responses from it.
	- represent a single user of a websocket application.
*/

type client struct {
	socket  *websocket.Conn // socket is the web socket for this client.
	receive chan []byte     // receive is a channel to receive messages from the client.
	room    *room           // room is the room this client is chatting in.
}

// read reads messages from the client and forwards them to the room
func (c *client) read() {
	defer c.socket.Close()

	// read messages from the client and forward them to the room continuously
	for {
		_, msg, err := c.socket.ReadMessage() // read message from the client
		if err != nil {
			return
		}

		c.room.forward <- msg // forward message to the room
	}
}

// write writes messages from the room to the client
func (c *client) write() {
	defer c.socket.Close()

	// write messages from the room to the client continuously
	for msg := range c.receive {
		err := c.socket.WriteMessage(websocket.TextMessage, msg) // write message to the client
		if err != nil {
			return
		}
	}
}

// // read reads messages from the client and forwards them to the room
// func (c *client) read() {
// 	defer c.socket.Close()

// 	// read messages from the client and forward them to the room continuously
// 	for {
// 		if _, msg, err := c.socket.ReadMessage(); err == nil {
// 			c.room.forward <- msg
// 		} else {
// 			break
// 		}
// 	}
// }

// // write writes messages from the room to the client
// func (c *client) write() {
// 	defer c.socket.Close()

// 	// write messages from the room to the client continuously
// 	for msg := range c.receive {
// 		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
// 			break
// 		}
// 	}
// }
