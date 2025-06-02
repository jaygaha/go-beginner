package main

import (
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

func main() {
	// File server for frontend/public
	publicFS := http.FileServer(http.Dir("frontend/public"))
	// File server for frontend/src
	srcFS := http.FileServer(http.Dir("frontend/src"))

	// Custom handler to route requests
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Serve /src/ requests from frontend/src
		if strings.HasPrefix(r.URL.Path, "/src/") {
			if filepath.Ext(r.URL.Path) == ".js" {
				w.Header().Set("Content-Type", "application/javascript")
			}
			srcFS.ServeHTTP(w, r)
			return
		}
		// Serve other requests from frontend/public
		publicFS.ServeHTTP(w, r)
	})

	http.Handle("/", handler)
	log.Println("Frontend server running on http://localhost:3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
