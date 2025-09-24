package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

const notesFilename = "notes.json"

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Title and content must not be empty.")
	}

	return Note{title, content, time.Now()}, nil
}

func (note Note) Print() {
	fmt.Printf("\nYour note '%s' has the following content:\n\n%s\n\n", note.Title, note.Content)
}

func (note Note) Save() error {
	json, err := json.Marshal(note)
	if err != nil {
		return err
	}

	return os.WriteFile(notesFilename, json, 0644)
}
