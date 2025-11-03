package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// showAll lists all tasks or filters them by status.
func showAll(args ListFilter) {
	file, err := os.ReadFile(FILEPATH)
	if err != nil {
		fmt.Println("No Tasks")
		return
	}

	var tasks []task
	if err := json.Unmarshal(file, &tasks); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	if args == "" {
		fmt.Println("Tasks : ")
		displayTask(tasks)
	} else {
		var filteredTasks []task
		for _, t := range tasks {
			if strings.EqualFold(string(t.Status), normalizeStatus(string(args))) {
				filteredTasks = append(filteredTasks, t)
			}
		}
		fmt.Println("Tasks : ")
		displayTask(filteredTasks)
	}
}

// addTask creates and saves a new task.
func addTask(description string) {
	newTask := createTask(description)
	file, err := os.ReadFile(FILEPATH)
	if err != nil {
		if os.IsNotExist(err) {
			tasks := []task{*newTask}
			saveJSON(tasks)
		}
		return
	}

	tasks := []task{}
	if err := json.Unmarshal(file, &tasks); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	tasks = append(tasks, *newTask)

	saveJSON(tasks)
	fmt.Println("Task Added With Id : ", newTask.Id)
	return
}

// deleteTask removes a task by its ID.
func deleteTask(id int) {
	file, err := os.ReadFile(FILEPATH)
	if err != nil {
		fmt.Println("No Tasks")
		return
	}

	tasks := []task{}
	if err := json.Unmarshal(file, &tasks); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	newTasks := []task{}
	found := false
	for _, t := range tasks {
		if t.Id != id {
			newTasks = append(newTasks, t)
		} else {
			found = true
		}
	}

	if !found {
		fmt.Println("Not Found")
		return
	}

	saveJSON(newTasks)
	fmt.Println("Task Deleted With Id : ", id)
	return
}

// updateTask updates the description of a task by ID.
func updateTask(id int, description string) {
	file, err := os.ReadFile(FILEPATH)
	if err != nil {
		fmt.Println("No Tasks")
		return
	}

	tasks := []task{}
	if err := json.Unmarshal(file, &tasks); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	newTasks := []task{}
	found := false
	for _, t := range tasks {
		if t.Id != id {
			newTasks = append(newTasks, t)
		} else {
			found = true
			t.Description = description
			t.Update_at = time.Now()
			newTasks = append(newTasks, t)
		}
	}

	if !found {
		fmt.Println("Not Found")
		return
	}

	saveJSON(newTasks)
	fmt.Println("Task Updated With Id : ", id)
	return
}

// markInProgressTask sets a task’s status to "IN_PROGRESS".
func markInProgressTask(id int) {
	file, err := os.ReadFile(FILEPATH)
	if err != nil {
		fmt.Println("No Tasks")
		return
	}

	tasks := []task{}
	if err := json.Unmarshal(file, &tasks); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	newTasks := []task{}
	found := false
	for _, t := range tasks {
		if t.Id != id {
			newTasks = append(newTasks, t)
		} else {
			found = true
			t.Status = in_progress
			t.Update_at = time.Now()
			newTasks = append(newTasks, t)
		}
	}

	if !found {
		fmt.Println("Not Found")
		return
	}

	saveJSON(newTasks)
	fmt.Println("Task Updated With Id : ", id)
	return
}

// markDoneTask sets a task’s status to "DONE".
func markDoneTask(id int) {
	file, err := os.ReadFile(FILEPATH)
	if err != nil {
		fmt.Println("No Tasks")
		return
	}

	tasks := []task{}
	if err := json.Unmarshal(file, &tasks); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	newTasks := []task{}
	found := false
	for _, t := range tasks {
		if t.Id != id {
			newTasks = append(newTasks, t)
		} else {
			found = true
			t.Status = done
			t.Update_at = time.Now()
			newTasks = append(newTasks, t)
		}
	}

	if !found {
		fmt.Println("Not Found")
		return
	}

	saveJSON(newTasks)
	fmt.Println("Task Updated With Id : ", id)
	return
}
