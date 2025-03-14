# Basic Miscellaneous Go Concepts

This section covers various fundamental Go concepts that are essential for Go programming.

## Type Casting

- Process of explicitly converting variables from one type to another
- Go doesn't support automatic type conversion

**Syntax:**

```go
v := typeName(otherTypeValue)
```

### Common Type Conversions:

- String to Number: `strconv.Atoi()`, `strconv.ParseFloat()`
- Number to String: `strconv.Itoa()`
- String to Boolean: `strconv.ParseBool()`
- Float to Int: `int(floatValue)`
- String to Bytes: `[]byte(stringValue)`

## Type Inference

- Automatic type determination by Go compiler
- Uses the value to determine variable type
- Works with variables, constants, and function parameters
- Not applicable to function return values

**Example:**

```go
var x = 42        // int
y := 3.14         // float64
name := "Go"      // string
```

## Type Assertions

- Process of extracting the concrete type value from an interface
- Used to access underlying type of an interface variable

**Syntax:**

```go
value := interfaceVar.(Type)
// Safe checking
if value, ok := interfaceVar.(Type); ok {
    // Use value
}
```

## Type Switches

- Compare types of interface values
- Similar to regular switch but cases specify types

**Syntax:**

```go
switch v := i.(type) {
case int:
    // v is an int
case string:
    // v is a string
default:
    // v is another type
}
```

## Maps and Make

### Maps

- Unordered collection of key-value pairs
- Keys must be unique
- Both keys and values can be of any type

**Basic Operations:**

```go
// Creation
m := make(map[KeyType]ValueType)

// Assignment
m[key] = value

// Deletion
delete(m, key)

// Check existence
value, exists := m[key]
```

### Make Function
- Creates maps, slices, and channels
- Initializes data structure with proper memory allocation

**Syntax:**

```go
make(map[KeyType]ValueType)           // Create map
make(map[KeyType]ValueType, capacity) // Create map with initial capacity
```

## How to Run

Execute the following command from this directory:

```bash
go run .
```

## Further Reading

- [Go Docs: Type Casting](https://golangdocs.com/type-casting-in-golang)
- [Go Variables: Type Inference](https://www.callicoder.com/golang-variables-zero-values-type-inference/#type-inference)
- [Types Assertions](https://go.dev/tour/methods/15)
- [Effective Go: Maps](https://go.dev/doc/effective_go#maps)
- [Effective Go: Allocation with make](https://go.dev/doc/effective_go#allocation_make)