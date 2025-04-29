# Chat with Username using Melody

This application demonstrates a real-time chat with username functionality using WebSockets in Go with the Melody framework. It allows multiple clients to connect, set their usernames, and exchange messages in real-time with user identification.

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
- **JSON**: For structured message format

### Architecture

1. **WebSocket Server**: Melody manages WebSocket connections and broadcasts
2. **Client Interface**: HTML/CSS interface with JavaScript for WebSocket communication
3. **Session Management**: Stores usernames for each connection
4. **Message Structure**: JSON format with username, message content, and timestamp

### Key Components

#### User Session Management

The application uses Melody's session management to store usernames for each connection:

```go
// When a client connects, extract username from URL query parameters
m.HandleConnect(func(s *melody.Session) {
    q := s.Request.URL.Query()
    user := q.Get("user")
    if user == "" {
        // Generate default username if none provided
        time := strconv.FormatInt(time.Now().Unix(), 10)
        user = "user_" + time
    }
    
    // Store username in the session
    s.Set("user", user)
})
```

#### Message Structure

Messages are structured as JSON objects containing the username, message content, and timestamp:

```go
// Message struct with user identification
type Message struct {
    User      string    `json:"user"`
    Msg       string    `json:"msg"`
    CreatedAt time.Time `json:"created_at"`
}
```

#### Message Handling

When a client sends a message, the server retrieves the username from the session, creates a structured message, and broadcasts it to all clients:

```go
// Handle incoming messages
m.HandleMessage(func(s *melody.Session, msg []byte) {
    // Get username from session
    user, exists := s.Get("user")
    if !exists {
        // Fallback if username not found
        time := strconv.FormatInt(time.Now().Unix(), 10)
        user = "user_" + time
    }
    
    // Create structured message with username
    message := Message{
        User:      user.(string),
        Msg:       string(msg),
        CreatedAt: time.Now(),
    }
    
    // Convert to JSON and broadcast
    jsonMsg, err := json.Marshal(message)
    if err != nil {
        return
    }
    m.Broadcast(jsonMsg)
})
```

#### Client-Side WebSocket

The client prompts for a username and establishes a WebSocket connection with the username as a URL parameter:

```javascript
// Prompt for username
const user = prompt("Please enter your username:");
if (!user) {
    alert("Username cannot be empty!");
    window.location.reload();
}

// Connect with username as URL parameter
const ws = new WebSocket(`ws://localhost:8800/ws?user=${encodeURIComponent(user)}`);
```

It handles incoming messages by parsing the JSON and displaying the username with the message:

```javascript
// Display messages with username
ws.onmessage = (event) => {
    const data = JSON.parse(event.data);
    const row = document.createElement("div");
    row.innerHTML = `${data.user}: ${data.msg} (${data.created_at})`;
    document.getElementById("chat-window").appendChild(row);
};
```

## Installation

### Steps

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd golang-beginner/cmd/20_web_frameworks/melody/session_chat
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
2. Enter a username when prompted
3. Type a message in the input field and click "Send"
4. The message will appear in all connected browser windows with your username

## Key Features

### Username Identification

- Users can set their own username when joining the chat
- If no username is provided, a default one is generated
- All messages are displayed with the sender's username

### Session Management

- Each WebSocket connection maintains its own session
- Username is stored in the session for message attribution
- Sessions persist for the duration of the connection

### Timestamped Messages

- Each message includes a timestamp showing when it was sent
- Timestamps are generated on the server for consistency

## Extending the Application

Here are some ways to extend this chat application:

### User Authentication

Implement proper user authentication to verify user identities:

```go
// Add authentication middleware
func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Verify user credentials
        // ...
        next.ServeHTTP(w, r)
    })
}
```

### Private Messaging

Implement private messaging between users:

```go
// Send message to specific user
func sendPrivateMessage(sender, recipient string, message []byte) {
    m.BroadcastFilter(message, func(s *melody.Session) bool {
        user, _ := s.Get("user")
        return user.(string) == recipient
    })
}
```

### Chat Rooms

Implement multiple chat rooms:

```go
// Join a specific room
m.HandleConnect(func(s *melody.Session) {
    room := s.Request.URL.Query().Get("room")
    s.Set("room", room)
})

// Send message to users in the same room
m.HandleMessage(func(s *melody.Session, msg []byte) {
    room, _ := s.Get("room")
    m.BroadcastFilter(msg, func(q *melody.Session) bool {
        qRoom, _ := q.Get("room")
        return qRoom == room
    })
})
```

## Additional Resources

- [Melody GitHub Repository](https://github.com/olahol/melody)
- [WebSocket API Documentation](https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API)
- [Go WebSocket Package](https://pkg.go.dev/golang.org/x/net/websocket)
- [JSON in Go Documentation](https://pkg.go.dev/encoding/json)