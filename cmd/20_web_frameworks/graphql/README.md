# GraphQL in Go

This tutorial guides beginners through creating a simple GraphQL server in Go using the `graphql-go` library. We'll build a basic API for managing a list of books, covering schema definition, resolvers, and querying.

## What is GraphQL?

GraphQL is a query language for APIs and a runtime for executing those queries with existing data. It allows clients to request only the data as needed, making it more efficient than traditional REST APIs.

### Basic Concepts
- **Schema**: Defines the structure of APIs, including types, queries, and mutations.
- **Resolvers**: Functions that provide data for fields in the schema.
- **Queries**: Requests for data from the server.

### Making Queries

To query the GraphQL server, you can use tools like GraphQL client like Postman, Apollo Client.
For example, to get all books:
```graphql
query {
  books {
    id
    title
    author
  }
}
```

## Getting Started

### Dependencies
- [graphql-go](https://github.com/graphql-go/graphql): A GraphQL library for Go.
- [gorilla/mux](https://github.com/gorilla/mux): A HTTP router and URL matcher for Go. 

Navigate to the `cmd/20_web_frameworks/graphql` directory and run:
```bash
go run tidy # to install dependencies
go run main.go # to run the server
```

[GraphQL Playground](http://localhost:8800/graphql) will be up and ready to use.

### Sample Project Setup

#### 1. Define the Data Model

`models/book.go` it includes the `Book` struct.

#### 2. Create the GraphQL Schema

`graphql/schema.go` defines the GraphQL schema, including types and resolvers.

#### 3. Handlers and Resolvers

`graphql/handlers.go` handles GraphQL requests.

#### 4. Running the Server

`main.go` sets up the server and starts listening on port 8800.

#### 5. Testing

[GraphQL Playground](http://localhost:8800/graphql)

#### 1 Create a Book (Mutation)
```graphql
mutation {
  createBook(title: "Book 1", author: "Author 1") {
    id
    title
    author
  }
}
```

#### 2 Get All Books (Query)
```graphql
query {
  books {
    id
    title
    author
  }
}
```

#### 3 Get a Book by ID (Query)
```graphql
query {
  book(id: 1) {
    id
    title
    author
  }
}
```
