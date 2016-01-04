package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorhill/cronexpr"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// Top Level/Primary Routes
var (
	frontEnd = "/"
)

const (
	// APIBase is the base path for API access
	APIBase = "/api/v1/"

	// JobsPath is the path to access jobs
	JobsPath = APIBase + "job"

	// TasksPath is the path to access tasks
	TasksPath = APIBase + "task"
)

var (
	// JobByID is the path to get specific job data
	JobByID = JobsPath + "/{id}"

	// TaskByID is the path to get specific task data
	TaskByID = TasksPath + "/{id}"
)

// ErrNoJobsFound given when a job isn't found
var ErrNoJobsFound = errors.New("no jobs found")

// ErrNoTasksFound given when a task isn't found
var ErrNoTasksFound = errors.New("no tasks found")

// FrontendHandler provides the handler for the main application
func FrontendHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, fmt.Sprintf("public/index.html"))
	}
}

// JobsRouteHandler provides the handler for jobs data
func JobsRouteHandler(ren *render.Render, conf *Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := NewDatabase(conf)
		if err != nil {
			log.Println(err)
		}
		defer db.Conn.Close()
		ren.JSON(w, http.StatusOK, map[string]interface{}{"jobs": db.GetJobs()})
	}
}

// NewJobRouteHandler creates a new job with the POST'd data
func NewJobRouteHandler(ren *render.Render, dispatcher *Dispatcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var nj Job

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&nj)
		if err != nil {
			ren.JSON(w, 400, map[string]string{"error": "unable to marshal the posted JSON"})
			return
		}
		defer r.Body.Close()
		switch {
		case nj.Name == "":
			ren.JSON(w, 400, map[string]string{"error": "missing or empty 'name' field"})
			return
		}

		dispatcher.NewJobChan <- nj
		ren.JSON(w, http.StatusOK, map[string]Job{"job": nj})
	}
}

// JobByIDRouteHandler provides the handler for jobs data for the given ID
func JobByIDRouteHandler(ren *render.Render, conf *Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		jid := vars["id"]

		jobID, err := strconv.Atoi(jid)
		if err != nil {
			log.Println(err)
		}

		db, err := NewDatabase(conf)
		if err != nil {
			log.Println(err)
		}
		defer db.Conn.Close()

		if t := db.GetJobByID(jobID); len(t) > 0 {
			ren.JSON(w, http.StatusOK, map[string]interface{}{"task": t})
		} else {
			ren.JSON(w, http.StatusOK, map[string]interface{}{"task": ErrNoJobsFound.Error()})
		}
	}
}

// JobDeleteByIDRouteHandler deletes the job data for the given ID
func JobDeleteByIDRouteHandler(ren *render.Render, conf *Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		jid := vars["id"]
		jobID, err := strconv.Atoi(jid)
		if err != nil {
			log.Println(err)
		}
		db, err := NewDatabase(conf)
		if err != nil {
			log.Println(err)
		}
		defer db.Conn.Close()

		db.DeleteTask(jobID)

		ren.JSON(w, http.StatusOK, map[string]interface{}{"task": jobID})
	}
}

// TasksRouteHandler provides the handler for tasks data
func TasksRouteHandler(ren *render.Render, conf *Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := NewDatabase(conf)
		if err != nil {
			log.Println(err)
		}
		defer db.Conn.Close()
		ren.JSON(w, http.StatusOK, map[string]interface{}{"tasks": db.GetTasks()})
	}
}

// NewTaskRouteHandler creates a new task with the POST'd data
func NewTaskRouteHandler(ren *render.Render, dispatcher *Dispatcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var nt Task

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&nt)
		if err != nil {
			ren.JSON(w, 400, map[string]string{"error": "unable to marshal the posted JSON"})
			return
		}
		defer r.Body.Close()

		switch {
		case nt.Name == "":
			ren.JSON(w, 400, map[string]string{"error": "missing or empty 'name' field"})
			return
		case nt.CMD == "":
			ren.JSON(w, 400, map[string]string{"error": "missing or empty 'cmd' field"})
			return
		case nt.Args == "":
			ren.JSON(w, 400, map[string]string{"error": "missing or empty 'args' field"})
			return
		case nt.Schedule == "":
			ren.JSON(w, 400, map[string]string{"error": "missing or empty 'schedule' field"})
			return
		}

		// validate that the entered cron string is valid.  Error if not.
		_, err = cronexpr.Parse(nt.Schedule)
		if err != nil {
			ren.JSON(w, 400, map[string]string{"error": "invalid cron format"})
			return
		}

		dispatcher.NewTaskChan <- nt
		ren.JSON(w, http.StatusOK, map[string]Task{"task": nt})
	}
}

// TaskByIDRouteHandler provides the handler for tasks data for the given ID
func TaskByIDRouteHandler(ren *render.Render, conf *Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		tid := vars["id"]

		taskID, err := strconv.Atoi(tid)
		if err != nil {
			log.Println(err)
		}

		db, err := NewDatabase(conf)
		if err != nil {
			log.Println(err)
		}
		defer db.Conn.Close()

		if t := db.GetTaskByID(taskID); len(t) > 0 {
			ren.JSON(w, http.StatusOK, map[string]interface{}{"task": t})
		} else {
			ren.JSON(w, http.StatusOK, map[string]interface{}{"task": ErrNoTasksFound.Error()})
		}
	}
}

// TaskDeleteByIDRouteHandler deletes the task data for the given ID
func TaskDeleteByIDRouteHandler(ren *render.Render, conf *Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		tid := vars["id"]
		taskID, err := strconv.Atoi(tid)
		if err != nil {
			log.Println(err)
		}
		db, err := NewDatabase(conf)
		if err != nil {
			log.Println(err)
		}
		defer db.Conn.Close()

		db.DeleteTask(taskID)

		ren.JSON(w, http.StatusOK, map[string]interface{}{"task": taskID})
	}
}
