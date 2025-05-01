package main

/*
Testing:
-> Go standard library provides a testing package (testing) and http/httptest package
 	- it allows you to test the behaviour of fiber routes and middleware
*/

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func NewServer() *fiber.App {
	app := fiber.New(fiber.Config{
		BodyLimit:       10 * 1024 * 1024, // 10MB
		ReadBufferSize:  8192,             // 8KB - increase the buffer size for reading request headers
		WriteBufferSize: 8192,             // 8KB - increase the buffer size for writing response headers
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			message := "Internal Server Error"
			// check if the error is a fiber error
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
				message = e.Message
			}

			// log the error
			fmt.Printf("Error: %s code: %v\n", err.Error(), code)
			// return the error
			return c.Status(code).JSON(fiber.Map{
				"message": message,
			})
		},
		// log the error,
	})

	// root endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Fiber!")
	})

	// error endpoint
	app.Get("/500", func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	})

	return app
}

func main() {
	app := NewServer()
	app.Listen(":3300")
}
