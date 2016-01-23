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

// Sender is an interface for sending data to NSQ
type Sender interface {
	Send() error
}

// Adder is an interface for adding data to the database
type Adder interface {
	Add() error
}

// QProducerConn connects to NSQ for sending data
func QProducerConn() (*nsq.Producer, error) {
	return nsq.NewProducer(fmt.Sprintf("%s:4150", queueHostFlag), nsqConfig)
}

func watchForNewJobs() error {
	q, err := nsq.NewConsumer("new_job", "add", nsqConfig)
	if err != nil {
		return err
	}

	var j *Job

	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		json.Unmarshal(message.Body, &j)

		db, err := NewDatabase(Conf)
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
	var t *Task
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		json.Unmarshal(message.Body, &t)

		db, err := NewDatabase(Conf)
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
