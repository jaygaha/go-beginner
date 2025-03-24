package main

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct {
	FileName string
}

// initializes a new storage with specified file name
func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{FileName: fileName}
}

// save item to the storage
func (s *Storage[T]) Save(item T) error {
	fileData, err := json.MarshalIndent(item, "", "") // formats the json data

	if err != nil {
		return err
	}

	return os.WriteFile(s.FileName, fileData, 0644)
}

// retrieves all items from the storage
func (s *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.FileName)

	if err != nil {
		return err
	}

	return json.Unmarshal(fileData, data)
}
