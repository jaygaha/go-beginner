# Advanced SQLite Implementation in Go

This project is the advanced implementation of previous SQLite implementation in Go. It focuses on advanced features and best practices for working with SQLite databases in Go.

## Features

- **Connection Pooling**: Configurable connection pool settings for optimal performance
- **Transaction Support**: Proper transaction management with commit and rollback
- **Context Usage**: All database operations support context for timeout and cancellation
- **Prepared Statements**: Efficient query execution with prepared statements
- **Error Handling**: Comprehensive error handling with custom error types
- **Structured Data**: Clean struct-based data representation

## Database Structure

The application uses a simple grocery database with the following schema:

```sql
CREATE TABLE IF NOT EXISTS groceries (
    id INTEGER PRIMARY KEY,
    name VARCHAR(255),
    price REAL,
    quantity INTEGER DEFAULT 0,
    is_available BOOLEAN DEFAULT false,
    created_at DATETIME,
    updated_at DATETIME DEFAULT NULL
)
```

## Data Model

The `Grocery` struct represents items in the database:

```go
type Grocery struct {
    ID          int        // Primary key
    Name        string     // Name of the grocery item
    Price       float64    // Price of the grocery item
    Quantity    int        // Quantity of the grocery item
    IsAvailable bool       // Availability status
    CreatedAt   time.Time  // Time when the grocery item was created
    UpdatedAt   *time.Time // Time when the grocery item was last updated (can be nil)
}
```

## API Reference

### Database Connection

```go
func ConnectDB() error
```

Establishes a connection to the SQLite database and sets up the schema. Configures connection pool settings for optimal performance.

### Adding Items

```go
func AddGrocery(ctx context.Context, name string, price float64, quantity int) (int64, error)
```

Adds a new grocery item to the database and returns the ID of the inserted item.

```go
func AddGroceriesInTransaction(ctx context.Context, groceries []Grocery) error
```

Adds multiple grocery items in a single transaction for better performance and data integrity.

### Retrieving Items

```go
func GetAllGroceries(ctx context.Context) ([]Grocery, error)
```

Retrieves all grocery items from the database.

```go
func GetGroceryByID(ctx context.Context, id int) (Grocery, error)
```

Retrieves a specific grocery item by ID.

### Updating Items

```go
func UpdateGrocery(ctx context.Context, id, quantity int) error
```

Updates the quantity of a grocery item. Uses transactions to ensure data integrity.

### Deleting Items

```go
func DeleteGrocery(ctx context.Context, id int) error
```

Removes a grocery item from the database.

### Display

```go
func DisplayGroceries(ctx context.Context) error
```

Prints all grocery items to the console in a formatted table.

## Usage Examples

### Connecting to the Database

```go
if err := ConnectDB(); err != nil {
    log.Fatalf("Failed to connect to database: %v", err)
}
defer DB.Close() // Close the database connection when the program exits
```

### Adding Items

```go
// Add a single item
id, err := AddGrocery(ctx, "Sugar", 1.23, 4)
if err != nil {
    log.Printf("Failed to add grocery: %v", err)
}

// Add multiple items in a transaction
err := AddGroceriesInTransaction(ctx, []Grocery{
    {Name: "Milk", Price: 10.0, Quantity: 3},
    {Name: "Bread", Price: 2.35, Quantity: 2},
    {Name: "Cereal", Price: 3.75, Quantity: 7},
})
if err != nil {
    log.Printf("Failed to add groceries in transaction: %v", err)
}
```

### Updating Items

```go
err := UpdateGrocery(ctx, 1, 5) // Update item with ID 1 to quantity 5
if err != nil {
    if errors.Is(err, ErrGroceryNotFound) {
        log.Println("Grocery not found")
    } else {
        log.Printf("Failed to update grocery: %v", err)
    }
}
```

### Retrieving Items

```go
// Get all items
groceries, err := GetAllGroceries(ctx)
if err != nil {
    log.Printf("Failed to get groceries: %v", err)
}

// Get a specific item
grocery, err := GetGroceryByID(ctx, 1)
if err != nil {
    if errors.Is(err, ErrGroceryNotFound) {
        log.Println("Grocery not found")
    } else {
        log.Printf("Failed to get grocery: %v", err)
    }
}
```

### Deleting Items

```go
err := DeleteGrocery(ctx, 1) // Delete item with ID 1
if err != nil {
    if errors.Is(err, ErrGroceryNotFound) {
        log.Println("Grocery not found")
    } else {
        log.Printf("Failed to delete grocery: %v", err)
    }
}
```

## Best Practices

1. **Always use contexts**: Provide contexts with appropriate timeouts for database operations.
2. **Use transactions for multiple operations**: When performing multiple related operations, use transactions to ensure data integrity.
3. **Close resources**: Always close database connections, statements, and rows when done with them.
4. **Handle errors properly**: Check for errors after each database operation and handle them appropriately.
5. **Use prepared statements**: For repeated queries, use prepared statements to improve performance.

## Testing

The project includes comprehensive tests for all database operations. Run the tests with:

```bash
go test -v
```

The tests use an in-memory SQLite database to avoid affecting the actual database file.

## Performance Considerations

- The connection pool is configured with sensible defaults (10 max open connections, 5 max idle connections).
- Transactions are used for batch operations to improve performance.
- Prepared statements are used for repeated queries to reduce parsing overhead.
- Contexts with timeouts prevent operations from hanging indefinitely.
