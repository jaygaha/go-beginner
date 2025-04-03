# Templates in Go

Go supports for creating dynamic content or customizing the output to the user. 
A template is a text document that contains variables or tags that will be replaced with actual values when the template is executed.

Go provides the `text/template` and `html/template` packages to create templates.

- [text/template](18_1_test/README.md)

## Template Syntax

Go templates use double curly braces `{{` and `}}` to delimit variables or tags.
Variables are replaced with their values, and tags are executed as Go code.

Here are some common tags:
- `{{.}}`: Prints the value of the current context.
- `{{.Field}}`: Prints the value of the field `Field` in the current context.
- `{{if .Condition}} ... {{end}}`: Executes the block if the condition is true.
- `{{range .}}... {{end}}`: Iterates over the elements of the current context.

## Basic Usage

```go
package main
import (
	"fmt"
	"os"
	"text/template"     
)
func main() {
	tmpl, err := template.New("test").Parse("Hello, {{.Name}}!")
	if err != nil {
		panic(err)
	}
	data := struct {
		Name string
	}
	{
		Name: "Jay",
	}
	err = tmpl.Execute(os.Stdout, data)
	if err!= nil {
		panic(err)
	}

    // Output: Hello, Jay!
}
```

## Further Reading

- [Chapter 9 Templates](https://jan.newmarch.name/golang/template/chapter-template.html)
- [Go by Example: Text Templates](https://gobyexample.com/text-templates)
