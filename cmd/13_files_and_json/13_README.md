# Files and JSON in Go

This directory contains examples of working with files and JSON in Go.

## JSON Overview

JSON (JavaScript Object Notation) is a lightweight data interchange format that is easy for humans to read and write and easy for machines to parse and generate. Go provides built-in support for JSON encoding and decoding through the `encoding/json` package.

### Key Concepts

#### 1. Marshaling and Unmarshaling

- **Marshaling**: Converting Go data structures to JSON
  ```go
  data, err := json.Marshal(value)
  ```

- **Unmarshaling**: Converting JSON to Go data structures
  ```go
  err := json.Unmarshal(jsonData, &value)
  ```

#### 2. JSON Tags

Go uses struct tags to control how struct fields are encoded/decoded:

```go
type Person struct {
    Name      string `json:"name"`              // Maps to "name" in JSON
    Age       int    `json:"age"`               // Maps to "age" in JSON
    Address   string `json:"address,omitempty"` // Omitted if empty
    SSN       string `json:"-"`                 // Ignored in JSON
}
```

#### 3. Encoder and Decoder

For streaming JSON data (useful for large files):

```go
// Encoding
encoder := json.NewEncoder(writer)
encoder.Encode(value)

// Decoding
decoder := json.NewDecoder(reader)
decoder.Decode(&value)
```

## File Operations

Go provides several packages for file operations:

- `os`: Basic file operations (create, open, read, write)
- `io`: Interfaces for I/O operations
- `bufio`: Buffered I/O for better performance

### Common File Operations

```go
// Reading a file
data, err := os.ReadFile("filename.txt")

// Writing to a file
err := os.WriteFile("filename.txt", data, 0644)

// Opening a file
file, err := os.Open("filename.txt")     // Read-only
file, err := os.Create("filename.txt")   // Create new file
```

## Examples in this Directory

1. **Basic JSON Marshaling/Unmarshaling**: Converting between Go data structures and JSON strings
2. **Reading JSON from Files**: Loading JSON data from files into Go structures
3. **Writing JSON to Files**: Saving Go structures as JSON in files
4. **JSON Streaming with Encoder/Decoder**: Efficient handling of JSON data streams

## How to Run

Execute the following command from this directory:

```bash
go run .
```

## Further Reading

- [Go JSON Documentation](https://pkg.go.dev/encoding/json)
- [Go File Operations](https://pkg.go.dev/os#pkg-overview)
- [JSON Specification](https://www.json.org/)
- [Go by Example: JSON](https://gobyexample.com/json)
- [Working with Files in Go](https://gobyexample.com/reading-files)