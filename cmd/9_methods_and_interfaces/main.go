package main

import "fmt"

type Student struct {
	Name   string
	Age    int
	Friend Friend
}

type Friend struct {
	Name         string
	Priority     string
	CommonThings []string
}

/*
Receiver function for the method

	Without *: Value receiver
	With *Student: Pointer receiver for the structs Student
		Pointer:
			- a pointer is a variable that stores the memory address of another variable.
			- the method can modify the value that its receiver points to
			- to avoid copying the value on each method call; more efficent if the receiver is a large stuct
*/
func (s Student) Print() {
	fmt.Println("Student Information")
	fmt.Printf("Name: %s\n", s.Name)
	fmt.Printf("Age: %d\n", s.Age)
	fmt.Printf("Friend: %s\n", s.Friend.Name)
}

func (s *Student) UpdateAge() {
	s.Age += 1
}

func FriendPriority(s *Student) string {
	return s.Friend.Priority
}

/*
Interface
*/
type Animal interface {
	GetName() string
}

type Mammal struct {
	Name string
}

type Bird struct {
	Name string
}

func (m *Mammal) GetName() string {
	return "Mammal Name: " + m.Name
}

func (b *Bird) GetName() string {
	return "Bird Name: " + b.Name
}

func PrintAmimalName(a Animal) {
	fmt.Println(a.GetName())
}

func main() {
	fmt.Println("Methods")
	student := Student{
		Name: "ジェイ",
		Age:  15,
		Friend: Friend{
			Name:         "John",
			Priority:     "High",
			CommonThings: []string{"football", "cycling"},
		},
	}

	/*
		Methods:
			- A method is a function with a special receiver argument.
			- The receiver appears in its own argument list between the func keyword and the method name.
			- the Print mehtod has a veceiver of type Student named s
	*/
	student.Print()

	fmt.Println("\nAge after 1 year")
	student.UpdateAge()
	student.Print()

	/*
		Methods are functions
			- a method is just a function with a receiver argument
			- in pointer receiver, need to add & sign
	*/
	fmt.Println("\nMethods are functions")
	fmt.Printf("Friend priority %s\n", FriendPriority(&student))

	/*
		Interfaces:

			- An interface type is defined as a set of method signatures.
			- a value of interface type can hold any value that implements those methods.
			- a way to define a set of methods that a struct or other type must implement to conform to that interface.
	*/
	fmt.Println("\nInterfaces")
	mammal := &Mammal{Name: "Elephant"}
	bird := &Bird{Name: "Eagle"}
	PrintAmimalName(mammal)
	PrintAmimalName(bird)
}
