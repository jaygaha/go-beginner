package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Country struct {
	Name      string `json:"name"` // `json:"name"` is a tag, have to match the json field name
	Capital   string `json:"capital"`
	Continent string `json:"continent,omitempty"` // omitempty is used to omit the field if it is empty
	Area      int    `json:"area"`
	// President string `json:"-"` // this field will be ignored from marshaling
	// PoliticalParties []string `json:"parties"` // renamed to "parties" from PoliticalParties
}

func ReadFile() {
	/*
		Files
			- Go provides a rich set of functions for working with files.
			- The file package lets you read and write files.
			- The os package provides functions for working with the operating system.
			- The io package provides interfaces and functions for working with streams of data.
			- The io/ioutil package provides functions for reading and writing files. (Deprecated)
			- The bufio package provides buffered I/O.

		JSON Tags
		 - to map a field of a struct to its JSON name, use the `json` tag.
		 - tag `omitempty` is used to omit the field if it is empty.
		 - the `-` (dash) is used to indicate the field should be ignored from marshaling.

	*/

	// 1. Open the file// 2. Read the file// 3. Close the file
	content, err := os.ReadFile("countries.json")
	if err != nil {
		panic(err)
	}

	countries := []Country{}

	err = json.Unmarshal(content, &countries) // &countries is a pointer to the slice of countries
	if err != nil {
		panic(err)
	}

	for _, country := range countries {
		fmt.Printf("%s is located in %s and its capital is %s with having a area of %d.\n", country.Name, country.Continent, country.Capital, country.Area)
	}
}
