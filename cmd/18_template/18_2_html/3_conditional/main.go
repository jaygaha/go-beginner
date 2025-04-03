package main

import (
	"html/template"
	"os"
	"strings"
)

type PageData struct {
	Title   string   // Title of the page
	Name    string   // Name of the user
	Users   []string // Slice(array) of user names
	IsAdmin bool     // Flag indicating if the user is an admin
}

const tmpl = `
		<html>
			<head>
				<title>{{.Title}}</title>
			</head>
			<body>
				<!-- Built-in printf formats strings -->
    			<p>{{printf "%s %s" "Hello" .Name}}</p>

				<!-- Displaying a list of users depending upon role -->
				{{if .IsAdmin}}
					<h1>Admin View</h1>
				{{else}}
					<h1>User View</h1>
				{{end}}

				<ul>
				<!-- Iterating over a slice of strings -->
					{{range .Users}}
						<li>{{. | upper}}</li> <!-- . represents the current element in the slice -->
					{{else}}
						<li>No users found</li> <!-- Executed if the slice is empty -->
					{{end}}
				</ul>
			</body>
		</html>
	`

func main() {
	/*
		Passing data to a template: Adding dynamic content to our HTML.
		Template with placeholders for dynamic data

		Conditional statements in templates:
			{{if condition}}
				// Code to execute if the condition is true
			{{else}}
				// Code to execute if the condition is false
			{{end}}

		Range statements in templates:
			{{range .SliceName}}
				// Code to execute for each element in the slice
			{{else}}
				// Code to execute if the slice is empty
			{{end}}

		Functions in templates:
			|: Pipes the output of one function to the input of another
			{{funcName arg1 arg2}}: Calls a function with given arguments
			{{len.SliceName}}: Returns the length of a slice
			{{index.SliceName index}}: Returns the element at the given index in a slice

			Custom functions:
			Define custom functions in Go and pass them to the template using Funcs method
			Example:
			Define a custom function in Go
			func double(x int) int {
				return x * 2
			}
			Pass the custom function to the template using Funcs method
			t := template.Must(template.New("tmpl").Funcs(template.FuncMap{"double": double}).Parse(tmpl))
			Use the custom function in the template
			{{double 5}}: Outputs 10


		Definitions:
			{{if .IsAdmin}}: Tests if IsAdmin is true
			{{range .Users}}: Loops over each item in Users
			{{.}}: In range, refers to current loop item
			{{end}}: Closes if or range blocks
			FuncMap: Maps function names to actual functions
			|: Pipeline operator, passes value to function
			len: Built-in function for length
			printf: Built-in function for formatted strings

	*/
	// Define custom function
	funcMap := template.FuncMap{
		"upper": strings.ToUpper,
	}

	t := template.New("tmpl").Funcs(funcMap)
	t = template.Must(t.Parse(tmpl))

	data := PageData{
		Title:   "My Web Page",
		Name:    "John Doe",
		Users:   []string{"hoge", "fuga", "suga"},
		IsAdmin: true, // false
	}
	err := t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
