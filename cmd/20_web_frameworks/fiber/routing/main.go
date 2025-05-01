package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

/*
Routing:
-> set up intuitive and easy to use routes for your application
-> define routes using HTTP methods such as GET, POST, PUT, DELETE, etc.
-> define routes using path parameters, query parameters, and request body
-> define routes using custom middlewares

// HTTP methods
func (app *App) Get(path string, handlers ...Handler) Router
func (app *App) Head(path string, handlers ...Handler) Router
func (app *App) Post(path string, handlers ...Handler) Router
func (app *App) Put(path string, handlers ...Handler) Router
func (app *App) Delete(path string, handlers ...Handler) Router
func (app *App) Connect(path string, handlers ...Handler) Router
func (app *App) Options(path string, handlers ...Handler) Router
func (app *App) Trace(path string, handlers ...Handler) Router
func (app *App) Patch(path string, handlers ...Handler) Router

// Add allows you to specifiy a method as value
func (app *App) Add(method, path string, handlers ...Handler) Router

// All will register the route on all HTTP methods
// Almost the same as app.Use but not bound to prefixes
func (app *App) All(path string, handlers ...Handler) Router
*/

func main() {
	// initializing the app
	app := fiber.New(fiber.Config{
		BodyLimit:       10 * 1024 * 1024, // 10MB
		ReadBufferSize:  8192,             // 8KB - increase the buffer size for reading request headers
		WriteBufferSize: 8192,             // 8KB - increase the buffer size for writing response headers
		// Custom error handler
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
	})

	// dynamic route
	app.Get("/greet/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		// c.SendString("Hello, " + name + "!")
		return c.JSON(fiber.Map{
			"message": "Hello, " + name + "!",
		})
	})

	// query parameters
	app.Get("/search", func(c *fiber.Ctx) error {
		query := c.Query("q")
		return c.JSON(fiber.Map{
			"message": "Searching for: " + query,
		})
	})

	// request body
	app.Post("/users", func(c *fiber.Ctx) error {
		type User struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		var user User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request body",
			})
		}

		return c.JSON(fiber.Map{
			"message": "User created successfully",
			"data":    user,
		})
	})

	/*
		middleware
		-> a function that is executed before the request is handled
		-> records information about the request and response
		-> authentication and authorization
		-> error handling
		-> data processing
	*/
	// global middleware
	app.Use(requestDetailsMiddleware)

	// route specific middleware
	middleware := func(c *fiber.Ctx) error {
		c.Set("X-Powered-By", "Go-Fiber")
		return c.Next()
	}

	// grouping routes
	api := app.Group("/api", middleware) // /api
	v1 := api.Group("/v1", middleware)   // /api/v1

	// v1 routes
	v1.Get("/posts", func(c *fiber.Ctx) error { // middleware for /api/v1
		return c.JSON(fiber.Map{"message": "Posts"})
	})

	v1.Get("/posts/:id", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Post details for ID: " + c.Params("id")})
	})

	/*
		Error Handling
		-> catch and handle errors; return appropriate error messages and status codes
	*/
	v1.Post("/posts", func(c *fiber.Ctx) error {
		var post struct {
			Title string `json:"title"`
			Body  string `json:"body"`
		}

		// parse the request body
		if err := c.BodyParser(&post); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request body",
			})
		}

		// validate the request body
		// post := req{}
		if post.Title == "" || post.Body == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Required fields are missing",
			})
		}

		// Do DB operations
		// ...

		// return the response
		return c.JSON(fiber.Map{
			"message": "Post created successfully",
		})
	})

	// error handling
	app.Get("/error", func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusNotFound, "Not found")
	})

	// start the server
	err := app.Listen(":3300")
	if err != nil {
		panic(err)
	}
}

func requestDetailsMiddleware(c *fiber.Ctx) error {
	start := time.Now() // start time

	// run the next handler
	err := c.Next()

	// calculate the duration
	duration := time.Since(start)

	// log the request details
	fmt.Printf("Request: %s %s %s %s %d %s\n", c.IP(), c.Method(), c.Path(), c.Protocol(), c.Response().StatusCode(), duration)

	return err
}
