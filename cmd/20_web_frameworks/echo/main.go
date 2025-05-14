package main

import (
	"html/template"

	"github.com/jaygaha/go-beginner/cmd/20_web_frameworks/echo/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	echo := echo.New()

	// -- custom error handler --
	echo.HTTPErrorHandler = handlers.CustomHTTPErrorHandler

	// -- middleware --
	// Logger middleware (logs requests to the console); it will register globally
	echo.Use(middleware.Logger())
	// Recover middleware (recovers from panics and logs them to the console)
	echo.Use(middleware.Recover())

	// init templates
	renderer := &handlers.TemplateRenderer{
		Templates: template.Must(template.ParseGlob("views/*.html")), // parse all html files in views folder
	}
	// set renderer
	echo.Renderer = renderer // register the custom renderer

	// routes
	echo.GET("/", handlers.HelloHandler)

	// greet user
	echo.GET("/greet/:name", handlers.GreetUserHandler)
	// route for when no name is provided; optional but good practice
	echo.GET("/greet", handlers.GreetUserHandler)

	// Query params
	echo.GET("search", handlers.SearchHandler)

	// Handling POST Requests and JSON Data
	echo.POST("/users", handlers.CreateUserHandler)

	//  Serving Static Files (CSS, JS, Images)
	// Serve static files from the "static" directory
	// The "/static" path in the URL will map to the "static" folder on disk
	echo.Static("/static", "static") // (url_prefix, filesystem_path)

	// Using Templates (HTML)
	echo.GET("/profile", handlers.UserProfileHandler)

	// Grouping Routes
	// Create a new group for admin routes
	adminGroup := echo.Group("/admin") // prefix all admin routes with /admin
	{
		adminGroup.GET("", handlers.AdminDashboardHandler)

		// admin/settings
		adminGroup.GET("/settings", handlers.AdminSettingsHandler)
	}

	// Handle errors
	echo.GET("/500", handlers.InternalServerErrorHandler)
	// panic handler
	echo.GET("/panic", handlers.PanicHandler)

	// start server
	echo.Logger.Fatal(echo.Start(":8800"))
}
