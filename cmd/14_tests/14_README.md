# Go Testing Examples

This directory contains examples of testing in Go, demonstrating different testing approaches and frameworks.

- Put the test code in a file whose name ends with `_test.go`.
- Write a function `TestXXX` with a single argument of type *testing.T. The test framework runs each such function.
- To indicate a failed test, call a failure function such as `t.Errorf` or `t.Fail`.


## Why Test?

- To ensure that code is working as expected
- To ensure that code continues to work as expected when it is changed
- To ensure that code works as expected when it is refactored

## Types of Tests

- **Unit Tests**: Tests which test a small part of a codebase, usually mocking out external dependencies
- **Integration Tests**: Tests which include external resources like APIs or databases
- **End to End Tests**: Tests which test the entire application flow

## Running Tests

```bash
go test            # run all tests in the current directory
go test ./...      # run all tests in the package and its subpackages
go test -run Test  # run a specific test
go test -v         # verbose output
go test -cover     # coverage report
```

### Benchmarks

Benchmarks are used to measure the performance of a function. They are written in the same file as the function they are testing. The name of the benchmark function must start with `Benchmark`. The benchmark function takes a single argument of type `*testing.B`. The benchmark framework runs the function repeatedly until it reaches a stable state.

```bash
go test -bench .
```

## Standard Go Testing vs Testify

This directory demonstrates two approaches to testing in Go:

1. **Standard Go Testing**: Using the built-in `testing` package
   - Requires manual assertions and comparisons
   - Example: `calculator_test.go`

2. **Testify Framework**: Using the popular `github.com/stretchr/testify` package
   - Provides more expressive assertions
   - Simplifies test writing with helper functions
   - Example: `testify/package_testify_test.go`