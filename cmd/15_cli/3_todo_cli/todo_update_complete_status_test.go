package main

import (
	"testing"
	"time"
)

/*
The test cases cover three key scenarios:
- Attempting to update a non-existent item (error case)
- Marking an existing item as completed
- Marking a completed item as incomplete

Each test case verifies:
- Proper error handling
- Correct status changes
- UpdatedAt timestamp updates
- Time validation to ensure updates occur within expected timeframes

The tests follow the table-driven pattern for consistency with existing tests in the codebase.
*/
func TestTodoListUpdateCompleteStatus(t *testing.T) {
	// Test cases
	tests := []struct {
		name           string
		items          TodoList
		id             int
		completed      bool
		expectedError  bool
		expectedStatus bool
	}{
		{
			name:           "Update non-existent item",
			items:          TodoList{},
			id:             1,
			completed:      true,
			expectedError:  true,
			expectedStatus: false,
		},
		{
			name: "Mark item as completed",
			items: TodoList{
				Todo{ID: 1, Title: "Test Todo", Completed: false},
			},
			id:             1,
			completed:      true,
			expectedError:  false,
			expectedStatus: true,
		},
		{
			name: "Mark completed item as incomplete",
			items: TodoList{
				Todo{ID: 1, Title: "Test Todo", Completed: true},
			},
			id:             1,
			completed:      false,
			expectedError:  false,
			expectedStatus: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Get initial time for comparison
			beforeUpdate := time.Now()

			// Update the todo item's complete status
			err := tt.items.UpdateCompleteStatus(tt.id, tt.completed)

			// Check error handling
			if tt.expectedError && err == nil {
				t.Error("Expected an error but got none")
			}
			if !tt.expectedError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			// If we expect no error, verify the updates
			if !tt.expectedError {
				// Find the updated item
				for _, item := range tt.items {
					if item.ID == tt.id {
						// Verify Completed status
						if item.Completed != tt.expectedStatus {
							t.Errorf("Expected Completed status %v, got %v", tt.expectedStatus, item.Completed)
						}

						// Verify UpdatedAt is set and within expected timeframe
						if item.UpdatedAt == nil {
							t.Error("Expected UpdatedAt to be set, but it was nil")
						} else if item.UpdatedAt.Before(beforeUpdate) || item.UpdatedAt.After(time.Now()) {
							t.Error("UpdatedAt time is not within expected timeframe")
						}
						break
					}
				}
			}
		})
	}
}
