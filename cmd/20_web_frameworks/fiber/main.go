package main

import (
	"github.com/gofiber/fiber/v2"
)

/*
Fiber
-> very popular lightweight and fast web framework
-> inspired by expressjs

Features:
-> High speed performance
	- built on top of Fasthttp offering zero memory allocation and optimized HTTP router
	- uses a custom routing algorithm to reduce routing overhead
-> Simple and intuitive API
  	- inspired by expressjs, which is intuitive and easy to use
-> Ecosystem
	- fiber offers a number of official and third-party middlewares, cors, logging, error handling, etc.
-> Flexible
	- combine with Go's standard library and third-party libraries
*/

func main() {
	// initializing the app with custom config to handle large headers
	app := fiber.New(fiber.Config{
		BodyLimit:       10 * 1024 * 1024, // 10MB
		ReadBufferSize:  8192,             // 8KB - increase the buffer size for reading request headers
		WriteBufferSize: 8192,             // 8KB - increase the buffer size for writing response headers
	})

	// defining a route
	// c is the context object
	app.Get("/", func(c *fiber.Ctx) error {
		c.SendString("Hello, Gophers!") // SendString is a method of the context object that sends a string as a response
		return nil
	})

	// start the server
	err := app.Listen(":3300")
	if err != nil {
		panic(err)
	}
}
