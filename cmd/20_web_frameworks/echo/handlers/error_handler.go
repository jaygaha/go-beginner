package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	message := "Internal Server Error"

	// Check if it's an HTTP error
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code

		// we can customize messages based on he.Code or he.Message
		if customMessage, ok := he.Message.(string); ok {
			message = customMessage
		} else {
			message = http.StatusText(code) // Use the default message for the given status code
		}
	} else {
		// For other types of errors, you can log them or handle them differently
		c.Logger().Error("Unknow error", err)
	}

	// Don't send the error response if response has already been committed
	// this is to prevent double responses
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead { // Issue #608 HEAD method
			err = c.NoContent(code)
		} else {
			// Handle custom error response using template
			err = c.Render(code, "error.html", map[string]any{
				"Code":    code,
				"Message": message,
			})

			// or you can use JSON
			// err = c.JSON(code, map[string]string{"error": message, "code": fmt.Sprintf("%d", code)})
		}

		if err != nil {
			c.Logger().Error(err)
		}
	}
}

func InternalServerErrorHandler(c echo.Context) error {
	return echo.NewHTTPError(http.StatusTeapot, "Internal Server Error")
}

func PanicHandler(c echo.Context) error {
	panic("A panic occurred")
}
