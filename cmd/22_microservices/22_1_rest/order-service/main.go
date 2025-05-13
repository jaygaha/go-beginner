package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

// book-service url
const bookServiceURL = "http://localhost:8801/books"

func main() {
	store := NewOrderStore()
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handleCreateOrder(w, r, store)
			return
		}

		if r.Method == http.MethodGet {
			handleGetOrders(w, r, store)
			return
		}

		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	fmt.Println("Order Service is running on :8802")
	if err := http.ListenAndServe(":8802", nil); err != nil {
		log.Fatal(err)
	}
}

func handleGetOrders(w http.ResponseWriter, r *http.Request, store *OrderStore) {
	customer := r.URL.Query().Get("customer")

	var orders []Order
	if customer != "" {
		// If customer parameter is provided, filter orders by customer
		orders = store.GetOrdersByCustomer(customer)
	} else {
		// Otherwise, return all orders
		orders = store.GetOrders()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func handleCreateOrder(w http.ResponseWriter, r *http.Request, store *OrderStore) {
	var order Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate book exists by querying Book Service for specific book ID
	resp, err := http.Get(fmt.Sprintf("%s/%d", bookServiceURL, order.BookID))
	if err != nil {
		http.Error(w, "Failed to contact Book Service", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	var books Book
	if err := json.NewDecoder(resp.Body).Decode(&books); err != nil {
		http.Error(w, "Failed to decode Book Service response", http.StatusInternalServerError)
		return
	}

	// create random order ID
	order.ID = rand.Intn(1000)

	// Order ID will be generated in the AddOrder method
	store.AddOrder(order)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}
