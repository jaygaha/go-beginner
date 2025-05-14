package handlers

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func HelloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Gopher!")
}

// greet user
func GreetUserHandler(c echo.Context) error {
	name := c.Param("name") // get name from path which is /greet/:name

	// capitalize first letter
	if len(name) > 0 {
		name = strings.ToUpper(name[:1]) + name[1:] // name[:1] is the first letter of name and name[1:] is the rest of name
	} else {
		name = "Guest"
	}

	return c.String(http.StatusOK, "Hello, "+name+"!")
}

func SearchHandler(c echo.Context) error {
	q := c.QueryParam("q") // get query param q from url which is /search?q=
	language := c.QueryParam("lang")

	if language == "" {
		language = "en" // default language
	}

	message := "Searching for: '" + q + "' in " + language

	return c.String(http.StatusOK, message)
}
