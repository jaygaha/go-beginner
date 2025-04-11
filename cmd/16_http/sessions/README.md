# Go Sessions and Cookies

This guide explains how to implement and manage sessions and cookies in Go web applications. It's designed for beginners who want to understand how to maintain state across HTTP requests.

## Introduction

HTTP is a stateless protocol, meaning each request is independent and the server doesn't maintain any information about previous requests. However, web applications often need to remember information about users across multiple requests. This is where cookies and sessions come in.

## Cookies

### What are Cookies?

Cookies are small pieces of data stored on the client's browser. They are sent with every request to the server, allowing the server to identify the client and maintain state.

### Setting Cookies

In Go, you can set cookies using the `http.SetCookie` function:

```go
cookie := http.Cookie{
    Name:     "go-cookie",
    Value:    "go-cookie-value",
    Expires:  time.Now().Add(24 * time.Hour), // 1 day
    Path:     "/",                            // all paths
    HttpOnly: true,                           // only accessible via HTTP
    Secure:   false,                          // only accessible via HTTPS
}

http.SetCookie(w, &cookie)
```

Important cookie attributes:

- Name : The name of the cookie
- Value : The value stored in the cookie
- Expires : When the cookie should expire
- Path : The URL path where the cookie is valid
- HttpOnly : If true, the cookie is inaccessible to JavaScript
- Secure : If true, the cookie is only sent over HTTPS

### Reading Cookies

To read cookies, you can use the `http.Request.Cookie` method:

```go
cookie, err := r.Cookie("go-cookie")
if err != nil {
    // Handle error (cookie not found or other error)
    return
}
// Use cookie.Value
```

### Deleting Cookies

To delete a cookie, set its expiration time to the past:

For example:

```go
cookie := http.Cookie{
    Name:    "go-cookie",
    Value:   "",
    Expires: time.Now().Add(-1 * time.Hour), // negative time
    Path:    "/",
}

http.SetCookie(w, &cookie)
```

## Sessions

### What are Sessions?

Sessions are server-side storage mechanisms that maintain state across multiple HTTP requests. Unlike cookies, session data is stored on the server, with only a session ID stored in a cookie on the client.

### Setting Up Sessions

Best way to set up sessions is to use the `gorilla/sessions` package. This package provides a simple and secure way to manage sessions in Go.

Install the package:

```bash
go get github.com/gorilla/sessions
```

Initialize the session store:

For example:

```go
// Key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
var SessionStore = sessions.NewCookieStore([]byte("secret-key"))

// Configure session options
SessionStore.Options = &sessions.Options{
    Path:     "/",
    MaxAge:   60 * 15, // 15 minutes
    HttpOnly: true,
    Secure:   false, // set to true if using HTTPS
}
```

### Creating a Session

To create a new session or retrieve an existing one, you can use the `sessions.SessionStore.New` method.

For example:

```go
// Get the session
session, _ := SessionStore.Get(r, "user-session")

// Set session values
session.Values["username"] = "user123"
session.Values["authenticated"] = true

// Save the session
if err := session.Save(r, w); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
}
```

### Accessing Session Data

To access session data, you can use the `sessions.Session.Values` map:

For example:

```go
// Get the session
session, _ := SessionStore.Get(r, "user-session")

// Check if a value exists and retrieve it
if username, ok := session.Values["username"].(string); ok {
    // Use username
}

// Check authentication status
if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
    http.Error(w, "Forbidden", http.StatusForbidden)
    return
}
```

### Destroying a Session

To destroy a session, you can set its expiration time to the past:

For example:

```go
// Get the session
session, _ := SessionStore.Get(r, "user-session")

// Revoke user authentication
session.Values["authenticated"] = false
session.Options.MaxAge = -1 // Delete the cookie

// Save the session
if err := session.Save(r, w); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
}
```

## Security Considerations

- **Use HTTPS** : Always use HTTPS in production to protect cookies and session IDs from being intercepted.
- **Secure Cookie Flag** : Set the Secure flag to true in production to ensure cookies are only sent over HTTPS.
- **HttpOnly Flag** : Set the HttpOnly flag to true to prevent JavaScript from accessing cookies, protecting against XSS attacks.
- **Strong Secret Keys** : Use strong, random secret keys for your session store.
- **Session Expiration** : Set appropriate expiration times for sessions to limit the window of opportunity for session hijacking.
- **CSRF Protection** : Implement Cross-Site Request Forgery protection for sensitive operations.

## Running the Example

To run the example, run the following command:

```bash
go run .
```

Server will start at [`http://localhost:8800`](http://localhost:8800). Open in your browser and navigate to the links to see the example in action.

The example demonstrates basic usage of cookies and sessions in Go.

- Setting, reading, and deleting cookies
- Creating sessions during login
- Accessing protected routes with session authentication
- Destroying sessions during logout

