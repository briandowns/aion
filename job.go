package main

import (
	"code.google.com/p/go-uuid/uuid"
)

// JobsDB is an interface so that we can use any number of
// databases on the backend to store jobs
type JobsDB interface{}

// Jobber
type Jobber interface {
	Start() error
	Stop() error
	Status() *JobStatus
}

// Job
type Job struct {
	ID    string
	Name  string
	Tasks []Task
}

// JobStatus
type JobStatus struct {
	Job
	Status string
}

// NewJob creates a new reference to Job
func (j *Job) NewJob(name string) *Job {
	return &Job{
		ID:   uuid.NewUUID().String(),
		Name: name,
	}
}
