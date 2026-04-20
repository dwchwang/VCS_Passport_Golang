package models

import "time"

type TaskStatus string
type TaskPriority string

const (
	StatusTodo       TaskStatus = "todo"
	StatusInProgress TaskStatus = "in_progress"
	StatusDone       TaskStatus = "done"

	PriorityLow    TaskPriority = "low"
	PriorityMedium TaskPriority = "medium"
	PriorityHigh   TaskPriority = "high"
)

type Task struct {
	ID          uint         `gorm:"primaryKey"            json:"id"`
	Title       string       `gorm:"not null"              json:"title"`
	Description string       `                             json:"description"`
	Status      TaskStatus   `gorm:"default:todo"          json:"status"`
	Priority    TaskPriority `gorm:"default:medium"        json:"priority"`
	ProjectID   uint         `gorm:"not null"              json:"project_id"`
	Project     *Project     `gorm:"foreignKey:ProjectID"  json:"project,omitempty"`
	AssigneeID  *uint        `                             json:"assignee_id"`
	Assignee    *User        `gorm:"foreignKey:AssigneeID" json:"assignee,omitempty"`
	DueDate     *time.Time   `                             json:"due_date"`
	CreatedAt   time.Time    `                             json:"created_at"`
	UpdatedAt   time.Time    `                             json:"updated_at"`
}
