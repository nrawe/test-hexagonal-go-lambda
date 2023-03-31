package tasks

import "time"

type TaskId string

type TaskStatus string

var (
	TaskStatusOpen     TaskStatus = "open"
	TaskStatusClosed              = "closed"
	TaskStatusUnderway            = "underway"
)

type Task struct {
	id      TaskId
	status  TaskStatus
	summary string
	created time.Time
	updated time.Time
}

type TaskList struct {
	tasks      []Task
	pagination Paginator
}

type Paginator struct {
	cursor TaskId
}

type TaskRepository interface {
	GetTaskList(
		after TaskId,
		before TaskId,
		limit int8,
		userId string,
	) (TaskList, error)
}
