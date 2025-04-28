package main

import (
	"errors"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

/*
Zerolog:
-> provides a fast and efficient logging library dedicated to JSON output for Go
-> designed for high-performance applications
-> uses a simple and intuitive API
-> supports structured logging and context-aware logging
-> Prety logging output
*/

func main() {
	// Basic Usage
	// standard logging
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()  // logger is a zerolog.Logger instance
	logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr}) // pretty logging

	logger.Info().Msg("Hello Gophers") // logs a message with the "info" level

	// structured logging
	logger.Info().
		Str("name", "Jay"). // adds a string field
		Int("age", 30).     // adds an integer field
		Msg("User information")

	// error logging
	log.Error().
		Err(nil).                // adds an error field
		Int("status_code", 500). // adds an integer field
		Msg("Error occurred")

	// JSON Logging
	log.Info().
		Str("operation", "login").
		Str("username", "user123").
		Msg("User logged in")

	/*
		Log Levels:
		-> Debug: used for detailed debugging information
		-> Info: used for general information about the application's state
		-> Warn: used for non-critical issues that may require attention
		-> Trace: used for tracing the execution flow
		-> Error: used for errors that occur during the application's execution
		-> Fatal: used for critical errors that cause the application to terminate
		-> Panic: used for unexpected errors that cause the application to panic
	*/
	log.Trace().Msg("This is a trace message")
	log.Debug().Msg("This is a debug message")
	log.Info().Msg("This is an info message")
	log.Warn().Msg("This is a warning message")
	log.Error().Msg("This is an error message")
	// log.Fatal().Msg("This is a fatal message") // exits the application
	// log.Panic().Msg("This is a panic message") // exits the application

	// Log level filtering
	//  -> Zerolog provides a level-based logging mechanism that allows you to filter logs based on their severity level
	// zerolog.SetGlobalLevel(zerolog.InfoLevel) // sets the global log level to info

	// warn level default
	zerolog.SetGlobalLevel(zerolog.WarnLevel) // sets the global log level to warn

	log.Info().Msg("This will not be logged because the log level is set to warn")

	// these will be logged because warn level and above
	log.Warn().Msg("This will be logged because the log level is set to warn")
	log.Error().Msg("This will be logged because the log level is set to warn")

	// resetting the global log level
	zerolog.SetGlobalLevel(zerolog.InfoLevel) // sets the global log level to info

	// Contextual Logging
	// -> allows to key-value pairs to be added to the log message
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix // sets the time format to Unix time

	log.Info().
		Str("username", "user123").
		Float32("balance", 100.50).
		Msg("User balance updated")

	logger = log.With().
		Str("service", "mail-service").
		Str("version", "1.0.0").
		Logger()

	logger.Info().
		Msg("Sending email")

	// Error Logging:
	err := errors.New("something went wrong")
	log.Err(err).Msg("Error occurred")

	// Save logs to a file
	file, err := os.Create("zerologs.txt")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open log file")
	}

	defer file.Close()

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: file})
	log.Info().Msg("This log will be saved to the file")
}
