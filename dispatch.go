package main

import (
	"log"
	"os/exec"
	"strings"

	"github.com/robfig/cron"
)

// Dispatcher holds the values that comprise the Aion dispatcher
type Dispatcher struct {
	Conf         *Config
	cron         *cron.Cron
	ResultChan   chan []byte
	NewJobChan   chan Job
	JobProcChan  chan Job
	NewTaskChan  chan Task
	TaskProcChan chan Task
}

// NewDispatcher creates a new refence of type Dispatcher
func NewDispatcher(conf *Config) *Dispatcher {
	return &Dispatcher{
		Conf:        conf,
		cron:        cron.New(),
		ResultChan:  make(chan []byte),
		NewJobChan:  make(chan Job),
		NewTaskChan: make(chan Task),
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
func (d *Dispatcher) AddExistingTasks() {
	db, err := NewDatabase(d.Conf)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Conn.Close()

	tasks := db.GetTasks()
	for _, task := range tasks {
		cmdFunc, err := d.generateTaskFunc(task.CMD, strings.Split(task.Args, ","))
		if err != nil {
			log.Println(err)
		}
		d.cron.AddFunc(task.Schedule, cmdFunc)
	}
}

// Run starts the dispatcher
func (d *Dispatcher) Run() error {
	db, err := NewDatabase(d.Conf)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Conn.Close()
	d.cron.Start()
	defer d.cron.Stop()

	for {
		select {
		case job := <-d.NewJobChan:
			db.AddJob(job)
		case task := <-d.NewTaskChan:
			db.AddTask(task)
		}
	}
}

// ResultWorkers starts the result workers
func (d *Dispatcher) ResultWorkers() {
}
