package main

import (
	"testing"
	"time"
)

/*
The test cases cover three key scenarios:

- Deleting from an empty list
- Attempting to delete a non-existent item
- Successfully deleting an existing item
*/
func TestTodoListDelete(t *testing.T) {
	// Test cases
	tests := []struct {
		name          string
		initialItems  TodoList
		idToDelete    int
		expectedError bool
		expectedList  TodoList
	}{
		{
			name:          "Delete from empty list",
			initialItems:  TodoList{},
			idToDelete:    1,
			expectedError: true,
			expectedList:  TodoList{},
		},
		{
			name: "Delete non-existent item",
			initialItems: TodoList{
				Todo{ID: 1, Title: "Existing Todo", Completed: false, CreatedAt: time.Now()},
			},
			idToDelete:    2,
			expectedError: true,
			expectedList: TodoList{
				Todo{ID: 1, Title: "Existing Todo", Completed: false, CreatedAt: time.Now()},
			},
		},
		{
			name: "Delete existing item",
			initialItems: TodoList{
				Todo{ID: 1, Title: "First Todo", Completed: false, CreatedAt: time.Now()},
				Todo{ID: 2, Title: "Second Todo", Completed: true, CreatedAt: time.Now()},
			},
			idToDelete:    1,
			expectedError: false,
			expectedList: TodoList{
				Todo{ID: 2, Title: "Second Todo", Completed: true, CreatedAt: time.Now()},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of the initial items for testing
			testList := make(TodoList, len(tt.initialItems))
			copy(testList, tt.initialItems)

			// Attempt to delete the item
			err := testList.Delete(tt.idToDelete)

			// Check error condition
			if (err != nil) != tt.expectedError {
				t.Errorf("expected error: %v, got error: %v", tt.expectedError, err != nil)
			}

			// Check list length
			if len(testList) != len(tt.expectedList) {
				t.Errorf("expected list length: %d, got: %d", len(tt.expectedList), len(testList))
			}

			// Check remaining items
			for i, expectedItem := range tt.expectedList {
				if i >= len(testList) {
					t.Errorf("missing expected item at index %d", i)
					continue
				}

				if testList[i].ID != expectedItem.ID {
					t.Errorf("expected ID %d at index %d, got %d", expectedItem.ID, i, testList[i].ID)
				}

				if testList[i].Title != expectedItem.Title {
					t.Errorf("expected Title %s at index %d, got %s", expectedItem.Title, i, testList[i].Title)
				}

				if testList[i].Completed != expectedItem.Completed {
					t.Errorf("expected Completed %v at index %d, got %v", expectedItem.Completed, i, testList[i].Completed)
				}
			}
		})
	}
}
