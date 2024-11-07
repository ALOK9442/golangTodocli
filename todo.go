package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	IsCompleted bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		CreatedAt:   time.Now(),
		IsCompleted: false,
		CompletedAt: nil,
	}

	*todos = append(*todos, todo)
	fmt.Println(todos)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("index not found")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todos *Todos) deleteTodo(index int) error {
	// t := todos
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}
	*todos = append(t[:index], t[index+1:]...)
	return nil
}

func (todos *Todos) editTodo(index int, title string) error {
	// t := todos
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}
	fmt.Println(t[index].Title)
	t[index].Title = title
	fmt.Println(t[index].Title)
	return nil
}

func (todos *Todos) isToggle(index int) error {
	// t := todos
	t := *todos
	if err := t.validateIndex(index); err != nil {
		fmt.Println("here")
		return err
	}
	// fmt.Println(t[index])
	fmt.Println("here2", t[index].IsCompleted)
	t[index].IsCompleted = !t[index].IsCompleted
	fmt.Println("here3", t[index].IsCompleted)
	if t[index].IsCompleted {
		fmt.Println("here4", t[index].IsCompleted)
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
		fmt.Println("here5", t[index].IsCompleted)
	} else {
		fmt.Println("here6", t[index].IsCompleted)
		t[index].CompletedAt = nil
		fmt.Println("here7", t[index].IsCompleted)
	}
	// fmt.Println(t[index].CompletedAt.Format(time.RFC1123))
	return nil
}

func (todos *Todos) Print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "completed", "created at", "completed at")
	for i, v := range *todos {
		completed := "❌"
		completedAt := ""
		if v.IsCompleted {
			completed = "✅"
			if v.CompletedAt != nil {
				completedAt = v.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(strconv.Itoa(i), v.Title, completed, v.CreatedAt.Format(time.RFC1123), completedAt)

	}
	table.Render()
}
