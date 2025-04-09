package handler

import (
	"log"
	"net/http"
	"regexp"
	"strings"
	"text/template"
)

type ValidationFormData struct {
	Username string
	Password string
	Error    string
}

// password regex with min characters 8 with number
var rxPassword = regexp.MustCompile(`^[A-Za-z\d]{8,}$`)

func ValidationFormHandler(w http.ResponseWriter, r *http.Request) {
	// render the form
	renderFile(w, "templates/form.tmpl", ValidationFormData{})
}

func ValidationFormSubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	data := ValidationFormData{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}

	// validation starts here
	// TrimSpace removes leading and trailing whitespace from a string
	if strings.TrimSpace(data.Username) == "" {
		data.Error = "Username is required"
		renderFile(w, "templates/form.tmpl", data)
		return
	}
	if len(data.Username) < 5 {
		data.Error = "Username must be at least 5 characters long"
		renderFile(w, "templates/form.tmpl", data)
		return
	}
	if !rxPassword.MatchString(data.Password) {
		data.Error = "Password must be at least 8 characters long and contain at least one number"
		renderFile(w, "templates/form.tmpl", data)
		return
	}

	// finally success
	renderFile(w, "templates/dashboard.tmpl", data)
}

func renderFile(w http.ResponseWriter, tmplFile string, data interface{}) {
	tmpl, err := template.ParseFiles(tmplFile)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Fatal("Encountered error: ", err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
	}
}
