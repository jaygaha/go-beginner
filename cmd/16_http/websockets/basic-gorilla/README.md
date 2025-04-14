# Go WebSockets

This tutorial introduces `WebSockets` in `Go` using the `Gorilla` WebSocket package. WebSockets provide full-duplex communication channels over a single TCP connection, enabling real-time applications with persistent connections.

This tutorial project demonstrates the power and simplicity of building real-time applications with `Go` and `WebSockets`. Feel free to experiment with the code and build upon it to create your own real-time applications!

## WebSockets vs HTTP

- **HTTP**: Stateless protocol where each request is independent
- **WebSockets**: Persistent connection maintained until client or server closes it
- **Benefits**: Reduced latency, improved performance, real-time bidirectional communication

## Use Cases

- Chat applications
- Multiplayer games
- Real-time data feeds (stock prices, weather updates)
- Live notifications
- Collaborative editing tools

## Project Structure

```
./
├── client.go         # Client implementation
├── go.mod            # Go module file
├── go.sum            # Go module checksum
├── main.go           # Main application entry point
├── server-room.go    # Room implementation for handling multiple clients
└── templates/
    └── chat.tmpl     # HTML template for the chat interface
```

## Key Components

### 1. Client (client.go)

The `client` struct represents a single connected user:

```go
type client struct {
    socket  *websocket.Conn // WebSocket connection
    receive chan []byte     // Channel to receive messages
    room    *room           // Reference to the room
}
```

Each client has two main methods:
- `read()`: Reads messages from the WebSocket and forwards them to the room
- `write()`: Writes messages from the room to the WebSocket

### 2. Room (server-room.go)

The `room` struct manages multiple clients and message broadcasting:

```go
type room struct {
    clients map[*client]bool // Map of connected clients
    forward chan []byte      // Channel for messages to broadcast
    join    chan *client     // Channel for new clients joining
    leave   chan *client     // Channel for clients leaving
}
```

The `run()` method handles client connections and message broadcasting using Go's select statement for concurrent operations.

### 3. Main Application (main.go)

The main application:
- Sets up HTTP routes
- Creates a template handler for the chat interface
- Initializes and runs the chat room
- Starts the HTTP server

## Running the Application

1. Clone the repository
2. Navigate to the project directory
3. Run the application:
   ```
   go run .
   ```
4. Open your browser and navigate to `http://localhost:8800`

## How It Works

1. When a user connects to `/rooms`, a WebSocket connection is established
2. A new client is created and added to the room
3. Messages sent by any client are broadcast to all connected clients
4. When a client disconnects, they are removed from the room

## Key Go Concepts Demonstrated

- **Goroutines**: Used for concurrent handling of client connections
- **Channels**: Used for communication between goroutines
- **Select Statement**: Used for handling multiple channel operations
- **HTTP Handlers**: Used for handling HTTP requests
- **Templates**: Used for rendering HTML

## Extending the Application

Here are some ideas to enhance this basic chat application:

1. Add user authentication
2. Implement multiple chat rooms
3. Add message persistence using a database
4. Add typing indicators
5. Support file sharing

## Resources for Learning More

- [Gorilla WebSocket Documentation](https://pkg.go.dev/github.com/gorilla/websocket)
- [Go Web Examples](https://gowebexamples.com/websockets/)

## Troubleshooting

- If you encounter connection issues, ensure no other application is using port 8800
- Check browser console for WebSocket errors
- Verify that the Gorilla WebSocket package is properly installed

Happy Coding! 🚀
