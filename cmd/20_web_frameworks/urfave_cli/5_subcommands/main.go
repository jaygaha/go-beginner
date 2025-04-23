package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

/*
Subcommands:

-> subcommands are used to group commands together.
-> subcommands are added to the `Commands` field of the `cli.App` struct.
-> subcommands are added as an array of `cli.Command` structs.

For example: `todo add`, `todo complete`, `git update` are subcommands of any group.
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
			// Subcommand example
			{
				Name:  "coffee",
				Usage: "type of coffee to make like black, latte or cappuccino",
				Subcommands: []*cli.Command{
					{
						Name:  "black",
						Usage: "make black coffee",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "bean",
								Value:   "arabica",
								Usage:   "type of bean to use like arabica, robusta or liberica",
								Aliases: []string{"b"},
							},
							&cli.StringFlag{
								Name:    "size",
								Value:   "medium",
								Usage:   "size of the coffee like small, medium or large",
								Aliases: []string{"s"},
							},
						},
						Action: func(c *cli.Context) error {
							fmt.Println("Making black coffee")

							// grab flags
							bean := c.String("bean")
							size := c.String("size")

							// default values
							if bean == "" {
								bean = "arabica"
							}
							if size == "" {
								size = "medium"
							}

							fmt.Printf("Using %s bean and %s size\n", bean, size)

							return nil
						},
					},
					{
						Name:  "latte",
						Usage: "make latte coffee",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "bean",
								Value:   "arabica",
								Usage:   "type of bean to use like arabica, robusta or liberica",
								Aliases: []string{"b"},
							},
							&cli.StringFlag{
								Name:    "size",
								Value:   "medium",
								Usage:   "size of the coffee like small, medium or large",
								Aliases: []string{"s"},
							},
							&cli.StringFlag{
								Name:    "milk",
								Value:   "whole",
								Usage:   "type of milk to use like whole, skim or almond",
								Aliases: []string{"m"},
							},
						},
						Action: func(c *cli.Context) error {
							fmt.Println("Making latte coffee")

							// grab flags
							bean := c.String("bean")
							size := c.String("size")
							milk := c.String("milk")

							// default values
							if bean == "" {
								bean = "arabica"
							}
							if size == "" {
								size = "medium"
							}
							if milk == "" {
								milk = "whole"
							}

							fmt.Printf("Using %s bean, %s size and %s milk\n", bean, size, milk)

							return nil
						},
					},
					{
						Name:  "cappuccino",
						Usage: "make cappuccino coffee",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "bean",
								Value:   "arabica",
								Usage:   "type of bean to use like arabica, robusta or liberica",
								Aliases: []string{"b"},
							},
							&cli.StringFlag{
								Name:    "size",
								Value:   "medium",
								Usage:   "size of the coffee like small, medium or large",
								Aliases: []string{"s"},
							},
						},
						Action: func(c *cli.Context) error {
							fmt.Println("Making cappuccino coffee")

							// grab flags
							bean := c.String("bean")
							size := c.String("size")
							// default values
							if bean == "" {
								bean = "arabica"
							}
							if size == "" {
								size = "medium"
							}

							fmt.Printf("Using %s bean and %s size\n", bean, size)
							return nil
						},
					},
				},
				// help command for coffee subcommand
				Action: func(c *cli.Context) error {
					fmt.Println("'coffee' command is used to make coffee")
					fmt.Println("Use 'coffee black' to make black coffee")
					fmt.Println("Use 'coffee latte' to make latte coffee")
					fmt.Println("Use 'coffee cappuccino' to make cappuccino coffee")
					fmt.Println("Use 'coffee --help' to see more information")
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
		$ go run main.go coffee black -b robusta -s large
		$ go run main.go coffee help

		Or build the binary and run it:
		$ go build
		$ ./cli coffee black -b robusta -s large
	*/
}
