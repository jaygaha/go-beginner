# Heimdall HTTP Client

## What is Heimdall?

Heimdall is a powerful HTTP client library for Go applications designed to make a large number of requests at scale with built-in resilience features. It was developed by Gojek and is available at [github.com/gojek/heimdall](https://github.com/gojek/heimdall).

With Heimdall, you can:

- Use a hystrix-like circuit breaker to control failing requests
- Add synchronous in-memory retries to each request, with customizable retry strategies
- Create clients with different timeouts for every request
- Make HTTP requests with a fluent interface for all HTTP methods

## Installation

```bash
go get -u github.com/gojek/heimdall/v7
```

## Examples in this Directory

This directory contains several examples demonstrating different features of the Heimdall HTTP client:

### 1. Basic GET Request

Location: [main.go](./main.go)

Demonstrates how to:
- Create a basic Heimdall HTTP client
- Configure timeout and retry settings
- Make a simple GET request
- Handle the response

```go
// Create a new client with heimdall
client := httpclient.NewClient(
    httpclient.WithHTTPTimeout(10*time.Second), // set the timeout to 10 seconds
    httpclient.WithRetryCount(3),               // set the number of retries to 3
)

// Make a GET request
resp, err := client.Get("https://dummyjson.com/quotes", nil)
```

### 2. POST Request Example

Location: [post/](./post/)

Demonstrates how to:
- Prepare JSON data for a POST request
- Set appropriate headers
- Make a POST request
- Handle the response

```go
// JSON data
jsonData := []byte(`{"body": "Comments by commenter", "postId": 3, "userId": 1}`)
headers := http.Header{
    "Content-Type": []string{"application/json"},
}

// Make a POST request
resp, err := client.Post("https://dummyjson.com/comments/add", bytes.NewBuffer(jsonData), headers)
```

### 3. Retry with Constant Backoff

Location: [retry-backoff/](./retry-backoff/)

Demonstrates how to:
- Create a constant backoff strategy with a fixed retry interval
- Configure an HTTP client with the backoff strategy
- Make HTTP requests with automatic retry capability

```go
// Create a constant backoff with 2-second retry interval
backoff := heimdall.NewConstantBackoff(2*time.Second, 1*time.Second)

// Configure HTTP client with the backoff strategy
client := httpclient.NewClient(
    httpclient.WithHTTPTimeout(10*time.Millisecond),      // Set request timeout
    httpclient.WithRetrier(heimdall.NewRetrier(backoff)), // Apply the backoff strategy
)
```

### 4. Retry with Exponential Backoff

Location: [retry-index-backoff/](./retry-index-backoff/)

Demonstrates how to:
- Implement exponential backoff retry strategy
- Configure parameters for exponential growth
- Set maximum retry intervals and total elapsed time

```go
// Set exponential backoff
backoff := heimdall.NewExponentialBackoff(
    500*time.Millisecond,    // Initial standby time (first interval of retry)
    5*time.Second,           // Maximum waiting time (maximum interval of retries)
    float64(15*time.Second), // Maximum elapsed time (maximum time for all retries)
    2.0,                     // Exponential increase rate (doubling the interval for each retry)
)

// Create a new client with heimdall
client := httpclient.NewClient(
    httpclient.WithHTTPTimeout(10*time.Second),           // set the timeout to 10 seconds
    httpclient.WithRetrier(heimdall.NewRetrier(backoff)), // set the retrier to backoff
)
```

## Comparison of Retry Strategies

| Strategy | Description | Best For |
|----------|-------------|----------|
| **No Retry** | Single attempt only | Simple, non-critical requests |
| **Constant Backoff** | Fixed time between retries | Predictable retry patterns, simple implementation |
| **Exponential Backoff** | Increasing time between retries | Reducing server load during outages, handling rate limits |

## When to Use Each Strategy

### Constant Backoff
- When you need a simple, predictable retry pattern
- For services with known recovery times
- When you want to limit the total number of retries in a fixed time window

### Exponential Backoff
- For external APIs that might experience temporary outages
- When dealing with rate-limited services
- For network requests over unreliable connections
- In distributed systems where components might be temporarily unavailable

## Benefits of Using Heimdall

- **Improved Reliability**: Automatically handles transient network failures
- **Reduced Error Handling**: Less manual error handling in your code
- **Controlled Retry Flow**: Prevents overwhelming servers with immediate retries
- **Configurable Behavior**: Easily adjust retry counts, intervals, and strategies
- **Performance**: Designed for applications that need to make many HTTP requests
- **Simplicity**: Clean API makes it easy to use for common HTTP operations

## Additional Resources

- [Heimdall GitHub Repository](https://github.com/gojek/heimdall)
- [Official Documentation](https://pkg.go.dev/github.com/gojek/heimdall/v7)
- [Gojek Engineering Blog](https://www.gojek.io/blog/)