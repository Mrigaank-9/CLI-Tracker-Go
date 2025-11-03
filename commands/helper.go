package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Write tasks back to JSON file
func saveJSON(tasks []task) {
	newFile, err := os.Create(FILEPATH)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	encoder := json.NewEncoder(newFile)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(tasks); err != nil {
		panic(err)
	}
}

// Call the functions for operations
func doOperation(operation Operations, tokens []string) {
	if operation == "LIST" {
		if len(tokens) == 2 {
			showAll("")
		} else {
			showAll(ListFilter(tokens[2]))
		}
	} else if operation == "UPDATE" {
		idStr := tokens[2]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Invalid ID:", idStr)
			return
		}

		description := ""
		for i := 3; i < len(tokens); i++ {
			if tokens[i] == `"` {
				continue
			}
			description += tokens[i] + " "
		}
		description = strings.TrimSpace(description)

		updateTask(id, description)
	} else if operation == "ADD" {
		description := ""
		for i := 2; i < len(tokens); i++ {
			if tokens[i] == `"` {
				continue
			}
			description += tokens[i] + " "
		}
		description = strings.TrimSpace(description)
		fmt.Println(description)
		addTask(description)
	} else if operation == "DELETE" {
		idStr := tokens[2]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Invalid ID:", idStr)
			return
		}
		deleteTask(id)
	} else if operation == "MARK_IN_PROGRESS" {
		idStr := tokens[2]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Invalid ID:", idStr)
			return
		}
		markInProgressTask(id)
	} else {
		idStr := tokens[2]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Invalid ID:", idStr)
			return
		}
		markDoneTask(id)
	}
}

// Extract Operation from CLI
func getOperation(tokens []string) string {
	operation := strings.ToUpper(strings.ReplaceAll(tokens[1], "-", "_"))
	fmt.Println(operation)

	switch Operations(operation) {
	case add, update, delete, list, mark_in_progress, mark_done:
		return operation
	default:
		return "INVALID"
	}
}

// Extract each token in the CLI
func extractToken(cli string) []string {
	tokens := []string{}
	current_token := ""

	for _, char := range cli {
		value := string(char)
		if value == " " {
			tokens = append(tokens, current_token)
			current_token = ""
		} else {
			current_token += value
		}
	}

	if current_token != " " {
		tokens = append(tokens, current_token)
	}
	if len(tokens) < 2 {
		return nil
	}
	return tokens
}

// Normalize the Status
func normalizeStatus(s string) string {
	s = strings.ToUpper(s)
	s = strings.ReplaceAll(s, "-", "")
	return s
}
