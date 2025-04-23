package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	// Basic implementation
	(&cli.App{}).Run(os.Args) // It works, but nothing is implemented, so a help message appears.

	// Check directories for more examples.
}
