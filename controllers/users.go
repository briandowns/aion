package controllers

import (
	"log"
	"net/http"

	"github.com/unrolled/render"

	"github.com/briandowns/aion/config"
	"github.com/briandowns/aion/database"
)

// UsersRouteHandler provides the handler for users data
func UsersRouteHandler(ren *render.Render, conf *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := database.NewDatabase(conf)
		if err != nil {
			log.Println(err)
		}
		defer db.Conn.Close()
		ren.JSON(w, http.StatusOK, map[string]interface{}{"tasks": db.GetUsers()})
	}
}
