package controllers

import "errors"

var (
	// ErrNoJobsFound is given when a job isn't found
	ErrNoJobsFound = errors.New("no jobs found")

	// ErrNoTasksFound is given when a task isn't found
	ErrNoTasksFound = errors.New("no tasks found")

	// ErrMissingNameField is given when the name field in a task is missing
	ErrMissingNameField = errors.New("missing or empty 'name' field")

	// ErrMissingCMDField is given when the cmd field in a task or job is missing
	ErrMissingCMDField = errors.New("missing or empty 'cmd' field")

	// ErrMissingArgsField is given when the name args in a task is missing
	ErrMissingArgsField = errors.New("missing or empty 'args' field")

	// ErrMissingScheduleField is given when the schedule field in a task is missing
	ErrMissingScheduleField = errors.New("missing or empty 'schedule' field")
)
