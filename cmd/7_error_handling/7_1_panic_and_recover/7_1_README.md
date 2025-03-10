# Panic and Recover in Go

## Overview
Go provides a mechanism for handling exceptional situations through `panic` and `recover`. Unlike traditional exception handling in other languages, Go's approach is more controlled and explicit.

## Key Concepts

### Defer
A defer statement defers the execution of a function until the surrounding function returns:

```go
defer fmt.Println("end") // This will execute last
fmt.Println("start")     // This will execute first
```

Deferred functions are executed in LIFO (Last-In-First-Out) order. They run even if the function panics.

### Panic
A panic is a runtime error that stops the normal execution flow:

```go
panic("something went wrong") // Stops execution and begins unwinding the stack
```

When a function panics:
1. Its execution stops immediately
2. Any deferred functions are executed
3. The panic propagates up the call stack

### Recover
The recover function allows a program to regain control after a panic:

```go
defer func() {
    if r := recover(); r != nil {
        fmt.Printf("Recovered from panic: %v\n", r)
    }
}()
```

Important notes about recover:
- It only works when called directly from a deferred function
- It returns the value passed to panic
- If no panic is occurring, recover returns nil

## Example Explained

The example in `main.go` demonstrates these concepts:

```go
func main() {
    // Defer demonstration
    defer fmt.Println("end") // Will execute at the end
    fmt.Println("start")
    
    // Recover demonstration
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Recovered from panic: %v\n", r)
        }
    }()
    
    // This function will panic
    raisePanic()
}

func raisePanic() {
    panic("something went wrong")
}
```

Execution flow:
1. "start" is printed
2. The `raisePanic()` function is called, which triggers a panic
3. The deferred anonymous function catches the panic with recover
4. "Recovered from panic: something went wrong" is printed
5. Finally, "end" is printed from the first defer statement

## Advanced Techniques

### Repanicking
Sometimes you might want to recover from a panic, perform some cleanup, and then panic again:

```go
defer func() {
    if r := recover(); r != nil {
        // Log the panic
        fmt.Printf("Logging panic: %v\n", r)
        // Cleanup resources
        
        // Repanic with the same or different error
        panic(r) // or panic("new error")
    }
}()
```

This is useful for logging, cleanup, or transforming the panic message.

### When to Use Panic and Recover

In Go, panic and recover should be used sparingly:

- **Use for exceptional conditions** that should never happen in normal operation
- **Not a substitute for error handling** - most errors should be returned as values
- **Appropriate for initialization failures** where the program cannot continue

## Running the Example

1. Navigate to this directory
2. Run the program:

```bash
go run main.go
```

Expected output:
```
start
Recovered from panic: something went wrong
end
```

## Further Reading

- [Defer, Panic, and Recover](https://go.dev/blog/defer-panic-and-recover)
- [Effective Go: Panic](https://go.dev/doc/effective_go#panic)