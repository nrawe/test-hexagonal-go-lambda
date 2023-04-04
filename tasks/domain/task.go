package domain

import (
	"time"

	"github.com/segmentio/ksuid"
)

type TaskStatus string

var (
	statusOpen       TaskStatus = "open"
	statusClosed     TaskStatus = "closed"
	statusInProgress TaskStatus = "underway"
)

type Task struct {
	createdAt time.Time
	createdBy string
	id        string
	status    TaskStatus
	summary   string
	updatedAt time.Time
}

func NewTask(userId string, summary string) Task {
	return Task{
		createdAt: time.Now(),
		createdBy: userId,
		id:        ksuid.New().String(),
		status:    statusOpen,
		summary:   summary,
		updatedAt: time.Now(),
	}
}

func (t Task) Id() string {
	return t.id
}

func (t Task) IsComplete() bool {
	return t.status == statusClosed
}

func (t Task) IsOpen() bool {
	return t.status == statusOpen
}

func (t Task) IsInProgress() bool {
	return t.status == statusInProgress
}

func (t Task) CreatedAt() time.Time {
	return t.createdAt
}

func (t Task) CreatedBy() string {
	return t.createdBy
}

func (t Task) Summary() string {
	return t.summary
}

func (t Task) UpdatedAt() time.Time {
	return t.updatedAt
}
