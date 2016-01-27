package controllers

import "errors"

// Top Level/Primary Routes
var (
	FrontEnd = "/"
)

const (
	// APIBase is the base path for API access
	APIBase = "/api/v1/"

	// JobsPath is the path to manage jobs
	JobsPath = APIBase + "job"

	// TasksPath is the path to manage tasks
	TasksPath = APIBase + "task"

	// CommandsPath is the path to manage tasks
	CommandsPath = APIBase + "command"

	// UserPath is the path to manage users
	UserPath = APIBase + AdminPath + "user"

	// AdminPath is the path to manage Aion
	AdminPath = APIBase + "admin/"
)

var (
	// JobByID is the path to get specific job data
	JobByID = JobsPath + "/{id}"

	// TaskByID is the path to get specific task data
	TaskByID = TasksPath + "/{id}"

	// CommandByID is the path to get specific task data
	CommandByID = CommandsPath + "/{id}"

	// UserByID is the path to get specific user data
	UserByID = UserPath + "/{id}"

	// APIStats is hte path to get API specific data
	APIStats = AdminPath + "api/stats"
)

var (
	// ErrNoJobsFound is given when a job isn't found
	ErrNoJobsFound = errors.New("no jobs found")

	// ErrNoTasksFound is given when a task isn't found
	ErrNoTasksFound = errors.New("no tasks found")

	// ErrNoCommandsFound is given when a command isn't found
	ErrNoCommandsFound = errors.New("no commands found")

	// ErrNoEntryFound is given when an entry isn't found in the database
	ErrNoEntryFound = errors.New("no entry found")

	// ErrMissingNameField is given when the name field in a task is missing
	ErrMissingNameField = errors.New("missing or empty 'name' field")

	// ErrMissingCMDField is given when the cmd field in a task or job is missing
	ErrMissingCMDField = errors.New("missing or empty 'cmd' field")

	// ErrMissingArgsField is given when the name args in a task is missing
	ErrMissingArgsField = errors.New("missing or empty 'args' field")

	// ErrMissingScheduleField is given when the schedule field in a task is missing
	ErrMissingScheduleField = errors.New("missing or empty 'schedule' field")
)
