# Go Concurrency

This section covers fundamental concurrency patterns in Go, demonstrating how to write concurrent programs effectively using goroutines, channels, and synchronization primitives.

## Key Concepts

### 1. Goroutines

Goroutines are lightweight threads managed by the Go runtime. They allow functions to run concurrently.

```go
go func() {
    // This function runs concurrently
}
```

### 2. Channels

Channels are typed conduits for communication between goroutines. They help implement the CSP (Communicating Sequential Processes) pattern.

#### Unbuffered Channels

- Block until both sender and receiver are ready
- Provide synchronization

```go
ch := make(chan string)
go func() {
    ch <- "message" // Blocks until receiver is ready
}()
msg := <-ch // Blocks until sender sends
```

#### Buffered Channels

- Have a capacity for storing messages
- Only block when buffer is full

```go
orders := make(chan string, 3) // Buffer size of 3
orders <- "order1" // Won't block unless buffer is full
```

### 3. Select Statement

Select allows a goroutine to wait on multiple channel operations, choosing one that's ready to proceed.

```go
select {
case msg1 := <-ch1:
    // Handle message from ch1
case msg2 := <-ch2:
    // Handle message from ch2
default:
    // Optional default case
}
```

### 4. Synchronization

#### WaitGroup

- Used to wait for multiple goroutines to complete
- Common for fan-out patterns

#### Mutex

- Provides mutual exclusion
- Prevents race conditions in shared memory access

## Best Practices

1. Don't communicate by sharing memory; share memory by communicating
2. Use buffered channels when you know the number of messages in advance
3. Always use proper synchronization when accessing shared resources
4. Close channels from the sender side
5. Check for channel closure when receiving

## Examples

Check the source files in this directory for detailed examples:

- `12_1_goroutines.go` - Basic goroutine usage
- `12_1_1_goroutines_waitgroup.go` - Synchronization with WaitGroup
- `12_1_1_goroutines_mutex.go` - Protecting shared resources with Mutex
- `12_2_channels.go` - Channel basics
- `12_2_1_unbuffered_channels.go` - Working with unbuffered channels
- `12_2_2_buffered_channels.go` - Working with buffered channels
- `12_2_3_select.go` - Using select for multiple channel operations


## How to Run

Execute the following command from this directory:

```bash
go run .
```
