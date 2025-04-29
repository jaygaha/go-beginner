# Basic Chat with Melody

This application demonstrates a simple real-time chat using WebSockets in Go with the Melody framework. It allows multiple clients to connect and exchange messages in real-time.

## What is Melody?

[Melody](https://github.com/olahol/melody) is a minimalist WebSocket framework for Go that makes it easy to build real-time applications. It provides a simple API for managing WebSocket connections and broadcasting messages.

Key features of Melody include:

- Simple API for handling WebSocket connections
- Broadcasting messages to all connected clients
- Session management for individual connections
- Support for binary and text messages
- Filters for selective broadcasting
- Middleware support

## How It Works

### Technologies Used

- **[Melody](https://github.com/olahol/melody)**: A minimalist WebSocket framework for Go
- **Go standard library**: For HTTP server and template rendering
- **HTML/CSS/JavaScript**: For the client interface

### Architecture

1. **WebSocket Server**: Melody manages WebSocket connections and broadcasts
2. **Client Interface**: Simple HTML/CSS interface with JavaScript for WebSocket communication

### Key Components

#### WebSocket Server

The application creates a Melody instance to manage WebSocket connections:

```go
m := melody.New() // Create new Melody instance
```

It sets up HTTP handlers for serving the chat interface and handling WebSocket connections:

```go
// Serve the chat interface
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "index.tmpl")
})

// Handle WebSocket connections
http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
    m.HandleRequest(w, r)
})
```

When a client sends a message, the server broadcasts it to all connected clients:

```go
// Broadcast messages to all clients
m.HandleMessage(func(s *melody.Session, msg []byte) {
    m.Broadcast(msg)
})
```

#### Client-Side WebSocket

The client establishes a WebSocket connection to the server:

```javascript
const ws = new WebSocket("ws://localhost:8800/ws");
```

It handles incoming messages from the server:

```javascript
ws.onmessage = (event) => {
    const row = document.createElement("div");
    row.innerHTML = event.data;
    document.getElementById("chat-window").appendChild(row);
};
```

And sends messages to the server when the user submits a message:

```javascript
sendButton.addEventListener("click", () => {
    const message = messageInput.value;
    if (message) {
        ws.send(message);
        messageInput.value = "";
    }
});
```

## Installation

### Steps

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd golang-beginner/cmd/20_web_frameworks/melody/basic_chat
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

4. Open your browser and navigate to:
   ```
   http://localhost:8800
   ```

## Usage

1. Open the application in multiple browser windows or tabs
2. Type a message in the input field and click "Send"
3. The message will appear in all connected browser windows

## Extending Melody

Melody provides several features that can be used to extend this basic chat application:

### Session Management

You can store and retrieve data for individual sessions:

```go
// Store data in a session
s.Set("username", "user123")

// Retrieve data from a session
username, _ := s.Get("username")
```

### Selective Broadcasting

You can broadcast messages to specific clients based on filters:

```go
// Broadcast to all clients except the sender
m.BroadcastOthers(msg, s)

// Broadcast with a filter
m.BroadcastFilter(msg, func(q *melody.Session) bool {
    return q.Get("room") == s.Get("room")
})
```

### Connection Events

You can handle various connection events:

```go
// Handle new connections
m.HandleConnect(func(s *melody.Session) {
    fmt.Println("New client connected")
})

// Handle disconnections
m.HandleDisconnect(func(s *melody.Session) {
    fmt.Println("Client disconnected")
})

// Handle errors
m.HandleError(func(s *melody.Session, err error) {
    fmt.Println("Error:", err)
})
```

## Use Cases

- Real-time chat applications
- Collaborative editing tools
- Live notifications
- Real-time dashboards
- Multiplayer games
- Live streaming applications

## Additional Resources

- [Melody GitHub Repository](https://github.com/olahol/melody)
- [WebSocket API Documentation](https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API)
- [Go WebSocket Package](https://pkg.go.dev/golang.org/x/net/websocket)