package main

import "fmt"

/*
Structs:
- blueprint for creating objects
- typed collections of fields; useful for grouping data together to form records
*/
// Employee struct type has age and gender fields
type Employee struct {
	age    int
	gender string
}

func main() {
	var e Employee
	e.age = 24
	e.gender = "male"
	fmt.Println(e) // {24 male}
	// fmt.Println(Employee{29, "female"}) // {29 female}
	var empAge, empGender = Employee{21, "female"}, Employee{19, "male"}
	fmt.Println(empAge, empGender) // {21 female} {19 male}

	// anonymous struct
	// useful when you don't need to use the struct type again
	var emp = struct {
		age    int
		gender string
	}{21, "female"}
	fmt.Println(emp.age, emp.gender) // 21 female

	// Custom Types
	// Reciever function
	myDepartment := newDepartment("R&D")
	fmt.Println(myDepartment) // {R&D map[]}

	// Note: Run both main and department at once
}
