package main

import (
	"sync"
)

// Jobber
type Jobber interface {
	Enable()
	Disable()
}

// JobStatus
type JobStatus struct {
	Job
	Status bool
	sync.Mutex
}

// NewJob creates a new reference to Job
func (j *Job) NewJob(name string) *Job {
	return &Job{
		Name: name,
	}
}

// Enable enables an unactive job
func (j *JobStatus) Enable() {
	switch j.Status {
	case true:
		return
	case false:
		j.Lock()
		defer j.Unlock()
		j.Status = true
	}
}

// Disable disables an active job
func (j *JobStatus) Disable() {
	switch j.Status {
	case false:
		return
	case true:
		j.Lock()
		defer j.Unlock()
		j.Status = false
	}
}
