package domain_test

import (
	"testing"

	"github.com/nrawe/tasks-kata/tasks/domain"
	"github.com/stretchr/testify/assert"
)

func TestTaskCreation(t *testing.T) {
	task := domain.NewTask("user-1", "my task")

	assert.NotEmpty(t, task.Id(), "task id")
	assert.NotEmpty(t, task.CreatedAt(), "task id")
	assert.NotEmpty(t, task.UpdatedAt(), "task id")
	assert.Equal(t, "user-1", task.CreatedBy(), "task owner set")
	assert.Equal(t, "my task", task.Summary(), "task summary set")
	assert.True(t, task.IsOpen(), "task is open")
	assert.False(t, task.IsComplete(), "task is not complete")
	assert.False(t, task.IsInProgress(), "task is not in progress")
}
