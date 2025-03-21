package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
	"time"
)

/*
The tests verify:
- Table formatting and header setup
- Row content formatting including completed status symbols (✅/❌)
- Proper timestamp formatting for CreatedAt and UpdatedAt fields
- Handling of nil UpdatedAt values
- Both empty and populated todo lists

The tests use table-driven testing pattern and capture stdout to verify the rendered table content matches expectations.
Each test case includes specific assertions for headers, item details, and formatting.
*/
func TestTodoListDisplay(t *testing.T) {
	// Create a fixed time for testing
	createdTime := time.Date(2025, 03, 21, 10, 0, 0, 0, time.UTC)
	updatedTime := time.Date(2025, 03, 21, 11, 0, 0, 0, time.UTC)

	// Test cases
	tests := []struct {
		name     string
		items    TodoList
		expected []string // Lines we expect to see in the output
	}{
		{
			name:  "Empty todo list",
			items: TodoList{},
			expected: []string{
				"Todo List",
				"ID", "Title", "Completed", "Created At", "Updated At", // Headers
			},
		},
		{
			name: "List with incomplete and complete items",
			items: TodoList{
				Todo{
					ID:        1,
					Title:     "Test Todo",
					Completed: false,
					CreatedAt: createdTime,
					UpdatedAt: nil,
				},
				Todo{
					ID:        2,
					Title:     "Completed Todo",
					Completed: true,
					CreatedAt: createdTime,
					UpdatedAt: &updatedTime,
				},
			},
			expected: []string{
				"Todo List",
				"ID", "Title", "Completed", "Created At", "Updated At", // Headers
				"1", "Test Todo", "❌", createdTime.Format(time.RFC1123), "", // First item
				"2", "Completed Todo", "✅", createdTime.Format(time.RFC1123), updatedTime.Format(time.RFC1123), // Second item
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stdout
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Call Display
			tt.items.Display()

			// Restore stdout
			w.Close()
			os.Stdout = old

			// Read captured output
			var buf bytes.Buffer
			buf.ReadFrom(r)
			output := buf.String()

			// Check if all expected strings are in the output
			for _, exp := range tt.expected {
				if !strings.Contains(output, exp) {
					t.Errorf("Expected output to contain %q, but it didn't\nOutput:\n%s", exp, output)
				}
			}
		})
	}
}
