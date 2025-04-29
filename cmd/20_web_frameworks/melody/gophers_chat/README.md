# Goofy Gophers - Real-time Cursor Tracking Chat

Goofy Gophers is a fun, interactive WebSocket application that displays each connected user as a gopher that follows their cursor movements in real-time. Built with Go and the Melody WebSocket framework, this application demonstrates the power of real-time web communications.

## Features

- **Real-time cursor tracking**: See other users' cursor positions represented as gophers
- **Unique user identification**: Each connected user gets assigned a unique ID
- **Connection notifications**: Users are notified when others connect or disconnect
- **Simple messaging**: Send messages to all other connected users

## How It Works

### Server-Side (Go + Melody)

The application uses the [Melody](https://github.com/olahol/melody) WebSocket framework to handle WebSocket connections and message broadcasting. Here's what happens on the server side:

1. **User Connection**: When a user connects, they're assigned a unique ID using an atomic counter
2. **Session Management**: User IDs are stored in the WebSocket session
3. **Message Broadcasting**: Messages and cursor positions are broadcast to all other connected users
4. **Disconnect Handling**: When a user disconnects, other users are notified

### Client-Side (HTML + JavaScript)

The client-side code handles:

1. **WebSocket Connection**: Establishes a connection to the server
2. **Command Processing**: Processes commands received from the server (iam, set, dis)
3. **Gopher Rendering**: Creates and positions gopher elements based on cursor positions
4. **Mouse Tracking**: Tracks mouse movements and sends coordinates to the server

## Protocol

The application uses a simple text-based protocol for communication:

- **Server to Client**:
  - `iam {id}` - Assigns an ID to the client
  - `user {id} disconnected` - Notifies when a user disconnects
  - `user {id}: {message}` - Relays a message from another user

- **Client to Server**:
  - `{id} {x} {y}` - Sends cursor position coordinates
  - Text messages - Sent as-is to be broadcast to other users

## Installation

### Prerequisites

- Go 1.18 or higher
- [Melody](https://github.com/olahol/melody) WebSocket framework

### Setup

1. Clone the repository or copy the files to your local machine
2. Install the Melody package:
   ```
   go get github.com/olahol/melody
   ```
3. Navigate to the project directory

## Usage

1. Start the server:
   ```
   go run main.go
   ```
2. Open your browser and navigate to `http://localhost:8800`
3. Open multiple browser windows to see multiple gophers interacting
4. Move your cursor to see your gopher move
5. Type messages to communicate with other connected users

## Code Structure

- **main.go**: Server-side code handling WebSocket connections and message broadcasting
- **index.tmpl**: Client-side HTML, CSS, and JavaScript for rendering gophers and handling user interactions

## How to Extend

Here are some ideas for extending the application:

1. **Custom Gopher Images**: Allow users to select different gopher images
2. **User Names**: Add support for user names instead of just IDs
3. **Rooms**: Implement separate rooms for different groups of users
4. **Persistence**: Add a database to store messages for later retrieval
5. **Additional Interactions**: Add more ways for gophers to interact (e.g., animations, emotes)

## Acknowledgments

- [Melody](https://github.com/olahol/melody) - WebSocket framework for Go
- [Go Gopher](https://blog.golang.org/gopher) - The Go mascot used in this project