package main

import (
	"testing"
)

/*
The test cases cover three important scenarios:

- Testing with an empty todo list
- Testing with a non-existent ID in a populated list
- Testing with a valid ID in a populated list
*/
func TestTodoListValidateId(t *testing.T) {
	// Test cases
	tests := []struct {
		name    string
		items   TodoList
		id      int
		wantErr bool
	}{
		{
			name:    "Empty list",
			items:   TodoList{},
			id:      1,
			wantErr: true,
		},
		{
			name:    "Non-existent ID",
			items:   TodoList{Todo{ID: 1, Title: "Test Todo"}},
			id:      2,
			wantErr: true,
		},
		{
			name:    "Valid ID",
			items:   TodoList{Todo{ID: 1, Title: "Test Todo"}},
			id:      1,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.items.ValidateId(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateId() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
