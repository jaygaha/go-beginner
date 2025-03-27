package main

func main() {
	items := TodoList{}

	storage := NewStorage[TodoList]("todo.json")
	storage.Load(&items)

	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&items)

	storage.Save(items)
}
