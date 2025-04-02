# Go Text Templates

This guide explains how to use Go's `text/template` package for generating text-based output from templates. It's designed for beginners who want to understand how templates work in Go.

## What are Templates?

Templates are text files containing placeholders that can be replaced with dynamic data. In Go, the `text/template` package provides a powerful way to generate text output based on templates.

Key concepts:
- **Templates** are text files with placeholders for dynamic data
- **Placeholders** are marked with double curly braces `{{ }}`
- **Data** is passed to templates to fill in the placeholders

## text/template vs html/template

Go provides two template packages:

1. **text/template**: Designed for general text output (files, reports, etc.)
2. **html/template**: Specifically for HTML output with built-in protection against XSS attacks

This guide focuses on `text/template`.

## Basic Template Syntax

### Placeholders

```go
{{.}}        // The entire data passed to the template
{{.Name}}    // Access a field named "Name" in the data
{{.Address.PostalCode}}  // Access nested fields
```

### Control Structures

```go
{{if condition}} ... {{else}} ... {{end}}  // Conditional statements
{{range .Items}} ... {{end}}               // Loop through items
{{with .Object}} ... {{end}}               // Create a new scope
```

### Functions

```go
{{.Name | upper}}  // Apply the "upper" function to .Name
{{if gt .Age 18}}  // "gt" function checks if .Age > 18
```

## Examples from this Project

### 1. Basic Template with Struct Data

```go
tmpl1 := "Hello {{.Name}}! You are {{.Age}} years old.\n"
p1 := Person{
    Name: "Hoge",
    Age:  35,
}
t1, err := template.New("t1").Parse(tmpl1)
if err != nil {
    panic(err)
}
err = t1.Execute(os.Stdout, p1)
```

This example:
1. Creates a template string with placeholders for Name and Age
2. Creates a Person struct with data
3. Parses the template
4. Executes the template with the Person data, writing output to stdout

### 2. Using Functions and Conditionals

```go
funcMap := template.FuncMap{
    "upper": strings.ToUpper,
}
tmpl2 := "Hello {{.Name | upper}}! You are {{.Age}} years old.{{if gt .Age 18}}You are an adult.{{else}}You are not an adult.{{end}}\n"

t2, err := template.New("geetings").Funcs(funcMap).Parse(tmpl2)
```

This example:
1. Creates a FuncMap to register the "upper" function
2. Uses the function with a pipe `|` to transform data
3. Uses the built-in `gt` function for comparison
4. Uses conditional statements to show different text based on Age

### 3. Template Files and Loops

The project includes a template file `cars.tmpl` that demonstrates:

```
Cars:
{{/* This is a comment */}}

Total Cars: {{ . | len -}}

{{ range . }}
Model: {{ .Model }}
Brand: {{ .Brand }}
Color: {{ .Color }}
Power: {{.Power }}
Build year: {{.BuildYear }}
Manufacturer country: {{.ManufacturerCountry }}
Is Electric: {{ if .IsElectric }}Yes{{ else }}No{{ end }}
Is Discontinued: {{ if .IsDiscontinued }}Yes{{ else }}No{{ end }}
---
{{ end }}
```

This template:
1. Uses `{{ . | len }}` to count the total number of cars
2. Uses `{{ range . }}...{{ end }}` to loop through each car
3. Accesses fields of each car inside the loop
4. Uses conditional statements to display "Yes" or "No" for boolean fields

## Loading Template Files

```go
tmpl, err := template.New(tmplFile).ParseFiles(tmplFile)
if err != nil {
    panic(err)
}

err = tmpl.Execute(os.Stdout, cars)
```

This code:
1. Creates a new template with the same name as the template file
2. Parses the template file
3. Executes the template with the cars data

## Common Template Functions

- `eq`: Equal comparison (`eq .Value1 .Value2`)
- `ne`: Not equal comparison
- `lt`: Less than comparison
- `le`: Less than or equal comparison
- `gt`: Greater than comparison
- `ge`: Greater than or equal comparison
- `and`: Logical AND
- `or`: Logical OR
- `not`: Logical NOT
- `len`: Length of array, slice, map, or string

## Custom Functions

You can add your own functions to templates using `template.FuncMap`:

```go
funcMap := template.FuncMap{
    "upper": strings.ToUpper,
    "add": func(a, b int) int { return a + b },
}

t := template.New("myTemplate").Funcs(funcMap)
```

## Security Considerations

The `text/template` package does not automatically escape content, which can lead to security issues when outputting user-provided data:

```go
tmpl3 := "Hello, {{.}}!\n"
p3 := "<script>alert('I am xss attack!!!')</script>"
```

This will output the script tag as-is. For web applications, use `html/template` instead, which automatically escapes HTML.

## Next Steps

- Try modifying the templates in this example
- Create your own templates for different data structures
- Explore the `html/template` package for web applications
- Look into template inheritance and composition for more complex templates

## Resources

- [Go text/template Package Documentation](https://golang.org/pkg/text/template/)
- [Go html/template Package Documentation](https://golang.org/pkg/html/template/)