# CLI Task Tracker (Go)

A simple command-line tool to manage your daily tasks using Go.
You can add, update, delete, and list tasks directly from your terminal.
project page URL - https://roadmap.sh/projects/task-tracker

---

## Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/Mrigaank-9/CLI-Tracker-Go.git
   cd CLI-Tracker-Go
   ```

2. Run the project:

   ```bash
   go run .\main.go
   ```

---

## Commands

### Adding a new task

```bash
> task-cli add "Buy groceries"
```

### Updating and deleting tasks

```bash
> task-cli update 1 "Buy groceries and cook dinner"
> task-cli delete 1
```

### Marking a task as in progress or done

```bash
> task-cli mark-in-progress 1
> task-cli mark-done 1
```

### Listing all tasks

```bash
> task-cli list
```

### Listing tasks by status

```bash
> task-cli list done
> task-cli list todo
> task-cli list in-progress
```

---

## Exit Command

```bash
> exit
```

This will close the CLI session.

---

## Data Storage

All tasks are stored locally in a `task.json` file located in the project directory.

---

## Example Output

```bash
ID: 1
Description: Buy groceries and cook dinner
Status: DONE
Created: 2025-11-03 14:20
Updated: 2025-11-03 15:10
```

---

## Project Structure

```bash
CLI-Tracker-Go/
│
├── commands/
│   ├── commands.go
│   ├── operations.go
│   ├── consts.go
│   ├── helper.go
│   └── models.go
│
├── task.json
├── main.go
└── go.mod
```
