# HTTP Clients in Go

This module demonstrates how to implement HTTP clients in Go, showcasing different HTTP methods and request handling patterns.

## Overview

The HTTP client implementation in Go allows you to make requests to HTTP servers and process their responses. This module covers:

- Creating and configuring HTTP clients
- Making requests with different HTTP methods (GET, POST, PUT, PATCH, DELETE)
- Handling responses and errors
- Working with JSON data
- Using context for timeout control

## Key Concepts

### HTTP Client Creation

Go provides two main ways to create HTTP clients:

1. **Using `http.DefaultClient`**: A pre-configured client suitable for basic use cases
2. **Creating a custom client**: Allows configuration of timeouts, transport options, etc.

```go
// Using default client
resp, err := http.DefaultClient.Do(req)

// Creating a custom client with timeout
client := &http.Client{
    Timeout: 10 * time.Second,
}
resp, err := client.Do(req)
```

### HTTP Methods

The module demonstrates all standard HTTP methods:

#### GET

Used to retrieve data from a server:

```go
// Create a GET request
req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.example.com/posts/42", nil)
if err != nil {
    // Handle error
}

// Send the request
resp, err := client.Do(req)
```

#### POST

Used to send data to create a resource on the server:

```go
// Prepare data to send
postData := postBody{
    UserID: 5,
    Title:  "Example Title",
    Body:   "Example Body",
}

// Marshal to JSON
postBodyJson, err := json.Marshal(postData)

// Create a POST request with the JSON data
req, err := http.NewRequestWithContext(
    ctx,
    http.MethodPost,
    "https://api.example.com/posts",
    bytes.NewBuffer(postBodyJson),
)

// Set content type header
req.Header.Set("Content-Type", "application/json")

// Send the request
resp, err := client.Do(req)
```

#### PUT

Used to update an existing resource on the server:

```go
// Create a PUT request with updated data
req, err := http.NewRequestWithContext(
    ctx,
    http.MethodPut,
    "https://api.example.com/posts/42",
    bytes.NewBuffer(updatedDataJson),
)
```

#### PATCH

Used to partially update an existing resource:

```go
// Create a PATCH request with partial data
req, err := http.NewRequestWithContext(
    ctx,
    http.MethodPatch,
    "https://api.example.com/posts/42",
    bytes.NewBuffer(partialDataJson),
)
```

#### DELETE

Used to delete a resource:

```go
// Create a DELETE request
req, err := http.NewRequestWithContext(
    ctx,
    http.MethodDelete,
    "https://api.example.com/posts/42",
    nil,
)
```

### Response Handling

After making a request, you need to handle the response:

```go
// Check for errors in making the request
if err != nil {
    // Handle network or client errors
    return
}

// Always close the response body when done
defer resp.Body.Close()

// Check the status code
if resp.StatusCode != http.StatusOK {
    // Handle non-200 status codes
    return
}

// Read and process the response body
var responseData SomeStruct
if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
    // Handle JSON decoding error
    return
}

// Use the response data
fmt.Printf("%+v\n", responseData)
```

### Using Context for Timeouts

Go's context package helps manage request timeouts and cancellations:

```go
// Create a context with timeout
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// Use the context when creating the request
req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
```

## Testing HTTP Clients

The module includes comprehensive tests that demonstrate how to test HTTP clients using Go's `httptest` package:

- Creating mock HTTP servers
- Verifying request methods, headers, and bodies
- Simulating different response scenarios
- Testing error handling

## Usage Examples

The main.go file contains complete examples of using HTTP clients with all the methods mentioned above. Run the examples with:

```bash
go run main.go
```

To run the tests:

```bash
go test -v
```