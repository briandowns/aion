package main

// Tasker
type Tasker interface {
	Execute() error
}

// NewTask
func NewTask(name, desc, exec, sched string) *Task {
	return &Task{
		Name:     name,
		Desc:     desc,
		Exec:     exec,
		Schedule: sched,
	}
}

func (t *Task) Execute() error {
	return nil
}
