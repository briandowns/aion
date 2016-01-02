package main

import (
	"log"
)

// JobManager
type JobManager struct {
	Conf     *Config
	JobChan  chan Job
	TaskChan chan Task
	ExitChan chan struct{}
}

// NewJobManager
func NewJobManager(conf *Config) *JobManager {
	return &JobManager{
		Conf:     conf,
		JobChan:  make(chan Job),
		TaskChan: make(chan Task),
		ExitChan: make(chan struct{}),
	}
}

// Run
func (j *JobManager) Run() {
	db, err := NewDatabase(j.Conf)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		select {
		case data := <-j.JobChan:
			db.AddJob(data)
		case data := <-j.TaskChan:
			db.AddTask(data)
		case <-j.ExitChan:
			return
		}
	}
}
