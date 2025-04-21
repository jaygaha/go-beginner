/*
Models:
-> It holds the data structure for the application.
-> It is used to map the data from the database to the application.
*/
package models

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
