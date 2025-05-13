package main

import (
	"math/rand"
	"sync"
)

// Order represents an order in the book store
type Order struct {
	ID        int    `json:"id"`
	BookID    int    `json:"book_id"`
	Quantity  int    `json:"quantity"`
	TotalCost int    `json:"total_cost"`
	Customer  string `json:"customer"`
}

// Book represents a book as received from the Book Service
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  int    `json:"price"`
}

// OrderStore represents an order store in memory
type OrderStore struct {
	orders map[int]Order
	mu     sync.RWMutex // for thread safety access
}

// NewOrderStore initializes a new OrderStore
func NewOrderStore() *OrderStore {
	return &OrderStore{
		orders: make(map[int]Order),
	}
}

// AddOrder adds a new order to the store
func (os *OrderStore) AddOrder(order Order) {
	os.mu.Lock()
	defer os.mu.Unlock()

	order.ID = rand.Intn(1000) // generate a random ID for the order
	os.orders[order.ID] = order
}

// GetOrders retrieves all orders from the store
func (os *OrderStore) GetOrders() []Order {
	os.mu.RLock()
	defer os.mu.RUnlock()

	orders := make([]Order, 0, len(os.orders))
	for _, order := range os.orders {
		orders = append(orders, order)
	}

	return orders
}

// GetOrdersByCustomer retrieves all orders for a specific customer
func (os *OrderStore) GetOrdersByCustomer(customer string) []Order {
	os.mu.RLock()
	defer os.mu.RUnlock()

	orders := make([]Order, 0)
	for _, order := range os.orders {
		if order.Customer == customer {
			orders = append(orders, order)
		}
	}

	return orders
}
