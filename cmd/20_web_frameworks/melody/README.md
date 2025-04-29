# Melody WebSocket Framework Examples

This directory contains a collection of examples demonstrating the use of [Melody](https://github.com/olahol/melody), a minimalist WebSocket framework for Go. Melody makes it easy to build real-time applications with features like connection management, broadcasting, and session handling.

## What is Melody?

Melody is a lightweight WebSocket framework for Go that simplifies the development of real-time applications. It provides a clean API for managing WebSocket connections and broadcasting messages.

### Key Features

- Simple API for handling WebSocket connections
- Broadcasting messages to all connected clients
- Session management for individual connections
- Support for binary and text messages
- Filters for selective broadcasting
- Middleware support

## Examples Overview

This repository contains the following examples, each demonstrating different aspects of the Melody framework:

### 1. [Basic Chat](./basic_chat/)

A simple real-time chat application where multiple clients can connect and exchange messages.

**Features:**
- Real-time message broadcasting to all connected clients
- Simple HTML/CSS interface
- Demonstrates basic WebSocket connection handling

**Run it:**
```bash
cd basic_chat
go run main.go
# Open http://localhost:8800 in your browser
```

### 2. [Session Chat](./session_chat/)

An enhanced chat application that includes username functionality for message identification.

**Features:**
- Username identification for messages
- Session management to store user information
- JSON-formatted messages with timestamps
- Improved UI with user identification

**Run it:**
```bash
cd session_chat
go run main.go
# Open http://localhost:8800 in your browser
```

### 3. [File Watch Chat](./filewatch_chat/)

A real-time file monitoring application that watches a text file for changes and broadcasts updates to all connected clients.

**Features:**
- Real-time file content synchronization
- Uses fsnotify for file system monitoring
- Demonstrates integration of file system events with WebSockets

**Run it:**
```bash
cd filewatch_chat
go run main.go
# Open http://localhost:8800 in your browser
# Edit file.txt to see changes in real-time
```

### 4. [Gophers Chat](./gophers_chat/)

A fun, interactive application that displays each connected user as a gopher that follows their cursor movements in real-time.

**Features:**
- Real-time cursor tracking with gopher avatars
- Unique user identification
- Connection/disconnection notifications
- Simple messaging between users

**Run it:**
```bash
cd gophers_chat
go run main.go
# Open http://localhost:8800 in your browser
# Open multiple windows to see multiple gophers interacting
```

## Getting Started

### Prerequisites

- Go 1.18 or higher
- [Melody](https://github.com/olahol/melody) WebSocket framework

### Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd golang-beginner/cmd/20_web_frameworks/melody
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Run any of the examples by navigating to their directory and executing:
   ```bash
   go run main.go
   ```

4. Open your browser and navigate to `http://localhost:8800`

## Learning Path

The examples are arranged in increasing order of complexity:

1. **Basic Chat**: Start here to understand the fundamentals of WebSocket communication with Melody
2. **Session Chat**: Learn how to manage user sessions and structured messages
3. **File Watch Chat**: Explore integration with file system events
4. **Gophers Chat**: See a more complex application with real-time cursor tracking and visual elements

## Additional Resources

- [Melody GitHub Repository](https://github.com/olahol/melody)
- [WebSocket RFC](https://tools.ietf.org/html/rfc6455)
- [Go WebSocket Documentation](https://pkg.go.dev/golang.org/x/net/websocket)