/*
	Gin
	-> lightweight web framework, high performance, high productivity
	-> provides a simple and efficient way to build web applications
	-> it supports routing, middleware, JSON validation, and more
	-> it is based on the net/http package, which provides a simple and efficient way to build web applications
	-> it is fast and efficient, and it is easy to use
	-> it is open source and free to use
	-> ideal for creating scalable, efficient, and maintainable web applications

	Simple book management API to demonstrate the usage of Gin
*/

package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type book struct {
	ID        string     `json:"id"` // UUID capitalized means public
	Isbn      string     `json:"isbn"`
	Title     string     `json:"title"`
	Author    string     `json:"author"`
	Quantity  int        `json:"quantity"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"` // nullable
	DeletedAt *time.Time `json:"deleted_at"` // nullable
}

// default book list
var books = []book{
	{ID: uuid.NewString(), Isbn: "0395489318", Title: "The Lord of the Rings", Author: "J.R.R. Tolkien", Quantity: 10, CreatedAt: time.Now()},
	{ID: uuid.NewString(), Isbn: "0590353403", Title: "harry Potter and the Sorcerer's Stone", Author: "J.K. Rowling", Quantity: 20, CreatedAt: time.Now()},
	{ID: uuid.NewString(), Isbn: "0395489319", Title: "The Hobbit", Author: "J.R.R. Tolkien", Quantity: 5, CreatedAt: time.Now()},
	{ID: uuid.NewString(), Isbn: "9783540315490", Title: "Marine biotechnology /", Author: "Le Gal, Yves", Quantity: 2, CreatedAt: time.Now()},
}

// get all books
// gin.Context is the context of the request and response
func getBooks(c *gin.Context) {
	if len(books) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No books found"}) // return a 404 status code
		return
	}

	c.IndentedJSON(http.StatusOK, books) // return the books list in JSON format
}

func createBook(c *gin.Context) {
	var newBook book

	// BindJSON binds the received JSON to newBook
	// it also handles errors and invalid data
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	// generate a new UUID for the book
	newBook.ID = uuid.NewString()
	// set the created_at field to the current time
	newBook.CreatedAt = time.Now()

	// add the book to the list
	books = append(books, newBook)

	c.IndentedJSON(http.StatusCreated, newBook) // return the new book in JSON format
}

// get a book by ID
func bookByID(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

// get a book by ID
func getBookById(c *gin.Context) {
	id := c.Param("id") // get the id parameter from the URL; :id is a placeholder for the id parameter
	// id, ok := c.GetQuery("id") // get the id parameter from the URL query string
	// if !ok {
	// 	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id query parameter"})
	// 	return
	// }

	book, err := bookByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

// Update a book
func updateBook(c *gin.Context) {
	id := c.Param("id")
	row, err := bookByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	var updatedBook book
	if err := c.BindJSON(&updatedBook); err != nil {
		return
	}
	// update time
	now := time.Now()
	updatedBook.UpdatedAt = &now
	// update the book in the list
	books = append(books, updatedBook) // append the book to the list, it will overwrite the old book data
	c.IndentedJSON(http.StatusOK, row)
}

// Delete a book
func deleteBook(c *gin.Context) {
	id := c.Param("id")
	book, err := bookByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	// update the deleted_at field to the current time
	now := time.Now()
	book.DeletedAt = &now

	// update the book in the list
	books = append(books, *book) // append the book to the list, it will overwrite the old book data
	c.IndentedJSON(http.StatusOK, book)
}

// User can borrow a book
func borrowBook(c *gin.Context) {
	id := c.Param("id")

	book, err := bookByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}
	if book.Quantity == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "book not available"})
		return
	}
	book.Quantity--
	// update the updated_at field to the current time
	now := time.Now()
	book.UpdatedAt = &now
	// update the book in the list
	books = append(books, *book) // append the book to the list, it will overwrite the old book data

	c.IndentedJSON(http.StatusOK, book)
}

// User can return a book
func returnBook(c *gin.Context) {
	id := c.Param("id")

	book, err := bookByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	book.Quantity++
	// update the updated_at field to the current time
	now := time.Now()
	book.UpdatedAt = &now
	// update the book in the list
	books = append(books, *book) // append the book to the list, it will overwrite the old book data

	c.IndentedJSON(http.StatusOK, book)
}

func main() {
	router := gin.Default() // create a new gin router instance
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	/*
		CURL USAGE
		curl --header "Content-Type: application/json" \
				--request POST \
				--data '{"isbn":"9781804619261","title":"Domain-Driven Design with Golang : Use Golang to Create Simple, Maintainable Systems to Solve Complex Business Problems", "author":"Boyle, Matthew.", "quantity":15}' \
				http://localhost:8800/books
	*/
	router.GET("/books/:id", getBookById)
	router.PUT("/books/:id", updateBook)
	/*
		CURL USAGE
		curl --header "Content-Type: application/json" \
				--request PUT \
				--data '{"isbn":"9781804619261","title":"Domain-Driven Design with Golang : Use Golang to Create Simple, Maintainable Systems to Solve Complex Business Problems", "author":"Boyle, Matthew.", "quantity":15}' \
				http://localhost:8800/books/dfd72ad7-25a9-49be-b8de-9f27669fc4a8
	*/
	router.DELETE("/books/:id", deleteBook)
	/*
		CURL USAGE
		curl --header "Content-Type: application/json" \
				--request DELETE \
				http://localhost:8800/books/dfd72ad7-25a9-49be-b8de-9f27669fc4a8
	*/

	router.PATCH("/books/borrow/:id", borrowBook)
	/*
		CURL USAGE
		curl --header "Content-Type: application/json" \
				--request PATCH \
				http://localhost:8800/books/borrow/dfd72ad7-25a9-49be-b8de-9f27669fc4a8
	*/
	router.PATCH("/books/return/:id", returnBook)
	/*
		CURL USAGE
		curl --header "Content-Type: application/json" \
				--request PATCH \

				http://localhost:8800/books/return/dfd72ad7-25a9-49be-b8de-9f27669fc4a8
	*/

	router.Run("localhost:8800") // start the server on port 8080
}
