# Basic Calculator CLI

A simple command-line calculator implemented in Go that performs basic arithmetic operations.

## Overview

This calculator is a command-line interface (CLI) application that allows users to perform basic arithmetic calculations. It accepts two numbers and an operator as command-line arguments and returns the result of the calculation.

## Features

- Supports basic arithmetic operations:
  - Addition (+)
  - Subtraction (-)
  - Multiplication (*)
  - Division (/)
- Input validation and error handling
- Protection against division by zero
- Formatted output with two decimal places

## Installation

```bash
go mod download
```

## Usage

Run the calculator using the following format:

```bash
go run main.go <number1> <operator> <number2>
```

Example usage:

```bash
go run main.go 5 + 3    # Addition
go run main.go 10 - 4   # Subtraction
go run main.go 6 * 2    # Multiplication
go run main.go 15 / 3   # Division
```

## Testing

The calculator includes a comprehensive test suite using table-driven tests. To run the tests:

```bash
go test
```

The test suite covers:
- Basic arithmetic operations
- Division by zero error handling
- Invalid operator validation

## Error Handling

The calculator handles several error cases:
- Invalid number inputs
- Missing or extra arguments
- Division by zero
- Invalid operators

## Example Output

```bash
$ go run main.go 5 + 3
5.00 + 3.00 = 8.00

$ go run main.go 10 / 2
10.00 / 2.00 = 5.00

$ go run main.go 5 / 0
Error: division by zero
```