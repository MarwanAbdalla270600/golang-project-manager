package task

import "time"

type TaskStatus string

const (
	StatusTodo       TaskStatus = "TODO"
	StatusInProgress TaskStatus = "IN_PROGRESS"
	StatusDone       TaskStatus = "DONE"
)

type Task struct {
	ID          uint       `json:"id"`
	Title       string     `json:"name"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	AssigneeID  int        `json:"assignee_id"`
	ProjectID   uint       `json:"project_id"`
	DueDate     time.Time  `json:"due_date"`
}
