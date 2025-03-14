package main

import (
	"fmt"
	"strconv"
)

func TypeCasting() {
	/*
		Type Casting
			- Go doesn't support automatic type conversion, but allows type casting which is the process of explicity changing the variable types

			Syntax:
			v := typeName(otherTypeValue)

			eg: i := int(3.1415) // casting to integer
	*/
	println("Type Casting")
	var i1 int = 69
	f1 := float64(i1)
	fmt.Println(f1) // 69

	/*
		Conversion between String and Number
			- strconv package has functions to convert strings to numbers
			- strconv.Atoi() converts string to integer
			- strconv.ParseFloat() converts string to float
			- strconv.ParseBool() converts string to boolean
	*/
	// string to integer
	var s2 string = "31415"
	i2, _ := strconv.Atoi(s2) // converting string to integer
	fmt.Println(i2)           // 31415

	// integer to string
	var i3 int = 31415
	s3 := strconv.Itoa(i3) // converting integer to string
	fmt.Println(s3)        // 31415

	// string to float
	var s4 string = "3.1415"
	f4, _ := strconv.ParseFloat(s4, 64) // converting string to float
	fmt.Println(f4)                     // 3.1415

	// string to boolean
	var s5 string = "true"
	b5, _ := strconv.ParseBool(s5) // converting string to boolean
	fmt.Println(b5)                // true

	// float to int
	f6 := 3.1415
	i6 := int(f6)   // converting float to integer
	fmt.Println(i6) // 3

	// int to float
	f7 := 31415
	i7 := float64(f7) // converting float to integer
	fmt.Println(i7)   // 31415

	// Strings and byte conversion
	var s8 string = "Hello Go-phers"
	var b8 []byte = []byte(s8) // converting string to byte

	fmt.Println(b8) // [72 101 108 108 111 32 71 111 45 80 104 101 114 115]

	ss := string(b8) // converting byte to string
	fmt.Println(ss)  // Hello Go-phers

	println("Type Casting End")
}
