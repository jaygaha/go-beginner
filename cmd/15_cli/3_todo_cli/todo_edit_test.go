package main

import (
	"errors"
	"testing"
	"time"
)

/*
The tests cover:
- Error handling for non-existent todo items
- Successful editing of existing todo items
- Verification of title updates
- Validation of UpdatedAt timestamp modifications

The test suite follows the table-driven testing pattern, maintaining consistency with existing tests in the codebase.
Each test case verifies the expected behavior, including proper error handling and state verification.
*/
func TestTodoListEdit(t *testing.T) {
	// Test cases
	tests := []struct {
		name     string
		items    TodoList
		id       int
		title    string
		expected error
	}{
		{
			name:     "Edit non-existent item",
			items:    TodoList{},
			id:       1,
			title:    "Updated Todo",
			expected: errors.New("todo item not found"),
		},
		{
			name: "Edit existing item",
			items: TodoList{
				Todo{ID: 1, Title: "Original Todo", Completed: false},
			},
			id:    1,
			title: "Updated Todo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Get initial time for comparison
			beforeEdit := time.Now()

			// Edit the todo item
			err := tt.items.Edit(tt.id, tt.title)

			// Check error cases
			if tt.expected != nil {
				if err == nil || err.Error() != tt.expected.Error() {
					t.Errorf("Expected error %v, got %v", tt.expected, err)
				}
				return
			}

			// For successful edits, verify the changes
			for _, item := range tt.items {
				if item.ID == tt.id {
					// Verify Title
					if item.Title != tt.title {
						t.Errorf("Expected Title %s, got %s", tt.title, item.Title)
					}

					// Verify UpdatedAt is set and within expected timeframe
					if item.UpdatedAt == nil {
						t.Error("Expected UpdatedAt to be set, got nil")
					} else if item.UpdatedAt.Before(beforeEdit) || item.UpdatedAt.After(time.Now()) {
						t.Error("UpdatedAt time is not within expected timeframe")
					}
					return
				}
			}
			t.Error("Todo item not found after edit")
		})
	}
}
