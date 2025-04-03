# Go HTML Templates

This guide explains how to use Go's `html/template` package for generating HTML output from templates. It's designed for beginners who want to understand how HTML templates work in Go.

## What is html/template?

The `html/template` package is a powerful tool for generating HTML dynamically in Go applications. It allows you to create HTML templates with placeholders that can be filled with data at runtime.

## html/template vs text/template

Go provides two template packages:

1. **text/template**: Designed for general text output (files, reports, etc.)
2. **html/template**: Specifically for HTML output with built-in protection against XSS attacks

The key difference is that `html/template` automatically escapes data to prevent cross-site scripting (XSS) attacks. This means that any potentially dangerous characters in your data will be converted to safe HTML entities.

## Basic Template Syntax

### Placeholders

```go
{{.}}        // The entire data passed to the template
{{.Name}}    // Access a field named "Name" in the data
{{.Title}}   // Access a field named "Title" in the data
```

### Control Structures

```go
{{if condition}} ... {{else}} ... {{end}}  // Conditional statements
{{range .Items}} ... {{end}}               // Loop through items
{{with .Object}} ... {{end}}               // Create a new scope
```

### Functions and Pipes

```go
{{.Name | upper}}  // Apply the "upper" function to .Name
{{if gt .Age 18}}  // "gt" function checks if .Age > 18
```

## Examples from this Project

### 1. Basic HTML Template

Rendering a simple template to the console.

From `1_basic/main.go`:

```go
// Create a new template named "basic" and parse our template string
t := template.Must(template.New("basic").Parse(tmpl))

// Execute the template and write output to os.Stdout
err := t.Execute(os.Stdout, nil)
```

This example creates a simple HTML template and outputs it to the console. The template doesn't use any dynamic data yet.

### 2. Passing Data to Templates

Adding logic to control template output and using functions to modify data.

From `2_passing_data/main.go`:

```go
// Create data to pass to the template
data := PageData{
    Title: "Welcome Page",
    Name:  "Hoge gopher",
}

// Execute with our data
err := t.Execute(os.Stdout, data)
```

This example shows how to pass a struct with data to a template. The template uses `{{.Title}}` and `{{.Name}}` to access the fields of the struct.

### 3. Conditionals and Loops

From `3_conditional/main.go`:

```go
// Conditional example
{{if .IsAdmin}}
    <h1>Admin View</h1>
{{else}}
    <h1>User View</h1>
{{end}}

// Loop example
{{range .Users}}
    <li>{{. | upper}}</li>
{{else}}
    <li>No users found</li>
{{end}}
```

This example demonstrates conditional rendering and looping through a slice of data. It also shows how to use the pipe operator to apply functions to data.

### 4. Template Layouts

Go templates support layouts and nested templates. There are two approaches:

#### Inline Templates

From `4_layout/1_inline/main.go`:

```go
// Define layout and content in code
const layoutTmpl = `{{define "layout"}}...{{end}}`
const contentTmpl = `{{define "title"}}...{{end}}{{define "content"}}...{{end}}`

// Parse both templates
tmpl := template.Must(template.New("layout").Parse(layoutTmpl))
tmpl = template.Must(tmpl.Parse(contentTmpl))

// Execute the layout template
tmpl.ExecuteTemplate(os.Stdout, "layout", nil)
```

#### External Template Files

From `4_layout/2_external/main.go`:

```go
// Parse the external template file
t, err := template.ParseFiles("layout.tmpl")

// Parse additional content template
t = template.Must(t.Parse(contentTmpl))

// Execute with data
t.Execute(os.Stdout, data)
```

This approach loads a template from an external file and combines it with content defined in code.

## Security Features

The `html/template` package automatically escapes HTML, JavaScript, CSS, and URLs to prevent XSS attacks. For example:

```go
data := struct{ Content string }{"<script>alert('XSS');</script>"}
```

When rendered with `{{.Content}}`, the output will be the escaped string:

```html
&lt;script&gt;alert(&#39;XSS&#39;);&lt;/script&gt;
```

This prevents the script from executing in the browser.

## Best Practices

1. **Use html/template for web content**: Always use `html/template` instead of `text/template` for HTML output.
2. **Organize templates**: For larger applications, organize templates into separate files.
3. **Create reusable layouts**: Use template composition to create reusable layouts.
4. **Validate input data**: Even though templates escape data, validate user input before passing it to templates.
5. **Handle errors**: Always check for errors when parsing and executing templates.

## Resources

- [Go html/template Package Documentation](https://golang.org/pkg/html/template/)