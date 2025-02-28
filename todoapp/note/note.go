package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
	"todoapp/inputs"
)

type Note struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

func NewNote() Note {
	title := inputs.ScanUserInput("Note title: ")
	content := inputs.ScanUserInput("Content: ")

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now().Format(time.DateTime),
	}
}

func (n Note) Display() error {
	emptyNode := Note{}

	if emptyNode == n {
		return errors.New("error - note is empty")
		// fmt.Println("Note is empty, there is nothing to desplay")
		// os.Exit(1)
	}

	fmt.Println()
	fmt.Println("Note Title:", n.Title)
	fmt.Println("Content:", n.Content)
	fmt.Println("Created:", n.CreatedAt)
	fmt.Println()
	return nil
}

func (n Note) WriteToFile() error {
	fileName := strings.ReplaceAll(n.Title, " ", "-")
	fileName = strings.ToLower(fileName) + ".json"

	note, err := json.Marshal(n)
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile("./data/"+fileName, note, 0644); err != nil {
		return err
	}
	return nil
}
