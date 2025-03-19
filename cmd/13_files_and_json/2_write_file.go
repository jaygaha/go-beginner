package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func WriteFile() {
	/*
		JSON encoding & decoding
		 - encoding/json package


	*/
	// Encode JSON
	fmt.Println("Encode JSON")
	e := EngineerEmployee{
		Name:   "Hoge",
		Age:    21,
		Gender: "",
		Active: true,
		Projects: []string{
			"Project 1",
			"Project 2",
			"Project 3",
		},
	}

	f, err := os.Create("employee.json")
	if err != nil {
		fmt.Println("Error creating file: ", err)
		panic(err)
	}
	defer f.Close()

	encoder := json.NewEncoder(f) // create a new encoder this will write to the file
	err = encoder.Encode(e)       // encode the data to JSON and write to the file

	if err != nil {
		fmt.Println("Error encoding JSON: ", err)
		panic(err)
	}

	fmt.Println("Successfully wrote to file")

	// Decode JSON
	fmt.Println("Decode JSON")
	var d EngineerEmployee
	f, err = os.Open("employee.json")

	if err != nil {
		fmt.Println("Error opening file: ", err)
		panic(err)
	}
	defer f.Close()

	decoder := json.NewDecoder(f) // create a new decoder this will read from the file
	err = decoder.Decode(&d)      // decode the data from JSON and write to the file

	if err != nil {
		fmt.Println("Error decoding JSON: ", err)
		panic(err)
	}

	fmt.Println("Successfully read from file:")
	fmt.Printf("%+v\n", e)

}

type EngineerEmployee struct {
	Name     string   `json:"name"`
	Address  string   `json:"-"`
	Age      int      `json:"age"`
	Gender   string   `json:"gender,omitempty"`
	Active   bool     `json:"active"`
	Projects []string `json:"projects"`
}
