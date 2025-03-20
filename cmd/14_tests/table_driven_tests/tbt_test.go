package tbt

import (
	"reflect"
	"testing"
)

func TestSplitStrings(t *testing.T) {
	got := SplitStrings("a:b:c", ":")
	want := []string{"a", "b", "c"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSplitStringsWrongSep(t *testing.T) {
	got := SplitStrings("a:b:c", "/")
	want := []string{"a:b:c"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSplitStringsNoSep(t *testing.T) {
	got := SplitStrings("abc", "/")
	want := []string{"abc"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

/*
In the above test cases, dublication occurs. For each test case only input, expected output and test name is changed

To avoid this, we can use table driven tests. All the input, expected output and test name are stored to a single test harness.
*/

func TestSplitStringsTBT(t *testing.T) {
	// struct to hold the test input, expected output
	// the tests stuct is usually locally defined to use same name in other cases
	// anonymous struct literal is used to avoid creating a new type
	tests := []struct {
		name  string
		input string
		sep   string
		want  []string
	}{
		{name: "simple", input: "a/b/c", sep: "/", want: []string{"a", "b", "c"}},
		{name: "wrong sep", input: "a/b/c", sep: ",", want: []string{"a/b/c"}},
		{name: "no sep", input: "abc", sep: "/", want: []string{"abc"}},
		// {name: "trailing sep", input: "a/b/c/", sep: "/", want: []string{"a", "b", "c"}}, // this trigger error
	}

	for _, test := range tests {
		got := SplitStrings(test.input, test.sep)

		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("got %v want %v", got, test.want)
		}
	}
}
