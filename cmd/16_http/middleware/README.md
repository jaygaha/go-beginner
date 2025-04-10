# HTTP Middleware in Go

This guide explains HTTP middleware in Go for beginners, with practical examples from this repository.

## What is Middleware?

In Go web applications, middleware is a function that processes HTTP requests before they reach their final handler, or processes responses before they're sent back to the client. Middleware functions can:

- Execute code before/after the handler
- Modify request and response objects
- Terminate the request-response cycle
- Call the next middleware in the chain

## Key Middleware Concepts

1. **Function Chaining**: Middleware can be chained together to create a pipeline of processing
2. **Request/Response Modification**: Middleware can read and modify both incoming requests and outgoing responses
3. **Early Termination**: Middleware can decide to stop the request from reaching subsequent handlers
4. **Reusability**: Well-designed middleware can be reused across different routes and applications

## Common Middleware Use Cases

- **Logging**: Recording request details for monitoring and debugging
- **Authentication/Authorization**: Verifying user identity and permissions
- **Request ID Generation**: Adding unique identifiers to requests for tracing
- **Error Handling**: Catching and handling panics to prevent application crashes
- **CORS**: Managing Cross-Origin Resource Sharing
- **Content Compression**: Compressing response data
- **Rate Limiting**: Controlling request frequency

## Middleware Implementation Patterns

### Function Wrapper Pattern

The most common pattern in Go is the function wrapper, where middleware takes a handler and returns a new handler:

```go
func MyMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Do something before calling the next handler
        
        next.ServeHTTP(w, r)
        
        // Do something after calling the next handler
    })
}
```

## Chaining Middleware

Middleware can be chained in different ways:

1. **Direct nesting**:
    ```go
    http.Handle("/", MyMiddleware(MyOtherMiddleware(http.HandlerFunc(MyHandler))))
    ```
2. **Using a middleware list**:
    ```go
    middlewareList := []Middleware{MyMiddleware, MyOtherMiddleware}
    http.Handle("/", ChainMiddleware(middlewareList, http.HandlerFunc(MyHandler)))


## Examples in This Repository

This repository contains several examples of middleware in action:
- `middleware/log.go`: Logs request details
- `middleware/authenticate.go`: Checks for valid authentication methods
- `uuid.go`: Generates a UUID for each request
- `middleware/recovery.go`: Recovers from panics

Handlers:
- `handlers/welcome.go`: A simple welcome handler
- `handlers/auth.go`: A handler for authentication
- `handlers/panic.go`: A handler that panics

## Best Practices

- Keep middleware focused : Each middleware should do one thing well
- Order matters : Consider the sequence of your middleware chain
- Handle errors appropriately : Don't let errors pass silently
- Be mindful of performance : Middleware runs on every request
- Use context for request-scoped values : Avoid global state

## Further Reading

- [Go's http.Handler interface](https://pkg.go.dev/net/http#Handler)
- [Context package documentation](https://pkg.go.dev/context)