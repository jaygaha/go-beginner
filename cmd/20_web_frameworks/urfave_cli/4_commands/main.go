package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

/*
Commands:
-> commands are used to group similar actions together.
-> commands can be added via an array that is attached to the `Commands` field of the `cli.App` struct.
-> in the arguments, the commands are passed in the order they are defined.
*/

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "hello",
				Usage: "say hello in your language",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "name",
						Value:   "Gopher",
						Usage:   "Name of the Gopher",
						Aliases: []string{"n"}, // aliases are used to shorten the flag name using `-` or `--`
					},
					&cli.StringFlag{
						Name:    "lang",
						Value:   "english",
						Usage:   "Language to greet in",
						Aliases: []string{"l"},
					},
				},
				Action: func(c *cli.Context) error {
					name := c.String("name") // name is the value of the name flag
					lang := c.String("lang") // lang is the value of the lang flag

					var greeting string

					switch lang {
					case "nepali":
					case "ne":
					case "hindi":
					case "hi":
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
			},
			{
				Name:  "bye",
				Usage: "say bye in your language",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "name",
						Value:   "Gopher",
						Usage:   "Name of the Gopher",
						Aliases: []string{"n"},
					},
					&cli.StringFlag{
						Name:    "lang",
						Value:   "english",
						Usage:   "Language to greet in",
						Aliases: []string{"l"},
					},
				},
				Action: func(c *cli.Context) error {
					name := c.String("name")
					lang := c.String("lang")
					var greeting string

					switch lang {
					case "nepali":
					case "ne":
						greeting = "फेरि भेटौँला" // text may render in the terminal as gibberish
					case "hindi":
					case "hi":
						greeting = "फिर मिलेंगे"
					case "spanish":
					case "sp":
						greeting = "Nos vemos de nuevo"
					case "french":
					case "fr":
						greeting = "À bientôt"
					case "japanese":
					case "jp":
						greeting = "また会いましょう"
					default:
						greeting = "See you again"
					}

					fmt.Printf("%s %s!\n", greeting, name)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

	/*
		Output:
			$ go run cmd/20_web_frameworks/urfave_cli/4_commands/main.go hello -n=Gopher -l=spanish
			$ Hello Gopher!
			$ go run cmd/20_web_frameworks/urfave_cli/4_commands/main.go bye -n=Gopher -l=fr
			$ See you again Gopher!
	*/
}
