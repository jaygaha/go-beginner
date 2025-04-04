package main

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

// Test import will return the names of the drivers that are registered with the database/sql package.
// It ensure that the mysql driver is registered with the database/sql package.
func TestImportDriver(t *testing.T) {
	assert.Equal(t, []string{"mysql"}, sql.Drivers())
}

// Test connecting to database
func TestConnectToDatabase(t *testing.T) {
	// Open a connection to the database
	// func Open(driverName, dataSourceName string) (*DB, error)
	// driverName: mysql
	// dataSourceName: user:password@tcp(localhost:3306)/dbName?param=value
	db, err := sql.Open("mysql", "gopher:golang123@tcp(localhost:3306)/go_test")
	assert.NoError(t, err) // NoError is a testify function that checks if the error is nil
	defer db.Close()

	// Ping the database to check if it is connected
	err = db.Ping()
	assert.NoError(t, err)
}

// Test specifying parameters in the connection string
func TestSpecifyingParametersInConnectionString(t *testing.T) {
	db, err := sql.Open("mysql", "gopher:golang123@tcp(localhost:3306)/go_test?parseTime=true&interpolateParams=true")
	// interpolateParams: true will interpolate the parameters into the query string instead of using placeholders (?)
	// parseTime: true will parse the time fields into time.Time type
	assert.NoError(t, err)
	defer db.Close()

	// Ping the database to check if it is connected
	err = db.Ping()
	assert.NoError(t, err)
}

// Test running the DDL statements
// DDL statment:
//   - Data Definition Language (DDL) statements are used to define the structure of the database.
//   - DDL statements are used to create, alter, and drop tables, views, and indexes.
func TestRunningDDLStatements(t *testing.T) {
	db, err := sql.Open("mysql", "gopher:golang123@tcp(localhost:3306)/go_test")
	assert.Nil(t, err) // Nil is a testify function that checks if the value is nil

	defer db.Close()

	// create a table
	// db.Exec: Execute a query; if _ given it will not return response.
	// func (db *DB) Exec(query string, args ...any) (Result, error)
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255))")
	assert.Nil(t, err)

	// alter a table
	_, err = db.Exec("ALTER TABLE users ADD COLUMN age INT")
	assert.NoError(t, err)

	// drop a table
	_, err = db.Exec("DROP TABLE IF EXISTS users")
	assert.Nil(t, err)
}

// Test running the DML statements
// DML statment:
//   - Data Manipulation Language (DML) statements are used to manipulate data in the database.
//   - DML statements are used to insert, update, and delete records from tables.
func TestRunningDMLStatements(t *testing.T) {
	db, err := sql.Open("mysql", "gopher:golang123@tcp(localhost:3306)/go_test")
	assert.NoError(t, err)
	defer db.Close()

	// create a table
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255))")
	assert.NoError(t, err)

	// insert a record
	record, err := db.Exec("INSERT INTO users (name) VALUES ('Hoge')")
	assert.NoError(t, err)

	rowAffected, err := record.RowsAffected()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), rowAffected)

	record, err = db.Exec("INSERT INTO users (name) VALUES ('Suga')")
	assert.NoError(t, err)

	rowAffected, err = record.RowsAffected()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), rowAffected)

	// test last inserted id
	lastInsertedId, err := record.LastInsertId()
	assert.NoError(t, err)
	assert.Equal(t, int64(2), lastInsertedId)

	// count all records
	// QueryRow executes a query that is expected to return at most one row.
	// func (db *DB) QueryRow(query string, args...any) *Row
	row := db.QueryRow("SELECT COUNT(*) FROM users")
	assert.NoError(t, row.Err())

	var count int
	row.Scan(&count)          // Scan copies the columns in the current row into the values pointed at by dest.
	assert.Equal(t, 2, count) // 2 records

	// select specific row
	row = db.QueryRow("SELECT id, name FROM users WHERE id = ?", lastInsertedId)
	assert.NoError(t, row.Err())

	var name string
	row.Scan(&lastInsertedId, &name) // Scan copies the columns in the current row into the values pointed at by dest.

	// assert
	assert.Equal(t, 2, int(lastInsertedId))
	assert.Equal(t, "Suga", name)

	// select all records
	// .Query: Query executes a query that returns rows, typically a SELECT. The args are for any placeholder parameters in the query.
	// func (db *DB) Query(query string, args...any) (*Rows, error)
	rows, err := db.Query("SELECT id, name FROM users")
	assert.NoError(t, err)

	names := []string{}

	// Iterate through results
	// rows.Next(): Next prepares the next result row for reading with the Scan method.
	// It returns true on success, false if there is no next result row or an error happened while preparing it.
	for rows.Next() {
		var id int
		// var name string

		// rows.Scan(): Scan copies the columns in the current row into the values pointed at by dest.
		err = rows.Scan(&id, &name)

		assert.NoError(t, err)
		names = append(names, name)
	}

	assert.Equal(t, []string{"Hoge", "Suga"}, names)

	// Update a record
	row = db.QueryRow("UPDATE users SET name =? WHERE id =?", "Piyo", lastInsertedId)
	assert.NoError(t, row.Err())

	// Check the name is updated
	row = db.QueryRow("SELECT name FROM users WHERE id =?", lastInsertedId)
	assert.NoError(t, row.Err())
	row.Scan(&name)
	assert.Equal(t, "Piyo", name)

	// delete a record
	row = db.QueryRow("DELETE FROM users WHERE id =?", lastInsertedId)
	assert.NoError(t, row.Err())

	// check the record is deleted
	row = db.QueryRow("SELECT COUNT(*) FROM users WHERE id =?", lastInsertedId)
	assert.NoError(t, row.Err())

	var count1 int
	row.Scan(&count1)
	assert.Equal(t, 0, count1) // There should be 0 record

	// drop a table
	_, err = db.Exec("DROP TABLE IF EXISTS users")
	assert.Nil(t, err)
}

// Using prepared statement
// Best practice is to use prepared statement.
func TestUsingPreparedStatement(t *testing.T) {
	db, err := sql.Open("mysql", "gopher:golang123@tcp(localhost:3306)/go_test")
	assert.NoError(t, err)
	defer db.Close()

	// create a table
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255))")
	assert.NoError(t, err)

	// insert a record
	stmt, err := db.Prepare("INSERT INTO users (name) VALUES (?)")
	assert.NoError(t, err)

	record, err := stmt.Exec("Hoge")
	assert.NoError(t, err)

	rowAffected, err := record.RowsAffected()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), rowAffected)

	// for rest insert we can use the same statement
	record, err = stmt.Exec("Suga")
	assert.NoError(t, err)

	rowAffected, err = record.RowsAffected()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), rowAffected)

	record, err = stmt.Exec("Fuga")
	assert.NoError(t, err)

	rowAffected, err = record.RowsAffected()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), rowAffected)

	// test last inserted id
	lastInsertedId, err := record.LastInsertId()
	assert.NoError(t, err)
	assert.Equal(t, int64(3), lastInsertedId)

	// close statement; it is necessary to close statement after using it. Otherwise it will leak the connection.
	stmt.Close()

	// Query row
	row := db.QueryRow("SELECT COUNT(*) FROM users")
	assert.NoError(t, row.Err())

	var count int
	row.Scan(&count)          // Scan copies the columns in the current row into the values pointed at by dest.
	assert.Equal(t, 3, count) // 2 records

	// Specific row
	smt, err := db.Prepare("SELECT id, name FROM users WHERE id =?")
	assert.NoError(t, row.Err())

	row = smt.QueryRow(lastInsertedId)
	assert.NoError(t, row.Err())

	var name string
	row.Scan(&lastInsertedId, &name) // Scan copies the columns in the current row into the values pointed at by dest.
	// assert
	assert.Equal(t, 3, int(lastInsertedId))
	assert.Equal(t, "Fuga", name)

	// close statement
	smt.Close()

	// Select all records
	smt, err = db.Prepare("SELECT id, name FROM users where id > ? ORDER BY id DESC")
	assert.NoError(t, err)

	rows, err := smt.Query(1) // 1 is the value of id >?
	assert.NoError(t, err)

	names := []string{}

	// Iterate through results
	for rows.Next() {
		var id int

		err = rows.Scan(&id, &name)
		assert.NoError(t, err)
		names = append(names, name)
	}

	assert.Equal(t, []string{"Fuga", "Suga"}, names)

	smt.Close()

	// Update a record
	stmt, err = db.Prepare("UPDATE users SET name =? WHERE id =?")
	assert.NoError(t, err)

	record, err = stmt.Exec("Piyo", lastInsertedId)
	assert.NoError(t, err)

	rowAffected, err = record.RowsAffected()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), rowAffected)

	// Check the name is updated
	row = db.QueryRow("SELECT name FROM users WHERE id =?", lastInsertedId)
	assert.NoError(t, row.Err())

	row.Scan(&name)
	assert.Equal(t, "Piyo", name)

	// close statement
	stmt.Close()

	// delete a record
	stmt, err = db.Prepare("DELETE FROM users WHERE id =?")
	assert.NoError(t, err)

	record, err = stmt.Exec(lastInsertedId)
	assert.NoError(t, err)

	rowAffected, err = record.RowsAffected()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), rowAffected)

	// check the record is deleted
	row = db.QueryRow("SELECT COUNT(*) FROM users WHERE id =?", lastInsertedId)
	assert.NoError(t, row.Err())

	var count1 int
	row.Scan(&count1)
	assert.Equal(t, 0, count1) // There should be 0 record

	// close statement
	stmt.Close()

	// drop a table
	_, err = db.Exec("DROP TABLE IF EXISTS users")
	assert.Nil(t, err)
}

// Using transaction
// This test demonstrates how to use transaction in a simple way.
func TestUsingTransaction(t *testing.T) {
	db, err := sql.Open("mysql", "gopher:golang123@tcp(localhost:3306)/go_test")
	assert.NoError(t, err)
	defer db.Close()

	// create a table
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255))")
	assert.NoError(t, err)

	/*
		Transaction:
			- A transaction is a group of SQL statements that are executed as a single unit.
			- A transaction is atomic, which means that either all of the statements in the transaction are executed or none of them are.
			- A transaction is consistent, which means that the database is left in a valid state after the transaction is completed.

			Syntax:
				START TRANSACTION;
				SQL statements;
				COMMIT;
				or
				START TRANSACTION;
				SQL statements;
				ROLLBACK;
				or
				BEGIN;
				SQL statements;
				COMMIT;

	*/
	// .Begin: Begin starts a transaction.
	tx, err := db.Begin()
	assert.NoError(t, err)

	// insert a record
	record, err := tx.Exec("INSERT INTO users (name) VALUES (?)", "Hoge")
	assert.NoError(t, err)

	rowAffected, err := record.RowsAffected()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), rowAffected)

	// .Rollback: Rollback rolls back the transaction.
	// if we want to rollback the transaction we can use Rollback method.
	// It means database will not be affected by the transaction.
	// It means we can rollback the transaction and the database will be in the same state as before the transaction.
	// It is useful when we want to test the transaction.
	err = tx.Rollback()
	assert.NoError(t, err)

	// check if the record is inserted
	row := db.QueryRow("SELECT COUNT(*) FROM users")
	assert.NoError(t, row.Err())

	var count int
	row.Scan(&count)
	assert.Equal(t, 0, count) // 0 records

	//.Commit: Commit commits the transaction.
	// if we want to commit the transaction we can use Commit method.
	tx, err = db.Begin()
	assert.NoError(t, err)

	// insert a record
	record, err = tx.Exec("INSERT INTO users (name) VALUES (?)", "Hoge")
	assert.NoError(t, err)

	rowAffected, err = record.RowsAffected()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), rowAffected)

	record, err = tx.Exec("INSERT INTO users (name) VALUES (?)", "Fuga")
	assert.NoError(t, err)

	rowAffected, err = record.RowsAffected()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), rowAffected)

	record, err = tx.Exec("INSERT INTO users (name) VALUES (?)", "Suga")
	assert.NoError(t, err)

	rowAffected, err = record.RowsAffected()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), rowAffected)

	// commit the transaction
	//.Commit: Commit commits the transaction which means that all the statements in the transaction are executed finally.
	err = tx.Commit()
	assert.NoError(t, err)

	// again initialize tx
	tx, err = db.Begin()
	assert.Nil(t, err)

	// check if the records are inserted
	row = tx.QueryRow("SELECT COUNT(*) FROM users")
	assert.NoError(t, row.Err())
	row.Scan(&count)
	assert.Equal(t, 3, count) // 3 records

	// check specific record
	row = tx.QueryRow("SELECT id, name FROM users WHERE id = ?", 2)
	assert.NoError(t, row.Err())

	var name string
	var id int
	row.Scan(&id, &name)
	assert.Equal(t, 2, id)
	assert.Equal(t, "Hoge", name)

	// select all records
	rows, err := tx.Query("SELECT id, name FROM users where id < ? ORDER BY id DESC", 4)
	assert.NoError(t, err)

	names := []string{}
	// Iterate through results
	for rows.Next() {
		var id int

		err = rows.Scan(&id, &name)
		assert.NoError(t, err)
		names = append(names, name)
	}

	assert.Equal(t, []string{"Fuga", "Hoge"}, names)

	err = tx.Commit()
	assert.Nil(t, err)

	// Update a record
	tx, err = db.Begin()
	assert.NoError(t, err)

	// First, let's check what ID we're working with
	var firstId int
	row = tx.QueryRow("SELECT id FROM users WHERE name = ? LIMIT 1", "Hoge")
	assert.NoError(t, row.Err())
	row.Scan(&firstId)

	// Now update using the correct ID
	record, err = tx.Exec("UPDATE users SET name = ? WHERE id = ?", "Piyo", firstId)
	assert.NoError(t, err)

	rowAffected, err = record.RowsAffected()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), rowAffected)

	err = tx.Commit()
	assert.NoError(t, err)

	// Check the name is updated
	row = db.QueryRow("SELECT name FROM users WHERE id = ?", firstId)
	assert.NoError(t, row.Err())

	row.Scan(&name)
	assert.Equal(t, "Piyo", name)

	// delete a record
	tx, err = db.Begin()
	assert.NoError(t, err)

	record, err = tx.Exec("DELETE FROM users WHERE name =?", "Piyo")
	assert.NoError(t, err)

	rowAffected, err = record.RowsAffected()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), rowAffected)

	err = tx.Commit()
	assert.NoError(t, err)

	// check the record is deleted
	row = db.QueryRow("SELECT COUNT(*) FROM users WHERE name =?", "Piyo")
	assert.NoError(t, row.Err())
	row.Scan(&count)
	assert.Equal(t, 0, count) // There should be 0 record

	// drop a table
	_, err = db.Exec("DROP TABLE IF EXISTS users")
	assert.Nil(t, err)
}

/*
Using context and transaction

Context in transaction is used to carry deadlines, cancellation signals, and other request-scoped values across API
boundaries and between processes.
Important:
	- Cancellation: you can cancel a transaction if it takes too long to complete or decides to abort the operation.
	- Timeouts: you can set a timeout for a transaction to complete within a certain amount of time.
	- Passing data: you can pass request-scoped values, such as user IDs, through the context to the transaction. Useful for logging and tracing purposes.
*/

func TestUsingContextAndTransaction(t *testing.T) {
	db, err := sql.Open("mysql", "gopher:golang123@tcp(localhost:3306)/go_test")
	assert.NoError(t, err)
	defer db.Close()

	// create a table
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255))")
	assert.NoError(t, err)

	// init context
	// context.Background() returns an empty context. It is used as the parent context for all other contexts.
	ctx := context.Background()

	// now init transaction
	// BeginTx starts a transaction.
	// The provided context is used until the transaction is committed or rolled back.
	// If the context is canceled, the transaction will be rolled back.
	// The provided TxOptions is optional and can be used to set the isolation level of the transaction.
	// If the TxOptions is nil, the default isolation level is used.
	// If the TxOptions is not nil, the isolation level is set to the value of the Isolation field.
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	assert.NoError(t, err)

	// insert a record
	record, err := tx.Exec("INSERT INTO users (name) VALUES (?)", "Hoge")
	assert.NoError(t, err)

	rowAffected, err := record.RowsAffected()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), rowAffected)

	// now rollback the transaction
	err = tx.Rollback()
	assert.NoError(t, err)

	// transaction has already been rolled back
	// so we can check if the record is inserted
	row := tx.QueryRow("SELECT COUNT(*) FROM users")
	assert.EqualError(t, row.Err(), "sql: transaction has already been committed or rolled back")
	assert.ErrorIs(t, sql.ErrTxDone, row.Err())

	// now init again
	ctx = context.Background()
	tx, err = db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	assert.NoError(t, err)

	// query row
	row = tx.QueryRow("SELECT COUNT(*) FROM users")
	assert.NoError(t, row.Err())

	var count int
	row.Scan(&count)
	assert.Equal(t, 0, count) // 0 records

	// now insert a record
	record, err = tx.Exec("INSERT INTO users (name) VALUES (?)", "Hoge")
	assert.NoError(t, err)

	rowAffected, err = record.RowsAffected()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), rowAffected)

	record, err = tx.Exec("INSERT INTO users (name) VALUES (?)", "Fuga")
	assert.NoError(t, err)

	rowAffected, err = record.RowsAffected()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), rowAffected)

	record, err = tx.Exec("INSERT INTO users (name) VALUES (?)", "Suga")
	assert.NoError(t, err)

	rowAffected, err = record.RowsAffected()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), rowAffected)

	// commit the transaction
	err = tx.Commit()
	assert.NoError(t, err)

	ctx = context.Background()
	// sql.LevelReadCommitted is the default isolation level for MySQL.
	// It means that the transaction will read the committed data only.
	// It means that the transaction will not see the data that is not committed yet.
	tx, err = db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	assert.NoError(t, err)

	// query row
	row = tx.QueryRow("SELECT COUNT(*) FROM users")
	assert.NoError(t, row.Err())
	row.Scan(&count)
	assert.Equal(t, 3, count) // 3 records

	row = tx.QueryRow("SELECT id, name FROM users WHERE id = ?", 2)
	assert.NoError(t, row.Err())

	var name string
	var id int

	row.Scan(&id, &name)
	assert.Equal(t, 2, id)
	assert.Equal(t, "Hoge", name)

	// select all records
	rows, err := tx.Query("SELECT id, name FROM users where id <? ORDER BY id DESC", 4)
	assert.NoError(t, err)

	names := []string{}

	// Iterate through results
	for rows.Next() {
		var id int
		err = rows.Scan(&id, &name)
		assert.NoError(t, err)
		names = append(names, name)
	}
	assert.Equal(t, []string{"Fuga", "Hoge"}, names)

	err = tx.Commit()
	assert.NoError(t, err)

	// Update a record
	tx, err = db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	assert.NoError(t, err)

	// First, let's check what ID we're working with
	var firstId int
	row = tx.QueryRow("SELECT id FROM users WHERE name =? LIMIT 1", "Hoge")
	assert.NoError(t, row.Err())

	row.Scan(&firstId)

	// Now update using the correct ID
	record, err = tx.Exec("UPDATE users SET name =? WHERE id =?", "Piyo", firstId)
	assert.NoError(t, err)

	rowAffected, err = record.RowsAffected()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), rowAffected)

	err = tx.Commit()
	assert.NoError(t, err)

	// Check the name is updated
	row = db.QueryRow("SELECT name FROM users WHERE id =?", firstId)
	assert.NoError(t, row.Err())

	row.Scan(&name)
	assert.Equal(t, "Piyo", name)

	// delete a record
	tx, err = db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	assert.NoError(t, err)

	record, err = tx.Exec("DELETE FROM users WHERE name =?", "Piyo")
	assert.NoError(t, err)

	rowAffected, err = record.RowsAffected()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), rowAffected)

	err = tx.Commit()
	assert.NoError(t, err)

	// check the record is deleted
	row = db.QueryRow("SELECT COUNT(*) FROM users WHERE name =?", "Piyo")
	assert.NoError(t, row.Err())

	row.Scan(&count)
	assert.Equal(t, 0, count) // There should be 0 record

	// drop a table
	_, err = db.Exec("DROP TABLE IF EXISTS users")
	assert.Nil(t, err)
}

/*
Connection Pool

  - Connection pool is a pool of connections to the database.

  - Connection pool is used to reuse connections to the database.

  - Connection pool is used to limit the number of connections to the database.

  - Important in high traffic applications where establishing a new connection to the database can be time consuming and resource intensive.

    Benefits:

  - Performance: resuing connections can improve performance.

  - Resource management: limiting the number of connections can help manage resources.

  - Concurrency: connection pool can help manage concurrency.

    database/sql package provides connection pool by default.
    sql package creates and maintains a pool of connections to the database automatically.
*/
func TestConnectionPool(t *testing.T) {
	db, err := sql.Open("mysql", "gopher:golang123@tcp(localhost:3306)/go_test")
	assert.NoError(t, err)

	defer db.Close()

	// asserting that the maximum number of idle connections in the pool is set to 0.
	assert.Zero(t, db.Stats().MaxOpenConnections)

	// set max idle connections
	// This ensures that the maximum number of idle connections in the pool is set to 100.
	db.SetMaxOpenConns(100)
	assert.Equal(t, 100, db.Stats().MaxOpenConnections)
}

// using context in connecting pool
func TestUsingContextConnectionPool(t *testing.T) {
	db, err := sql.Open("mysql", "gopher:golang123@tcp(localhost:3306)/go_test")
	assert.NoError(t, err)
	defer db.Close()

	// create a table
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255))")
	assert.NoError(t, err)

	// handle connection
	ctx := context.Background()
	// Conn returns a single connection from the pool.
	// The returned connection is only used by the caller and should be returned to the pool by calling Conn.Close.
	conn, err := db.Conn(ctx)
	assert.NoError(t, err)

	defer conn.Close()

	// insert a record
	// ExecContext executes a query without returning any rows.
	// The args are for any placeholder parameters in the query.
	record, err := conn.ExecContext(ctx, "INSERT INTO users (name) VALUES (?)", "Hoge")
	assert.NoError(t, err)

	rowAffected, err := record.RowsAffected()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), rowAffected)

	// query row
	// QueryRowContext executes a query that is expected to return at most one row.
	// QueryRowContext always returns a non-nil value. Errors are deferred until Row's Scan method is called.
	// If the query selects no rows, the *Row's Scan will return ErrNoRows.
	// Otherwise, the *Row's Scan scans the first selected row and discards the rest.
	row := conn.QueryRowContext(ctx, "SELECT COUNT(*) FROM users")
	assert.NoError(t, row.Err())

	var count int
	row.Scan(&count)
	assert.Equal(t, 1, count) // 1 record

	// specific record
	row = conn.QueryRowContext(ctx, "SELECT id, name FROM users WHERE id =?", 1)
	assert.NoError(t, err)

	var name string
	var id int
	row.Scan(&id, &name)
	assert.Equal(t, 1, id)
	assert.Equal(t, "Hoge", name)

	// select all records
	// QueryContext executes a query that returns rows, typically a SELECT.
	// The args are for any placeholder parameters in the query.
	rows, err := conn.QueryContext(ctx, "SELECT id, name FROM users where id <? ORDER BY id DESC", 4)
	assert.NoError(t, err)

	names := []string{}

	// Iterate through results
	for rows.Next() {
		var id int
		err = rows.Scan(&id, &name)
		assert.NoError(t, err)
		names = append(names, name)
	}
	assert.Equal(t, []string{"Hoge"}, names)

	// at last drop a  test table
	_, err = db.Exec("DROP TABLE IF EXISTS users")
	assert.Nil(t, err)
}

// Using transaction in connecting pool
func TestUsingTransactionConnectionPool(t *testing.T) {
	db, err := sql.Open("mysql", "gopher:golang123@tcp(localhost:3306)/go_test")
	assert.NoError(t, err)
	defer db.Close()

	// create a table
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255))")
	assert.NoError(t, err)

	// handle connection & transaction
	ctx := context.Background()
	conn, err := db.Conn(ctx)
	tx, err := conn.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	assert.NoError(t, err)

	defer conn.Close()

	// insert a record
	record, err := tx.ExecContext(ctx, "INSERT INTO users (name) VALUES (?)", "Hoge")
	assert.NoError(t, err)
	rowAffected, err := record.RowsAffected()

	assert.NoError(t, err)
	assert.Equal(t, int64(1), rowAffected)

	err = tx.Commit()
	assert.NoError(t, err)

	// query row
	ctx = context.Background()
	tx, err = conn.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})

	row := tx.QueryRowContext(ctx, "SELECT COUNT(*) FROM users")
	assert.NoError(t, row.Err())
	var count int
	row.Scan(&count)
	assert.Equal(t, 1, count) // 1 record

	// specific record
	row = tx.QueryRowContext(ctx, "SELECT id, name FROM users WHERE id =?", 1)
	assert.NoError(t, err)
	var name string
	var id int
	row.Scan(&id, &name)
	assert.Equal(t, 1, id)
	assert.Equal(t, "Hoge", name)

	// select all records
	rows, err := tx.QueryContext(ctx, "SELECT id, name FROM users where id <? ORDER BY id DESC", 4)
	assert.NoError(t, err)

	names := []string{}
	// Iterate through results
	for rows.Next() {
		var id int
		err = rows.Scan(&id, &name)
		assert.NoError(t, err)
		names = append(names, name)
	}
	assert.Equal(t, []string{"Hoge"}, names)

	err = tx.Commit()
	assert.NoError(t, err)

	// at last drop a  test table
	_, err = db.Exec("DROP TABLE IF EXISTS users")
	assert.Nil(t, err)
}
