package main

import (
	"net/http"

	"github.com/jaygaha/go_beginner/cmd/16_http/forms/handler"
)

func main() {
	// routes
	http.HandleFunc("/", handler.InlineFormHandler)
	http.HandleFunc("/inline-submit", handler.InlineFormSubmitHandler)

	// form validation
	http.HandleFunc("/form-validation", handler.ValidationFormHandler)
	http.HandleFunc("/form-validate", handler.ValidationFormSubmitHandler)

	// form validation
	http.HandleFunc("GET /contacts", handler.ContactValidationHandler)
	http.HandleFunc("POST /contacts", handler.ContactSubmitHandler)
	// inline handleFunc
	http.HandleFunc("/confirmation", func(w http.ResponseWriter, r *http.Request) {
		// Set content type to plain text
		w.Header().Set("Content-Type", "text/plain")
		// print message
		w.Write([]byte("Thank you for your submission!"))
		return
	})

	// start server
	http.ListenAndServe(":8080", nil)
}
