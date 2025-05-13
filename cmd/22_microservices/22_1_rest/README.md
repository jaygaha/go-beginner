# RESTful Microservices in Go

This tutorial introduces microservices in Go by building an **Online Bookstore** with two microservices: a **Book Service** (handles book data) and an **Order Service** (handles orders). We'll use HTTP REST APIs for communication, store data in memory, and focus on simplicity to teach microservices concepts clearly.

## Tutorial Overview

We'll build:
1. **Book Service**: A RESTful API for managing book data.
2. **Order Service**: A RESTful API for managing orders like creating orders, checks availability.
3. **Communication**: HTTP GET requests from Order Service to Book Service to validate book IDs.

## Prerequisites
- Go 1.18 or later.
- Basic understanding of Go programming (structs, functions, HTTP basics).

## Architecture

Our microservices architecture consists of:

```
┌─────────────────┐      HTTP Request      ┌─────────────────┐
│                 │ ─────────────────────> │                 │
│   Book Service  │                        │  Order Service  │
│   (Port 8801)   │ <─────────────────────┤   (Port 8802)   │
│                 │      HTTP Response     │                 │
└─────────────────┘                        └─────────────────┘
```

- **Book Service**: Manages the book inventory (CRUD operations)
- **Order Service**: Handles order creation and management
- **Communication**: REST API calls between services

## Data Models

### Book Service

```go
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  int    `json:"price"`
}
```

### Order Service

```go
type Order struct {
	ID        int    `json:"id"`
	BookID    int    `json:"book_id"`
	Quantity  int    `json:"quantity"`
	TotalCost int    `json:"total_cost"`
	Customer  string `json:"customer"`
}
```

## API Endpoints

### Book Service (Port 8801)

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/books` | GET | Get all books |
| `/books` | POST | Add a new book |
| `/books/{id}` | GET | Get a book by ID |

### Order Service (Port 8802)

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/orders` | GET | Get all orders (can filter by customer) |
| `/orders` | POST | Create a new order |

## Implementation Details

### Book Service

- Uses an in-memory map to store books
- Thread-safe operations with sync.RWMutex
- Simple HTTP handlers for CRUD operations
- Generates random IDs for new books

### Order Service

- Uses an in-memory map to store orders
- Communicates with Book Service to validate book existence
- Thread-safe operations with sync.RWMutex
- Filters orders by customer

### Inter-Service Communication

The Order Service communicates with the Book Service using HTTP GET requests to validate that a book exists before creating an order:

```go
// Validate book exists by querying Book Service for specific book ID
resp, err := http.Get(fmt.Sprintf("%s/%d", bookServiceURL, order.BookID))
```

## Running the Services

1. Start the Book Service:
   ```bash
   cd book-service
   go run .
   ```

2. Start the Order Service (in a separate terminal):
   ```bash
   cd order-service
   go run .
   ```

3. Test the Book Service:
   ```bash
   # Add a book
   curl -X POST http://localhost:8801/books -H "Content-Type: application/json" -d '{"title":"The Go Programming Language","author":"Alan A. A. Donovan","price":30}'
   
   # Get all books
   curl http://localhost:8801/books
   ```

4. Test the Order Service:
   ```bash
   # Create an order (replace {book_id} with an actual book ID)
   curl -X POST http://localhost:8802/orders -H "Content-Type: application/json" -d '{"book_id":{book_id},"quantity":1,"customer":"John Doe"}'
   
   # Get all orders
   curl http://localhost:8802/orders
   ```

[Note] Sample JSON collection included in the repository. [Check here](collection_postman_go-rest-microservice.json)

## Benefits of This Architecture

1. **Separation of Concerns**: Each service has a specific responsibility
2. **Independent Deployment**: Services can be deployed independently
3. **Technology Flexibility**: Different services can use different technologies
4. **Scalability**: Services can be scaled independently based on demand
5. **Resilience**: Failure in one service doesn't bring down the entire system

## Potential Improvements

1. **Database Integration**: Replace in-memory storage with a persistent database
2. **Service Discovery**: Implement service discovery for dynamic service locations
3. **API Gateway**: Add an API gateway to handle cross-cutting concerns
4. **Authentication**: Implement authentication and authorization
5. **Containerization**: Package services in Docker containers
6. **Monitoring**: Add logging and monitoring capabilities
7. **Circuit Breaker**: Implement circuit breaker pattern for resilience

## Conclusion

This simple implementation demonstrates the core concepts of microservices architecture in Go. By focusing on clear separation of concerns and simple HTTP communication, we've built a foundation that can be extended with more advanced patterns and technologies as needed.