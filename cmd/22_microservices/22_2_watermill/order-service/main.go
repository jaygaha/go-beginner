package order_service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

// External references to shared in-memory Pub/Sub and BookStore
var pubSub *gochannel.GoChannel

// SetPubSub sets the shared pubSub instance from the main program
func SetPubSub(ps *gochannel.GoChannel) {
	pubSub = ps
}

func Main() {
	// start the order service
	store := NewOrderStore()

	// subscribe to BookCreated events
	bookCreatedSub, err := pubSub.Subscribe(context.Background(), "books.created")
	if err != nil {
		log.Fatalf("failed to subscribe to book_created: %v", err)
	}
	// start a goroutine to handle BookCreated events
	go processBookCreatedEvents(bookCreatedSub, store)

	// HTTP handler for creating orders
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		handleCreatedOrder(w, r, store)
	})

	log.Println("Order service listening on :8802")
	log.Fatal(http.ListenAndServe(":8802", nil))
}

func processBookCreatedEvents(messages <-chan *message.Message, store *OrderStore) {
	for msg := range messages {
		var book Book
		if err := json.Unmarshal(msg.Payload, &book); err != nil {
			log.Printf("failed to unmarshal book_created event: %v", err)
			msg.Nack() // nack the message to retry later
			continue
		}

		// add the book to the store
		store.AddBookId(book.ID)

		msg.Ack() // ack the message after processing
	}
}

func handleCreatedOrder(w http.ResponseWriter, r *http.Request, store *OrderStore) {
	var order Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// check if the book exists in the store
	if !store.BookExists(order.BookID) {
		http.Error(w, "book not found", http.StatusNotFound)
		return
	}

	// generate a unique order ID
	order.ID = fmt.Sprintf("O%d", rand.Intn(1000))

	// store the order
	store.AddOrder(order)

	// respond with the created order
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}
