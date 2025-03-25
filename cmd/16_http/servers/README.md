# HTTP Servers in Go

This package demonstrates various HTTP server implementations in Go, showcasing different request handling patterns and server configurations.

## Multiplexing requests handlers
 - Go facilitates multiplexing request handlers.
 - handles multiple requests concurrently using a single server instance using goroutine.
 - In the context of web servers, multiplexing allows a single server to handle multiple HTTP requests simultaneously. 

## Features

### 1. Basic Welcome Server
- **Endpoint**: `/`
- **Method**: GET
- **Response**: Plain text welcome message
- **Example Response**: `Welcome to go-phers world!`

### 2. Hello Handler with Query Parameters
- **Endpoint**: `/hello`
- **Method**: GET
- **Query Parameters**: `name` (string)
- **Example Request**: `/hello?name=John`
- **Example Response**: `Hello, John!`

### 3. Hello Handler with POST Form Data
- **Endpoint**: `/hello`
- **Method**: POST
- **Content-Type**: `application/x-www-form-urlencoded`
- **Form Parameters**: `name` (string, required)
- **Response Codes**:
  - 201: Successfully processed request
  - 422: When name is empty
- **Example Request**:
  ```
  POST /hello
  Content-Type: application/x-www-form-urlencoded

  name=John
  ```
- **Example Response**: `Hello, John!`

### 4. Multiple Server Configuration
- Demonstrates running multiple HTTP servers on different ports
- Each server includes context-aware welcome messages
- Server addresses are configurable (default: :8010 and :8011)
- **Example Response**: `Welcome to go-phers world! from server :8010`

## Usage

1. Import the package in your Go application
2. Initialize the desired handlers
3. Configure server settings as needed
4. Start the HTTP server(s)

Refer to the test cases for detailed usage examples of each handler.