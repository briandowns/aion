package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/gorhill/cronexpr"

	"github.com/briandowns/aion/config"
	"github.com/briandowns/aion/database"
	"github.com/briandowns/aion/dispatcher"
)

// TasksRouteHandler provides the handler for tasks data
func TasksRouteHandler(ren *render.Render, conf *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := database.NewDatabase(conf)
		if err != nil {
			log.Println(err)
		}
		defer db.Conn.Close()
		ren.JSON(w, http.StatusOK, map[string]interface{}{"tasks": db.GetTasks()})
	}
}

// NewTaskRouteHandler creates a new task with the POST'd data
func NewTaskRouteHandler(ren *render.Render, dispatcher *dispatcher.Dispatcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var nt database.Task

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&nt)
		if err != nil {
			ren.JSON(w, 400, map[string]error{"error": err})
			return
		}
		defer r.Body.Close()

		switch {
		case nt.Name == "":
			ren.JSON(w, 400, map[string]error{"error": ErrMissingNameField})
			return
		case nt.CMD == "":
			ren.JSON(w, 400, map[string]error{"error": ErrMissingCMDField})
			return
		case nt.Args == "":
			ren.JSON(w, 400, map[string]error{"error": ErrMissingArgsField})
			return
		case nt.Schedule == "":
			ren.JSON(w, 400, map[string]error{"error": ErrMissingScheduleField})
			return
		}

		// validate that the entered cron string is valid.  Error if not.
		_, err = cronexpr.Parse(nt.Schedule)
		if err != nil {
			ren.JSON(w, 400, map[string]error{"error": err})
			return
		}

		dispatcher.SenderChan <- &nt
		ren.JSON(w, http.StatusOK, map[string]database.Task{"task": nt})
	}
}

// TaskByIDRouteHandler provides the handler for tasks data for the given ID
func TaskByIDRouteHandler(ren *render.Render, conf *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		tid := vars["id"]

		taskID, err := strconv.Atoi(tid)
		if err != nil {
			log.Println(err)
		}

		db, err := database.NewDatabase(conf)
		if err != nil {
			log.Println(err)
		}
		defer db.Conn.Close()

		if t := db.GetTaskByID(taskID); len(t) > 0 {
			ren.JSON(w, http.StatusOK, map[string]interface{}{"task": t})
		} else {
			ren.JSON(w, http.StatusBadRequest, map[string]interface{}{"task": ErrNoTasksFound.Error()})
		}
	}
}

// TaskDeleteByIDRouteHandler deletes the task data for the given ID
func TaskDeleteByIDRouteHandler(ren *render.Render, conf *config.Config, dispatch *dispatcher.Dispatcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		tid := vars["id"]
		taskID, err := strconv.Atoi(tid)
		if err != nil {
			log.Println(err)
		}
		db, err := database.NewDatabase(conf)
		if err != nil {
			log.Println(err)
		}
		defer db.Conn.Close()

		task := db.GetTaskByID(taskID)
		if len(task) > 0 {
			dispatch.RemoveTaskChan <- task[0]	
			db.DeleteTask(taskID)
		} else {
			ren.JSON(w, http.StatusOK, map[string]interface{}{"task": taskID})
		}

		ren.JSON(w, http.StatusOK, map[string]interface{}{"task": taskID})
	}
}