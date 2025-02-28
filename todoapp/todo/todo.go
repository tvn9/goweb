package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
	"todoapp/inputs"
)

type Todo struct {
	Content   string `json:"text"`
	CreatedAt string `json:"created_at"`
}

func NewTodo() Todo {
	todoText := inputs.ScanUserInput("Todo Content: ")
	return Todo{
		Content:   todoText,
		CreatedAt: time.Now().Format(time.DateTime),
	}
}

func (t Todo) Display() error {
	emptyTodo := Todo{}

	if emptyTodo == t {
		return errors.New("error - Todo is empty")
	}

	fmt.Println()
	fmt.Println("Todo:", t.Content)
	fmt.Println("Created:", t.CreatedAt)
	fmt.Println()
	return nil
}

func (t Todo) WriteToFile() error {
	const fileName = "./data/todo.json"

	todo, err := json.Marshal(t)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile(fileName, todo, 0644); err != nil {
		fmt.Printf("error: %v", err)
		return err
	}
	return nil
}
