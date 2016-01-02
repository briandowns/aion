package main

// Job holds what's needs to represent a task
type Job struct {
	ID    int    `sql:"auto_increment" gorm:"primary_key" json:"id"`
	Name  string `gorm:"column:name" json:"name"`
	Tasks []int  `gorm:"column:tasks" json:"tasks"`
}

// Task holds what's needs to represent a task
type Task struct {
	ID       int    `sql:"auto_increment" gorm:"primary_key" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	Desc     string `gorm:"column:desc" json:"desc"`
	Exec     string `gorm:"column:exec" json:"exec"`
	Schedule string `gorm:"column:schedule" json:"schedule"`
}
