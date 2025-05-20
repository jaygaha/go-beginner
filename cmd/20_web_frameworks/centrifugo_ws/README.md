# Centrifugo WebSocket

This is a simple real-time chat application built with Go and Centrifugo WebSockets. It demonstrates how to create a basic chat system where multiple users can exchange messages in real-time.

## What is this project?

This project demonstrates:

- How to use WebSockets for real-time communication
- How to implement a chat application using Centrifugo
- Basic client-server architecture for real-time applications

## Project Structure

```
./
├── client/
│   └── index.html    # Frontend chat interface
├── go.mod            # Go module file
├── go.sum            # Go module dependencies
└── main.go           # Server implementation
```

## What is Centrifugo?

[Centrifugo](https://centrifugal.dev/) is a real-time messaging server that simplifies WebSocket implementation. It handles connection management, channel subscriptions, and message broadcasting, allowing developers to focus on application logic rather than WebSocket complexities.

## How to Run the Application

### Prerequisites

- Go 1.16 or higher
- Web browser

### Steps

1. Clone the repository or navigate to this directory

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Run the server:
   ```
   go run main.go
   ```

4. Open your browser and navigate to:
   ```
   http://localhost:8800
   ```

## How It Works

### Server Side (main.go)

1. **Centrifugo Node Setup**: The server initializes a Centrifugo node with logging configuration.

2. **Connection Handling**: When clients connect, the server sets up handlers for:
   - Subscriptions to channels
   - Publishing messages
   - Client disconnections

3. **Authentication**: A simple middleware adds user credentials to each connection.

4. **Static File Serving**: The server serves the HTML/JS client from the `client` directory.

### Client Side (index.html)

1. **Centrifuge Client**: The JavaScript initializes a Centrifuge client that connects to the WebSocket server.

2. **Connection Events**: The client handles various connection states (connecting, connected, disconnected).

3. **Channel Subscription**: The client subscribes to the "chat" channel to send and receive messages.

4. **Message Publishing**: When a user types a message and clicks "Send", the client publishes the message to the channel.

## Key Concepts

- **Channels**: Messages are published to named channels ("chat" in this example). Clients subscribe to channels to receive messages.

- **Publications**: When a client sends a message, it's published to a channel and delivered to all subscribers.

- **Real-time Updates**: All connected clients receive messages instantly without page refreshes.

## Extending the Application

Here are some ideas to enhance this basic chat application:

1. Add user authentication
2. Implement private messaging
3. Add message persistence (database storage)
4. Create multiple chat rooms
5. Add typing indicators
6. Implement read receipts

## Troubleshooting

- If the connection fails, ensure the server is running and accessible at localhost:8800
- Check browser console for JavaScript errors
- Verify that the WebSocket endpoint is correctly configured in the client

## Learn More

- [Centrifugo Documentation](https://centrifugal.dev/docs/getting-started/introduction)
- [WebSockets MDN Documentation](https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API)