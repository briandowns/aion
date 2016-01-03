package main

// Tasker
type Tasker interface {
	Execute() error
}

// NewTask
func NewTask(name, desc, cmd, args, sched string) *Task {
	return &Task{
		Name:     name,
		Desc:     desc,
		CMD:      cmd,
		Args:     args,
		Schedule: sched,
	}
}

func (t *Task) Execute() error {
	return nil
}
