# Logging in Go

This project demonstrates how to implement logging in a Go application using both the standard library's `log` package and the third-party `zap` package. It includes examples of logging to a file and using structured logging with `zap`.

## Getting Started

### Prerequisites

- Go 1.24.0 or later
- `zap` package

### Logging Details

- Standard Library Logging : Uses custom loggers to write messages to `log.txt` with different log levels (INFO, WARNING, ERROR).
- Zap Logging : Demonstrates structured logging with `zap` , outputting JSON-formatted logs to `zap_log.log`.

### Logging Levels

- Info : General information about the application's operation.
- Warning : Non-critical issues that may require attention.
- Error : Critical issues that may affect the application's functionality.