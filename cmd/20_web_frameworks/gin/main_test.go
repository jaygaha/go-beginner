package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// Setup test router with all the routes defined in main
func setupRouter() *gin.Engine {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create a new router instance
	router := gin.Default()

	// Register all routes
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.GET("/books/:id", getBookById)
	router.PUT("/books/:id", updateBook)
	router.DELETE("/books/:id", deleteBook)
	router.PATCH("/books/borrow/:id", borrowBook)
	router.PATCH("/books/return/:id", returnBook)

	return router
}

// Helper function to reset books slice to initial state before each test
func resetBooks() {
	books = []book{
		{ID: uuid.NewString(), Isbn: "0395489318", Title: "The Lord of the Rings", Author: "J.R.R. Tolkien", Quantity: 10, CreatedAt: time.Now()},
		{ID: uuid.NewString(), Isbn: "0590353403", Title: "harry Potter and the Sorcerer's Stone", Author: "J.K. Rowling", Quantity: 20, CreatedAt: time.Now()},
		{ID: uuid.NewString(), Isbn: "0395489319", Title: "The Hobbit", Author: "J.R.R. Tolkien", Quantity: 5, CreatedAt: time.Now()},
		{ID: uuid.NewString(), Isbn: "9783540315490", Title: "Marine biotechnology /", Author: "Le Gal, Yves", Quantity: 2, CreatedAt: time.Now()},
	}
}

// Test getting all books
func TestGetBooks(t *testing.T) {
	// Reset books to initial state
	resetBooks()

	// Setup router
	router := setupRouter()

	// Create a new HTTP request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books", nil)

	// Serve the HTTP request
	router.ServeHTTP(w, req)

	// Assert HTTP status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response body
	var response []book
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Assert no error in parsing
	assert.Nil(t, err)

	// Assert response contains 4 books
	assert.Len(t, response, 4)
}

// Test creating a new book
func TestCreateBook(t *testing.T) {
	// Reset books to initial state
	resetBooks()

	// Setup router
	router := setupRouter()

	// Create a new book payload
	newBook := map[string]interface{}{
		"isbn":     "9781234567890",
		"title":    "Test Book",
		"author":   "Test Author",
		"quantity": 5,
	}

	// Convert payload to JSON
	payload, _ := json.Marshal(newBook)

	// Create a new HTTP request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	// Serve the HTTP request
	router.ServeHTTP(w, req)

	// Assert HTTP status code
	assert.Equal(t, http.StatusCreated, w.Code)

	// Parse response body
	var response book
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Assert no error in parsing
	assert.Nil(t, err)

	// Assert book properties
	assert.Equal(t, "9781234567890", response.Isbn)
	assert.Equal(t, "Test Book", response.Title)
	assert.Equal(t, "Test Author", response.Author)
	assert.Equal(t, 5, response.Quantity)
	assert.NotEmpty(t, response.ID)
}

// Test getting a book by ID
func TestGetBookById(t *testing.T) {
	// Reset books to initial state
	resetBooks()

	// Get the ID of the first book
	bookID := books[0].ID

	// Setup router
	router := setupRouter()

	// Create a new HTTP request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books/"+bookID, nil)

	// Serve the HTTP request
	router.ServeHTTP(w, req)

	// Assert HTTP status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response body
	var response book
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Assert no error in parsing
	assert.Nil(t, err)

	// Assert book ID matches
	assert.Equal(t, bookID, response.ID)
}

// Test borrowing a book
func TestBorrowBook(t *testing.T) {
	// Reset books to initial state
	resetBooks()

	// Get the ID of the first book
	bookID := books[0].ID
	initialQuantity := books[0].Quantity

	// Setup router
	router := setupRouter()

	// Create a new HTTP request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/books/borrow/"+bookID, nil)

	// Serve the HTTP request
	router.ServeHTTP(w, req)

	// Assert HTTP status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response body
	var response book
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Assert no error in parsing
	assert.Nil(t, err)

	// Assert quantity decreased by 1
	assert.Equal(t, initialQuantity-1, response.Quantity)
}

// Test returning a book
func TestReturnBook(t *testing.T) {
	// Reset books to initial state
	resetBooks()

	// Get the ID of the first book
	bookID := books[0].ID
	initialQuantity := books[0].Quantity

	// Setup router
	router := setupRouter()

	// Create a new HTTP request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/books/return/"+bookID, nil)

	// Serve the HTTP request
	router.ServeHTTP(w, req)

	// Assert HTTP status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response body
	var response book
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Assert no error in parsing
	assert.Nil(t, err)

	// Assert quantity increased by 1
	assert.Equal(t, initialQuantity+1, response.Quantity)
}

// Test updating a book
func TestUpdateBook(t *testing.T) {
	// Reset books to initial state
	resetBooks()

	// Get the ID of the first book
	bookID := books[0].ID

	// Setup router
	router := setupRouter()

	// Create updated book payload
	updatedBook := map[string]interface{}{
		"id":       bookID,
		"isbn":     "0395489318",
		"title":    "The Lord of the Rings: Updated Edition",
		"author":   "J.R.R. Tolkien",
		"quantity": 15,
	}

	// Convert payload to JSON
	payload, _ := json.Marshal(updatedBook)

	// Create a new HTTP request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/books/"+bookID, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	// Serve the HTTP request
	router.ServeHTTP(w, req)

	// Assert HTTP status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response body
	var response book
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Assert no error in parsing
	assert.Nil(t, err)

	// Assert book ID matches
	assert.Equal(t, bookID, response.ID)
}

// Test deleting a book
func TestDeleteBook(t *testing.T) {
	// Reset books to initial state
	resetBooks()

	// Get the ID of the first book
	bookID := books[0].ID

	// Setup router
	router := setupRouter()

	// Create a new HTTP request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/books/"+bookID, nil)

	// Serve the HTTP request
	router.ServeHTTP(w, req)

	// Assert HTTP status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response body
	var response book
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Assert no error in parsing
	assert.Nil(t, err)

	// Assert book ID matches
	assert.Equal(t, bookID, response.ID)

	// Assert DeletedAt is not nil
	assert.NotNil(t, response.DeletedAt)
}
