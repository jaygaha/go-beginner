package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

/*
Flags:
-> flags are used to pass arguments to the application.
-> flags can¥n be  added via an array that is attached to the `Flags` field of the `cli.App` struct.
-> in the arguments, the arguments are passed in the order they are defined.
*/

func main() {
	app := &cli.App{
		// flags are added to the `Flags` field of the `cli.App` struct
		Flags: []cli.Flag{
			// first flag is the name of the user
			&cli.StringFlag{
				Name:  "name",
				Value: "Gopher",
				Usage: "Name of the Gopher",
			},
			// lang flag is the language to greet in
			&cli.StringFlag{
				Name:  "lang",
				Value: "english",
				Usage: "Language to greet in",
			},
		},
		// action is the function that is called when the application is run
		Action: func(c *cli.Context) error {
			name := c.String("name") // name is the value of the name flag
			lang := c.String("lang") // lang is the value of the lang flag

			var greeting string

			switch lang {
			case "nepali":
			case "ne":
				greeting = "नमस्ते"
			case "spanish":
			case "sp":
				greeting = "Hola"
			case "french":
			case "fr":
				greeting = "Bonjour"
			case "japanese":
			case "jp":
				greeting = "こんにちは"
			default:
				greeting = "Hello"
			}

			fmt.Printf("%s %s!\n", greeting, name)
			return nil
		},
		// multiple actions can be added to the `Action` field of the `cli.App` struct
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

	/*

		Usage:
		```
		$ go build -o hello
		$ ./hello --name=Gopher --lang=ne
		नमस्ते Gopher!
		```
	*/
}
