package main

import (
	"fmt"
	"os"
	"todoapp/datamanager"
	"todoapp/note"
	"todoapp/todo"
)

func main() {
	fmt.Println("Welcome to Note app")
	note := note.NewNote()

	fmt.Println("Welcome to Todo app")
	todo := todo.NewTodo()

	outputData(note)
	outputData(todo)

}

func outputData(d datamanager.Datahandler) {
	if err := d.WriteToFile(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := d.Display(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
