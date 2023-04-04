package tasks1

import "errors"

var (
	ErrorUnauthorized  = errors.New("unauthorized")
	ErrorInvalidLimit  = errors.New("limit must be between 1 and 10")
	ErrorInvalidCursor = errors.New("after and before cannot both be specified")
)

type taskApp struct {
	tr     TaskRepository
	userId string
}

// NewApp returns a pointer to a task application instance.
func NewApp(userId string, tr TaskRepository) taskApp {
	return taskApp{
		tr:     tr,
		userId: userId,
	}
}

// GetTaskList is a use case for the
func (a taskApp) GetTaskList(limit int8, after, before TaskId) (TaskList, error) {
	if a.userId == "" {
		return TaskList{}, ErrorUnauthorized
	}

	if limit == 0 {
		limit = 5
	}

	if limit < 1 || limit > 10 {
		return TaskList{}, ErrorInvalidLimit
	}

	if after != "" && before != "" {
		return TaskList{}, ErrorInvalidCursor
	}

	return a.tr.GetTaskList(
		after,
		before,
		limit,
		a.userId,
	)
}
