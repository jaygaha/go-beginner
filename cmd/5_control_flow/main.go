package main

import "fmt"

func main() {
	// Control Flow: if, else if, else, switch, for, range, break, continue
	fmt.Println("\nControl Flow")
	/** 1 if, else if, else
	Go's if statements are like its for loops; the expression need not be surrounded by parentheses ( ) but the braces { } are required.
	*/

	fmt.Println("\n1.1: if")
	score := 81

	// 1.1 if
	if score >= 80 {
		fmt.Println("A")
	}

	// 1.2 if else
	fmt.Println("\n1.2: if else")
	score = 71
	if score >= 80 {
		fmt.Println("A")
	} else {
		fmt.Println("B")
	}

	// 1.3 if else if else
	fmt.Println("\n1.3: if else if else")
	score = 40
	if score >= 80 {
		fmt.Println("A")
	} else if score >= 60 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	// and, or, not
	score = 81

	// and
	if score >= 60 && score <= 100 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}

	// or
	score = 45
	practicalScode := 80
	if score >= 80 || practicalScode >= 85 {
		fmt.Println("Excelent")
	} else if score >= 60 || practicalScode >= 70 {
		fmt.Println("Good")
	} else {
		fmt.Println("Need to study more")
	}

	// 2 switch
	fmt.Println("\n2.1: switch")
	score = 69

	switch score {
	case 80:
		fmt.Println("A")
	case 60:
		fmt.Println("B")
	case 40:
		fmt.Println("C")
	default:
		fmt.Println("D")
	}

	/**
	3 Loop
	Go has only one looping construct, the for loop. The basic for loop has three components separated by semicolons:

	the init statement: executed before the first iteration
	the condition expression: evaluated before every iteration
	the post statement: executed at the end of every iteration
	*/
	/**
	 * 3.1 for
	 * for condition
	 */
	fmt.Println("\n3.1: for")
	i := 1
	// Go while loop
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// forever loop
	for {
		fmt.Println("forever")
		break
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	/**
	 * 3.2 range
	  	The range form of the for loop iterates over a slice or map.

		When ranging over a slice, two values are returned for each iteration.
		The first is the index, and the second is a copy of the element at that index.
		You can choose which value to receive by assigning to the blank identifier _.
	 * for i, v := range s { }
	 * for i := range s { }
	 * for _, v := range s { }
	 * for i, v := range s { }
	*/
	fmt.Println("\n3.2: for range")
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	for j := range 5 {
		fmt.Println(j)
	}

	// range skip: assign `_` to skip index or value
	// example:
	// for i, _ := range pow
	// for _, v := range pow

	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

	/**
	4 break, continue
	The break statement immediately terminates execution of the innermost for loop.
	The continue statement begins the next top-level loop iteration.
	*/
	fmt.Println("\n4.1: break")

	for i := 1; i <= 10; i++ {
		// terminates the loop when i is equal to 6
		if i == 6 {
			break
		}
		fmt.Print(i) // 1 2 3 4 5
	}

	fmt.Println("\n4.2: continue")
	for i := 1; i <= 10; i++ {
		// skip the iteration when i is equal to 5
		if i == 5 {
			continue
		}
		fmt.Print(i) // 1 2 3 4 6 7 8 9 10
	}

}
