package main

import (
	"fmt"

	"example.com/notes/note"
	"example.com/notes/prompt"
)

func main() {
	noteTitle := prompt.PromptAndGetStr("Note title")
	noteContent := prompt.PromptAndGetStr("Note content")

	noteVar, err := note.New(noteTitle, noteContent)
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}

	noteVar.Print()

	err = noteVar.Save()
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}
}
