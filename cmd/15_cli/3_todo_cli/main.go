package main

func main() {
	items := TodoList{}

	storage := NewStorage[TodoList]("todo.json")
	storage.Load(&items)

	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&items)

	storage.Save(items)

	// items.Add("Buy milk")
	// items.Add("Buy eggs")
	// items.Add("Buy bread")
	// // fmt.Printf("%+v\n\n", items)
	// // items.Delete(2)
	// items.UpdateCompleteStatus(1, true)
	// // fmt.Printf("%+v\n\n", items)

	// time.Sleep(2 * time.Second)
	// items.Add("Buy cheese")
	// items.UpdateCompleteStatus(1, false)
	// items.UpdateCompleteStatus(3, true)
	// items.Edit(1, "Buy yogurt")

	// items.Display()

	// storage.Add(items)
}
