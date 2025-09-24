package main

import (
	"fmt"

	"example.com/todos/todo"
)

type outputtable interface {
	Save() error
	Print()
}

func main() {
	todoStr := promptAndGetStr("To-do")

	todoVar, err := todo.New(todoStr)
	if err != nil {
		fmt.Println("ERROR", err)
	}

	outputData(todoVar)
}

func outputData(data outputtable) {
	err := data.Save()
	if err != nil {
		fmt.Println("ERROR", err)
	}

	data.Print()
}
