package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Grocery represents a grocery item in the database
type Grocery struct {
	ID          int        // Primary key
	Name        string     // Name of the grocery item
	Price       float64    // Price of the grocery item
	Quantity    int        // Quantity of the grocery item
	IsAvailable bool       // Availability status
	CreatedAt   time.Time  // Time when the grocery item was created
	UpdatedAt   *time.Time // Time when the grocery item was last updated (can be nil)
}

// DB is the global database connection pool
var DB *sql.DB

// ErrGroceryNotFound is returned when a grocery item is not found
var ErrGroceryNotFound = errors.New("grocery not found")

func main() {
	// Connect to the database
	if err := ConnectDB(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer DB.Close() // Close the database connection when the program exits

	// Create a context with timeout for database operations
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Demonstrate transaction usage for adding multiple items
	if err := AddGroceriesInTransaction(ctx, []Grocery{
		{Name: "Milk", Price: 10.0, Quantity: 3},
		{Name: "Bread", Price: 2.35, Quantity: 2},
		{Name: "Cereal", Price: 3.75, Quantity: 7},
	}); err != nil {
		log.Printf("Failed to add groceries in transaction: %v", err)
	}

	// Add a single grocery item
	if _, err := AddGrocery(ctx, "Sugar", 1.23, 4); err != nil {
		log.Printf("Failed to add grocery: %v", err)
	}

	// Update a grocery item after consumption
	if err := UpdateGrocery(ctx, 1, 2); err != nil {
		log.Printf("Failed to update grocery: %v", err)
	}

	// Delete a grocery item
	if err := DeleteGrocery(ctx, 1); err != nil {
		log.Printf("Failed to delete grocery: %v", err)
	}

	// Display all grocery items
	if err := DisplayGroceries(ctx); err != nil {
		log.Printf("Failed to display groceries: %v", err)
	}
}

// ConnectDB establishes a connection to the SQLite database and sets up the schema
func ConnectDB() error {
	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite3", "./grocery.db")
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	// Configure connection pool settings
	db.SetMaxOpenConns(10)           // Maximum number of open connections
	db.SetMaxIdleConns(5)            // Maximum number of idle connections
	db.SetConnMaxLifetime(time.Hour) // Maximum lifetime of a connection

	// Verify the connection is working
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	// Create table if it doesn't exist
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS groceries (
		id INTEGER PRIMARY KEY,
		name VARCHAR(255),
		price REAL,
		quantity INTEGER DEFAULT 0,
		is_available BOOLEAN DEFAULT false,
		created_at DATETIME,
		updated_at DATETIME DEFAULT NULL
	)`)

	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	DB = db
	return nil
}

// AddGrocery adds a new grocery item to the database
func AddGrocery(ctx context.Context, name string, price float64, quantity int) (int64, error) {
	// Insert a row into the table
	result, err := DB.ExecContext(
		ctx,
		`INSERT INTO groceries (name, price, quantity, created_at) VALUES (?, ?, ?, ?)`,
		name, price, quantity, time.Now(),
	)
	if err != nil {
		return 0, fmt.Errorf("failed to add grocery: %w", err)
	}

	// Get the ID of the inserted row
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert ID: %w", err)
	}

	return id, nil
}

// AddGroceriesInTransaction adds multiple grocery items in a single transaction
func AddGroceriesInTransaction(ctx context.Context, groceries []Grocery) error {
	// Begin a transaction
	tx, err := DB.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Prepare the statement once for efficiency
	stmt, err := tx.PrepareContext(
		ctx,
		`INSERT INTO groceries (name, price, quantity, created_at) VALUES (?, ?, ?, ?)`,
	)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	// Execute the statement for each grocery item
	for _, g := range groceries {
		_, err := stmt.ExecContext(ctx, g.Name, g.Price, g.Quantity, time.Now())
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to add grocery '%s': %w", g.Name, err)
		}
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// GetAllGroceries retrieves all grocery items from the database
func GetAllGroceries(ctx context.Context) ([]Grocery, error) {
	// Select all rows from the table
	rows, err := DB.QueryContext(ctx, `SELECT * FROM groceries`)
	if err != nil {
		return nil, fmt.Errorf("failed to query groceries: %w", err)
	}
	defer rows.Close() // Close the rows when the function exits

	var groceries []Grocery

	// Iterate over the rows
	for rows.Next() {
		var grocery Grocery
		// Scan the row into the grocery struct
		err := rows.Scan(
			&grocery.ID,
			&grocery.Name,
			&grocery.Price,
			&grocery.Quantity,
			&grocery.IsAvailable,
			&grocery.CreatedAt,
			&grocery.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan grocery: %w", err)
		}

		groceries = append(groceries, grocery)
	}

	// Check for errors after iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return groceries, nil
}

// GetGroceryByID retrieves a specific grocery item by ID
func GetGroceryByID(ctx context.Context, id int) (Grocery, error) {
	var grocery Grocery

	// Query the database for the grocery item
	row := DB.QueryRowContext(ctx, `SELECT * FROM groceries WHERE id = ?`, id)

	// Scan the row into the grocery struct
	err := row.Scan(
		&grocery.ID,
		&grocery.Name,
		&grocery.Price,
		&grocery.Quantity,
		&grocery.IsAvailable,
		&grocery.CreatedAt,
		&grocery.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Grocery{}, ErrGroceryNotFound
		}
		return Grocery{}, fmt.Errorf("failed to get grocery: %w", err)
	}

	return grocery, nil
}

// UpdateGrocery updates the quantity of a grocery item
func UpdateGrocery(ctx context.Context, id, quantity int) error {
	// Begin a transaction
	tx, err := DB.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Check if the grocery item exists
	var exists bool
	err = tx.QueryRowContext(ctx, `SELECT EXISTS(SELECT 1 FROM groceries WHERE id = ?)`, id).Scan(&exists)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to check if grocery exists: %w", err)
	}

	if !exists {
		tx.Rollback()
		return ErrGroceryNotFound
	}

	// Update the grocery item
	result, err := tx.ExecContext(
		ctx,
		`UPDATE groceries SET quantity = ?, updated_at = ? WHERE id = ?`,
		quantity, time.Now(), id,
	)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update grocery: %w", err)
	}

	// Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		tx.Rollback()
		return ErrGroceryNotFound
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// DeleteGrocery removes a grocery item from the database
func DeleteGrocery(ctx context.Context, id int) error {
	// Delete the row from the table
	result, err := DB.ExecContext(ctx, `DELETE FROM groceries WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("failed to delete grocery: %w", err)
	}

	// Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return ErrGroceryNotFound
	}

	return nil
}

// DisplayGroceries prints all grocery items to the console
func DisplayGroceries(ctx context.Context) error {
	groceries, err := GetAllGroceries(ctx)
	if err != nil {
		return fmt.Errorf("failed to get groceries: %w", err)
	}

	if len(groceries) == 0 {
		fmt.Println("No grocery items found")
		return nil
	}

	// Print header
	fmt.Printf("%-5s %-15s %-10s %-10s %-12s %-25s %-25s\n", "ID", "Name", "Price", "Quantity", "Available", "Created At", "Updated At")
	fmt.Println("--------------------------------------------------------------------------------------------------------")

	// Print each grocery item
	for _, grocery := range groceries {
		updatedAt := ""
		if grocery.UpdatedAt != nil {
			updatedAt = grocery.UpdatedAt.Format("2006-01-02 15:04:05")
		}

		fmt.Printf("%-5d %-15s $%-9.2f %-10d %-12t %-25s %-25s\n",
			grocery.ID,
			grocery.Name,
			grocery.Price,
			grocery.Quantity,
			grocery.IsAvailable,
			grocery.CreatedAt.Format("2025-03-27 15:04:05"), // Format created_at
			updatedAt,
		)
	}

	return nil
}
