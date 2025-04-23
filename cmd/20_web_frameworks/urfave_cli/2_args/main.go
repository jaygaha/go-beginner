package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

/*
	Arguments:
	-> arguments lookup by calling the `Args` function on `cli.Context`

	Usage:
	```
	$ go build -o hello
	$ ./hello world
	Hello world!
	```

*/

func main() {
	app := &cli.App{
		Name:  "hello",
		Usage: "Greet someone",
		Action: func(c *cli.Context) error {
			fmt.Printf("Hello %s!\n", c.Args().Get(0))
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
