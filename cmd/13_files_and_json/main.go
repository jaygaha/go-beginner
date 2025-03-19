package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	/*
		JSON
			- JSON is a popular format for storing and transporting data.
			- Go provides built-in support for JSON encoding and decoding.
			- JSON is a text-based data interchange format.
			- built-in JSON support in the encoding/json package.

			Two main concepts:
			1. Marshaling: converting a Go value (data structure) to JSON
				syntax: json.Marshal(v interface{}) ([]byte, error)
			2. Unmarshaling: converting a JSON value to a Go value (data structure)
				syntax: json.Unmarshal(data []byte, v interface{}) error
	*/

	// Marshal
	fmt.Println("Marshal")
	jsonData := map[string]any{
		"name":    "John",
		"age":     30,
		"city":    "New York",
		"married": false,
	}

	marshalOutput := Marshal(jsonData)

	fmt.Printf("Output: %s\n", marshalOutput)

	// Unmarshal
	fmt.Println("\nUnmarshal")
	jsonData2 := `{"name":"ジョン","age":38,"city":"Tokyo","married":true, "hobbies": ["climbing", "cycling", "running"]}`

	var unmarshalOutput map[string]any // data type in Go is a generic type, so we can use any type

	err := json.Unmarshal([]byte(jsonData2), &unmarshalOutput)

	if err != nil {
		panic(err)
	}

	for k, v := range unmarshalOutput {
		fmt.Printf("%s: %v\n", k, v)
	}

	// Read JSON from a file
	fmt.Println("\nRead JSON File")
	ReadFile()

	// Encoder and Decoder
	fmt.Println("\nEncoder and Decoder")
	WriteFile()
}

func Marshal(jsonData any) []byte {
	output, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}
	return output
}
