# Fiber Web Framework

This directory contains examples and tutorials for working with the Fiber web framework in Go. Fiber is a popular, lightweight, and fast web framework inspired by Express.js that makes building web applications in Go simple and intuitive.

## What is Fiber?

Fiber is a Go web framework built on top of Fasthttp, offering:

- **High Performance**: Zero memory allocation and optimized HTTP routing
- **Express-like Syntax**: Intuitive API design familiar to Express.js developers
- **Rich Ecosystem**: Official and third-party middlewares for common tasks
- **Flexibility**: Easy integration with Go's standard library and third-party packages

## Directory Structure

```
/fiber
├── main.go         # Basic Fiber application setup
├── routing/        # Routing examples and patterns
└── testing/        # Testing Fiber applications
```

## Getting Started

### Installation

```bash
go get github.com/gofiber/fiber/v2
```

### Basic Usage (main.go)

The main.go file demonstrates how to create a simple Fiber application:

```go
package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize with custom config
	app := fiber.New(fiber.Config{
		BodyLimit:       10 * 1024 * 1024, // 10MB
		ReadBufferSize:  8192,             // 8KB
		WriteBufferSize: 8192,             // 8KB
	})

	// Define a route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Gophers!")
	})

	// Start the server
	app.Listen(":3300")
}
```

## Routing Examples

The `routing/` directory contains examples of different routing patterns in Fiber:

### Basic Routing

```go
// HTTP method-based routing
app.Get("/", handler)      // GET
app.Post("/users", handler) // POST
app.Put("/users", handler)  // PUT
app.Delete("/users", handler) // DELETE
```

### Dynamic Routes

```go
// Route with path parameter
app.Get("/greet/:name", func(c *fiber.Ctx) error {
	name := c.Params("name")
	return c.JSON(fiber.Map{
		"message": "Hello, " + name + "!",
	})
})
```

### Query Parameters

```go
// Route with query parameters
app.Get("/search", func(c *fiber.Ctx) error {
	query := c.Query("q")
	return c.JSON(fiber.Map{
		"message": "Searching for: " + query,
	})
})
```

### Request Body Parsing

```go
// Parsing JSON request body
app.Post("/users", func(c *fiber.Ctx) error {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var user User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User created successfully",
		"data":    user,
	})
})
```

## Middleware

Fiber supports middleware for request processing:

### Global Middleware

```go
// Apply middleware to all routes
app.Use(func(c *fiber.Ctx) error {
	// Log request details
	fmt.Printf("Request: %s %s\n", c.Method(), c.Path())
	return c.Next()
})
```

### Route-Specific Middleware

```go
// Middleware for specific routes
middleware := func(c *fiber.Ctx) error {
	c.Set("X-Powered-By", "Go-Fiber")
	return c.Next()
}

app.Get("/protected", middleware, handler)
```

### Route Grouping

```go
// Group routes with shared middleware
api := app.Group("/api", middleware)
v1 := api.Group("/v1", anotherMiddleware)

// Routes become: /api/v1/posts
v1.Get("/posts", postsHandler)
```

## Error Handling

Fiber provides built-in error handling:

### Custom Error Handler

```go
app := fiber.New(fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		message := "Internal Server Error"
		
		// Check if it's a Fiber error
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
			message = e.Message
		}
		
		return c.Status(code).JSON(fiber.Map{
			"message": message,
		})
	},
})
```

### Returning Errors

```go
// Return a custom error
app.Get("/error", func(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusNotFound, "Not found")
})
```

## Testing Fiber Applications

The `testing/` directory demonstrates how to test Fiber applications using Go's standard testing package and the `httptest` package.

### Basic Test Example

```go
func TestHome(t *testing.T) {
	app := NewServer()

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response, _ := app.Test(request)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	// Read response body
	body := make([]byte, response.ContentLength)
	_, err := response.Body.Read(body)

	// EOF is expected when the entire body has been read
	assert.Equal(t, io.EOF, err)
	// Verify response content
	assert.Equal(t, "Welcome to Fiber!", string(body))
}
```

### Testing Error Responses

```go
func TestInternalServerError(t *testing.T) {
	app := NewServer()
	request := httptest.NewRequest(http.MethodGet, "/500", nil)
	response, _ := app.Test(request)
	assert.Equal(t, http.StatusInternalServerError, response.StatusCode)

	// Read response body
	body := make([]byte, response.ContentLength)
	_, err := response.Body.Read(body)
	assert.Equal(t, io.EOF, err)
	
	// Verify JSON response
	assert.Equal(t, "{\"message\":\"Internal Server Error\"}", string(body))
}
```

## Best Practices

1. **Organize Routes**: Group related routes together for better maintainability
2. **Use Middleware**: Leverage middleware for cross-cutting concerns like logging and authentication
3. **Proper Error Handling**: Always handle errors and return appropriate status codes
4. **Testing**: Write tests for your routes to ensure they work as expected
5. **Configuration**: Use Fiber's configuration options to optimize performance

## Additional Resources

- [Official Fiber Documentation](https://docs.gofiber.io/)
- [Fiber GitHub Repository](https://github.com/gofiber/fiber)
- [Fiber Examples](https://github.com/gofiber/recipes)

## Running the Examples

To run any of the examples in this directory:

```bash
cd cmd/20_web_frameworks/fiber
go run main.go
# or
cd routing
go run main.go
```

The server will start on port 3300. You can access it at http://localhost:3300.