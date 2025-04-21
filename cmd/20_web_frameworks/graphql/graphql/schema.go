/*
GraphQL
-> an open spec for describing data and exposing operations over it
-> it make the app more flexible and responsive
-> it is a query language for APIs and a runtime for fulfilling queries with your existing data
-> it is a type system to define the shape of data and a query language to fetch that data
-> it is a client-server architecture

GraphQL vs REST
-> GraphQL returns only the data that is requested by the client instead of returning all the data
-> It only exposes a single endpoint, which client can perform different operations like query, mutation, subscription
-> Strongly typed and schema driven

3 types
-> Query: to fetch data, client can dicid which data to fetch

	Example:
		query {
			users {
				id
				name
				email
			}
		}

-> Mutation: to create, update, delete data

	-> follows the same syntax as query; always needs to mention keyword "mutation"

	Example:
		mutation {
			createUser(input: {
				name: "John Doe",
				email: "johndoe@domain.com"
			}) {
				id
				name
				email
			}
		}
	}

-> Subscription: to subscribe to data changes

Understanding the flow:

 1. Data Model: Define structs (e.g., Book) to represent your data.
 2. GraphQL Schema: Define GraphQL types and fields based on your data model.
 3. Resolvers: Implement functions to resolve the GraphQL queries and mutations.
 4. GraphQL Server: Create a GraphQL server that exposes your schema and resolvers.
 5. Client: Use a GraphQL client (e.g., Apollo Client) to make queries and mutations to your GraphQL server.
*/
package graphql

import (
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/jaygaha/go-beginner/cmd/20_web_frameworks/graphql/models"
)

// in-memory database
var books = []*models.Book{}

/*
Object Types:
*/
// define the Book GraphQL object type
var bookType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Book",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"author": &graphql.Field{
			Type: graphql.String,
		},
	},
},
)

/*
Mutations:
-> Used to modify the data in the database like create, update, delete.
*/

var bookInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "BookInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"author": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})

/*
Queries:
-> Used to fetch the data from the database.
*/
// define the root query GraphQL object type
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"books": &graphql.Field{
			Type: graphql.NewList(bookType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return books, nil
			},
		},
		"book": &graphql.Field{
			Type: bookType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["id"].(int)
				if !ok {
					return nil, errors.New("invalid book id")
				}

				for _, book := range books {
					if book.ID == id {
						return book, nil
					}
				}
				return nil, errors.New("book not found")
			},
		},
	},
})

// define the root mutation GraphQL object type
var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createBook": &graphql.Field{
			Type: bookType,
			Args: graphql.FieldConfigArgument{
				"input": &graphql.ArgumentConfig{
					Type: bookInputType,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				input, ok := p.Args["input"].(map[string]interface{})
				if !ok {
					return nil, errors.New("invalid input")
				}

				book := &models.Book{
					ID:     len(books) + 1,
					Title:  input["title"].(string),
					Author: input["author"].(string),
				}

				books = append(books, book)
				return book, nil
			},
		},
	},
})

// define the GraphQL schema
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})

/*
Above defines the GraphQL schema for the application.
	-> A Book type with id, title, and author fields
	-> A BookInput type for mutation inputs
	-> Queries: books (list all books) and book (get a book by ID)
	-> Mutation: createBook (create a new book)
	-> An in-memory books slice for data storage
*/
