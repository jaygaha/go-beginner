# Error Handling in Go

## Overview
Error handling is a critical aspect of writing robust Go programs. Unlike many other languages that use exceptions, Go handles errors as values that can be returned from functions, checked, and processed. This approach encourages explicit error checking and handling.

## Key Concepts

### The Error Interface
In Go, errors are values that implement the built-in `error` interface:

```go
type error interface {
    Error() string
}
```

Any type that implements the `Error()` method, which returns a string, satisfies this interface.

## Error Handling Approaches

Go provides several approaches to handle errors, each with its own use cases:

### 1. Constructing Basic Errors

The simplest way to create errors is using the `fmt.Errorf()` function or `errors.New()`:

```go
// Using fmt.Errorf
return 0, fmt.Errorf("cannot divide '%d' by zero", numerator)

// Using errors.New
return 0, errors.New("cannot divide by zero")
```

This approach is suitable for simple error messages that don't require additional context.

### 2. Expected (Sentinel) Errors

Sentinel errors are predefined error variables that can be checked using `errors.Is()`:

```go
// Define a sentinel error
var errDivisionByZero = errors.New("cannot divide by zero")

// Return the predefined error
return 0, errDivisionByZero

// Check for specific error
if errors.Is(err, errDivisionByZero) {
    // Handle specific error
}
```

This approach is useful when you need to check for specific error conditions throughout your code.

### 3. Custom Error Types

For more complex error scenarios, you can create custom error types by implementing the `error` interface:

```go
// Define custom error type
type DivisionError struct {
    Numerator   int
    Denominator int
    Message     string
}

// Implement the Error() method
func (e *DivisionError) Error() string {
    return fmt.Sprintf("%d / %d = %s", e.Numerator, e.Denominator, e.Message)
}

// Create and return custom error
return 0, &DivisionError{
    Numerator:   numerator,
    Denominator: denominator,
    Message:     "cannot divide by zero",
}

// Check and extract custom error
var divErr *DivisionError
if errors.As(err, &divErr) {
    // Access error fields: divErr.Numerator, divErr.Denominator, etc.
}
```

Custom error types allow you to include additional context and data with your errors.

## Best Practices

1. **Always check errors**: Never ignore returned errors in Go.
   
```go
result, err := someFunction()
if err != nil {
    // Handle error
}
```

2. **Use `errors.Is()` for sentinel errors**: When checking for specific predefined errors.

3. **Use `errors.As()` for custom error types**: When you need to extract information from custom error types.

4. **Return early on errors**: The common pattern is to check for errors and return immediately.

5. **Provide context**: Make error messages descriptive and include relevant information.

## Running the Example

1. Navigate to this directory
2. Read the comments in `main.go` for explanations
3. Run the program:

```bash
go run main.go
```

## Further Reading

- [Go Error Handling Best Practices](https://go.dev/blog/error-handling-and-go)
- [Working with Errors in Go 1.13](https://go.dev/blog/go1.13-errors)