package main

import (
	"sync"

	"github.com/briandowns/aion/models"
)

// JobStatuser
type JobStatuser interface {
	Enable()
	Disable()
}

// JobStatus holds the
type JobStatus struct {
	models.Job
	Status bool
	sync.Mutex
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
