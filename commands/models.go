package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// task represents a single task with description, status, and timestamps
type task struct {
	Id          int
	Description string
	Status      ListFilter
	Created_at  time.Time
	Update_at   time.Time
}

// createTask creates and returns a new task with a unique ID like Constructor
func createTask(description string) *task {
	file, err := os.ReadFile(FILEPATH)
	id := 1
	if err != nil {
		if os.IsNotExist(err) {
			id = 1
		}
	} else {
		tasks := []task{}
		if err := json.Unmarshal(file, &tasks); err != nil {
			fmt.Println("Error decoding JSON:", err)
		}

		for _, t := range tasks {
			if id < t.Id {
				id = t.Id
			}
		}
		id++
	}
	newTask := task{
		Id:          id,
		Description: description,
		Status:      todo,
		Created_at:  time.Now(),
		Update_at:   time.Now(),
	}

	return &newTask
}

// displayTask prints all tasks in a formatted output
func displayTask(tasks []task) {
	for _, t := range tasks {
		fmt.Printf(
			"\nID: %d\nDescription: %s\nStatus: %s\nCreated: %s\nUpdated: %s\n\n",
			t.Id, t.Description, t.Status,
			t.Created_at.Format("2006-01-02 15:04"),
			t.Update_at.Format("2006-01-02 15:04"),
		)
	}
}
