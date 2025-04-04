# MySQL with Go: Beginner's Guide

MySQL is one of the most popular open-source relational database management systems. This guide will help you understand how to use MySQL with Go, covering basic to intermediate concepts.

We are learning this guide using TDD (Test-Driven Development) approach. Please check [test file](./main_test.go)

## Prerequisites

1. **Install MySQL Driver**: You need to install the MySQL driver for Go:
   
   ```bash
   go get github.com/go-sql-driver/mysql
   ```

2. **Import the Driver**: In your Go code, import the MySQL driver:
   
   ```go
   import (
       "database/sql"
       _ "github.com/go-sql-driver/mysql"
   )
   ```
   
   The underscore (`_`) before the import means we're importing the package for its side effects (registering the MySQL driver) without directly using its exported names.

## Connecting to MySQL Database

### Basic Connection

```go
db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/dbname")
if err != nil {
    panic(err)
}
defer db.Close()

// Test the connection
err = db.Ping()
if err != nil {
    panic(err)
}
```

### Connection String Format

The connection string follows this format:
```
username:password@protocol(host:port)/dbname?param=value
```

### Connection Parameters

You can add parameters to the connection string:

```go
db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/dbname?parseTime=true&interpolateParams=true")
```

Common parameters:
- `parseTime=true`: Automatically converts DATETIME fields to Go's `time.Time`
- `interpolateParams=true`: Interpolates query parameters client-side
- `timeout=30s`: Connection timeout
- `readTimeout=30s`: Read timeout
- `writeTimeout=30s`: Write timeout

## Basic Database Operations

### Creating Tables (DDL Statements)

DDL statements are used to create, alter, and drop tables.

```go
// Create a table
_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255))")
if err != nil {
    panic(err)
}

// Alter a table
_, err = db.Exec("ALTER TABLE users ADD COLUMN age INT")
if err != nil {
    panic(err)
}

// Drop a table
_, err = db.Exec("DROP TABLE IF EXISTS users")
if err != nil {
    panic(err)
}
```

### Data Manipulation (DML Statements)

DML statements are used to insert, update, and delete data.

#### Inserting Data

```go
// Insert a record
result, err := db.Exec("INSERT INTO users (name) VALUES (?)", "John")
if err != nil {
    panic(err)
}

// Get the number of affected rows
rowsAffected, err := result.RowsAffected()
if err != nil {
    panic(err)
}

// Get the last inserted ID
lastID, err := result.LastInsertId()
if err != nil {
    panic(err)
}
```

#### Querying Data

```go
// Query a single row
row := db.QueryRow("SELECT id, name FROM users WHERE id = ?", 1)

var id int
var name string
err = row.Scan(&id, &name)
if err != nil {
    panic(err)
}

// Query multiple rows
rows, err := db.Query("SELECT id, name FROM users WHERE id > ? ORDER BY id", 0)
if err != nil {
    panic(err)
}
defer rows.Close()

for rows.Next() {
    var id int
    var name string
    err = rows.Scan(&id, &name)
    if err != nil {
        panic(err)
    }
    fmt.Printf("ID: %d, Name: %s\n", id, name)
}

if err = rows.Err(); err != nil {
    panic(err)
}
```

#### Updating Data

```go
result, err := db.Exec("UPDATE users SET name = ? WHERE id = ?", "Jane", 1)
if err != nil {
    panic(err)
}

rowsAffected, err := result.RowsAffected()
if err != nil {
    panic(err)
}
```

#### Deleting Data

```go
result, err := db.Exec("DELETE FROM users WHERE id = ?", 1)
if err != nil {
    panic(err)
}

rowsAffected, err := result.RowsAffected()
if err != nil {
    panic(err)
}
```

## Using Prepared Statements

Prepared statements improve performance and security by separating SQL logic from data:

```go
// Create a prepared statement
stmt, err := db.Prepare("INSERT INTO users (name) VALUES (?)")
if err != nil {
    panic(err)
}
defer stmt.Close()

// Execute the statement multiple times with different values
result, err := stmt.Exec("John")
if err != nil {
    panic(err)
}

result, err = stmt.Exec("Jane")
if err != nil {
    panic(err)
}

// Prepared statements for queries
stmt, err = db.Prepare("SELECT id, name FROM users WHERE id = ?")
if err != nil {
    panic(err)
}
defer stmt.Close()

row := stmt.QueryRow(1)
var id int
var name string
err = row.Scan(&id, &name)
if err != nil {
    panic(err)
}
```

## Transactions

Transactions ensure that a group of operations either all succeed or all fail:

```go
// Start a transaction
tx, err := db.Begin()
if err != nil {
    panic(err)
}

// Perform operations within the transaction
_, err = tx.Exec("INSERT INTO users (name) VALUES (?)", "John")
if err != nil {
    tx.Rollback() // Roll back on error
    panic(err)
}

_, err = tx.Exec("INSERT INTO users (name) VALUES (?)", "Jane")
if err != nil {
    tx.Rollback() // Roll back on error
    panic(err)
}

// Commit the transaction if all operations succeeded
err = tx.Commit()
if err != nil {
    panic(err)
}
```

## Using Context with Database Operations

Context allows you to control timeouts and cancellation:

```go
// Create a context with a timeout
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// Use the context with a transaction
tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
if err != nil {
    panic(err)
}

// Execute queries with context
row := tx.QueryRowContext(ctx, "SELECT COUNT(*) FROM users")
var count int
err = row.Scan(&count)
if err != nil {
    tx.Rollback()
    panic(err)
}

// Commit with context awareness
err = tx.Commit()
if err != nil {
    panic(err)
}
```

## Connection Pooling

Go's `database/sql` package automatically manages a connection pool:

```go
// Configure the connection pool
db.SetMaxOpenConns(100)    // Maximum number of open connections
db.SetMaxIdleConns(10)     // Maximum number of idle connections
db.SetConnMaxLifetime(time.Hour) // Maximum lifetime of a connection
```

### Using a Single Connection from the Pool

```go
ctx := context.Background()

// Get a connection from the pool
conn, err := db.Conn(ctx)
if err != nil {
    panic(err)
}
defer conn.Close() // Return the connection to the pool

// Use the connection
result, err := conn.ExecContext(ctx, "INSERT INTO users (name) VALUES (?)", "John")
if err != nil {
    panic(err)
}
```

## Best Practices

1. **Always close resources**: Use `defer` to close rows, statements, and connections.
2. **Use prepared statements**: They improve security and performance.
3. **Handle errors properly**: Check errors after every database operation.
4. **Use transactions** for operations that must succeed or fail together.
5. **Configure connection pool** appropriately for your application's needs.
6. **Use context** for timeout and cancellation control.
7. **Validate input data** before sending it to the database.
8. **Use parameter placeholders** (`?`) instead of string concatenation to prevent SQL injection.

## Common Errors and Solutions

1. **Connection refused**: Check if MySQL server is running and accessible.
2. **Access denied**: Verify username, password, and permissions.
3. **Unknown database**: Ensure the database exists.
4. **Too many connections**: Adjust connection pool settings or MySQL's `max_connections`.
5. **Query timeout**: Optimize queries or adjust timeout settings.

MySQL with Go provides a powerful combination for building robust database applications. This guide covers the basics to get you started, but there's much more to explore as you build more complex applications.

## Further Reading
- [SQL package documentation](https://pkg.go.dev/database/sql)
- [MySQL driver](https://github.com/go-sql-driver/mysql)