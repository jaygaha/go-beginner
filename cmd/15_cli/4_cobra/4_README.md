# Toolbox - A Cobra CLI Learning Project

## Overview

Toolbox is a command-line application built with [Cobra](https://github.com/spf13/cobra), a powerful library for creating modern CLI applications in Go. This project serves as a learning example for understanding how to structure and implement CLI tools using the Cobra framework.

## What is Cobra?

Cobra is a library that provides a simple interface to create powerful modern CLI interfaces similar to git & go tools. Cobra is also an application that generates applications and command files to quickly develop a Cobra-based application.

Key features of Cobra include:
- Easy subcommand-based CLIs: `app server`, `app fetch`, etc.
- Fully POSIX-compliant flags (including short & long versions)
- Nested subcommands
- Global, local and cascading flags
- Intelligent suggestions (`app srver` → `Did you mean app server?`)
- Automatic help generation for commands and flags
- Automatic help flag recognition of `-h`, `--help`, etc.
- Automatically generated shell autocomplete for your application
- Automatically generated man pages for your application

## Project Structure

```
.
├── LICENSE
├── cmd/
│   ├── info/
│   │   ├── diskUsage.go   # Disk usage command implementation
│   │   └── info.go        # Info command definition
│   ├── net/
│   │   ├── net.go         # Net command definition
│   │   └── ping.go        # Ping command implementation
│   └── root.go            # Root command definition
├── go.mod                 # Go module definition
├── go.sum                 # Go module checksum
└── main.go                # Application entry point
```

## Installation

### Prerequisites

- Go 1.24 or higher

### Steps

1. Install the required dependencies:
   ```bash
   go install
   ```
   
2. Build the application
   ```bash
   go build -o toolbox
   ```

3. Run the application
   ```bash
   ./toolbox
   ```

## Usage

The Toolbox CLI provides several commands for different functionalities:

### Root Command

```bash
./toolbox
```

This will display the help information for the application.

### Info Command

The `info` command provides information about the toolbox:

```bash
./toolbox info # it will show the help information for the info command
```

#### Disk Usage Subcommand

Check the disk usage of the current directory:

```bash
./toolbox info diskUsage # it will show the disk usage of the current directory
```

### Net Command

The `net` command contains network-related utilities:

```bash
./toolbox net # it will show the help information for the net command
```

#### Ping Subcommand

Ping a URL and get the response status:

```bash
./toolbox net ping -u jaygaha.com.np # it will ping the domain and show the response http status
```

Options:
- `-u, --url`: URL to ping (required)

## Adding New Commands

To add a new command to the toolbox:

1. Create a new Go file for your command in the appropriate directory under `cmd/`
2. Define your command using the Cobra framework
3. Add your command to the appropriate parent command in the `init()` function

Example:

```go
package mycommand

import (
    "github.com/spf13/cobra"
)

var MyCmd = &cobra.Command{
    Use:   "mycommand",
    Short: "A brief description of your command",
    Long:  `A longer description of your command`,
    Run: func(cmd *cobra.Command, args []string) {
        // Your command logic here
    },
}

func init() {
    // Add flags if needed
    MyCmd.Flags().StringVarP(&someVar, "flag", "f", "", "Description of flag")
}
```

Then add your command to the appropriate parent command:

```go
// In the parent command's init function
parentCmd.AddCommand(mycommand.MyCmd)
```

## Dependencies

- [github.com/spf13/cobra](https://github.com/spf13/cobra): The main CLI framework
- [github.com/ricochet2200/go-disk-usage/du](https://github.com/ricochet2200/go-disk-usage): Used for disk usage information

## Learning Resources

- [Cobra GitHub Repository](https://github.com/spf13/cobra)
- [Cobra Documentation](https://cobra.dev/)
- [Go Command-line Interfaces (CLIs)](https://go.dev/solutions/clis)
- [How to write beautiful Golang CLI](https://www.youtube.com/watch?v=SSRIn5DAmyw)