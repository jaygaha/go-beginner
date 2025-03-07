package main

import "fmt"

func main() {
	/*
		Function
			A function is a block of code that performs a specific task.
			Go main feature in function is that function can return multiple values
	*/
	// without arguments
	fmt.Println("Function")
	sayHelloWorld()

	// with arguments
	fmt.Println()
	fmt.Println("Function with arguments")
	sayHello("ジェイ")
	num1 := 8
	num2 := 7
	resultMul := multiplication(num1, num2)
	fmt.Printf("%dx%d = %d\n", num1, num2, resultMul)

	// multiple return values
	fmt.Println()
	fmt.Println("Multiple return values")
	var numerator int = 65
	var denominator int = 8

	result, remainder, err := intDivision(numerator, denominator)
	fmt.Printf("%d / %d = %d with remainder %d. Errors = %s\n", numerator, denominator, result, remainder, err)

	// Use of blank identifier _ to ignore return values
	_, _, err = intDivision(10, 0)
	fmt.Println(err)

	// Variadic function:
	// fmt.Println() is a variadic function
	fmt.Println()
	fmt.Println("Variadic function")
	sum(5, 2)         // 7
	sum(78, 99, 2, 1) // 180
	nums := []int{1, 2, 3, 4}
	sum(nums...) // 10

	// Defer
	fmt.Println()
	fmt.Println("Defer function")
	deferred()

	// Closures
	fmt.Println()
	fmt.Println("Closures")
	sum1 := closureSum()

	for i := range 3 {
		fmt.Println(sum1(i)) // Will print: 0, 1, 3 (accumulating sum)
	}

	newSum := closureSum()
	fmt.Println(newSum(5)) // Will print: 5 (new closure instance)

	// Recursion
	fmt.Println()
	fmt.Println("Recursion")
	fmt.Println(factorial(5)) // 120

	// anonymous function: annonymous function can be recursive, but requires explicit declaring a variable
	var fib func(n int) int
	fib = func(n int) int {
		if n <= 1 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(9)) // 34
}

func sayHelloWorld() {
	fmt.Println("Hello Go-phers!!")
}

func sayHello(name string) {
	fmt.Println("Hello", name)
}

/*
Go requires explicit returns, i.e. it won't automatically return the value of the last expression.
If multiple consecutive parameters of the same type, we can omit the type name for all but the last.
*/
func multiplication(num1, num2 int) int {
	return num1 * num2
}

/*
Multiple return values

	In a Go function, you can name the return values, just like the input parameters.
	These named return values start with zero values for their types. If the function returns without arguments,
	it uses the current values of these named return parameters.
*/
func intDivision(numerator, denominator int) (int, int, error) {
	var err error

	if denominator == 0 {
		err = fmt.Errorf("division by zero not allowed")
		return 0, 0, err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator

	return result, remainder, err
}

/*
Variadic function

	Variadic function is a function that can take any number of arguments.
	Variadic function is declared by using ... after the type of the last parameter.
*/
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

/*
Defer function

	Defer function is a function that is called after the function that contains the defer statement returns.
	defer keyword is used to defer a function call; it delays the execution of the function until the surrounding function returns.
	- useful for closing files, releasing resources, etc.
	- useful for cleanup activity like closing db connections, freeing up a mutex, or closing a file
*/
func deferred() {
	defer fmt.Println("a defer function")
	fmt.Println("a non-defer function")
}

/*
Closures

	Annonymous function which can form a closure around the code it references.
	Useful when you want to define a function inline without having to name it.
*/
func closureSum() func(int) int {
	// closure is a function that references variables from outside its body.
	// The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
	// For example, the function adder returns a closure. Each closure is bound to its own sum variable.

	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

/*
Recursion

	Recursion is a function that calls itself.
	Recursion is useful when you want to define something in terms of itself.
*/
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
