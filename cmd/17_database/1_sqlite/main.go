package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// struct to represent a row in the table
type Grocery struct {
	Id          int
	Name        string
	Price       float64
	Quantity    int
	isAvailable bool
	CreatedAt   time.Time
	UpdatedAt   *time.Time // can be nil
}

// gloabl db: database connection
var DB *sql.DB

func main() {
	// connect to the database
	ConnectDB()

	defer DB.Close() // close the database connection when the program exits

	// add a grocery item
	AddGrocery("Milk", 10.0, 3)
	AddGrocery("Bread", 2.35, 2)
	AddGrocery("Cerels", 3.75, 7)

	// get all grocery items
	// DisplayGroceries()

	// Add new item
	AddGrocery("Sugar", 1.23, 4)

	// print the grocery items
	// fmt.Print("\n\nAdded: \n\n")
	// DisplayGroceries()

	// update a grocery item after consumption
	UpdateGrocery(1, 2)

	// get all grocery items
	// fmt.Print("\n\nUpdated: \n\n")
	// DisplayGroceries()

	// remove a grocery item
	DeleteGrocery(1)

	// get all grocery items
	// fmt.Print("\n\nDeleted: \n\n")
	DisplayGroceries()
}

// function to connect to the database
func ConnectDB() {
	db, err := sql.Open("sqlite3", "./grocery.db") // sqlite3: driver name, ./grocery.db: database file name
	if err != nil {
		panic(err)
	}

	// create table if not exists
	stetement, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS groceries (
		id INTEGER PRIMARY KEY,
		name VARCHAR(255),
		price REAL,
		quantity INTEGER DEFAULT 0,
		is_available BOOLEAN DEFAULT false,
		created_at DATETIME,
		updated_at DATETIME DEFAULT NULL
	)`)
	_, err = stetement.Exec()

	if err != nil {
		panic(err)
	}

	DB = db
}

// function to add a grocery item
func AddGrocery(name string, price float64, quantity int) sql.Result {
	// insert a row into the table
	statement, _ := DB.Prepare(`INSERT INTO groceries (name, price, quantity, created_at) VALUES (?, ?, ?, ?)`)
	row, err := statement.Exec(name, price, quantity, time.Now())

	if err != nil {
		panic(err)
	}

	return row
}

// function to get all grocery items
func GetAllGroceries() []Grocery {
	// select all rows from the table
	rows, err := DB.Query(`SELECT * FROM groceries`)
	if err != nil {
		panic(err)
	}
	defer rows.Close() // close the rows when the function exits

	var groceries []Grocery

	// iterate over the rows
	for rows.Next() {
		var grocery Grocery
		// scan the row into the grocery struct
		err := rows.Scan(&grocery.Id, &grocery.Name, &grocery.Price, &grocery.Quantity, &grocery.isAvailable, &grocery.CreatedAt, &grocery.UpdatedAt)

		if err != nil {
			panic(err)
		}

		groceries = append(groceries, grocery)
	}

	return groceries
}

// function to update a grocery item
func UpdateGrocery(id, quantity int) {
	// select the row from the table
	row, err := DB.Query(`SELECT * FROM groceries WHERE id = ?`, id)

	if err != nil {
		panic(err)
	}

	// check if the row exists
	hasRow := row.Next()
	row.Close() // close the row immediately after checking

	if !hasRow {
		panic("Grocery not found")
	}

	// update the row
	statement, _ := DB.Prepare(`UPDATE groceries SET quantity =?, updated_at =? WHERE id =?`)
	_, err = statement.Exec(quantity, time.Now(), id)

	if err != nil {
		panic(err)
	}

	statement.Close() // close the statement after execution
}

// delete a grocery item
func DeleteGrocery(id int) {
	// delete the row from the table
	statement, _ := DB.Prepare(`DELETE FROM groceries WHERE id =?`)
	_, err := statement.Exec(id)

	if err != nil {
		panic(err)
	}
}

// display all grocery items
func DisplayGroceries() {
	groceries := GetAllGroceries()

	for _, grocery := range groceries {
		fmt.Println(grocery.Id, grocery.Name, grocery.Price, grocery.Quantity, grocery.isAvailable, grocery.CreatedAt, grocery.UpdatedAt)
	}
}
