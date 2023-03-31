package tasks

import (
	"net/http"

	"github.com/nrawe/test-hexagonal-go-lambda/internal"
)

// Problem templates for use by the application.
var (
	ProblemBadRequest = internal.Problem{
		Title:  "Bad Request",
		Type:   "https://cooltasksapi.com/problems/bad-request",
		Status: http.StatusBadRequest,
	}
)

// TaskListQuery query represents a request to return a paginated set of
// tasks for the given user.
type TaskListQuery struct {
	After  TaskId
	Before TaskId
	Limit  int8
	UserId string
}

func (o TaskListQuery) Validate() error {
	return nil
}

// Application main interface for working with tasks.
type Application struct {
	Repository TaskRepository
}

// Run is the input port for the application, allowing for the dispatch of many
// different operations related to tasks.
func (a Application) Run(op internal.Operation) (interface{}, error) {
	err := op.Validate()

	if err != nil {
		return ProblemBadRequest, err
	}

	if op, ok := op.(TaskListQuery); ok {
		return a.runTaskListQuery(op)
	}

	return ProblemBadRequest, nil
}

func (a Application) runTaskListQuery(op TaskListQuery) (interface{}, error) {
	return a.Repository.GetTaskList(
		op.After,
		op.Before,
		op.Limit,
		op.UserId,
	)
}
