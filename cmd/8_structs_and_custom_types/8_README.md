# Structs and Custom Types in Go

This directory contains examples demonstrating the use of structs and custom types in Go.

## Files

- `main.go`: Demonstrates basic struct usage, anonymous structs, and custom type initialization
- `department.go`: Shows how to create custom types with receiver functions

## Concepts Covered

### Structs
- Structs are blueprints for creating objects
- They are typed collections of fields, useful for grouping data together to form records
- Example: `Employee` struct with age and gender fields

### Anonymous Structs
- Structs that are defined and used without creating a named type
- Useful when you don't need to reuse the struct type

### Custom Types
- Creating your own data types with specific behaviors
- Example: `department` type with a constructor function

## How to Run

Execute the following command from this directory:

```bash
go run .
```

## Further Reading

- [Go Structs](https://go.dev/tour/moretypes/2)
- [Go by Example: Structs](https://gobyexample.com/structs)