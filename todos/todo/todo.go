package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Todo string `json:"todo"`
}

const todoFilename = "todo.json"

func New(todo string) (Todo, error) {
	if todo == "" {
		return Todo{}, errors.New("Invalid to-do")
	}

	return Todo{todo}, nil
}

func (todo Todo) Print() {
	fmt.Println(todo.Todo)
}

func (todo Todo) Save() error {
	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	return os.WriteFile(todoFilename, json, 0644)
}
