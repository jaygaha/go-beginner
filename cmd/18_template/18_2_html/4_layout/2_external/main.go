package main

import (
	"html/template"
	"os"
)

type PageData struct {
	Title   string // Page title
	Message string // Message to display
	Year    int    // copyright year
}

/*
	{{template "title" .}}: A placeholder that expects a "title" block to be defined elsewhere. The . passes the current data context.
	{{template "content" .}}: A placeholder for the main content block, also using the current data context.
	{{template "footer".}}: A placeholder for the footer block, also using the current data context.
*/
// Base layout template
const contentTmpl = `
{{define "title"}}{{.Title}}{{end}} <!-- Define title block -->

{{define "content"}} <!-- Define content block -->
    <h1>Welcome</h1>
	<h2>{{.Title}}</h2>
    <p>{{.Message}}</p>
{{end}}

{{define "footer"}} <!-- Define footer block -->
    <p>Copyright &copy; {{.Year}}</p>
{{end}}
`

func main() {
	// parse the external template
	// template.ParseFiles reads the template from a file
	t, err := template.ParseFiles("layout.tmpl")
	if err != nil {
		panic(err)
	}

	// Parse the content template and add it to the same template set
	// Must() ensures parsing succeeds or panics
	t = template.Must(t.Parse(contentTmpl))

	// createdata to pass to the template
	data := PageData{
		Title:   "Home Page",
		Message: "Welcome to my website",
		Year:    2025,
	}

	// Execute the template with the data
	// output goes to the os.Stdout (console)
	err = t.ExecuteTemplate(os.Stdout, "layout.tmpl", data)
	// err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
