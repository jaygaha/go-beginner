# Go HTTP Web Programming

This directory contains examples of HTTP web programming in Go, demonstrating both client and server implementations, form handling, and middleware concepts.

## Directory Structure

- [`/servers`](./servers) - HTTP server implementations
- [`/clients`](./clients) - HTTP client implementations
- [`/forms`](./forms) - Form handling and validation
- [`/middleware`](./middleware) - HTTP middleware patterns
- [`/sessions`](./sessions) - Session & Cookie management [TODO]
- [`/websockets`](./websockets) - WebSocket support [TODO]
- [`/uploads`](./uploads) - File upload handling [TODO]

## HTTP Servers (`/servers`)

The servers directory demonstrates how to create HTTP servers in Go, handling different types of requests and routing patterns.

### Features

- Basic HTTP server setup
- Request routing with http.ServeMux
- Handling different HTTP methods
- Working with query parameters, path parameters, and request bodies
- Running multiple servers concurrently

**[Further Reading](servers/README.md)**

## HTTP Clients (`/clients`)

The clients directory demonstrates how to make HTTP requests to external services using Go's `net/http` package.

### Features

- Creating and configuring HTTP clients
- Making requests with different HTTP methods (GET, POST, PUT, PATCH, DELETE)
- Handling responses and errors
- Working with JSON data
- Using context for timeout control

**[Further Reading](clients/README.md)**

## Form Handling (`/forms`)

The forms directory demonstrates how to handle HTML forms in Go, including form rendering, submission processing, and validation.

### Features

- Rendering HTML forms with templates
- Processing form submissions
- Form validation techniques
- Handling file uploads
- Displaying validation errors

**[Further Reading](forms/README.md)**

## Middleware (`/middleware`)

The middleware directory demonstrates how to implement HTTP middleware in Go, which is a software design pattern that allows you to intercept and modify HTTP requests and responses.

### Features

- Implementing middleware for logging, authentication, and rate limiting
- Chaining multiple middleware functions
- Passing data between middleware functions
- Handling errors in middleware

**[Further Reading](middleware/README.md)**

## Key Concepts

### HTTP Methods

- GET : Retrieve data from the server
- POST : Submit data to the server
- PUT : Update an entire resource
- PATCH : Partially update a resource
- DELETE : Remove a resource
- OPTIONS : Get information about the resource
- HEAD : Get information about the resource without the body

### Status Codes

- 2xx : Success (200 OK, 201 Created, etc.)
- 3xx : Redirection (301 Moved Permanently, 302 Found, etc.)
- 4xx : Client Error (400 Bad Request, 404 Not Found, etc.)
- 5xx : Server Error (500 Internal Server Error, etc.)

### Content Types

- `application/json` : JSON data
- `application/x-www-form-urlencoded` : Form data
- `multipart/form-data` : Form data with file uploads
- `text/html` : HTML content
- `text/plain` : Plain text

## Best Practices
- Use appropriate HTTP methods for each request
- Validate input data to prevent security vulnerabilities
- Handle errors gracefully and provide meaningful error messages
- Use appropriate status codes and content types
- Follow RESTful API design principles
- Use context for timeout control
- Use middleware for common functionality (e.g. logging, authentication)

## Further Reading

- [Go net/http Package Documentation](https://pkg.go.dev/net/http)
- [Go by Example: HTTP Clients](https://gobyexample.com/http-clients)
- [Go by Example: HTTP Servers](https://gobyexample.com/http-servers)
- [Middleware pattern in Go](https://drstearns.github.io/tutorials/gomiddleware/)
