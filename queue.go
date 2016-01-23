package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/briandowns/aion/database"

	"github.com/bitly/go-nsq"
)

var (
	newJobChan    = make(chan *database.Job)
	newTaskChan   = make(chan *database.Task)
	newResultChan = make(chan *database.Result)
)

var nsqConfig = nsq.NewConfig()

// Sender is an interface for sending data to NSQ
type Sender interface {
	Send(db *database.Database) error
}

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
		return nil
	}))
	err = q.ConnectToNSQD(fmt.Sprintf("%s:4150", queueHostFlag))
	if err != nil {
		return err
	}

	return nil
}
