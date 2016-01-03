package main

import (
	"os/exec"
	"strings"

	"github.com/robfig/cron"
)

// Dispatcher
type Dispatcher struct {
	Conf *Config
	cron *cron.Cron
	ResultChan  chan []byte
	NewTaskChan chan Task
	NewJobChan chan Job
}

// NewDispatcher creates a new refence of type Dispatcher
func NewDispatcher(conf *Config) *Dispatcher {
	return &Dispatcher{
		Conf:       conf,
		cron: cron.New(),
		ResultChan: make(chan []byte),
	}
}

// generateTaskFunc generates a function suitable for use in the scheduler
func (d *Dispatcher) generateTaskFunc(cmd string, args []string) (func(), error) {
	var f func()
	cmdOut, err := exec.Command(cmd, args...).Output()
	if err != nil {
		return nil, err
	}
	d.ResultChan <- cmdOut
	return f, nil
}

// AddExistingTasks adds active tasks from the database to the scheduler
func () AddExistingTasks() {
	db, err := NewDatabase(d.Conf)
	if err != nil {
		return err
	}
	defer db.Conn.Close()

	tasks := db.GetTasks()
	for _, task := range tasks {
		cmdFunc, err := d.generateTaskFunc(task.CMD, strings.Split(task.Args, ","))
		if err != nil {
			return err
		}
		d.cron.AddFunc(task.Schedule, cmdFunc)
	}
}

// Run
func (d *Dispatcher) Run() error {
	d.cron.Start()
	defer d.cron.Stop()

	for {
		select {
		case task := <-d.NewTaskChan:
			//
		}
	}
}
