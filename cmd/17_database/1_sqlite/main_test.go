package main

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB() *sql.DB {
	// Use in-memory SQLite database for testing
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	// Create table for testing
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

	if err != nil {
		panic(err)
	}

	return db
}

func TestConnectDB(t *testing.T) {
	// Save the original DB value
	originalDB := DB

	// Call ConnectDB
	ConnectDB()

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

func TestAddGrocery(t *testing.T) {
	// Setup test database
	DB = setupTestDB()
	defer DB.Close()

	// Test adding a grocery item
	result := AddGrocery("Milk", 10.0, 3)

	// Verify the result is not nil
	if result == nil {
		t.Fatal("Expected non-nil result from AddGrocery")
	}

	// Get the ID of the inserted row
	id, err := result.LastInsertId()
	if err != nil {
		t.Fatalf("Failed to get last insert ID: %v", err)
	}

	// Verify the ID is correct
	if id != 1 {
		t.Errorf("Expected ID 1, got %d", id)
	}

	// Verify the item exists in the database
	groceries := GetAllGroceries()

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

	// Add a grocery item to update
	result := AddGrocery("Milk", 10.0, 3)

	// Get the ID of the inserted row
	id, err := result.LastInsertId()
	if err != nil {
		t.Fatalf("Failed to get last insert ID: %v", err)
	}

	// Update the grocery item using the actual ID
	UpdateGrocery(int(id), 5)

	// Verify the item was updated
	groceries := GetAllGroceries()

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

	// Add a grocery item to delete
	result := AddGrocery("Milk", 10.0, 3)

	// Get the ID of the inserted row
	id, err := result.LastInsertId()
	if err != nil {
		t.Fatalf("Failed to get last insert ID: %v", err)
	}

	// Verify item exists before deletion
	groceries := GetAllGroceries()

	if len(groceries) != 1 {
		t.Errorf("Expected 1 grocery item before deletion, got %d", len(groceries))
	}

	// Delete the grocery item using the actual ID
	DeleteGrocery(int(id))

	// Verify the item was deleted
	groceries = GetAllGroceries()

	if len(groceries) != 0 {
		t.Errorf("Expected 0 grocery items after deletion, got %d", len(groceries))
	}
}

func TestGetAllGroceries(t *testing.T) {
	// Setup test database
	DB = setupTestDB()
	defer DB.Close()

	// Add multiple grocery items
	AddGrocery("Milk", 10.0, 3)
	AddGrocery("Bread", 2.35, 2)
	AddGrocery("Cereal", 3.75, 7)

	// Get all grocery items
	groceries := GetAllGroceries()

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

func TestDisplayGroceries(t *testing.T) {
	// This is mostly a visual function, so we'll just verify it doesn't panic
	// Setup test database
	DB = setupTestDB()
	defer DB.Close()

	// Add a grocery item
	AddGrocery("Milk", 10.0, 3)

	// Call DisplayGroceries and verify it doesn't panic
	DisplayGroceries()
	// If we get here, the test passes
}
