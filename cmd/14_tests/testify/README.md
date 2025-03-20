# Testify Framework Examples

This directory demonstrates the use of the [Testify](https://github.com/stretchr/testify) testing framework in Go, which provides a more expressive and convenient way to write tests compared to the standard Go testing package.

## What is Testify?

Testify is a popular testing toolkit for Go that extends the standard testing package with additional functionality:

- **Assertion Functions**: Provides a rich set of assertion methods that make tests more readable and expressive
- **Mock Objects**: Helps create mock objects for testing
- **Suite Testing**: Supports test suite functionality for setup and teardown

## Files in this Directory

- `package_testify.go`: Contains a `Student` struct and a `FilterUniqueStudents` function that filters unique student names from a list
- `package_testify_test.go`: Contains tests for the `FilterUniqueStudents` function using the testify framework

## Example Usage

The tests in this directory demonstrate how to use the `assert` package from testify to simplify test assertions:

```go
// With testify
assert.Equal(t, expected, FilterUniqueStudents(students))

// Without testify (commented out in the test file)
actual := FilterUniqueStudents(students)
if !reflect.DeepEqual(expected, actual) {
    t.Fail()
}
```

## Running the Tests

To run the tests in this directory:

```bash
cd cmd/14_tests
go test ./testify
```

Or with verbose output:

```bash
go test -v ./testify
```