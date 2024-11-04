package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	todos := Todos{}
	todos.add("book eating")
	todos.add("book reading")
	todos.isToggle(1)
	todos.Print()
}