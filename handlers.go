package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorhill/cronexpr"
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
	JobsPath = APIBase + "jobs"

	// TasksPath is the path to access tasks
	TasksPath = APIBase + "tasks"
)

// FrontendHandler provides the handler for the main application
func FrontendHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, fmt.Sprintf("public/index.html"))
	}
}

// JobsRouteHandler provides the handler for jobs data
func JobsRouteHandler(ren *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ren.JSON(w, http.StatusOK, map[string]interface{}{"jobs": ""})
	}
}

// NewJobsRouteHandler creates a new job with the POST'd data
func NewJobsRouteHandler(ren *render.Render, dispatcher *Dispatcher) http.HandlerFunc {
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

// NewTasksRouteHandler creates a new task with the POST'd data
func NewTasksRouteHandler(ren *render.Render, dispatcher *Dispatcher) http.HandlerFunc {
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
