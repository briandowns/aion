package main

import (
	"time"
)

// Job holds what's needs to represent a job
type Job struct {
	ID    int    `sql:"auto_increment" gorm:"column:id" gorm:"primary_key" json:"id"`
	Name  string `gorm:"column:name" json:"name"`
	Tasks string `gorm:"column:tasks" json:"tasks"`
}

// Task holds what's needs to represent a task
type Task struct {
	ID       int    `sql:"auto_increment" gorm:"column:id" gorm:"primary_key" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	Desc     string `gorm:"column:desc" json:"desc"`
	CMD      string `gorm:"column:cmd" json:"cmd"`
	Args     string `gorm:"column:args" json:"args"`
	Schedule string `gorm:"column:schedule" json:"schedule"`
}

// User holds what's needs to represent a user
type User struct {
	ID           int    `sql:"auto_increment" gorm:"column:id" gorm:"primary_key" json:"id"`
	Username     string `gorm:"column:username" json:"username"`
	Password     string `gorm:"column:password" json:"password"`
	PermissionID int    `gorm:"column:permission_id" json:"permission_id"`
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
}
