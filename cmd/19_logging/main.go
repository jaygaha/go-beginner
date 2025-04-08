package main

import (
	"log"
	"os"

	"go.uber.org/zap"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	logFile       *os.File
)

func init() {
	// Logging to a file
	// Using custom logger
	var err error
	logFile, err = os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}

	// output to file
	// log.SetOutput(f)
	InfoLogger = log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(logFile, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

/*
	Logging:
		- logging is the process of recording and storing information about events, actions, or errors that occur in a system or application
		- it provides insights into the behavior, performance, and usage of the system, allowing for troubleshooting, monitoring, and analysis
		- logging can be categorized into different levels, such as debug, info, warning, error, and fatal, each with a specific purpose and importance

		What not to log:
			- sensitive information, such as passwords, credit card numbers, or social security numbers
			- information that is not relevant to the current context or situation, such as temporary files or system configurations

	log package:
		- built-in package for logging in Go
*/

func main() {
	// standard logger
	log.Println("Hello World")
	// output: 2025/04/07 13:22:42 Hello World

	// fatal logger; it will exit the program
	// log.Fatalln("Error: Fatal")
	// 2025/04/07 13:26:17 Error: Fatal
	// exit status 1

	// panic logger; it will exit the program
	// log.Panicln("Error: Panic")
	// output: 2025/04/07 13:27:09 Error: Panic
	// panic: Error: Panic
	// other info about the error

	// fatal log with format
	// log.Fatalf("Error: %s", "Fatalf")
	// output: 2025/04/07 13:28:16 Error: Fatalf
	// exit status 1

	// panic log with format
	// log.Panicf("Error: %s", "Panicf") // output: 2025/04/07 13:22:42 Error: Error

	// Custom Logger
	// this ensure that the log file remain open for loggin throughout the program execution
	defer logFile.Close()

	InfoLogger.Println("This is an info message")
	WarningLogger.Println("This is a warning message")
	ErrorLogger.Println("This is an error message")

	/*
		Third Party Loggers:
		 - Go doesn't support leveled logging
		 - zap package

			Common Log Levels:
			 - Trace: lowest level of logging, used for detailed information that is only relevant during development or debugging
			 - Debug: used for detailed information that is only relevant during development or debugging
			 - Info: used for general information about the application, such as startup messages or configuration changes
			 - Warn: used for non-critical errors that may indicate a problem, but the application can continue to run
			 - Error: used for critical errors that may cause the application to fail or stop functioning
			 - Fatal: highest level of logging, used for critical errors that cause the application to stop running immediately
			 	- exit after logging the error
			 - Panic: used for critical errors that cause the application to stop running immediately
			 	- exit after logging the error
			 	- defer a function that recovers from the panic and logs the error
			 	- the function that recovers from the panic should be called using defer

			Log Flags:
			 - log.Ldate: log the date in the local time zone: 2009/01/23
			 - log.Ltime: log the time in the local time zone: 01:23:23
			 - log.Lmicroseconds: log the time in microseconds: 01:23:23.123123 (NOTE: Ltime is required)
	*/

	// create a configuration (builder)
	config := zap.NewProductionConfig()

	// configure the destination path
	config.OutputPaths = []string{"zap_log.log"}

	// Build the logger
	logger, err := config.Build()

	// Go doesn't support exceptions, so we check for the error, existing the app if needed
	if err != nil {
		log.Fatalf("Error building logger: %s", err)
	}

	// flushes buffer, if any
	defer logger.Sync()

	// finally log a message at the INFO level
	logger.Info("This is a zap log message")
	// zap support structured logging the result is a JSON object
	logger.Debug("Debug message", zap.String("key", "value"))

	// TODO: explore the other log levels
}
