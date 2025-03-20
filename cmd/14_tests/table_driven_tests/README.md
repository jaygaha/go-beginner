# Table-Driven Tests in Go

Table-driven tests are a powerful testing pattern in Go that allows you to write more maintainable and concise test cases. Instead of writing multiple similar test functions, you can combine related test cases into a single test function using a slice of test cases.

## Benefits

- **Reduced Code Duplication**: Eliminates repetitive test code by consolidating similar test cases
- **Better Maintainability**: Makes it easier to add new test cases without writing new test functions
- **Improved Readability**: Clearly shows the relationship between inputs and expected outputs
- **Consistent Error Handling**: Uses the same validation logic across all test cases

## Example

Let's look at how to convert traditional test cases into a table-driven test:

### Traditional Test Cases

```go
func TestSplitStrings(t *testing.T) {
    got := SplitStrings("a:b:c", ":")
    want := []string{"a", "b", "c"}

    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v want %v", got, want)
    }
}

func TestSplitStringsWrongSep(t *testing.T) {
    got := SplitStrings("a:b:c", "/")
    want := []string{"a:b:c"}

    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v want %v", got, want)
    }
}
```

### Table-Driven Test Version

```go
func TestSplitStringsTBT(t *testing.T) {
    tests := []struct {
        name  string
        input string
        sep   string
        want  []string
    }{
        {name: "simple", input: "a/b/c", sep: "/", want: []string{"a", "b", "c"}},
        {name: "wrong sep", input: "a/b/c", sep: ",", want: []string{"a/b/c"}},
        {name: "no sep", input: "abc", sep: "/", want: []string{"abc"}},
    }

    for _, test := range tests {
        got := SplitStrings(test.input, test.sep)
        if !reflect.DeepEqual(got, test.want) {
            t.Fatalf("got %v want %v", got, test.want)
        }
    }
}
```

## Best Practices

1. **Descriptive Test Names**: Give each test case a clear, descriptive name that indicates what's being tested
2. **Anonymous Struct**: Use an anonymous struct to define the test cases structure locally within the test function
3. **Complete Test Cases**: Include edge cases, error cases, and boundary conditions
4. **Clear Input/Output**: Make the relationship between inputs and expected outputs obvious
5. **Consistent Error Messages**: Use consistent error message formatting across all test cases

## When to Use Table-Driven Tests

Table-driven tests are particularly useful when:

- You have multiple test cases with similar structure
- Test cases differ only in input data and expected output
- You need to test various edge cases and error conditions
- You want to make it easy to add new test cases in the future

## Running the Tests

To run the tests in this directory:

```bash
go test -v
```

Use the `-v` flag to see the details of each test case execution.