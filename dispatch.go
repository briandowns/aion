package main

import (
	"os/exec"

	"github.com/robfig/cron"
)

// Dispatcher
type Dispatcher struct {
	Conf       *Config
	ResultChan chan []byte
}

// NewDispatcher
func NewDispatcher(conf *Config) *Dispatcher {
	return &Dispatcher{
		Conf:       conf,
		ResultChan: make(chan []byte),
	}
}

// generateTaskFunc
func (d *Dispatcher) generateTaskFunc(cmd string, args []string) (func(), error) {
	var f func()
	cmdOut, err := exec.Command(cmd, args...).Output()
	if err != nil {
		return nil, err
	}
	d.ResultChan <- cmdOut
	return f, nil
}

// Run
func (d *Dispatcher) Run() error {
	db, err := NewDatabase(d.Conf)
	if err != nil {
		return err
	}
	c := cron.New()
	tasks := db.GetTasks()
	for _, task := range tasks {
		cmdFunc, err := d.generateTaskFunc(task.Exec, []string{})
		if err != nil {
			return err
		}
		c.AddFunc(task.Schedule, cmdFunc)
	}
	c.Start()
	defer c.Stop()
	return nil
}
