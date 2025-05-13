package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	store := NewBookStore() // create a new book store

	// routes
	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handleGetBooks(w, store)
		case http.MethodPost:
			handleAddBook(w, r, store)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// route for getting a book by ID
	http.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handleGetBookById(w, r, store)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Book service is running on port 8801...")
	// start the server
	if err := http.ListenAndServe(":8801", nil); err != nil {
		log.Fatal(err)
	}
}

func handleGetBooks(w http.ResponseWriter, store *BookStore) {
	books := store.GetBooks() // get all books from the store

	w.Header().Set("Content-Type", "application/json") // set response header
	json.NewEncoder(w).Encode(books)                   // encode books as JSON and write to response
}

func handleAddBook(w http.ResponseWriter, r *http.Request, store *BookStore) {
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil { // decode request body to book
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// generate a random int ID
	randID := rand.Intn(1000)

	book.ID = randID
	store.AddBook(book) // add book to the store

	w.Header().Set("Content-Type", "application/json") // set response header
	w.WriteHeader(http.StatusCreated)                  // set response status code
	json.NewEncoder(w).Encode(book)                    // encode book as JSON and write to response
}

func handleGetBookById(w http.ResponseWriter, r *http.Request, store *BookStore) {
	// extract book ID from request URL path eg. /books/1
	id := strings.TrimPrefix(r.URL.Path, "/books/") // get book ID from request URL path
	if id == "" {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	// convert book ID to int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	book, exists := store.GetBook(idInt) // get book from store by ID
	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json") // set response header
	json.NewEncoder(w).Encode(book)                    // encode book as JSON and write to response
}
