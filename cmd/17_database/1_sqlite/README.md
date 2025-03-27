# SQLite with Go: Beginner's Guide

SQLite is a lightweight, disk-based database that doesn't require a separate server process. It's perfect for small to medium-sized applications, embedded systems, and for learning purposes.

## Setting Up SQLite in Go

1. **Install SQLite Driver**: Ensure you have the SQLite driver for Go installed. You can do this using:
   
   ```bash
   go get github.com/mattn/go-sqlite3
   ```

2. **Connect to SQLite Database**:
   
   - Use the `sql.Open` function to connect to your SQLite database.
   - Example:
     ```go
     db, err := sql.Open("sqlite3", "./grocery.db")
     if err != nil {
         panic(err)
     }
     defer db.Close()
     ```

## Basic Operations

### Creating Tables

Use SQL commands to create tables. Example:

```go
statement, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS groceries (
    id INTEGER PRIMARY KEY,
    name VARCHAR(255),
    price REAL,
    quantity INTEGER DEFAULT 0,
    is_available BOOLEAN DEFAULT false,
    created_at DATETIME,
    updated_at DATETIME DEFAULT NULL
)`) 
_, err = statement.Exec()
```

### Inserting Records

Insert records using SQL `INSERT` command:

```go
statement, _ := db.Prepare(`INSERT INTO groceries (name, price, quantity, created_at) VALUES (?, ?, ?, ?)`) 
_, err := statement.Exec("Milk", 10.0, 3, time.Now())
```

### Updating Records

Update records using SQL `UPDATE` command:

```go
statement, _ := db.Prepare(`UPDATE groceries SET quantity =?, updated_at =? WHERE id =?`) 
_, err = statement.Exec(2, time.Now(), 1)
```

### Deleting Records

Delete records using SQL `DELETE` command:

```go
statement, _ := db.Prepare(`DELETE FROM groceries WHERE id =?`) 
_, err := statement.Exec(1)
```

## Advantages of SQLite

- **Lightweight**: No server required, easy to set up.
- **Cross-Platform**: Works on various operating systems.
- **ACID Compliance**: Ensures reliable transactions.

## Disadvantages of SQLite
- **Limited Concurrency**: Not suitable for high-write applications.
- **Scalability**: Better suited for smaller databases.

SQLite is a great choice for beginners due to its simplicity and ease of use. It's ideal for learning database concepts and for applications where a full-fledged database server is not necessary.

## Testing

The project includes comprehensive tests for all database operations. Run the tests with:

```bash
go test -v
```