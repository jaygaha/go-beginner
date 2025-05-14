# Echo Web Framework

Echo is a high-performance, extensible, minimalist web framework for Go. This guide demonstrates how to use Echo to build web applications with Go, covering key concepts and features with practical examples.

## Features

- **High Performance**: Built on a fast HTTP router with zero dynamic memory allocation
- **Middleware Support**: Robust middleware system for request/response processing
- **Context-Based API**: Clean context-based request handling
- **Data Binding**: Automatic binding of request data to Go structs
- **Template Rendering**: Support for HTML templates
- **Error Handling**: Customizable error handling
- **Static File Serving**: Easy serving of static assets
- **Route Grouping**: Logical organization of routes

## Directory Structure

```
/echo
├── main.go           # Application entry point and server setup
├── go.mod            # Go module definition
├── go.sum            # Go module checksums
├── handlers/         # Request handlers
│   ├── admin_handler.go    # Admin route handlers
│   ├── error_handler.go    # Error handling
│   ├── hello_handler.go    # Basic route handlers
│   └── user_handler.go     # User-related handlers
├── static/          # Static assets
│   └── style.css    # CSS styles
└── views/           # HTML templates
    ├── error.html   # Error page template
    └── profile.html # User profile template
```

## Installation

1. Initialize a Go module:

```bash
go mod init yourmodule
```

2. Install Echo:

```bash
go get github.com/labstack/echo/v4
```

## Basic Usage

### Creating a Server

```go
package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	// Create a new Echo instance
	echo := echo.New()

	// Define a route
	echo.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Start the server
	echo.Logger.Fatal(echo.Start(":8800"))
}
```

## Key Concepts

### Routing

Echo provides a powerful routing system that supports various HTTP methods and path parameters:

```go
// Basic route
echo.GET("/", HelloHandler)

// Path parameter
echo.GET("/greet/:name", GreetUserHandler)

// Query parameters
echo.GET("search", SearchHandler)

// POST request handling
echo.POST("/users", CreateUserHandler)
```

### Path Parameters

Access path parameters using `c.Param()`:

```go
func GreetUserHandler(c echo.Context) error {
	name := c.Param("name") // get name from path which is /greet/:name
	
	// capitalize first letter
	if len(name) > 0 {
		name = strings.ToUpper(name[:1]) + name[1:]
	} else {
		name = "Guest"
	}

	return c.String(http.StatusOK, "Hello, "+name+"!")
}
```

### Query Parameters

Access query parameters using `c.QueryParam()`:

```go
func SearchHandler(c echo.Context) error {
	q := c.QueryParam("q") // get query param q from url which is /search?q=
	language := c.QueryParam("lang")

	if language == "" {
		language = "en" // default language
	}

	message := "Searching for: '" + q + "' in " + language

	return c.String(http.StatusOK, message)
}
```

### Middleware

Echo supports middleware for request/response processing:

```go
// Logger middleware (logs requests to the console)
echo.Use(middleware.Logger())

// Recover middleware (recovers from panics and logs them to the console)
echo.Use(middleware.Recover())
```

### Request Data Binding

Bind request data to Go structs:

```go
// User struct for JSON binding
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Handler for creating a new user (receives JSON)
func CreateUserHandler(c echo.Context) error {
	u := new(User)

	// bind the incoming JSON request to the User struct
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload: " + err.Error()})
	}

	// Process the user data...
	return c.JSON(http.StatusCreated, map[string]string{"message": "User created"})
}
```

### Template Rendering

Render HTML templates:

```go
// TemplateRenderer is a custom renderer for Echo framework
type TemplateRenderer struct {
	Templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

// In main.go:
renderer := &handlers.TemplateRenderer{
	Templates: template.Must(template.ParseGlob("views/*.html")),
}
echo.Renderer = renderer

// In handler:
func UserProfileHandler(c echo.Context) error {
	data := map[string]any{
		"Title":       "Go Beginner",
		"Name":        "Jay Gaha",
		"Age":         24,
		"CurrentTime": time.Now().Format(time.RFC1123),
	}

	return c.Render(http.StatusOK, "profile.html", data)
}
```

### Static Files

Serve static files:

```go
// Serve static files from the "static" directory
echo.Static("/static", "static") // (url_prefix, filesystem_path)
```

### Route Groups

Organize routes logically:

```go
// Create a new group for admin routes
adminGroup := echo.Group("/admin") // prefix all admin routes with /admin
{
	adminGroup.GET("", handlers.AdminDashboardHandler)
	adminGroup.GET("/settings", handlers.AdminSettingsHandler)
}
```

### Error Handling

Customize error handling:

```go
// Set custom error handler
echo.HTTPErrorHandler = handlers.CustomHTTPErrorHandler

// In error_handler.go:
func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	message := "Internal Server Error"

	// Check if it's an HTTP error
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		// Get custom message or use default
		if customMessage, ok := he.Message.(string); ok {
			message = customMessage
		} else {
			message = http.StatusText(code)
		}
	}

	// Render error template
	err = c.Render(code, "error.html", map[string]any{
		"Code":    code,
		"Message": message,
	})
}
```

## Running the Application

To run the application:

```bash
go run main.go
```

Then visit `http://localhost:8800` in your browser.

## Further Reading

- [Echo Official Documentation](https://echo.labstack.com/)
- [Go net/http Package Documentation](https://pkg.go.dev/net/http)
- [Go Templates Documentation](https://pkg.go.dev/html/template)
- [Echo GitHub Repository](https://github.com/labstack/echo)