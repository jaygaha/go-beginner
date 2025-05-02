# Basic RESTful API in Go

This is a simple implementation of a RESTful API using Go's standard library. This project demonstrates how to create a basic CRUD (Create, Read, Update, Delete) API without using any third-party frameworks.

## What is a RESTful API?

REST (Representational State Transfer) is an architectural style for designing networked applications. RESTful APIs use HTTP requests to perform CRUD operations:

- **C**reate: POST
- **R**ead: GET
- **U**pdate: PUT
- **D**elete: DELETE

## Project Structure

```
rest/
├── handlers/
│   └── item_handler.go  # Contains the CRUD operation handlers
├── routes/
│   └── api.go           # Routes HTTP requests to appropriate handlers
└── main.go              # Entry point of the application
```

## Implementation Details

### Data Model

This example uses a simple `Bag` struct with in-memory storage:

```go
type Bag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	SKU  int    `json:"sku"`
}
```

### API Endpoints

All endpoints are accessible via the `/bags` route with different HTTP methods:

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET    | /bags    | Get all bags |
| GET    | /bags?id=1 | Get a specific bag by ID |
| POST   | /bags    | Create a new bag |
| PUT    | /bags?id=1 | Update a bag by ID |
| DELETE | /bags?id=1 | Delete a bag by ID |

## How to Run

1. Make sure you have Go installed on your system
2. Navigate to the project directory
3. Run the server:

```bash
go run main.go
```

The server will start on port 8800.

## API Usage Examples

### Create a Bag (POST)

```bash
curl -X POST http://localhost:8800/bags \
  -H "Content-Type: application/json" \
  -d '{"name":"Backpack","sku":12345}'
```

### Get All Bags (GET)

```bash
curl http://localhost:8800/bags
```

### Get a Specific Bag (GET)

```bash
curl http://localhost:8800/bags?id=1
```

### Update a Bag (PUT)

```bash
curl -X PUT http://localhost:8800/bags?id=1 \
  -H "Content-Type: application/json" \
  -d '{"id":1,"name":"Updated Backpack","sku":54321}'
```

### Delete a Bag (DELETE)

```bash
curl -X DELETE http://localhost:8800/bags?id=1
```

## Key Concepts

### 1. HTTP Methods

The API uses standard HTTP methods to perform different operations:
- `GET`: Retrieve data
- `POST`: Create new data
- `PUT`: Update existing data
- `DELETE`: Remove data

### 2. JSON Encoding/Decoding

The API uses Go's `encoding/json` package to:
- Decode JSON from request bodies (`json.NewDecoder(r.Body).Decode(&bag)`)
- Encode data as JSON in responses (`json.NewEncoder(w).Encode(bag)`)

### 3. Concurrency Safety

The implementation uses a mutex (`sync.Mutex`) to ensure thread safety when accessing the shared in-memory data store.

### 4. HTTP Status Codes

The API returns appropriate HTTP status codes:
- 200 OK: Successful GET and PUT requests
- 201 Created: Successful POST requests
- 204 No Content: Successful DELETE requests
- 400 Bad Request: Invalid input
- 404 Not Found: Resource not found
- 405 Method Not Allowed: Unsupported HTTP method

## Next Steps

To enhance this basic implementation, consider:

1. Adding input validation
2. Implementing persistent storage (database)
3. Adding authentication and authorization
4. Implementing pagination for large collections
5. Adding logging and error handling
6. Creating middleware for cross-cutting concerns

This simple example provides a foundation for understanding how RESTful APIs work in Go without the complexity of third-party frameworks.