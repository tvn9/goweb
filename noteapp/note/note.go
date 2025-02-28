package note

import (
	"encoding/json"
	"fmt"
	"noteapp/inputs"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

func NewNote() Note {
	title, err := inputs.GetNoteTitle("Note title: ")
	if err != nil {
		fmt.Println(err)
	}

	content, err := inputs.GetNoteContent("Content: ")
	if err != nil {
		fmt.Println(err)
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now().Format(time.DateTime),
	}
}

func (n Note) DisplayNote() {
	fmt.Println(n.Title)
	fmt.Println(n.Content)
	fmt.Println(n.CreatedAt)
}

func (n Note) SaveNoteToFile() error {
	fileName := strings.ReplaceAll(n.Title, " ", "-")
	fileName = strings.ToLower(fileName) + ".json"

	jsonDoc, err := json.Marshal(n)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, jsonDoc, 0644)
}
