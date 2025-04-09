package handler

import (
	"log"
	"net/http"
	"regexp"
	"strings"
	"text/template"
)

var rxEmail = regexp.MustCompile(".+@.+\\..+")

type ContactFormDataStrct struct {
	Email   string
	Message string
	Errors  map[string]string
}

func ContactValidationHandler(w http.ResponseWriter, r *http.Request) {
	// render the form
	renderTemplate(w, "templates/contact-form.tmpl", ContactFormDataStrct{})
}

func ContactSubmitHandler(w http.ResponseWriter, r *http.Request) {
	// Step 1: validate the form
	cfd := ContactFormDataStrct{
		Email:   r.FormValue("email"),
		Message: r.FormValue("message"),
	}
	if cfd.Validate() == false {
		renderTemplate(w, "templates/contact-form.tmpl", cfd)
		return
	}

	// Step 2: Do as you like for the content like saving to db, sending email, etc.
	// TODO

	// Step 3: Redirect to a thank you page
	http.Redirect(w, r, "/confirmation", http.StatusSeeOther)
}

func (cfd *ContactFormDataStrct) Validate() bool {
	// initialize the errors map
	cfd.Errors = make(map[string]string)

	match := rxEmail.MatchString(cfd.Email)
	if !match {
		cfd.Errors["Email"] = "Please provide a valid email address"
	}

	if strings.Trim(cfd.Message, " ") == "" {
		cfd.Errors["Message"] = "Please provide a message"
	}

	return len(cfd.Errors) == 0
}

func renderTemplate(w http.ResponseWriter, tmplFile string, data any) {
	tmpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Print(err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
	}
}
