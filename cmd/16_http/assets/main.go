package main

import (
	"fmt"
	"net/http"
)

/*
Assets & Files
 - In Go, how to serve static files like CSS, JS, images, etc. from a specific directory
*/

func main() {
	// FileServer: serves files from the given directory
	// Useful for serving static files like CSS, JS, images, etc.
	// here assets directory is the root directory for the files
	fs := http.FileServer(http.Dir("assets"))

	// StripPrefix: removes the given prefix from the request URL
	// Useful for serving files from a subdirectory
	// assets directory is overridden by the static directory
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// start the server
	err := http.ListenAndServe(":8800", nil)

	if err != nil {
		fmt.Println(err)
	}

	/*
		Demo:
		- run web server
		- curl to the assets directory
		 - $ curl -s localhost:8800/static/css/style.css
		 - $ curl -s localhost:8800/static/images/golang.png
	*/
}
