# File Watching Chat with Melody

This application demonstrates real-time file monitoring using WebSockets in Go. It watches a text file for changes and broadcasts those changes to all connected clients instantly. This creates a real-time collaborative experience where any change to the monitored file is immediately visible to all connected users.

## How It Works

### Technologies Used

- **[Melody](https://github.com/olahol/melody)**: A minimalist WebSocket framework for Go
- **[fsnotify](https://github.com/fsnotify/fsnotify)**: A cross-platform file system notifications library for Go
- **Go standard library**: For HTTP server and file operations

### Architecture

1. **File Watcher**: Uses fsnotify to monitor changes to `file.txt`
2. **WebSocket Server**: Melody manages WebSocket connections and broadcasts
3. **Client Interface**: Simple HTML/CSS interface to display file contents in real-time

### Key Components

#### File Monitoring

The application uses fsnotify to watch for changes to the specified file:

```go
w, _ := fsnotify.NewWatcher()
// ...
w.Add(file) // Add file to watcher
```

When a file write event is detected, the application reads the updated content and broadcasts it to all connected clients:

```go
go func() {
    for {
        select {
        case event := <-w.Events:
            if event.Op&fsnotify.Write == fsnotify.Write {
                content, _ := os.ReadFile(file)
                m.Broadcast(content)
            }
        }
    }
}()
```

#### WebSocket Communication

Melody handles WebSocket connections and message broadcasting:

```go
m := melody.New() // Create new Melody instance

// Handle new connections
m.HandleConnect(func(s *melody.Session) {
    content, _ := os.ReadFile(file)
    s.Write(content) // Send current file content to new client
})
```

## Installation

### Steps

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd golang-beginner/cmd/20_web_frameworks/melody/filewatch_chat
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

1. Connect to the application through your web browser
2. The current content of `file.txt` will be displayed
3. Edit the `file.txt` file using any text editor
4. Watch as changes appear instantly in all connected browsers

## Use Cases

- Real-time collaborative text editing
- Live configuration monitoring
- Log file viewers
- Status monitoring dashboards
- Simple chat applications

## How to Extend

- Add authentication for secure file access
- Implement multiple file monitoring
- Add user identification to track who made changes
- Implement a text editor in the web interface
- Add versioning or history of file changes