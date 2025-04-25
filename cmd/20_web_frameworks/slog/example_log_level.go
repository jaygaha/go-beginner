package main

import (
	"log/slog"
	"os"
)

func LogLevel() {

	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	// check environment variable
	env := os.Getenv("ENV")
	if env == "production" {
		opts.Level = slog.LevelInfo // set log level to info
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handler)
	slog.SetDefault(logger) // set the default logger

	// log messages output will be in JSON format
	slog.Debug("Debug message")
	slog.Info("Info message")
	slog.Warn("Warn message")
	slog.Error("Error message")
}
