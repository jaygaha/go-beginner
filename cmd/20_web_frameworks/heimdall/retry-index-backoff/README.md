# Heimdall HTTP Client - Exponential Backoff Example

This example demonstrates how to implement exponential backoff retry strategy using the [Heimdall](https://github.com/gojek/heimdall) HTTP client library for Go.

## What is Exponential Backoff?

Exponential backoff is a retry strategy that gradually increases the waiting time between retries. This approach is useful for:

- Reducing load on servers during temporary failures
- Increasing the likelihood of success for transient errors
- Implementing a more efficient retry mechanism compared to constant intervals

With exponential backoff, each retry waits longer than the previous one, typically doubling the wait time with each attempt.

## Example Implementation

This example demonstrates how to configure Heimdall with an exponential backoff strategy:

```go
// Set index backoff
backoff := heimdall.NewExponentialBackoff(
    500*time.Millisecond,    // Initial standby time (first interval of retry)
    5*time.Second,           // Maximum waiting time (maximum interval of retries)
    float64(15*time.Second), // Maximum elapsed time (maximum time for all retries)
    2.0,                     // Exponential increase rate (doubling the interval for each retry)
)

// create a new client with heimdall
client := httpclient.NewClient(
    httpclient.WithHTTPTimeout(10*time.Second),           // set the timeout to 10 seconds
    httpclient.WithRetrier(heimdall.NewRetrier(backoff)), // set the retrier to backoff
)
```

## Key Components

### Exponential Backoff Configuration

- **Initial Interval**: 500ms - The first retry will wait for 500 milliseconds
- **Maximum Interval**: 5s - No retry will wait longer than 5 seconds, regardless of the exponential calculation
- **Maximum Elapsed Time**: 15s - The total time spent on retries will not exceed 15 seconds
- **Multiplier**: 2.0 - Each retry interval will be twice as long as the previous one

### Retry Sequence Example

With the above configuration, the retry intervals would follow this pattern:

1. First retry: 500ms wait
2. Second retry: 1s wait (500ms × 2)
3. Third retry: 2s wait (1s × 2)
4. Fourth retry: 4s wait (2s × 2)
5. Fifth retry: 5s wait (capped at maximum interval)

Retries would stop after approximately 15 seconds of total elapsed time.

## Usage

To run this example:

```bash
go run main.go
```

## When to Use Exponential Backoff

Exponential backoff is particularly useful in scenarios such as:

- API calls to external services that might experience temporary outages
- Network requests over unreliable connections
- Distributed systems where components might be temporarily unavailable
- Rate-limited APIs where you need to back off when approaching limits

## Comparison with Constant Backoff

Unlike constant backoff (which uses the same interval for all retries), exponential backoff provides a more intelligent approach by:

- Starting with shorter intervals to quickly retry for transient errors
- Gradually increasing intervals to reduce load during persistent issues
- Providing a maximum cap to prevent excessive waiting