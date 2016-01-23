package main

import (
	"log"

	"github.com/briandowns/aion/config"
	"github.com/briandowns/aion/database"
)

// JobManager
type JobManager struct {
	Conf     *config.Config
	JobChan  chan database.Job
	TaskChan chan database.Task
	ExitChan chan struct{}
}

// NewJobManager creates a new job manager
func NewJobManager(conf *config.Config) *JobManager {
	return &JobManager{
		Conf:     conf,
		JobChan:  make(chan database.Job),
		TaskChan: make(chan database.Task),
		ExitChan: make(chan struct{}),
	}
}

// Run runs the job manager
func (j *JobManager) Run() {
	db, err := database.NewDatabase(j.Conf)
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
