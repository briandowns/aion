package main

import (
	"log"

	"github.com/briandowns/aion/database"
)

// JobManager
type JobManager struct {
	Conf     *Config
	JobChan  chan database.Job
	TaskChan chan database.Task
	ExitChan chan struct{}
}

// NewJobManager
func NewJobManager(conf *Config) *JobManager {
	return &JobManager{
		Conf:     conf,
		JobChan:  make(chan database.Job),
		TaskChan: make(chan database.Task),
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
