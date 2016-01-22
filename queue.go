package main

import (
	"encoding/json"
	"fmt"

	"github.com/bitly/go-nsq"
)

var (
	newJobChan    = make(chan *Job)
	newTaskChan   = make(chan *Task)
	newResultChan = make(chan *Result)
)

var nsqConfig = nsq.NewConfig()

// Sender is an interface for sending data to NSQ
type Sender interface {
	Send() error
}

// QProducerConn connects to NSQ for sending data
func QProducerConn() (*nsq.Producer, error) {
	return nsq.NewProducer(fmt.Sprintf("%s:4150", queueHostFlag), nsqConfig)
}

// Send sends a new job to NSQ
func (j *Job) Send() error {
	w, err := QProducerConn()
	if err != nil {
		return nil
	}

	s, err := json.Marshal(j)
	if err != nil {
		return err
	}

	err = w.Publish("new_job", s)
	if err != nil {
		return err
	}
	w.Stop()

	return nil
}

// Send sends a new task to NSQ
func (t *Task) Send() error {
	w, err := QProducerConn()
	if err != nil {
		return nil
	}

	s, err := json.Marshal(t)
	if err != nil {
		return err
	}

	err = w.Publish("new_task", s)
	if err != nil {
		return err
	}
	w.Stop()

	return nil
}

// Send sends a new result to NSQ
func (r *Result) Send() error {
	w, err := QProducerConn()
	if err != nil {
		return nil
	}

	s, err := json.Marshal(r)
	if err != nil {
		return err
	}

	err = w.Publish("new_result", s)
	if err != nil {
		return err
	}
	w.Stop()

	return nil
}
