package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/briandowns/aion/database"
	"github.com/briandowns/aion/dispatcher"

	"github.com/bitly/go-nsq"
)

var (
	newJobChan    = make(chan *database.Job)
	newTaskChan   = make(chan *database.Task)
	newResultChan = make(chan *database.Result)
)

var nsqConfig = nsq.NewConfig()

// Adder is an interface for adding data to the database
type Adder interface {
	Add() error
}

func watchForNewJobs() error {
	q, err := nsq.NewConsumer("new_job", "add", nsqConfig)
	if err != nil {
		return err
	}

	var j *database.Job

	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		json.Unmarshal(message.Body, &j)

		db, err := database.NewDatabase(Conf)
		if err != nil {
			log.Println(err)
		}
		db.AddJob(*j)
		return nil
	}))
	err = q.ConnectToNSQD(fmt.Sprintf("%s:4150", queueHostFlag))
	if err != nil {
		return err
	}

	return nil
}

func watchForNewTasks() error {
	q, err := nsq.NewConsumer("new_task", "add", nsqConfig)
	if err != nil {
		return err
	}
	var t *database.Task
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		json.Unmarshal(message.Body, &t)

		db, err := database.NewDatabase(Conf)
		if err != nil {
			log.Println(err)
		}
		db.AddTask(*t)
		log.Print("made it this far...\n")
		dispatch := dispatcher.NewDispatcher(Conf)
		dispatch.TaskProcChan <- *t
		return nil
	}))
	err = q.ConnectToNSQD(fmt.Sprintf("%s:4150", queueHostFlag))
	if err != nil {
		return err
	}

	return nil
}

func watchForNewResults() error {
	q, err := nsq.NewConsumer("new_result", "add", nsqConfig)
	if err != nil {
		return err
	}
	var r *database.Result
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		json.Unmarshal(message.Body, &r)

		db, err := database.NewDatabase(Conf)
		if err != nil {
			log.Println(err)
		}
		db.AddResult(*r)
		return nil
	}))
	err = q.ConnectToNSQD(fmt.Sprintf("%s:4150", queueHostFlag))
	if err != nil {
		return err
	}

	return nil
}

func watchForNewCommands() error {
	q, err := nsq.NewConsumer("new_command", "add", nsqConfig)
	if err != nil {
		return err
	}
	var c *database.Command
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		json.Unmarshal(message.Body, &c)

		db, err := database.NewDatabase(Conf)
		if err != nil {
			log.Println(err)
		}
		db.AddCommand(*c)
		return nil
	}))
	err = q.ConnectToNSQD(fmt.Sprintf("%s:4150", queueHostFlag))
	if err != nil {
		return err
	}

	return nil
}
