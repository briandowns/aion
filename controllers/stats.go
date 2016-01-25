package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/thoas/stats"
)

// AdminAionAPIServerStats returns Aion API server statistics
func AdminAionAPIServerStats(mw *stats.Stats) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		stats := mw.Data()

		b, err := json.Marshal(stats)
		if err != nil {
			log.Println(err)
		}

		w.Write(b)
	}
}
