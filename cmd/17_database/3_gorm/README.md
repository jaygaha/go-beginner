# Golang Beginner Project - GORM Example

This project demonstrates basic CRUD operations using GORM, an ORM library for Golang, with a SQLite database.

## ORM

- object relational mapping; technique used to convert data between relational database and object-oriented programming languages
- instead of using db native lanuage such as sql, postgres, mysql, etc; use orm library to interact with db
- orm libs: gorm, sqlx, etc

**Key features:**

- mapping struct to table
- CRUD operations
- query builder
- relationships
- migrations
- transactions
- logging

**Advantages:**

- easier to use
- faster to develop
- easier to maintain

**Disadvantages:**

- less performant
- less secure
- complex to understand

## Project Structure

- **main.go**: Contains the main logic for CRUD operations on the `Book` struct.
- **sqlite_db/db_connection.go**: Contains the database connection logic.

## Requirements

- Go 1.24.0 or later
- GORM library
- SQLite database

## Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/jaygaha/go-beginner.git
   cd go-beginner/cmd/17_database/3_gorm # Navigate to the project directory
   go mod download # Download dependencies
   ```

## Usage

1. Run the application:
   
   ```bash
   go run main.go
    ```
2. The application will perform the following operations:
   
   - Create a new book entry.
   - Read the book entry.
   - Update the book entry.
   - Delete the book entry.

## Notes

- Ensure that the SQLite database file ( gorm.db ) is accessible and writable.
- The database connection logic is modularized in the sqlite_db package for reusability.