package main

import (
	"os"
	"strings"
	"text/template"
)

type Car struct {
	Model               string
	Brand               string
	Color               string
	Power               int
	BuildYear           int
	ManufacturerCountry string
	IsElectric          bool
	IsDiscontinued      bool
}

type Person struct {
	Name string
	Age  int
}

func main() {
	/*
		Template
			- Most common use case is to generate HTML pages
			- template is a text that contains one or more placeholders for dynamic data

		text/template
			- text/template is a package that provides a way to generate text based on templates
			- Difference from html/template package
				- text/template is designed for text, not HTML, normally used for text based output like files, reports, etc.
				- html/template is designed for HTML, normally used for web pages which provide additional features like forms, images, etc.
					and also provides additiinal safety features to help prevent cross-site scripting (XSS) attacks.

		tmpl:
		 	- tmpl is package that provides a way to generate text based on templates

		Basic Syntax:
		 	- Placeholders
				- {{.}} is a placeholder for dynamic data
				- {{.Name}} is a placeholder for dynamic data with name "Name"
				- {{.Address.postal_code}} is a placeholder for dynamic data with name "Name" and field "Age"

			- Control Structures
				- {{if}} is a conditional statement
				- {{range}} is a loop statement
				- {{with}} is a block statement

		Example:

	*/
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
	if err != nil {
		panic(err)
	}

	// control structures & functions
	// gt is a function that returns true if the first argument is greater than the second argument
	// custom functions: using template.FuncMap to register a function
	funcMap := template.FuncMap{
		"upper": strings.ToUpper,
	}
	tmpl2 := "Hello {{.Name | upper}}! You are {{.Age}} years old.{{if gt .Age 18}}You are an adult.{{else}}You are not an adult.{{end}}\n"
	p2 := Person{
		Name: "fuge",
		Age:  17,
	}

	// Funcs() is a method that returns a new template with the given functions.
	t2, err := template.New("geetings").Funcs(funcMap).Parse(tmpl2)
	if err != nil {
		panic(err)
	}
	err = t2.Execute(os.Stdout, p2)
	if err != nil {
		panic(err)
	}

	//  do not escape the xss attack; not safe
	tmpl3 := "Hello, {{.}}!\n"                           // {{.}} is a placeholder for dynamic data
	p3 := "<script>alert('I am xss attack!!!')</script>" // <script>alert('xss attack')</script> is a dynamic data
	t3, err := template.New("xss").Parse(tmpl3)
	if err != nil {
		panic(err)
	}
	err = t3.Execute(os.Stdout, p3)
	if err != nil {
		panic(err)
	}

	// external template file for the loop
	cars := []Car{
		{
			Model:               "A6",
			Brand:               "Audi",
			Color:               "Black",
			Power:               200,
			BuildYear:           2010,
			ManufacturerCountry: "Germany",
			IsElectric:          false,
			IsDiscontinued:      false,
		},
		{
			Brand:               "Toyota",
			Model:               "Corolla",
			Color:               "White",
			Power:               120,
			BuildYear:           2005,
			ManufacturerCountry: "Japan",
			IsElectric:          false,
			IsDiscontinued:      false,
		},
		{
			Brand:               "BMW",
			Model:               "X5",
			Color:               "Red",
			Power:               300,
			BuildYear:           2015,
			ManufacturerCountry: "Germany",
			IsElectric:          false,
			IsDiscontinued:      false,
		},
		// electric cars
		{
			Brand:               "Tesla",
			Model:               "Model 3",
			Color:               "White",
			Power:               300,
			BuildYear:           2020,
			ManufacturerCountry: "USA",
			IsElectric:          true,
			IsDiscontinued:      false,
		},
		{
			Brand:               "BYD",
			Model:               "Tucson",
			Color:               "Black",
			Power:               200,
			BuildYear:           2021,
			ManufacturerCountry: "China",
			IsElectric:          true,
			IsDiscontinued:      false,
		},
		// discontinued cars
		{
			Brand:               "Volkswagen",
			Model:               "Passat",
			Color:               "Black",
			Power:               180,
			BuildYear:           2000,
			ManufacturerCountry: "Germany",
			IsElectric:          false,
			IsDiscontinued:      true,
		},
	}

	var tmplFile = "cars.tmpl" // tmpl file is a template file
	// template.New() creates a new template with the given name.
	// template.ParseFiles() parses the named files and associates the resulting templates with it.
	tmpl, err := template.New(tmplFile).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, cars)
	if err != nil {
		panic(err)
	}
}
