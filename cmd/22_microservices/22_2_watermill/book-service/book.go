package book_service

import "sync"

// Book represents a book in the bookstore.
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// BookStore represents a bookstore with a collection of books in memory.
type BookStore struct {
	books map[string]Book
	mu    sync.RWMutex
}

// NewBookStore creates a new BookStore.
func NewBookStore() *BookStore {
	return &BookStore{
		books: make(map[string]Book),
	}
}

// AddBook adds a new book to the bookstore.
func (bs *BookStore) AddBook(book Book) {
	bs.mu.Lock()
	defer bs.mu.Unlock()

	bs.books[book.ID] = book
}
