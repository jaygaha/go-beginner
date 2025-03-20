package package_testify

// model
type Student struct {
	Name     string
	standard int
}

/*
Filter out unique students from a list of students
*/
func FilterUniqueStudents(students []Student) []string {
	var uniqueStudents []string
	check := make(map[string]int) // map of string to int

	for _, student := range students {
		check[student.Name] = 1
	}

	for name := range check {
		uniqueStudents = append(uniqueStudents, name)
	}

	return uniqueStudents
}
