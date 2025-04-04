package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	/*
		database/sql:
			- This is a standard library package that provides a generic interface for interacting with databases.

		To integrate with mysql database we need to use external package or driver.:
			- go get github.com/go-sql-driver/mysql

		- The sql.Open() function takes two arguments: the driver name and the data source name (DSN).
		- The driver name is the name of the driver to use, such as "mysql" for MySQL.
		- The DSN is a string that contains the connection information for the database, such as the hostname, port, username, password, and database name.
	*/

	fmt.Println("For more information, check test cases.\nThis tutorial is based on TDD apprach about the MySQL integration.")
}
