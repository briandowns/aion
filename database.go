package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Database holds db conf and a connection
type Database struct {
	User string
	Pass string
	Host string
	DB   string
	Conn *gorm.DB
}

// NewDatabase creates a new Database object
func NewDatabase(user, pass, host, db string) (*Database, error) {
	d := &Database{
		User: user,
		Pass: pass,
		Host: host,
		DB:   db,
	}
	if err := d.connect(); err != nil {
		return nil, err
	}
	return d, nil
}

// Connect will provide the caller with a db connection
func (d *Database) connect() error {
	db, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=%s&charset=utf8&parseTime=True&loc=Local",
			d.User, d.Pass, d.Host, 3306, d.DB, "60s"))
	if err != nil {
		return err
	}
	db.LogMode(false)
	d.Conn = &db
	return nil
}

// AddJob adds a new job record to the database
func (d *Database) AddJob(j Job) {
	d.Conn.NewRecord(j)
	d.Conn.Create(&j)
	d.Conn.NewRecord(j)
}

// GetJobs gets all jobs from the database
func (d *Database) GetJobs() ([]Job, error) {
	var data []Job
	d.Conn.Find(&data)
	return data, nil
}

// AddTask adds a new task record to the database
func (d *Database) AddTask(t Task) {
	d.Conn.NewRecord(t)
	d.Conn.Create(&t)
	d.Conn.NewRecord(t)
}

// GetTasks gets all tasks from the database
func (d *Database) GetTasks() ([]Task, error) {
	var data []Task
	d.Conn.Find(&data)
	return data, nil
}

// Setup ...sets up the database
func (d *Database) Setup() {
	log.Println("Aion database setup starting...")
	d.Conn.CreateTable(&Job{})
	d.Conn.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Job{})
	d.Conn.CreateTable(&Task{})
	d.Conn.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Task{})
	d.Conn.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Job{}, &Task{})
	log.Println("Complete!")
}
