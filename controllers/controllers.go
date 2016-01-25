package controllers

import "errors"

// Top Level/Primary Routes
var (
	frontEnd = "/"
)

const (
	// APIBase is the base path for API access
	APIBase = "/api/v1/"

	// JobsPath is the path to manage jobs
	JobsPath = APIBase + "job"

	// TasksPath is the path to manage tasks
	TasksPath = APIBase + "task"

	// AdminPath is the path to manage Aion
	AdminPath = APIBase + "admin/"

	// UserPath is the path to manage users
	UserPath = APIBase + AdminPath + "user"
)

var (
	// JobByID is the path to get specific job data
	JobByID = JobsPath + "/{id}"

	// TaskByID is the path to get specific task data
	TaskByID = TasksPath + "/{id}"

	// UserByID is the path to get specific user data
	UserByID = UserPath + "/{id}"

	// APIStats is hte path to get API specific data
	APIStats = UserPath + "api/stats"
)

// ErrNoJobsFound given when a job isn't found
var ErrNoJobsFound = errors.New("no jobs found")

// ErrNoTasksFound given when a task isn't found
var ErrNoTasksFound = errors.New("no tasks found")
