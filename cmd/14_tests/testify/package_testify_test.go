package package_testify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterUniqueStudents(t *testing.T) {
	students := []Student{
		{Name: "Hoge"},
		{Name: "Fuge", standard: 7},
		{Name: "Hoge", standard: 6},
		{Name: "Piyo", standard: 7},
		{"Suzu", 7},
	}

	expected := []string{"Hoge", "Fuge", "Piyo", "Suzu"}

	assert.Equal(t, expected, FilterUniqueStudents(students))

	// Without testify
	// actual := FilterUniqueStudents(students)

	// // reflect.DeepEqual() is a function that compares two values for equality.
	// if !reflect.DeepEqual(expected, actual) {
	// 	t.Fail()
	// }
}

func TestNegativeFilterUniqueStudents(t *testing.T) {
	students := []Student{
		{Name: "Hoge"},
		{Name: "Fuge", standard: 7},
		{Name: "Hoge", standard: 6},
		{Name: "Piyo", standard: 7},
		{"Suzu", 7},
	}

	expected := []string{
		"Piyo",
		"Suzu",
	}

	assert.NotEqual(t, expected, FilterUniqueStudents(students))
}

// check the testify documentation for more examples and try to implement
