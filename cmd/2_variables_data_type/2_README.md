# Constants, variables and data types

## Constants

- Constants are declared using the `const` keyword.
- The `const` statement declares a list of constants; as in function argument lists, the type is last
- Constants can be character, string, boolean, or numeric values.

```go
// Syntax
const <constant_name> <data_type> = <value>
// Examples
const pi = 3.14 # float64
const str = "Hello World" # string
const role, grade int = 10, 2 # int
const isActive = true # bool
// Numeric Constants
const (
  Big = 1 << 100 # 1 << 100 is a bit shift operation
  Small = Big >> 99 # 1 << 100 >> 99 is a bit shift operation
)

```

## Variables

- Variable is the name given to a memory location to store a value of a spefic type
- Variables are declared using the `var` keyword.
- The `var` statement declares a list of variables; as in function argument lists, the type is last

```go
// Syntax
var variableName <data_type>

// Examples
var str = "Hello World" # string
var role, grade int = 10, 2 # int
var isActive = true # bool
f := 10 # short hand for var f int = 10
```

## Basic Data Types

- Golang has the following basic data types:
  - bool
  - string
  - int
  - int8
  - int16
  - int32 // alias for rune
  - int64
  - uint
  - uint8 // alias for byte
  - uint16
  - uint32
  - uint64
  - float32
  - float64
  - complex64
  - complex128
- The `int`, `uint`, and `uintpr` types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems.

## Zero values

- Zero values are the default values for variables.

```go
var i int # 0
var f float64 # 0
var b bool # false
var s string # ""
```

## Type conversion

- Type conversion is the process of converting a value of one type to another.
- The expression `T(v)` converts the value `v` to the type `T`.

```go
var i int = 13
var f float64 = float64(i) # 13
var u uint = uint(f) # 13

// OR
i := 13
f := float64(i) # 13
u := uint(f) # 13
```

## Type inference

- Type inference is the process of automatically determining the type of a variable.
- While declaring a variable, without specifying an explicit type (either short hand syntax or var = expression syntax), the variable is given a type by the value that is assigned to it.

```go
var i int
j := i // j is an int
v := 10 // v is an int; data type is inferred
```