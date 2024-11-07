package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	List   bool
	Add    string
	Edit   string
	Del    int
	Toggle int
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}
	flag.StringVar(&cf.Add, "add", "", "Add a todo in the list")
	flag.BoolVar(&cf.List, "list", false, "List all TODO")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo in the list by Index")
	flag.IntVar(&cf.Del, "del", -1, "Delete a todo from the list by Index")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify a todo to toggle by Index")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.Print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Toggle != -1:
		todos.isToggle(cf.Toggle)
	case cf.Del != -1:
		todos.deleteTodo(cf.Del)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("error: invalid form to edit, use format id:new_title")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		todos.editTodo(index, parts[1])

	default:
		fmt.Println("invalid command")
	}
}
