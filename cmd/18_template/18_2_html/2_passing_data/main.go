package main

import (
	"html/template"
	"os"
)

type PageData struct {
	Title string
	Name  string
}

func main() {
	/*
				Passing data to a template: Adding dynamic content to our HTML.

				// Template with placeholders for dynamic data
		    	// {{.FieldName}} accesses fields from the data struct

				Definitions:

				{{.Title}}: Dot notation accesses struct fields
				struct: A custom data type to organize related data
				os.Stdout: Standard output (your terminal/console)
				os.Stderr: Standard error output (for error messages)
	*/

	const tmpl = `
<!DOCTYPE html>
<html>
<head>
    <title>{{.Title}}</title> <!-- Inserts Title field -->
</head>
<body>
    <h1>Hello, {{.Name}}!</h1> <!-- Inserts Name field -->
</body>
</html>`

	t := template.Must(template.New("data").Parse(tmpl))

	// Create data to pass to the template
	data := PageData{
		Title: "Welcome Page",
		Name:  "Hoge gopher",
	}

	// Execute with our data instead of nil
	err := t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
