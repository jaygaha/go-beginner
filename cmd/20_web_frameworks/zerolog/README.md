# Zerolog - Fast and Simple Logging for Go

## What is Zerolog?

[Zerolog](https://github.com/rs/zerolog) is a fast and efficient logging library for Go, designed for high-performance applications. It focuses on structured logging and outputs logs in JSON format by default, making it easy to parse and analyze logs.

---

## Key Features

- **High Performance:** Minimal allocations and fast execution.
- **Structured Logging:** Add key-value pairs to your logs for better context.
- **JSON Output:** Logs are output in JSON by default, but pretty (human-readable) output is also supported.
- **Multiple Log Levels:** Supports levels like Trace, Debug, Info, Warn, Error, Fatal, and Panic.
- **Contextual Logging:** Easily add context (like user ID, request ID) to your logs.
- **Flexible Output:** Write logs to files, console, or other destinations.

---

## Logging Levels

- Trace: Detailed tracing information
- Debug: Debugging information
- Info: General application events
- Warn: Non-critical issues
- Error: Errors during execution
- Fatal: Critical errors (application will exit)
- Panic: Unexpected errors (application will panic)

You can set the global log level to control what gets logged:
```go
zerolog.SetGlobalLevel(zerolog.InfoLevel)
```
---

## Installation

To install Zerolog, use the following command:

```bash
go get github.com/rs/zerolog
```

## Basic Usage

```go
import (
    "os"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
)

func main() {
    // Create a logger with timestamp and pretty output
    logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
    logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})

    // Log a simple info message
    logger.Info().Msg("Hello Gophers")

    // Structured logging with fields
    logger.Info().
        Str("name", "Jay").
        Int("age", 30).
        Msg("User information")
}
```

## Contextual Logging

Add context to your logs for better traceability:

```go
logger := log.With().
Str("service", "mail-service").
Str("version", "1.0.0").
Logger()

logger.Info().Msg("Sending email")
```

## Logging to a File

```go
file, err := os.Create("zerologs.txt")
if err != nil {
    log.Fatal().Err(err).Msg("Failed to open log file")
}
defer file.Close()

log.Logger = log.Output(zerolog.ConsoleWriter{Out: file})
log.Info().Msg("This log will be saved to the file")
```

## Resources
- [Zerolog GitHub Repository](https://github.com/rs/zerolog)
- [Zerolog Documentation](https://pkg.go.dev/github.com/rs/zerolog)