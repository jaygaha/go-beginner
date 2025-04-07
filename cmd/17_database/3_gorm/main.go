package main

import (
	"fmt"

	"github.com/jaygaha/go-beginner/tree/main/cmd/17_database/3_gorm/sqlite_db"
	"gorm.io/gorm"
)

type Book struct {
	ID     int    `gorm:"primaryKey"`
	Isbn   string `gorm:"unique"`
	Title  string
	Author string
}

// create a book
func CreateBook(db *gorm.DB, book Book) (Book, error) {
	result := db.Create(&book)
	if result.Error != nil {
		return Book{}, result.Error
	}

	return book, nil
}

// read a book
func GetBook(db *gorm.DB, id int) (Book, error) {
	var book Book
	// .First: get the first row that matches the condition
	// db.First(&book, id)
	// db.Where("id = ?", id).First(&book)
	// db.Where("id =? AND author =?", id, "John").First(&book)
	result := db.First(&book, id)
	if result.Error != nil {
		return Book{}, result.Error
	}

	return book, nil
}

// update a book
func UpdateBook(db *gorm.DB, book Book) (Book, error) {
	result := db.Save(&book)
	if result.Error != nil {
		return Book{}, result.Error
	}

	return book, nil
}

// delete a book
func DeleteBook(db *gorm.DB, id int) error {
	result := db.Delete(&Book{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func main() {
	/*
		ORM:
			- object relational mapping; technique used to convert data between relational database and object-oriented programming languages
			- instead of using db native lanuage such as sql, postgres, mysql, etc; use orm library to interact with db
			- orm libs: gorm, sqlx, etc

		Key features:
			- mapping struct to table
			- CRUD operations
			- query builder
			- relationships
			- migrations
			- transactions
			- logging

		Advantages:
			- easier to use
			- faster to develop
			- easier to maintain

		Disadvantages:
			- less performant
			- less secure
			- complex to understand
	*/

	// connect to sqlite
	db, err := sqlite_db.ConnectToSQLite()
	if err != nil {
		panic(err)
	}

	// migrate the schema
	err = db.AutoMigrate(&Book{})
	if err != nil {
		panic(err)
	}

	// CRUD operations
	// create
	newBook := Book{
		Isbn:   "1234567890",
		Title:  "The Lord of the Rings",
		Author: "J.R.R. Tolkien",
	}

	row, err := CreateBook(db, newBook)
	if err != nil {
		panic(err)
	}

	fmt.Println("New book aded: ", row)

	// read
	book, err := GetBook(db, row.ID)
	if err != nil {
		panic(err)
	}

	fmt.Println("Book found: ", book)

	// update
	book.Title = "The Lord of the Rings: The Fellowship of the Ring"
	updatedBook, err := UpdateBook(db, book)
	if err != nil {
		panic(err)
	}
	fmt.Println("Book updated: ", updatedBook)

	// delete
	err = DeleteBook(db, updatedBook.ID)
	if err != nil {
		panic(err)
	}

	fmt.Println("Book deleted")
}
