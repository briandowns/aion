package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"code.google.com/p/go-uuid/uuid"

	"github.com/unrolled/render"
)

// Top Level/Primary Routes
var (
	frontEnd = "/"
)

const (
	APIBase = "/api/v1/"
)

var (
	JobsPath  = APIBase + "jobs"
	TasksPath = JobsPath + "/tasks"
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
func NewJobsRouteHandler(ren *render.Render) http.HandlerFunc {
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
		nj.ID = uuid.NewUUID().String()
		ren.JSON(w, http.StatusOK, map[string]Job{"job": nj})
	}
}

// TasksRouteHandler provides the handler for tasks data
func TasksRouteHandler(ren *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ren.JSON(w, http.StatusOK, map[string]interface{}{"tasks": ""})
	}
}

// NewTasksRouteHandler creates a new task with the POST'd data
func NewTasksRouteHandler(ren *render.Render) http.HandlerFunc {
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
		case nt.Exec == "":
			ren.JSON(w, 400, map[string]string{"error": "missing or empty 'exec' field"})
			return
		case nt.Schedule == "":
			ren.JSON(w, 400, map[string]string{"error": "missing or empty 'schedule' field"})
			return
		}
		nt.ID = uuid.NewUUID().String()
		ren.JSON(w, http.StatusOK, map[string]Task{"task": nt})
	}
}
