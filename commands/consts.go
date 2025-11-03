package commands

// FILEPATH for JSON
const FILEPATH = "task.json"

// Operations represents all possible CLI task operations
type Operations string

const (
	update           Operations = "UPDATE"
	list             Operations = "LIST"
	delete           Operations = "DELETE"
	add              Operations = "ADD"
	mark_in_progress Operations = "MARK_IN_PROGRESS"
	mark_done        Operations = "MARK_DONE"
)

// ListFilter defines task status filters for listing tasks.
type ListFilter string

const (
	done        ListFilter = "DONE"
	todo        ListFilter = "TODO"
	in_progress ListFilter = "INPROGRESS"
)
