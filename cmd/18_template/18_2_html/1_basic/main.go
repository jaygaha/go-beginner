package main

import (
	"html/template"
	"os"
)

type Entry struct {
	Name string
	Done bool
}
type TodoList struct {
	User    string
	Entries []Entry
}

func main() {
	/*
		html/template: this package provides a way to generate HTML dynamically.
		It allows you to define templates with placeholders and then fill them in with data.
		Here's an example of how to use the html/template package:
	*/

	// Basic template defined directly in Go code
	const tmpl = `
<!DOCTYPE html>
<html>
<head>
    <title>Go First Template</title>
</head>
<body>
    <h1>Hello, Go-phers!</h1>
</body>
</html>`

	// Create a new template named "basic" and parse our template string
	// template.Must ensures any parsing errors cause a panic (stops execution)
	t := template.Must(template.New("basic").Parse(tmpl))

	// Execute the template and write output to os.Stdout (the console)
	// nil is passed as data since this template has no dynamic content
	err := t.Execute(os.Stdout, nil)
	if err != nil {
		panic(err) // Stop execution if there's an error
	}

	/*
		Definitions:
			template.New("name"): Creates a new template with a given name
			Parse(): Converts the template string into a usable template
			template.Must(): Wraps parsing and panics on error
			Execute(): Renders the template with provided data
	*/
}
