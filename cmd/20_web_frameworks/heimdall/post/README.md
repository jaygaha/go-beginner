# Heimdall HTTP Client - POST Example

## Making POST Requests with Heimdall

The example in `main.go` demonstrates how to:

1. Create a Heimdall HTTP client with custom timeout and retry settings
2. Prepare JSON data for a POST request
3. Set appropriate headers
4. Make the POST request
5. Handle the response

### Code Explanation

```go
// Create a new client with heimdall
client := httpclient.NewClient(
    httpclient.WithHTTPTimeout(10*time.Second), // set the timeout to 10 seconds
    httpclient.WithRetryCount(3),               // set the number of retries to 3
)
```

This creates a new Heimdall HTTP client with a 10-second timeout and configured to retry failed requests up to 3 times.

```go
// JSON data
jsonData := []byte(`{"body": "Comments by commenter", "postId": 3, "userId": 1}`)
headers := http.Header{
    "Content-Type": []string{"application/json"},
}
```

Here we prepare the JSON data to send and set the appropriate `Content-Type` header to indicate we're sending JSON data.

```go
// make a POST request
resp, err := client.Post("https://dummyjson.com/comments/add", bytes.NewBuffer(jsonData), headers)
```

This line makes the actual POST request to the specified URL, sending the JSON data with the configured headers.

## Response Handling

The example shows how to properly handle the response:

```go
// read the response
body, err := io.ReadAll(resp.Body)
if err != nil {
    log.Fatalf("Error reading response body: %v", err)
}

// print the response
fmt.Println(string(body))
```

This reads the response body and prints it to the console. The response from the example API looks like:

```json
{"id":341,"body":"Comments by commenter","postId":3,"user":{"id":1,"username":"emilys","fullName":"Emily Johnson"}}
```

## Key Benefits of Using Heimdall

1. **Reliability**: Automatic retries help handle transient network issues
2. **Performance**: Configurable timeouts prevent requests from hanging indefinitely
3. **Scalability**: Designed for applications that need to make many HTTP requests
4. **Simplicity**: Clean API makes it easy to use for common HTTP operations

## Running the Example

To run this example:

```bash
go run main.go
```