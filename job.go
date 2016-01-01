package main

import (
	"code.google.com/p/go-uuid/uuid"
)

// Jobber
type Jobber interface {
	Start() error
	Stop() error
	Status() *JobStatus
}

// Job
type Job struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Tasks []Task `json:"tasks"`
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
