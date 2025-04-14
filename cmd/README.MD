# Go Learning Examples

This directory contains various Go examples to help you learn the language step by step.

## Examples Overview

1. **[Introduction](1_introduction/1_README.md)**
   - Introduction to Go
   - Installation & setup
   - Initialization & running a Go program

2. **[Constant, Variables, and Types](2_variables_data_type/2_README.md)**
   - Basic data types
   - Variable declaration
   - Constants

3. **[Printing & formatting](3_print_formatter/3_README.md)**
    - Printing to console
    - Formatting strings

4. **[Arrays, Slices, and Maps](4_array_and_slices/4_README.md)**
    - Fixed-length arrays
    - Dynamic slices
    - Key-value maps

5. **[Control Flow](5_control_flow/5_README.md)**
    - If/else statements
    - Switch cases
    - Loops (for, range)
    - Control statements (break, continue)

6. **[Functions](6_functions/6_README.md)**
    - Basic function syntax
    - Multiple return values
    - Variadic functions
    - Defer statements
    - Closures
    - Recursion
    - Anonymous functions

7. **[Error Handling](7_error_handling/7_README.md)**
    - The error interface
    - Constructing basic errors
    - Expected (sentinel) errors
    - Custom error types
    - Error handling best practices
    - **[Panic and Recover](7_error_handling/7_1_panic_and_recover/7_1_README.md)**

8. **[Structs](8_structs_and_custom_types/8_README.md)**
    - Basic function syntax
    - Custom Types

9. **[Methods and Interface](9_methods_and_interfaces/9_README.md)**
    - Methods
    - Interface

10. **[Packages and Modules](10_packages_and_modules/10_README.md)**
    - Package organization
    - Module management
    - Local/remote imports
    - Semantic versioning

11. **[Basic Miscellaneous](11_basic_misc/11_README.md)**
    - Type Casting
    - Type Inference
    - Type Assertions
    - Type Switches
    - Maps and Make
    - Pointers
    - Generics

12. **[Concurrency](12_concurrency/12_README.md)**
    - Goroutines
    - Channels (Buffered and Unbuffered)
    - Select Statement
    - Synchronization (WaitGroup, Mutex)

13. **[File I/O and Working with JSON](13_files_and_json/13_README.md)**
    - File I/O
    - Working with JSON
    - JSON Marshalling and Unmarshalling
    - JSON Encoding and Decoding

14. **[Testing](14_tests/14_README.md)**
    - Writing tests
    - Test Driven Development
    - Package testing
    - Table-driven tests

15. **[Command line interfaces](15_cli/15_README.md)**
    - Basic CLI
    - Flags
    - Subcommands
    - Parsing arguments

16. **Web Development**
    - [HTTP Server](16_http/servers/README.md)
        - basic server
        - multiplexing
        - multiple server configuration
    - [HTTP Client](16_http/clients/README.md)
        - basic client
        - request and response
        - error handling
    - [Middleware](16_http/middleware/README.md)
        - basic middleware
        - error handling
        - logging
        - authentication
    - [Sessions](16_http/sessions/README.md)
        - basic session
        - session management
        - cookies
    - [Assets](16_http/assets/README.md)
    - [Password Hashing](16_http/hashing/README.md)
17. **Database**
    - [SQLite](17_database/1_sqlite/README.md)
        - basic database
        - CRUD operations
        - error handling
        - advanced
            - [Advanced](17_database/1_sqlite/advanced/README.md)
                - CRUD operations
                - Context
                - error handling
                - best practices
                - transactions
                - testing
    - [MySQL](17_database/2_mysql/README.md)
        - connection
        - CRUD operations using TDD approach
        - transactions
        - prepared statements
        - context
        - best practices
18. **[Templates](18_template/README.md)**
    - [Text templates](18_template/18_1_text/README.md)
    - [HTML templates](18_template/18_2_html/README.md)
19. **[Logging](19_logging/README.md)**
    - basic logging
    - log using third party

## How to Run Examples

1. Navigate to the specific example directory
2. Read the comments in the code for explanation
3. Run the program using:
   ```bash
   go run main.go
    ```

## __Tips for Beginners__

- Take time to understand each concept before moving to the next
- Try to modify the examples to experiment with the code
- Use `go doc` command to read documentation
- Practice writing your own programs using learned concepts