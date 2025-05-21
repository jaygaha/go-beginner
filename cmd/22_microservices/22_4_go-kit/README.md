# Go-Kit Microservice Tutorial

This project demonstrates how to build a microservice using the Go-Kit framework. It implements a simple feedback service with basic CRUD operations.

## What is Go-Kit?

Go-Kit is a toolkit for building microservices in Go. It provides a set of packages and best practices that help you structure your application according to clean architecture principles. Go-Kit is based on three main components:

1. **Service** - Contains your business logic
2. **Endpoint** - Defines the request/response format for each method
3. **Transport** - Handles the communication protocol (HTTP, gRPC, etc.)

## Project Structure

```
/22_4_go-kit
├── feedback/               # Feedback service implementation
│   ├── endpoint.go         # Endpoint definitions
│   ├── entity.go           # Data models
│   ├── logic.go            # Business logic implementation
│   ├── repo.go             # Repository for data storage
│   ├── reqresp.go          # Request/response structures
│   ├── server.go           # HTTP transport layer
│   └── service.go          # Service interface definition
├── go-kit-microservice.json # Postman collection
├── go.mod                  # Go module file
├── go.sum                  # Go module checksums
└── main.go                 # Application entry point
```

## Features

This microservice provides the following endpoints:

- `POST /feedbacks` - Add a new feedback
- `GET /feedbacks/{id}` - Get a specific feedback by ID
- `GET /feedbacks` - Get all feedbacks

## Getting Started

### Prerequisites

- Go 1.16 or higher

### Running the Service

1. Clone the repository
2. Navigate to the project directory
3. Run the service:

```bash
go run main.go
```

The service will start on port 8800 by default. You can specify a different port using the `-http` flag:

```bash
go run main.go -http=:8080
```

## API Usage

### Add Feedback

```bash
curl -X POST http://localhost:8800/feedbacks \
  -H "Content-Type: application/json" \
  -d '{"message": "My feedback to your app"}'
```

### Get Feedback by ID

```bash
curl -X GET http://localhost:8800/feedbacks/{id}
```

### Get All Feedbacks

```bash
curl -X GET http://localhost:8800/feedbacks
```

## Postman Collection

A Postman collection is included in the project (`go-kit-microservice.json`). You can import this file into Postman to test the API endpoints.

## Key Concepts

### Service Layer

The service layer defines the interface and implements the business logic. In this project, the `Service` interface is defined in `service.go` and implemented in `logic.go`.

### Endpoint Layer

The endpoint layer converts the service methods into endpoint functions that can be used by the transport layer. Each endpoint function takes a request object and returns a response object.

### Transport Layer

The transport layer handles the HTTP communication. It defines how requests are decoded and how responses are encoded. In this project, the transport layer is implemented in `server.go` and uses the Gorilla Mux router.

### Repository Layer

The repository layer handles data storage. In this project, it's a simple in-memory implementation, but in a real-world application, it would typically connect to a database.

## Further Reading

- [Go-Kit GitHub Repository](https://github.com/go-kit/kit)
- [Go-Kit Documentation](https://gokit.io/)