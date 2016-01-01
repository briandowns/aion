package main

// Jobber
type Jobber interface {
	Start() error
	Stop() error
	Status() *JobStatus
}

// JobStatus
type JobStatus struct {
	Job
	Status string
}

// NewJob creates a new reference to Job
func (j *Job) NewJob(name string) *Job {
	return &Job{
		Name: name,
	}
}
