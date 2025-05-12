# GraphQL with `gqlgen` in Go

This tutorial demonstrates how to build a GraphQL API in Go using the [`gqlgen`](https://github.com/99designs/gqlgen) library. gqlgen is a Go library for building GraphQL servers without boilerplate code.

## What is GraphQL?

GraphQL is a query language for APIs and a runtime for executing those queries against your data. It provides a complete and understandable description of the data in your API and gives clients the power to ask for exactly what they need.

## What is gqlgen?

gqlgen is a Go library that helps you build GraphQL servers in Go with minimal boilerplate. It takes a schema-first approach to developing GraphQL servers by generating code based on your schema definition.

Key features of gqlgen:
- Schema-first development
- Type-safe code generation
- Minimal boilerplate
- Flexibility in resolver implementation
- Built-in validation and error handling

## Project Structure

```
.
├── go.mod             # Go module definition
├── go.sum             # Go module checksums
├── gqlgen.yml         # gqlgen configuration file
├── graph/             # Generated GraphQL code
│   ├── generated.go   # Auto-generated GraphQL server code
│   ├── model/         # Generated model definitions
│   │   └── models_gen.go
│   ├── resolver.go    # Custom resolver implementation
│   ├── schema.graphqls # GraphQL schema definition
│   └── schema.resolvers.go # Generated resolvers
├── server.go          # Main server entry point
└── tools.go           # Tools dependencies
```

## Getting Started

### Prerequisites

- Go 1.16 or later
- Basic understanding of GraphQL concepts

### Installation

1. Clone the repository or create a new directory for your project
2. Initialize a Go module:
   ```bash
   go mod init your-module-name
   ```
3. Install gqlgen: Add `gqlgen` as a tool dependency for code generation.
    ```bash
    printf '//go:build tools\npackage tools\nimport _ "github.com/99designs/gqlgen"' | gofmt > tools.go
    go mod tidy
    ```
4. Initialize gqlgen: Run the init command to create the project structure.
   ```bash
   go run github.com/99designs/gqlgen init
   ```

   This generates:
    - **gqlgen.yml**: Configuration file
    - **graph/schema.graphqls**: Schema file
    - **graph/generated/**: Generated code
    - **graph/model/**: Generated models
    - **graph/resolver.go**: Root resolver
    - g**raph/schema.resolvers.go**: Resolver stubs
    - **server.go**: Server entry point

### Schema Definition

The GraphQL schema is defined in `graph/schema.graphqls`. This example defines a simple book API:

```graphql
# A Book type with id, title, author, and published fields
type Book {
  id: ID!
  title: String!
  author: String!
  published: Boolean!
}

# A Query to fetch all books
type Query {
  books: [Book!]!
}

# 
input NewBook {
  title: String!
  author: String!
}

# A Mutation to add a new book using a NewBook input type
type Mutation {
  addBook(input: NewBook!): Book!
}
```

### Configuration

The `gqlgen.yml` file configures how gqlgen generates code. Key configuration options include:

- `schema`: Location of schema files
- `exec`: Where generated server code should go
- `model`: Where generated models should go
- `resolver`: Where resolver implementations should go

### Running the Server

To run the GraphQL server:

```bash
go run server.go
```

This will start a server on http://localhost:8800/ with a GraphQL playground interface where you can test your API.

### Making Queries

Once the server is running, you can access the GraphQL playground at http://localhost:8800/ and make queries like:

```graphql
query GetBooks {
  books {
    id
    title
    author
    published
  }
}
```

Or mutations:

```graphql
mutation AddBook {
  addBook(input: {title: "The Go Programming Language", author: "Alan A. A. Donovan & Brian W. Kernighan"}) {
    id
    title
    author
  }
}
```

## Customizing Resolvers

The resolver implementations are in `graph/schema.resolvers.go`. This is where you implement the business logic for your GraphQL API.

## Learn More

- [gqlgen Documentation](https://gqlgen.com/)
- [GraphQL Official Documentation](https://graphql.org/)
- [Go Documentation](https://golang.org/doc/)

## Next Steps

- Add authentication
- Connect to a database
- Implement more complex resolvers
- Add subscriptions for real-time updates