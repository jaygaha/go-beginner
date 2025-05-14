package handlers

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// User struct for JSON binding
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Handler for creating a new user (receives JSON)
func CreateUserHandler(c echo.Context) error {
	u := new(User) // u is a pointer to User struct which is empty

	// bind the incoming JSON request to the User struct
	// if there is an error, return it
	if err := c.Bind(u); err != nil {
		// If binding fails, return an error response
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload: " + err.Error()})
	}

	// DO something with the user data like validation, saving to DB, etc.

	message := fmt.Sprintf("User created: Name: %s, Age: %d", u.Name, u.Age)

	return c.JSON(http.StatusCreated, map[string]string{"message": message})
}

// TemplateRenderer is a custom renderer for Echo framework
type TemplateRenderer struct {
	Templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

func UserProfileHandler(c echo.Context) error {
	data := map[string]any{
		"Title":       "Go Beginner",
		"Name":        "Jay Gaha",
		"Age":         24,
		"CurrentTime": time.Now().Format(time.RFC1123),
	}

	// render the profile.html template with the data
	return c.Render(http.StatusOK, "profile.html", data)
}
