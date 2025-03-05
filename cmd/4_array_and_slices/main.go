package main

import "fmt"

func main() {
	// 1: Array
	/**
	1: Arrays
		- fixed length
		- same type
		- can't change the length
		- indexable
		- contiguous memory allocation
	*/
	fmt.Println("Arrays")

	// var ages [3]int = [3]int{12, 13, 14} // fixed length
	var ages = [3]int{12, 13, 14}
	// ages := [...]int32{12, 13, 14} // ... means the length will be the number of elements in the array

	names := [4]string{"a", "b", "e", "d"}

	fmt.Println(ages, len(ages))   // [12 13 14] 3
	fmt.Println(names, len(names)) // [a b e d] 4
	fmt.Println(&names[0])         // & for memory location: 0xc0000180c0

	// 2: Slices: use arrays under the hood that mean you can manually change the size of the array
	fmt.Println("Slices")

	names[2] = "c"
	fmt.Println(names) // [a b c d]

	var scores = []int{100, 200, 300}
	fmt.Println(scores, len(scores)) // [100 200 300] 3

	// change the value of the slice
	scores[2] = 500

	// append
	scores = append(scores, 400)
	fmt.Println(scores, len(scores)) // [100 200 500 400] 4

	// 3: Slice range: loop through the slice
	fmt.Println("Slice range")

	rangeOne := names[1:3]  // [start:end]
	rangeTwo := names[2:]   // [start:] start to end of the array
	rangeThree := names[:2] // [:end] start to end of the array

	fmt.Println(rangeOne)   // [b c]
	fmt.Println(rangeTwo)   // [c d]
	fmt.Println(rangeThree) // [a b]

	// append to slice
	rangeOne = append(rangeOne, "z")
	fmt.Println(rangeOne) // [b c z]

	for i, v := range names {
		fmt.Printf("Name %v: %v\n", i, v) // a, b c z
	}

	// 4: Slice make: create a slice
	fmt.Println("Slice make")
	var sliceMake []int32 = make([]int32, 3)

	fmt.Println(sliceMake) // [0 0 0]

	// 5: Map: key-value pair
	fmt.Println("Map")

	var studentScores map[string]uint8 = map[string]uint8{
		"a": 77,
		"b": 30,
		"c": 56,
	}

	fmt.Println(studentScores)      // map[a:77 b:30 c:56]
	fmt.Println(studentScores["b"]) // 30

	var name, isPassed = studentScores["c"]

	fmt.Println(name, isPassed) // c true
}
