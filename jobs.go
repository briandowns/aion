package main

// JobsDB is an interface so that we can use any number of
// databases on the backend to store jobs
type JobsDB interface{}

// Jobber
type Jobber interface {
	Start()
	Stop()
	Status()
}

// Job
type Job struct{}

// NewJob creates a new reference to Job
func (j *Job) NewJob() *Job {
	return &Job{}
}
