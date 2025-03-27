package main

import (
	"context"
	"database/sql"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB() *sql.DB {
	// Use in-memory SQLite database for testing
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	// Configure connection pool settings for testing
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(30 * time.Minute)

	// Create table for testing
	ctx := context.Background()
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
		panic(err)
	}

	return db
}

func TestAddGrocery(t *testing.T) {
	// Setup test database
	DB = setupTestDB()
	defer DB.Close()

	// Create context
	ctx := context.Background()

	// Test adding a grocery item
	id, err := AddGrocery(ctx, "Milk", 10.0, 3)
	if err != nil {
		t.Fatalf("Failed to add grocery: %v", err)
	}

	// Verify the ID is correct
	if id != 1 {
		t.Errorf("Expected ID 1, got %d", id)
	}

	// Verify the item exists in the database
	groceries, err := GetAllGroceries(ctx)
	if err != nil {
		t.Fatalf("Failed to get groceries: %v", err)
	}

	if len(groceries) != 1 {
		t.Errorf("Expected 1 grocery item, got %d", len(groceries))
	}

	// Verify the item properties
	if groceries[0].Name != "Milk" {
		t.Errorf("Expected name 'Milk', got '%s'", groceries[0].Name)
	}
	if groceries[0].Price != 10.0 {
		t.Errorf("Expected price 10.0, got %f", groceries[0].Price)
	}
	if groceries[0].Quantity != 3 {
		t.Errorf("Expected quantity 3, got %d", groceries[0].Quantity)
	}
}

func TestUpdateGrocery(t *testing.T) {
	// Setup test database
	DB = setupTestDB()
	defer DB.Close()

	// Create context
	ctx := context.Background()

	// Add a grocery item to update
	_, err := AddGrocery(ctx, "Milk", 10.0, 3)
	if err != nil {
		t.Fatalf("Failed to add grocery: %v", err)
	}

	// Update the grocery item
	err = UpdateGrocery(ctx, 1, 5)
	if err != nil {
		t.Fatalf("Failed to update grocery: %v", err)
	}

	// Verify the item was updated
	groceries, err := GetAllGroceries(ctx)
	if err != nil {
		t.Fatalf("Failed to get groceries: %v", err)
	}

	if len(groceries) != 1 {
		t.Errorf("Expected 1 grocery item, got %d", len(groceries))
	}

	// Verify the updated quantity
	if groceries[0].Quantity != 5 {
		t.Errorf("Expected quantity 5, got %d", groceries[0].Quantity)
	}

	// Verify updated_at field is set
	if groceries[0].UpdatedAt == nil {
		t.Error("Expected UpdatedAt to be set, but it's nil")
	}
}

func TestDeleteGrocery(t *testing.T) {
	// Setup test database
	DB = setupTestDB()
	defer DB.Close()

	// Create context
	ctx := context.Background()

	// Add a grocery item to delete
	_, err := AddGrocery(ctx, "Milk", 10.0, 3)
	if err != nil {
		t.Fatalf("Failed to add grocery: %v", err)
	}

	// Verify item exists before deletion
	groceries, err := GetAllGroceries(ctx)
	if err != nil {
		t.Fatalf("Failed to get groceries: %v", err)
	}

	if len(groceries) != 1 {
		t.Errorf("Expected 1 grocery item before deletion, got %d", len(groceries))
	}

	// Delete the grocery item
	err = DeleteGrocery(ctx, 1)
	if err != nil {
		t.Fatalf("Failed to delete grocery: %v", err)
	}

	// Verify the item was deleted
	groceries, err = GetAllGroceries(ctx)
	if err != nil {
		t.Fatalf("Failed to get groceries: %v", err)
	}

	if len(groceries) != 0 {
		t.Errorf("Expected 0 grocery items after deletion, got %d", len(groceries))
	}
}

func TestGetAllGroceries(t *testing.T) {
	// Setup test database
	DB = setupTestDB()
	defer DB.Close()

	// Create context
	ctx := context.Background()

	// Add multiple grocery items
	_, err := AddGrocery(ctx, "Milk", 10.0, 3)
	if err != nil {
		t.Fatalf("Failed to add grocery: %v", err)
	}
	_, err = AddGrocery(ctx, "Bread", 2.35, 2)
	if err != nil {
		t.Fatalf("Failed to add grocery: %v", err)
	}
	_, err = AddGrocery(ctx, "Cereal", 3.75, 7)
	if err != nil {
		t.Fatalf("Failed to add grocery: %v", err)
	}

	// Get all grocery items
	groceries, err := GetAllGroceries(ctx)
	if err != nil {
		t.Fatalf("Failed to get groceries: %v", err)
	}

	// Verify the correct number of items
	if len(groceries) != 3 {
		t.Errorf("Expected 3 grocery items, got %d", len(groceries))
	}

	// Verify the items are in the correct order (by ID)
	if groceries[0].Name != "Milk" {
		t.Errorf("Expected first item to be 'Milk', got '%s'", groceries[0].Name)
	}
	if groceries[1].Name != "Bread" {
		t.Errorf("Expected second item to be 'Bread', got '%s'", groceries[1].Name)
	}
	if groceries[2].Name != "Cereal" {
		t.Errorf("Expected third item to be 'Cereal', got '%s'", groceries[2].Name)
	}
}

func TestConnectDB(t *testing.T) {
	// Test the ConnectDB function
	// Save the original DB value
	originalDB := DB

	// Call ConnectDB
	err := ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	// Verify DB is not nil
	if DB == nil {
		t.Error("Expected DB to be set, but it's nil")
	}

	// Close the database connection
	if DB != nil {
		DB.Close()
	}

	// Restore the original DB value
	DB = originalDB
}

func TestDisplayGroceries(t *testing.T) {
	// This is mostly a visual function, so we'll just verify it doesn't error
	// Setup test database
	DB = setupTestDB()
	defer DB.Close()

	// Create context
	ctx := context.Background()

	// Add a grocery item
	_, err := AddGrocery(ctx, "Milk", 10.0, 3)
	if err != nil {
		t.Fatalf("Failed to add grocery: %v", err)
	}

	// Call DisplayGroceries and verify it doesn't error
	err = DisplayGroceries(ctx)
	if err != nil {
		t.Fatalf("DisplayGroceries failed: %v", err)
	}
	// If we get here, the test passes
}
