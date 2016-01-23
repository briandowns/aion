package main

import (
	"log"
	"os/exec"
	"strings"

	"github.com/briandowns/aion/config"
	"github.com/briandowns/aion/database"
	"github.com/robfig/cron"
)

// Dispatcher holds the values that comprise the Aion dispatcher
type Dispatcher struct {
	Conf         *config.Config
	cron         *cron.Cron
	ResultChan   chan []byte
	JobProcChan  chan database.Job
	TaskProcChan chan database.Task
	SenderChan   chan Sender
}

// NewDispatcher creates a new refence of type Dispatcher
func NewDispatcher(conf *config.Config) *Dispatcher {
	return &Dispatcher{
		Conf:       conf,
		cron:       cron.New(),
		ResultChan: make(chan []byte),
		SenderChan: make(chan Sender),
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
	db, err := database.NewDatabase(d.Conf)
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
	db, err := database.NewDatabase(d.Conf)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Conn.Close()
	d.cron.Start()
	defer d.cron.Stop()

	for {
		select {
		case data := <-d.SenderChan:
			if err := data.Send(db); err != nil {
				log.Println(err)
			}
		}
	}
}

// ResultWorkers starts the result workers
func (d *Dispatcher) ResultWorkers() {
}
