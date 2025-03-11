package main

// Blueprint for the department object
// Custom data type
type department struct {
	name      string
	employees map[string]float64
}

// make new departments
// Custom Types
func newDepartment(name string) department {
	d := department{
		name:      name,
		employees: map[string]float64{},
	}

	return d
}
