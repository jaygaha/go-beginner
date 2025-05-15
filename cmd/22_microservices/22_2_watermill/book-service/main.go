package book_service

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

// Shared in-memory Pub/Sub for both services
var pubSub *gochannel.GoChannel

// SetPubSub sets the shared pubSub instance from the main program
func SetPubSub(ps *gochannel.GoChannel) {
	pubSub = ps
}

// Capital function name to export Main function for bookService data
func Main() {

	// Start the Book Service
	store := NewBookStore()

	// HTTP handler for adding a book
	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		handleAddBook(w, r, store, pubSub)
	})

	// Start the HTTP server
	log.Println("Book Service is running on :8801")
	log.Fatal(http.ListenAndServe(":8801", nil))
}

func handleAddBook(w http.ResponseWriter, r *http.Request, store *BookStore, publisher message.Publisher) {
	var book Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// generate unique ID for the book
	book.ID = fmt.Sprintf("B%d", rand.Intn(1000))

	// add the book to the store
	store.AddBook(book)

	// publish the event
	bookJson, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Failed to marshal book", http.StatusInternalServerError)
		return
	}

	msg := message.NewMessage(watermill.NewUUID(), bookJson)

	if err := publisher.Publish("books.created", msg); err != nil {
		log.Printf("Failed to publish event: %v", err)
		http.Error(w, "Failed to publish event", http.StatusInternalServerError)
		return
	}

	// respond with the added book
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}
