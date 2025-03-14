package main

import "fmt"

func MakeMaps() {
	/*
		Maps
			- unordered collection of key-value pairs
			- key is unique
			- value can be of any type
			- key can be of any type
	*/
	println("Maps")
	var m1 map[string]int = map[string]int{
		"John": 21,
		"Doe":  23,
		"Foo":  24,
		"Bar":  19,
	}

	// Go doesn't keep order of the map, so the output will be different
	fmt.Printf("m1:\t%v\n", m1) // map[Bar:19 Doe:23 Foo:24 John:21]

	// accesing map
	// syntax: m[key]
	fmt.Printf("Doe:\t%v\n", m1["Doe"]) // Doe:23

	// adding to map
	// syntax: m[key] = value
	m1["Baz"] = 17
	fmt.Printf("adding m1:\t%v\n", m1) // map[Bar:19 Baz:17 Doe:23 Foo:24 John:21]

	// updating map
	// syntax: m[key] = value
	m1["John"] = 22
	fmt.Printf("updating m1:\t%v\n", m1) // map[Bar:19 Baz:17 Doe:23 Foo:24 John:22]

	// deleting from map
	// syntax: delete(map, key)
	delete(m1, "Foo")
	fmt.Printf("deleting Foo:\t%v\n", m1) // map[Bar:19 Baz:17 Doe:23 John:22]

	// Check for specific elements in a map
	// syntax: _, ok := m[key]
	_, ok := m1["Foo"]
	fmt.Printf("Foo exists?:\t%v\n", ok) // Foo exists?:    false

	// Iterating over a map
	// syntax: for key, value := range map
	// key: value pair
	fmt.Println("Iterating over a map")
	for key, value := range m1 {
		fmt.Printf("key:\t%v, value:\t%v\n", key, value)
	}

	/*
		Make
			- used to create a map, slice, channel depending on the arguments passed to it

			Syntax:
			make(map[type of key]type of value)
	*/
	println("\nMake")
	m2 := make(map[string]int) // creating a empty map
	m2["orage"] = 21
	m2["apple"] = 30
	m2["banana"] = 15
	fmt.Printf("m2:\t%v\n", m2)

	m3 := make(map[string]string, 5) // creating a empty map with capacity 5
	m3["country"] = "Japan"
	m3["populations"] = "120000000"
	m3["capital"] = "Tokyo"
	m3["currency"] = "yen"
	m3["language"] = "Japanese"
	m3["food"] = "ramen"
	fmt.Printf("m3:\t%v\n", m3)

	fmt.Println("Maps and Make End")
}
