package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Database holds db conf and a connection
type Database struct {
	Conf *Config
	Conn *gorm.DB
}

// NewDatabase creates a new Database object
func NewDatabase(conf *Config) (*Database, error) {
	d := &Database{
		Conf: conf,
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
			d.Conf.Database.DBUser, d.Conf.Database.DBPass, d.Conf.Database.DBHost, 3306, d.Conf.Database.DBName, "60s"))
	if err != nil {
		return err
	}
	db.LogMode(true)
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
func (d *Database) GetJobs() []Job {
	var data []Job
	d.Conn.Find(&data)
	return data
}

// GetJobByID gets the job for the given ID
func (d *Database) GetJobByID(id int) []Job {
	var data []Job
	d.Conn.Where("id = ?", id).Find(&data)
	return data
}

// AddTask adds a new task record to the database
func (d *Database) AddTask(t Task) {
	d.Conn.NewRecord(t)
	d.Conn.Create(&t)
	d.Conn.NewRecord(t)
}

// GetTasks gets all tasks from the database
func (d *Database) GetTasks() []Task {
	var data []Task
	d.Conn.Find(&data)
	return data
}

// GetTaskByID gets the task for the given ID
func (d *Database) GetTaskByID(id int) []Task {
	var data []Task
	d.Conn.Where("id = ?", id).Find(&data)
	return data
}

// DeleteTask deletes a task
func (d *Database) DeleteTask(id int) {
	d.Conn.Delete(&Task{ID: id})
}

// Setup ...sets up the database
func (d *Database) Setup() {
	log.Println("Aion database setup starting...")
	d.Conn.CreateTable(&Job{})
	d.Conn.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Job{})

	d.Conn.CreateTable(&Task{})
	d.Conn.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Task{})

	d.Conn.CreateTable(&User{})
	d.Conn.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})
	d.Conn.Model(&User{}).AddForeignKey("permission_id", "permissions(id)", "RESTRICT", "RESTRICT")

	d.Conn.CreateTable(&Permission{})
	d.Conn.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Permission{})

	d.Conn.CreateTable(&Result{})
	d.Conn.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Result{})
	d.Conn.Model(&Result{}).AddForeignKey("task_id", "tasks(id)", "RESTRICT", "RESTRICT")
	d.Conn.Model(&Result{}).AddIndex("idx_start_end", "start_time", "end_time")

	d.Conn.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Job{}, &Task{}, &User{}, &Permission{}, &Result{})
	log.Println("Complete!")
}
