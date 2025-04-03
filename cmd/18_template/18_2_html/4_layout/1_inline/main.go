package main

import (
	"html/template"
	"os"
)

type PageData struct {
	Title string // Page title
}

// Base layout template
const layoutTmpl = `
{{define "layout"}}
		<html>
			<head>
				<title>{{template "title" .}}</title> <!-- Include title block -->
			</head>
			<body>
				{{template "content" .}} <!-- Inline template -->
			</body>
		</html>
{{end}}
	`

// Specific content template
const contentTmpl = `
{{define "title"}}Home Page{{end}} <!-- Define title block -->

{{define "content"}} <!-- Define content block -->
    <h1>Welcome</h1>
    <p>This is the home page content</p>
{{end}}`

func main() {
	// parse both templates
	tmpl := template.Must(template.New("layout").Parse(layoutTmpl))
	tmpl = template.Must(tmpl.Parse(contentTmpl))

	// Execute the layout template, which includes the others
	tmpl.ExecuteTemplate(os.Stdout, "layout", nil)
}
