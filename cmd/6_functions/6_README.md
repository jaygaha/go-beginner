# Functions in Go

This section covers various aspects of functions in Go, demonstrating the language's powerful function capabilities.

## Basic Concepts

### Function Definition

In Go, a function is a block of code that performs a specific task. Functions are defined using the `func` keyword, followed by the function name, parameters, and return type.

```go
func functionName(parameter1 type, parameter2 type) returnType {
    // function body
    return value
}
```

### Simple Functions

A basic function without parameters or return values:

```go
func sayHelloWorld() {
    fmt.Println("Hello Go-phers!!")
}
```

### Functions with Arguments

- Go requires explicit returns, i.e. it won't automatically return the value of the last expression.

Functions can accept parameters:

```go
func sayHello(name string) {
    fmt.Println("Hello", name)
}
```

### Functions with Return Values

Functions can return values:

```go
func multiplication(num1, num2 int) int {
    return num1 * num2
}
```

Note: Go requires explicit returns; it won't automatically return the value of the last expression.

## Advanced Function Features

### Multiple Return Values

One of Go's distinctive features is the ability for functions to return multiple values:

```go
func intDivision(numerator, denominator int) (int, int, error) {
    // function body
    return result, remainder, err
}
```

Note: This is particularly useful for returning both results and error status.

### Named Return Values

In Go, you can name the return values in the function signature:

```go
func divide(x, y float64) (result float64, err error) {
    // result and err are already declared
    if y == 0 {
        err = errors.New("division by zero")
        return // returns the current values of result and err
    }
    result = x / y
    return // returns result and nil for err
}
```

### Variadic Functions

Variadic functions can accept a variable number of arguments:

```go
func sum(nums ...int) {
    // nums is a slice of int
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}
```

You can call variadic functions with individual arguments or by "unpacking" a slice with the `...` operator:

```go
sum(1, 2, 3)       // Individual arguments
nums := []int{1, 2, 3, 4}
sum(nums...)        // Unpacking a slice
```

### Defer Statement

The `defer` statement delays the execution of a function until the surrounding function returns:

```go
func deferred() {
    defer fmt.Println("a defer function") // Executes last
    fmt.Println("a non-defer function")  // Executes first
}
```

Note: Defer is commonly used for cleanup operations like closing files or releasing resources.

### Closures

Go supports closures, which are anonymous functions that can access variables from the enclosing function:

```go
func closureSum() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}
```

Note: Each closure maintains its own reference to the variables it captures.

### Recursion

Go supports recursive functions - functions that call themselves:

```go
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}
```

### Anonymous Functions

Go allows you to define functions without naming them:

```go
var fib func(n int) int
fib = func(n int) int {
    if n <= 1 {
        return n
    }
    return fib(n-1) + fib(n-2)
}
```

## Best Practices

1. **Keep functions focused**: Each function should perform a single, well-defined task.
2. **Use meaningful names**: Function names should clearly indicate what the function does.
3. **Handle errors**: Return errors explicitly rather than using panic.
4. **Document your functions**: Use comments to explain what the function does, especially for exported functions.
5. **Limit function length**: If a function becomes too long, consider breaking it into smaller functions.

## Running the Example

1. Navigate to this directory
2. Read the comments in `main.go` for explanations
3. Run the program:

```bash
cd cmd/6_functions
go run main.go
```

Happy coding 🚀