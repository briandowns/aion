package main

import (
	"code.google.com/p/go-uuid/uuid"
)

// Tasker
type Tasker interface {
	Execute() error
}

// Task
type Task struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	Exec     string `json:"exec"`
	Schedule string `json:"schedule"`
}

// NewTask
func NewTask(name, desc, exec, sched string) *Task {
	return &Task{
		ID:       uuid.NewUUID().String(),
		Name:     name,
		Desc:     desc,
		Exec:     exec,
		Schedule: sched,
	}
}

func (t *Task) Execute() error {
	return nil
}
