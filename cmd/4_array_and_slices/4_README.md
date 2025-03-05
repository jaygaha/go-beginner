# Arrays, Slices, and Maps

This example demonstrates the fundamental concepts of arrays, slices, and maps in Go programming language.

## Arrays

Arrays in Go have the following characteristics:
- Fixed length
- Same type elements
- Cannot change length after declaration
- Indexable
- Contiguous memory allocation

### Array Declaration and Usage
```go
// Fixed length array
var ages = [3]int{12, 13, 14}

// Array with inferred length
names := [4]string{"a", "b", "e", "d"}
```

## Slices

Slices are built on top of arrays but provide more flexibility:
- Dynamic length
- Can be extended using append
- Reference type

### Basic Slice Operations
```go
// Creating a slice
var scores = []int{100, 200, 300}

// Modifying elements
scores[2] = 500

// Appending elements
scores = append(scores, 400)
```

### Slice Ranges

Slices support different range operations:
```go
rangeOne := names[1:3]   // Elements from index 1 to 2
rangeTwo := names[2:]    // Elements from index 2 to end
rangeThree := names[:2]  // Elements from start to index 1
```

### Creating Slices with make

You can create slices with predefined length using make:
```go
var sliceMake []int32 = make([]int32, 3) // Creates [0 0 0]
```

## Maps

Maps are key-value pairs in Go:
- Dynamic key-value storage
- Keys must be unique
- Fast lookups

### Map Declaration and Usage
```go
var studentScores map[string]uint8 = map[string]uint8{
    "a": 77,
    "b": 30,
    "c": 56,
}

// Accessing values
score := studentScores["b"]    // Gets value: 30

// Checking if key exists
value, exists := studentScores["c"]
```

## Running the Example

To run this example:
```bash
go run main.go
```