package database

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/bitly/go-nsq"
)

var nsqConfig = nsq.NewConfig()

// Job holds what's needs to represent a job
type Job struct {
	ID    int    `sql:"auto_increment" gorm:"column:id" gorm:"primary_key" json:"id"`
	Name  string `gorm:"column:name" json:"name"`
	Desc  string `gorm:"column:desc" json:"desc"`
	Tasks string `gorm:"column:tasks" json:"tasks"`
}

// NewJob creates a new reference to Job
func (j *Job) NewJob(name string) *Job {
	return &Job{
		Name: name,
	}
}

// Add adds a job to the database
func (j *Job) Add(db *Database) error {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	q, err := nsq.NewConsumer("new_job", "add", nsqConfig)
	if err != nil {
		return err
	}

	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		json.Unmarshal(message.Body, &j)
		log.Printf("Got a message: %+v", j)
		db.AddJob(*j)
		wg.Done()
		return nil
	}))
	err = q.ConnectToNSQD(fmt.Sprintf("%s:4150", db.Conf.QueueHost))
	if err != nil {
		return err
	}
	wg.Wait()

	return nil
}

// Send sends a new job to NSQ
func (j *Job) Send(db *Database) error {
	w, err := nsq.NewProducer(fmt.Sprintf("%s:4150", db.Conf.QueueHost), nsqConfig)
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

// Task holds what's needs to represent a task
type Task struct {
	ID       int    `sql:"auto_increment" gorm:"column:id" gorm:"primary_key" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	Desc     string `gorm:"column:desc" json:"desc"`
	CMD      string `gorm:"column:cmd" json:"cmd"`
	Args     string `gorm:"column:args" json:"args"`
	Schedule string `gorm:"column:schedule" json:"schedule"`
	Func     func() `sql:"-" json:"-"`
}

// Add adds a task to the database
func (t *Task) Add(db *Database) error {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	q, err := nsq.NewConsumer("new_task", "add", nsqConfig)
	if err != nil {
		return err
	}

	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		json.Unmarshal(message.Body, &t)
		log.Printf("Got a message: %+v", t)
		db.AddTask(*t)
		wg.Done()
		return nil
	}))
	err = q.ConnectToNSQD("192.168.99.100:4150")
	if err != nil {
		return err
	}
	wg.Wait()

	return nil
}

// Send sends a new task to NSQ
func (t *Task) Send(db *Database) error {
	w, err := nsq.NewProducer(fmt.Sprintf("%s:4150", db.Conf.QueueHost), nsqConfig)
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

// User holds what's needs to represent a user
type User struct {
	ID        int    `sql:"auto_increment" gorm:"column:id" gorm:"primary_key" json:"id"`
	Username  string `gorm:"column:username" json:"username"`
	Password  string `gorm:"column:password" json:"password"`
	FirstName string `gorm:"column:first_name" json:"first_name"`
	LastName  string `gorm:"column:last_name" json:"last_name"`
	Email     string `gorm:"column:email" json:"email"`
	Enabled   bool   `gorm:"column:Enabled" json:"Enabled"`
}

// Permission holds what's needs to represent a permission
type Permission struct {
	ID          int    `sql:"auto_increment" gorm:"column:id" gorm:"primary_key" json:"id"`
	Permission  int    `gorm:"column:permission" json:"permission"`
	Description string `gorm:"column:description" json:"description"`
}

// Result holds what's needs to represent a result
type Result struct {
	ID        int       `sql:"auto_increment" gorm:"column:id" gorm:"primary_key" json:"id"`
	TaskID    int       `gorm:"column:task_id" json:"task_id"`
	StartTime time.Time `gorm:"column:start_time" json:"start_time"`
	EndTime   time.Time `gorm:"column:end_time" json:"end_time"`
	Result    []byte    `sql:"type:longblob" gorm:"column:result" json:"result"`
	Error     string    `gorm:"column:error" json:"error"`
}

// Send sends a new result to NSQ
func (r *Result) Send(db *Database) error {
	w, err := nsq.NewProducer(fmt.Sprintf("%s:4150", db.Conf.QueueHost), nsqConfig)
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

// Add adds a task to the database
func (r *Result) Add(db *Database) error {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	q, err := nsq.NewConsumer("new_result", "add", nsqConfig)
	if err != nil {
		return err
	}

	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		json.Unmarshal(message.Body, &r)
		log.Printf("Got a message: %+v", r)
		db.AddResult(*r)
		wg.Done()
		return nil
	}))
	err = q.ConnectToNSQD("192.168.99.100:4150")
	if err != nil {
		return err
	}
	wg.Wait()

	return nil
}

// UserSession represents an active session for an Aion user
type UserSession struct {
	SessionKey   string    `gorm:"primary_key" gorm:"column:session_key" json:"session_key"`
	UserID       int       `gorm:"column:user_id" json:"user_id"`
	LoginTime    time.Time `gorm:"column:login_time" json:"login_time"`
	LastSeenTime time.Time `gorm:"column:last_seen_time" json:"last_seen_time"`
}
