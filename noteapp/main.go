package main

import (
	"fmt"
	"noteapp/note"
)

func main() {
	note := note.NewNote()

	note.DisplayNote()

	if err := note.SaveNoteToFile(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Saving the note succeeded!")
	}

}
