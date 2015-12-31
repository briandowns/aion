package main

// Tasker
type Tasker interface{}

// Task
type Task struct{}

// NewTask
func NewTask() *Task {
	return &Task{}
}
