package tbt

import "strings"

/*
Table Driven Tests
- Table Driven Tests are a way to test a function with multiple inputs and expected outputs
*/

/*
- Split slices s into all substrings separated by sep
- returns a slice of the substrings between those separators.
*/
func SplitStrings(s, sep string) []string {
	var result []string
	i := strings.Index(s, sep) // this will return the index of the first occurrence of sep in s

	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}

	return append(result, s)
}
