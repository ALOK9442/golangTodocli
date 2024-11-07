package main

import "fmt"

func main() {
	fmt.Println("Hello, User!")
	todos := Todos{}
	storage := NewStorage[Todos]("jsondata.json")
	storage.Load(&todos)
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)
	storage.Save(todos)
}