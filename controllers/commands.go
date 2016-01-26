package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"

	"github.com/briandowns/aion/config"
	"github.com/briandowns/aion/database"
	"github.com/briandowns/aion/dispatcher"
)

// CommandsRouteHandler provides the handler for commands data
func CommandsRouteHandler(ren *render.Render, conf *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := database.NewDatabase(conf)
		if err != nil {
			log.Println(err)
		}
		defer db.Conn.Close()
		ren.JSON(w, http.StatusOK, map[string]interface{}{"commands": db.GetCommands()})
	}
}

// NewCommandRouteHandler creates a new command with the POST'd data
func NewCommandRouteHandler(ren *render.Render, dispatcher *dispatcher.Dispatcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var c database.Command

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&c)
		if err != nil {
			ren.JSON(w, 400, map[string]error{"error": err})
			return
		}
		defer r.Body.Close()

		switch {
		case c.CMD == "":
			ren.JSON(w, 400, map[string]error{"error": ErrMissingCMDField})
			return
		case c.Args == "":
			ren.JSON(w, 400, map[string]error{"error": ErrMissingArgsField})
			return
		}

		dispatcher.SenderChan <- &c
		ren.JSON(w, http.StatusOK, map[string]database.Command{"command": c})
	}
}

// CommandByIDRouteHandler provides the handler for commands data for the given ID
func CommandByIDRouteHandler(ren *render.Render, conf *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		tid := vars["id"]

		commandID, err := strconv.Atoi(tid)
		if err != nil {
			log.Println(err)
		}

		db, err := database.NewDatabase(conf)
		if err != nil {
			log.Println(err)
		}
		defer db.Conn.Close()

		if t := db.GetCommandByID(commandID); len(t) > 0 {
			ren.JSON(w, http.StatusOK, map[string]interface{}{"command": t})
		} else {
			ren.JSON(w, http.StatusOK, map[string]interface{}{"command": ErrNoCommandsFound.Error()})
		}
	}
}

// CommandDeleteByIDRouteHandler deletes the command data for the given ID
func CommandDeleteByIDRouteHandler(ren *render.Render, conf *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		tid := vars["id"]
		commandID, err := strconv.Atoi(tid)
		if err != nil {
			log.Println(err)
		}
		db, err := database.NewDatabase(conf)
		if err != nil {
			log.Println(err)
		}
		defer db.Conn.Close()

		db.DeleteCommand(commandID)

		ren.JSON(w, http.StatusOK, map[string]interface{}{"command": commandID})
	}
}