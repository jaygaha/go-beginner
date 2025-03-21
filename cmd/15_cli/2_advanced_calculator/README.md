# Advanced Calculator CLI

An interactive command-line calculator implemented in Go that supports continuous calculations and enhanced features.

## Overview

This advanced calculator builds upon the basic calculator by providing an interactive shell-like interface. Users can perform multiple calculations in a single session.

## Features

- Interactive command-line interface
- Continuous calculation mode
- Support for basic arithmetic operations:
  - Addition (+)
  - Subtraction (-)
  - Multiplication (*)
  - Division (/)
- Special commands:
  - 'exit' or 'quit' to end the session
- Enhanced error handling and input validation
- Protection against division by zero
- Formatted output with two decimal places

## Installation

```bash
go mod download
```

## Usage

Start the interactive calculator:

```bash
go run main.go
```

Once started, you'll see a prompt where you can enter calculations. Follow the instructions to perform calculations.

## Testing

Run the test suite:

```bash
go test
```

The test suite covers:
- Basic arithmetic operations
- Command parsing and execution
- Error handling scenarios

## Error Handling

The calculator handles various error cases:
- Invalid number inputs
- Invalid operators
- Division by zero
- Malformed expressions
- Invalid commands