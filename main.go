package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	todos := Todos{}
	storage := NewStorage[Todos]("jsondata.json")
	storage.Load(&todos)
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)
	storage.Save(todos)
	// todos.Print()
}