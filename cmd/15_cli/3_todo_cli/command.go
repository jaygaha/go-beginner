package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// CmdFlags is a struct that contains the command line flags
type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Update string
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{} // create a new CmdFlags struct

	// parse the command line flags
	flag.StringVar(&cf.Add, "add", "", "Add a new todo. Format: 'new title'")
	flag.IntVar(&cf.Del, "del", 0, "Delete a todo by ID")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by ID and new text. Format: ID:new_text")
	flag.StringVar(&cf.Update, "update", "", "Update a status of todo by ID and new status. Format: ID:0/1")
	flag.BoolVar(&cf.List, "list", false, "List all todos")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(items *TodoList) {
	switch {
	case cf.List:
		items.Display()

	case cf.Add != "":
		items.Add(cf.Add)

	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)

		if len(parts) != 2 {
			fmt.Println("Invalid format for edit. Use ID:new_text")
			os.Exit(1)
		}

		id, err := strconv.Atoi(parts[0]) // convert string to int
		if err != nil {
			fmt.Println("Invalid ID. ID must be an integer")
			os.Exit(1)
		}

		items.Edit(id, parts[1])

	case cf.Update != "":
		parts := strings.SplitN(cf.Update, ":", 2)

		if len(parts) != 2 {
			fmt.Println("Invalid format for edit. Use ID:new_text")
			os.Exit(1)
		}

		id, err := strconv.Atoi(parts[0]) // convert string to int
		if err != nil {
			fmt.Println("Invalid ID. ID must be an integer")
			os.Exit(1)
		}

		completed, err := strconv.ParseBool(parts[1]) // convert string to bool
		if err != nil {
			fmt.Println("Invalid completed status. Completed status must be an 0 or 1")
			os.Exit(1)
		}

		items.UpdateCompleteStatus(id, completed)
	case cf.Del > 0:
		items.Delete(cf.Del)

	default:
		fmt.Println("No command provided")
	}
}
