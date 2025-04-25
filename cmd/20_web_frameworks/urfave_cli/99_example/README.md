# urfave/cli Example Application

This example demonstrates how to build command-line applications in Go using the [urfave/cli](https://github.com/urfave/cli) library.

## Overview

This application showcases two common CLI commands:

1. **Kill Command**: Terminate processes by PID or name
2. **Volumes Command**: List all mounted volumes with detailed information

## Features Demonstrated

- Creating a CLI application with multiple commands
- Adding command aliases for convenience
- Defining and validating command flags
- Handling command arguments
- Implementing command actions
- Error handling and validation
- Formatting and displaying output (JSON)

## Installation

1. Clone the repository
2. Install dependencies:

```bash
go mod download
```

## Usage

### Build the application:

```bash
go build -o cli-demo main.go
```

### Run the application:

```bash
# Show help
./cli-demo --help

# Kill a process by PID
./cli-demo kill --pid 1234
# or using alias
./cli-demo k -p 1234

# Kill a process by name
./cli-demo kill --name chrome
# or using alias
./cli-demo k -n chrome

# List all mounted volumes
./cli-demo volumes
# or using alias
./cli-demo v
```

## Code Structure

- **Main Application** : Sets up the CLI app with commands and flags
- **Kill Command** : Demonstrates process management with flag validation
- **Volumes Command** : Shows how to gather system information and format as JSON

## Key Concepts

- **Commands** : Logical groupings of functionality
- **Flags** : Options that modify command behavior
- **Actions** : Functions that execute when commands are invoked
- **Validation** : Ensuring proper command usage

## Practice

Feel free to modify and extend this example to suit your needs.

## Resources
- [urfave/cli Documentation](https://cli.urfave.org/v2/getting-started/)
- [gopsutil Documentation](https://github.com/shirou/gopsutil)
- [How to Build Cli in Go](https://hackajob.com/talent/blog/how-to-build-cli-in-go)
