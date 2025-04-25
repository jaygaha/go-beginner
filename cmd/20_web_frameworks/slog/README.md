# Go `slog` Package

The `slog` package, introduced in Go 1.21, provides **structured logging** for Go applications. Unlike the traditional `log` package, `slog` allows you to log not just plain strings, but also key-value pairs, making your logs easier to filter, search, and analyze.

---

## Key Features

- **Structured Logging:** Log messages with additional context using key-value pairs.
- **Flexible Output:** Choose between formats like JSON or plain text.
- **Performance:** Optimized for handling large volumes of logs efficiently.
- **Custom Levels:** Supports standard and custom log levels (e.g., Debug, Info, Warn, Error).
- **Contextual Logging:** Attach context (like request IDs) to logs for better traceability.

---

## Main Components

- **Logger:** The main object used to log messages.
- **Handler:** Determines how and where log messages are written (e.g., to a file, console, or in JSON format).

---

## Why Use slog Over log ?

- **Structured data**: Easier to parse and analyze.
- **Customizable**: Choose your format and destination.
- **Better performance**: Handles large logs efficiently.
- **More control**: Supports custom log levels and error handling.

---

## Basic Usage

```go
import (
    "log/slog"
)

func main() {
    slog.Info("Application started")
    slog.Error("An error occurred", "error", err)
}
```

## Adding Attributes

You can add extra information to your logs using key-value pairs:

```go
slog.Info("User login", "user_id", 123, "role", "admin")
```

## Custom Handlers

You can customize how logs are formatted and where they are sent:

```go
import (
    "os"
    "log/slog"
)

func main() {
    handler := slog.NewJSONHandler(os.Stdout, nil)
    logger := slog.New(handler)
    logger.Info("Logging in JSON format")
}
```

## Grouping Attributes

Group related attributes for better organization:

```go
slog.Info("User details",
    slog.Group("user",
        slog.String("name", "Alice"),
        slog.Int("age", 30),
    ),
)
```

## Error Handling

Log errors with additional context:

```go
if err != nil {
    slog.Error("Failed to open file", "error", err, "retry", false)
}
```

## Contextual Logging

Attach context (like request or user IDs) to your logs:

```go
ctx := context.Background()
logger := slog.With("request_id", "abc123")
logger.InfoContext(ctx, "Processing request")
```

## Resources
- [Official slog Documentation](https://pkg.go.dev/log/slog)
- [Go Blog: Structured Logging with slog](https://go.dev/blog/slog)