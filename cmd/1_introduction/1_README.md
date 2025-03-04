# Introduction

Golang is a high-level, statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson.

## Six main points

1. a compiled language.
2. a statically typed language.
3. strongly typed language; operation depends on variable types
4. fast compilation time.
5. built-in concurrency with goroutines.
6. simplicity & garbage collection built into the design.

## Installtion

Installation on different platforms can be done from official [site](https://go.dev/doc/install).

## Structure

Golang is a package-based language.

1. Packages
   - A package is a collection of source files.
2. Modules
   - A module is a collection of packages.

## Initialization

Open terminal in desired project directory and run the following command.

```bash
go mod init <module-name>
```

`go.mod` file will be generated.

Sample:

```go
module go_tutorials

go 1.24
```

## Example

Create a file named `main.go` under `cmd/packageName` directory and add the following code.

```go
// need to define package name at top of the file
package main

import "fmt"

func main() {
	fmt.Println("Hello, World")
}
```

Every go file is a part of the package. In the above example, `main` is the package name. The `main` function is the entry point of the program. 

## Compilation

### Development

Run the following command:

```bash
go run <file-name>
```

### Production

Run the following command to compile the code.

```bash
go build -o <output-file-name> <file-name>
```

Above code produces an executable file. `main` file will be created.

#### Run

Run the following command to run the executable file.

```bash
./main
```

