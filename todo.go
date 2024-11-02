package main

import (
	"errors"
	"fmt"
	"time"
)

type Todo struct {
	Title       string
	isCompleted bool
	createdAt   time.Time
	completedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		createdAt:   time.Now(),
		isCompleted: false,
		completedAt: nil,
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
		return err
	}
	// fmt.Println(t[index])
	t[index].isCompleted = !t[index].isCompleted
	if t[index].isCompleted {
		completionTime := time.Now()
		t[index].completedAt = &completionTime
	} else {
		t[index].completedAt = nil
	}
	fmt.Println(t[index].completedAt.Format(time.RFC1123))
	return nil
}
