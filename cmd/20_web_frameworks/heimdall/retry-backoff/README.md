# Heimdall HTTP Client - Retry with Backoff Example

## Overview

This example demonstrates how to implement retry mechanisms with constant backoff using the Heimdall HTTP client library. Heimdall is a powerful HTTP client for Go applications that need to make a large number of requests at scale with built-in resilience features.

## What is Retry with Backoff?

Retry with backoff is a technique used to retry failed HTTP requests with a controlled delay between attempts:

- **Retry**: Automatically attempts the request again after a failure
- **Backoff**: Implements a waiting period between retry attempts
- **Constant Backoff**: Uses a fixed time interval between retries (as shown in this example)

## Example Code Explanation

The example in `main.go` demonstrates how to:

1. Create a constant backoff strategy with a fixed retry interval
2. Configure an HTTP client with the backoff strategy
3. Make HTTP requests with automatic retry capability

### Key Components

```go
// Create a constant backoff with 2-second retry interval
backoff := heimdall.NewConstantBackoff(2*time.Second, 1*time.Second)

// Configure HTTP client with the backoff strategy
client := httpclient.NewClient(
    httpclient.WithHTTPTimeout(10*time.Millisecond),      // Set request timeout
    httpclient.WithRetrier(heimdall.NewRetrier(backoff)), // Apply the backoff strategy
)
```

### Parameters Explained

- `NewConstantBackoff(2*time.Second, 1*time.Second)`
  - First parameter: The constant retry interval (2 seconds)
  - Second parameter: Initial backoff duration (1 second)

- `WithHTTPTimeout(10*time.Millisecond)`
  - Sets a very short timeout to demonstrate retry behavior
  - In production, use a more reasonable timeout value

## Usage

To run this example:

```bash
go run main.go
```

## Other Backoff Strategies

Heimdall supports different backoff strategies:

1. **Constant Backoff**: Fixed time between retries (shown in this example)
2. **Exponential Backoff**: Increasing time between retries
3. **Custom Backoff**: Implement your own backoff strategy

## Benefits of Using Heimdall with Retries

- **Improved Reliability**: Automatically handles transient network failures
- **Reduced Error Handling**: Less manual error handling in your code
- **Controlled Retry Flow**: Prevents overwhelming servers with immediate retries
- **Configurable Behavior**: Easily adjust retry counts and intervals