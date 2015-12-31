package main

import (
	"fmt"
	"net/http"

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

// TasksRouteHandler provides the handler for tasks data
func TasksRouteHandler(ren *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ren.JSON(w, http.StatusOK, map[string]interface{}{"tasks": ""})
	}
}
