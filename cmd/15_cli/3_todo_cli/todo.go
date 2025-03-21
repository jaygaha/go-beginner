package main

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

// Todo: struct to represent a todo item
type Todo struct {
	ID        int
	Title     string
	Completed bool
	CreatedAt time.Time
	UpdatedAt *time.Time // Pointer to time.Time; it can be nil
}

type TodoList []Todo // Slice of Todo to hold the todo items

func (items *TodoList) Add(title string) {

	// If items is empty assign id to 1 else assign id to the last item id + 1
	id := 1
	if len(*items) > 0 {
		lastItem := (*items)[len(*items)-1]
		id = lastItem.ID + 1
	}

	todoItem := Todo{
		ID:        id,
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
		UpdatedAt: nil,
	}

	// Add the todo item to the list
	// *items: this pointer allows to access and modify the original list stored in the memory
	*items = append(*items, todoItem)
}

/*
Validate if the todo item exists in the list
items: this is a pointer to the list of todo items
id: this is the id of the todo item to validate
*/
func (items *TodoList) ValidateId(id int) error {
	for _, item := range *items {
		if item.ID == id {
			return nil // nil means that the todo item exists in the list
		}
	}

	err := errors.New("todo item not found")
	fmt.Println(err)

	return err
}

/*
Delete a todo item from the list
items: this is a pointer to the list of todo items
id: this is the id of the todo item to delete
*/
func (items *TodoList) Delete(id int) error {
	// Validate if the todo item exists in the list
	err := items.ValidateId(id)
	if err != nil {
		return err
	}

	// Delete the todo item from the list
	for i, item := range *items {
		if item.ID == id {
			/*
				slices.Delete:
					- items: this is a pointer to the list of todo items
					- i: this is the index of the todo item to delete
					- i+1: this is the index of the next todo item to delete
			*/
			*items = slices.Delete((*items), i, i+1)

			return nil
		}
	}

	return nil
}

/*
Update complete status of a todo item
- If completed status is true, update as complete as true(completed) and update the updatedAt time and vice versa
*/
func (items *TodoList) UpdateCompleteStatus(id int, completed bool) error {
	// Validate if the todo item exists in the list
	err := items.ValidateId(id)
	if err != nil {
		return err
	}

	// Update the complete status of the todo item
	for i, item := range *items {
		if item.ID == id {
			(*items)[i].Completed = completed
			updatedAt := time.Now()
			(*items)[i].UpdatedAt = &updatedAt

			return nil
		}
	}

	return nil
}

/*
Edit a todo item
*/
func (items *TodoList) Edit(id int, title string) error {
	// Validate if the todo item exists in the list
	err := items.ValidateId(id)
	if err != nil {
		return err
	}
	/*
		Update the title of the todo item
	*/
	for i, item := range *items {
		if item.ID == id {
			(*items)[i].Title = title
			updatedAt := time.Now()
			(*items)[i].UpdatedAt = &updatedAt

			return nil
		}
	}

	return nil
}

/*
Print the todo list
  - use of external package table package to print the todo list
*/
func (items *TodoList) Display() {
	fmt.Println("Todo List")
	table := table.New(os.Stdout)                                            // Create a new table & os.Stdout is the output stream
	table.SetRowLines(false)                                                 // Disable row lines
	table.SetHeaders("ID", "Title", "Completed", "Created At", "Updated At") // Set the headers of the table

	for _, item := range *items {
		completed := "❌"
		UpdatedAt := ""

		if item.Completed {
			completed = "✅"
		}

		if item.UpdatedAt != nil {
			UpdatedAt = item.UpdatedAt.Format(time.RFC1123)
		}

		table.AddRow(strconv.Itoa(item.ID), item.Title, completed, item.CreatedAt.Format(time.RFC1123), UpdatedAt) // Add a row to the table
	}

	table.Render() // Render the table
}
