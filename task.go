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
	ID   string
	Name string
	Desc string
	Exec string
}

// NewTask
func NewTask(name, desc, exec string) *Task {
	return &Task{
		ID:   uuid.NewUUID().String(),
		Name: name,
		Desc: desc,
		Exec: exec,
	}
}

func (t *Task) Execute() error {
	return nil
}
