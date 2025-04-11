# Serving Static Assets in Go

This guide explains how to serve static files (CSS, JavaScript, images, etc.) in Go web applications. It covers the basics of file serving, common patterns, and best practices.

Web applications often need to serve static assets like CSS stylesheets, JavaScript files, images, and other resources. Go's standard library provides simple and efficient ways to serve these files through the `net/http` package.

## Basic File Serving

The simplest way to serve static files in Go is using the `http.FileServer` handler:

```go
// Create a file server that serves files from the "assets" directory
fileServer := http.FileServer(http.Dir("assets"))

// Register the file server to handle all requests
http.Handle("/", fileServer)

// Start the server
http.ListenAndServe(":8080", nil)

```

This will serve files from the "assets" directory at the root URL path. For example, a file at `assets/css/style.css` would be accessible at `http://localhost:8080/css/style.css`.

###  Serving from a Subdirectory

Often, you'll want to serve static files from a specific URL path prefix, like `/static/`. This is where `http.StripPrefix` comes in:

```go
// Create a file server that serves files from the "assets" directory
fileServer := http.FileServer(http.Dir("assets"))

// Serve files under the "/static/" path
http.Handle("/static/", http.StripPrefix("/static/", fileServer))

// Start the server
http.ListenAndServe(":8080", nil)
```

With this setup, a file at `assets/css/style.css` would be accessible at `http://localhost:8080/static/css/style.css`.

The `StripPrefix` function removes the specified prefix from the URL path before passing the request to the file server. This is necessary because the file server expects the URL path to match the file system path relative to the root directory.

## Handling File Not Found Errors

By default, `http.FileServer` returns a 404 Not Found error when a file is not found. You can customize this behavior by wrapping the file server with your own handler:

```go
// Custom file server handler
func systemFileServer(fs http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Try to serve the file
        fs.ServeHTTP(w, r)
        
        // Check if the response was a 404
        if w.(*responseWriter).status == http.StatusNotFound {
            // Serve a custom 404 page
            http.ServeFile(w, r, "pages/404.html")
        }
    })
}

// Use the custom file server
fileServer := http.FileServer(http.Dir("assets"))
http.Handle("/static/", systemFileServer(http.StripPrefix("/static/", fileServer)))
```

*Note: This example requires a custom responseWriter that tracks the response status code.*

## Caching

It is a good practice to enable caching for static assets to improve performance. You can do this by setting the appropriate headers in your response:

```go
func cachedFileServer(fs http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Set caching headers
        w.Header().Set("Cache-Control", "max-age=86400") // Cache for 1 day
        
        // Serve the file
        fs.ServeHTTP(w, r)
    })
}

// Use the cached file server
fileServer := http.FileServer(http.Dir("assets"))
http.Handle("/static/", cachedFileServer(http.StripPrefix("/static/", fileServer)))
```

## Best Practices

- Consider using CDN for static assets.
- Avoid serving sensitive files.
- Use dedicated directory for static assets.

## Examples

Check `main.go` for examples

### Directory Structure

```plaintext
├── main.go
└── assets
    ├── css
    │   └── style.css
    └── images
        └── golang.png

### Use frameworks 

For more advanced use cases, consider using a Go web framework like Gin, Echo or other available options.