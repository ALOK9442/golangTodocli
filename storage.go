package main

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct {
	Filename string
}

func NewStorage[T any](Filename string) *Storage[T] {
	return &Storage[T]{Filename: Filename}
}

func (s *Storage[T]) Save(data T) error {
	dataInput, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.Filename, dataInput, 0644)
}

func (s *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.Filename) 
	if err !=nil {
		return err
	}
	return json.Unmarshal(fileData, data)
}