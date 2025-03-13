# Go Packages and Modules

## Module Management

Go modules provide dependency management through `go.mod`:

```go
module github.com/jaygaha/golang-beginner/cmd/10_packages_and_modules

go 1.24.0
```

## Custom Package Examples

- **math**: Contains arithmetic operations
- **greet**: Handles greeting messages

## Usage Examples

```go
// Import local package
import "project_pkg/greet"
import "project_pkg/math"

// Import third-party package
import "rsc.io/quote"
```

## Key Concepts

1. Module = Collection of related packages
2. Semantic versioning (v1.2.3)
3. Minimal version selection

## How to Run

Execute the following command from this directory:

```bash
go run .
```

## Further Reading

- [Go Packages](https://go.dev/tour/basics/1)
- [Managing dependencies](https://go.dev/doc/modules/managing-dependencies)