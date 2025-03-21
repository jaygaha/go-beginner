package main

import (
	"testing"
	"time"
)

/*
The test cases cover these key scenarios:
- ID auto-increment logic for both empty and non-empty lists
- Correct initialization of todo item fields (Title, Completed status)
- Proper timestamp handling (CreatedAt within expected timeframe, UpdatedAt initially nil)
- List management through multiple test cases

The test suite uses table-driven tests with two main cases:
1. Adding first item to an empty list
2. Adding item to a non-empty list

Each test case verifies all fields of the created todo item, ensuring the Add function works as expected.
*/
func TestTodoListAdd(t *testing.T) {
	// Test cases
	tests := []struct {
		name     string
		items    TodoList
		title    string
		expected Todo
	}{
		{
			name:  "Add first item to empty list",
			items: TodoList{},
			title: "First Todo",
			expected: Todo{
				ID:        1,
				Title:     "First Todo",
				Completed: false,
				UpdatedAt: nil,
			},
		},
		{
			name: "Add item to non-empty list",
			items: TodoList{
				Todo{ID: 1, Title: "Existing Todo", Completed: false},
			},
			title: "Second Todo",
			expected: Todo{
				ID:        2,
				Title:     "Second Todo",
				Completed: false,
				UpdatedAt: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Get initial time for comparison
			beforeAdd := time.Now()

			// Add the todo item
			tt.items.Add(tt.title)

			// Get the added item
			addedItem := tt.items[len(tt.items)-1]

			// Verify ID
			if addedItem.ID != tt.expected.ID {
				t.Errorf("Expected ID %d, got %d", tt.expected.ID, addedItem.ID)
			}

			// Verify Title
			if addedItem.Title != tt.expected.Title {
				t.Errorf("Expected Title %s, got %s", tt.expected.Title, addedItem.Title)
			}

			// Verify Completed status
			if addedItem.Completed != tt.expected.Completed {
				t.Errorf("Expected Completed %v, got %v", tt.expected.Completed, addedItem.Completed)
			}

			// Verify CreatedAt is set and within expected timeframe
			if addedItem.CreatedAt.Before(beforeAdd) || addedItem.CreatedAt.After(time.Now()) {
				t.Error("CreatedAt time is not within expected timeframe")
			}

			// Verify UpdatedAt is nil
			if addedItem.UpdatedAt != nil {
				t.Error("Expected UpdatedAt to be nil")
			}
		})
	}
}
