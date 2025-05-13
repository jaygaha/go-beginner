package main

import "sync"

// Book represents a book in the book store
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  int    `json:"price"`
}

// BookStore represents a book store in memory
type BookStore struct {
	books map[int]Book
	mu    sync.RWMutex // for thread safety access
}

// NewBookStore initializes a new BookStore
func NewBookStore() *BookStore {
	return &BookStore{
		books: make(map[int]Book),
	}
}

// AddBook adds a new book to the store
func (bs *BookStore) AddBook(book Book) {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	bs.books[book.ID] = book
}

// GetBooks returns all books in the store
func (bs *BookStore) GetBooks() []Book {
	bs.mu.RLock()
	defer bs.mu.RUnlock()

	books := make([]Book, 0, len(bs.books))
	for _, book := range bs.books {
		books = append(books, book)
	}

	return books
}

// GetBook retrieves a book by its ID
func (bs *BookStore) GetBook(id int) (Book, bool) {
	bs.mu.RLock()
	defer bs.mu.RUnlock()

	book, ok := bs.books[id]

	return book, ok
}
