package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
)

/*
log/slog
-> stuctured logging added from Go 1.21
-> use key-value pairs to parse. filter, search, etc.

Main features:
1. Support for structured logs
-> record not only a string, but also a set of key-value pairs
2. Flexible customization
-> set the format like JSON, text, etc.
3. Performance optimization
-> efficiently handle large logs
4. Rogert's log analysis
-> set diffrent loggers for each context

Levels:

Levels are just integers so not limited to the predefined levels

Main levels:
1. Debug
2. Info
3. Error
4. Warn

Main components:
1. Logger
-> central object that logs messages
2. Handler
-> responsible for writing log messages to the desired format and destination

Difference between slog and log:
-> log only supports string while slog supports any type of value
-> log support limited customization while slog supports more customization
-> log non corresponging standard log levels while slog supports any level
-> slog performance is better than log
-> slog handle errors in structured logs better than log
*/

func main() {
	// Baisc example
	slog.Info("hello gophers") // level=info msg="hello gophers"
	// output: 2025/04/23 11:25:28 INFO hello gohers
	slog.Debug("hello gophers") // level=debug msg="hello gophers"
	// it will not print anything
	slog.Error("hello gophers") // level=error msg="hello gophers"
	slog.Warn("hello gophers")  // level=warn msg="hello gophers"
	// Adding attributes
	slog.Info("hello gophers", "count", 10) // level=info msg="hello gophers" count=10

	// Handler
	jsonHandler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(jsonHandler)
	logger.Info("hello gophers") // {"time":"2025-04-23T14:26:27.058255+09:00","level":"INFO","msg":"hello gophers"}

	// Custom handler
	options := &slog.HandlerOptions{
		AddSource: true, // add source file and line number
		Level:     slog.LevelDebug,
	}
	textHandler := slog.NewTextHandler(os.Stdout, options)
	logger2 := slog.New(textHandler)
	logger2.Info("hello gophers")
	logger2.Debug("hello gophers", "filename", "test.txt")

	/*
		Adding attributes:
		1. slog.String()
		2. slog.Int()
		3. slog.Bool()
		4. slog.Any()
		5. slog.Group()
		-> attributes can be colleced into groups
	*/
	// Grouping attributes
	slog.Info("hello gophers", slog.Group("user",
		slog.String("name", "John"),
		slog.Int("age", 30),
	)) // level=info msg="hello gophers" user.name=John user.age=30

	// multiple groups
	slog.Info("hello gophers",
		slog.Group("user",
			slog.String("name", "John"),
			slog.Int("age", 30),
			slog.Bool("admin", true),
		),
		slog.Group("company",
			slog.String("name", "MyOwn Inc."),
			slog.String("address", "123 Secondary St."),
		),
	) // level=info msg="hello gophers" user.name=John user.age=30 company.name=Acme Inc. company.address=123 Main St.

	/*
		Error Handling
	*/
	err := errNotExist()
	if err != nil {
		slog.Error("failed to open file", "error", err, "operation", "errNotExist", "retry", false) // level=error msg="failed to open file"
	}

	/*
		Context
		-> add context to the log message
		-> add request id, user id, etc.
	*/
	// Add context to the log message
	ctx := context.Background()
	logger3 := slog.With(
		"request_id", "12345",
		"user_id", "12345",
	)

	logger3.InfoContext(ctx, "context") // level=info msg="context" request_id=12345 user_id=12345

	// Example
	fmt.Println("Example")
	LogLevel()
}

func errNotExist() error {
	return errors.New("no such file or directory")
}
