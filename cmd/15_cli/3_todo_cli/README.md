# Todo CLI Application

A simple command-line todo list application written in Go that helps you manage your tasks efficiently. The application stores tasks in JSON format and provides a user-friendly interface for task management.

## Features

- Add new todo items with descriptions
- Delete existing todo items by ID
- Edit todo item descriptions
- Mark todo items as complete/incomplete
- Display todo list in a formatted table view
- Persistent storage using JSON file
- Command-line flags for all operations

## Installation

1. Make sure you have Go installed on your system
2. Clone the repository
3. Navigate to the project directory
4. Run the application:
   ```bash
   go run .
   ```

## Usage

The application supports the following commands:

```bash
# Add a new todo
go run . -add "Buy groceries"

# List all todos
go run . -list

# Edit a todo (format: ID:new_text)
go run . -edit "1:Buy organic groceries"

# Update todo status (format: ID:status)
go run . -update "1:1"  # 1 for complete, 0 for incomplete

# Delete a todo
go run . -del 1
```

## Project Structure

- `main.go`: Contains the main application logic and entry point
- `todo.go`: Implements the TodoList type and its methods
- `command.go`: Handles command-line flag parsing and execution
- `command_test.go`: Contains test cases for command handling

## Testing

The project includes comprehensive tests for both the todo list operations and command handling. To run the tests:

```bash
go test ./...
```

The test suite covers:
- Todo list operations (add, edit, delete, update)
- Command-line flag parsing
- Command execution logic
- Error handling

## Contributing

Feel free to submit issues and enhancement requests!

## Further Reading

- [Command-line Interfaces (CLIs)](https://go.dev/solutions/clis)
- [How to Build A CLI Todo App in Go](https://codingwithpatrik.dev/posts/how-to-build-a-cli-todo-app-in-go/)