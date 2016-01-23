package main

import (
	"log"

	"github.com/briandowns/aion/models"
)

// JobManager
type JobManager struct {
	Conf     *Config
	JobChan  chan models.Job
	TaskChan chan models.Task
	ExitChan chan struct{}
}

// NewJobManager
func NewJobManager(conf *Config) *JobManager {
	return &JobManager{
		Conf:     conf,
		JobChan:  make(chan models.Job),
		TaskChan: make(chan models.Task),
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
