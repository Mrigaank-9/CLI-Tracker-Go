package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// getCli parses a single CLI command input (like "task-cli add 'Buy milk'"),
// validates its structure, determines the operation type, and triggers the corresponding action.
func getCli(cli string) {
	tokens := extractToken(cli)
	fmt.Println(tokens)
	if tokens == nil {
		fmt.Println("Invalid CLI Command")
		return
	}
	if strings.ToLower(tokens[0]) != "task-cli" {
		fmt.Println("Invalid CLI Command")
		return
	}
	operation := getOperation(tokens)
	if operation == "INVALID" {
		fmt.Println("Wrong Operation")
		return
	}

	doOperation(Operations(operation), tokens)
}

// Command starts an interactive CLI loop that continuously reads user input,
// parses commands like `task-cli add`, `task-cli list`, etc., and executes them
// until the user types "exit".
func Command() {

	for {
		fmt.Print("> ")
		var cli string
		reader := bufio.NewReader(os.Stdin)
		cli, _ = reader.ReadString('\n')
		cli = strings.TrimSpace(cli)

		if cli == "exit" {
			break
		}
		getCli(cli)
	}
}
