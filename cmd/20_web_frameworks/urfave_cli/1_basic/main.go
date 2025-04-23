package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "hello",
		Usage: "fight the loneliness!",
		Action: func(c *cli.Context) error {
			fmt.Println("Hello Gopher!")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
