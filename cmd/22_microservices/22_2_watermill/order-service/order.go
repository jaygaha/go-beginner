package order_service

import (
	"fmt"
	"sync"
)

// Book represents a book from the book service event.
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// Order represents an order in the bookstore.
type Order struct {
	ID       string `json:"id"`
	BookID   string `json:"book_id"`
	Quantity int    `json:"quantity"`
}

// OrderStore manages orders and known books.
type OrderStore struct {
	orders  map[string]Order
	bookIds map[string]bool
	mu      sync.RWMutex
}

// NewOrderStore init a new OrderStore.
func NewOrderStore() *OrderStore {
	return &OrderStore{
		orders:  make(map[string]Order),
		bookIds: make(map[string]bool),
	}
}

// AddOrder adds a new order to the store.
func (s *OrderStore) AddOrder(order Order) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.orders[order.ID] = order
}

// AddBookId adds a book ID to the known books list.
func (s *OrderStore) AddBookId(bookId string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.bookIds[bookId] = true
}

// BookExists checks if a book ID exists in the known books list.
func (s *OrderStore) BookExists(bookId string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	fmt.Printf("Books %v\n", s.bookIds)
	return s.bookIds[bookId]
}
