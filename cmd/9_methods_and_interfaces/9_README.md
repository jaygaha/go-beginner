# Methods and Interfaces in Go

This example demonstrates how methods and interfaces work in Go programming language.

## Methods

In Go, a method is a function with a special receiver argument. The receiver appears between the `func` keyword and the method name.

### Value Receivers vs Pointer Receivers

- **Value Receiver**: `func (s Student) Print()`
  - Receives a copy of the struct
  - Cannot modify the original struct
  - Good for read-only operations

- **Pointer Receiver**: `func (s *Student) UpdateAge()`
  - Receives a pointer to the struct
  - Can modify the original struct
  - More efficient for large structs (avoids copying)
  - Necessary when you need to modify the receiver

## Interfaces

An interface in Go is a set of method signatures. A type implements an interface by implementing its methods.

### Key Points About Interfaces

- Interfaces define behavior, not data
- Types implement interfaces implicitly (no `implements` keyword needed)
- A value of interface type can hold any value that implements those methods
- Interfaces allow for polymorphism - different types can be treated the same if they implement the same interface

## Example Code Explanation

In this example:

1. We define a `Student` struct with methods:
   - `Print()` (value receiver)
   - `UpdateAge()` (pointer receiver)

2. We define an `Animal` interface with a single method:
   - `GetName() string`

3. We create two types that implement the `Animal` interface:
   - `Mammal`
   - `Bird`

4. We demonstrate polymorphism by passing different types to the same function `PrintAnimalName()`

## Try It Yourself

1. Add a new method to the `Student` struct
2. Create a new type that implements the `Animal` interface
3. Modify the `UpdateAge()` method to take a parameter for the number of years to add

## How to Run

Execute the following command from this directory:

```bash
go run main.go
```

## Further Reading

- [Methods and interfaces](https://go.dev/tour/methods/1)