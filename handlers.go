package main

import (
	"fmt"
	"net/http"
)

// Top Level/Primary Routes
var (
	frontEnd = "/"
)

/*const (
	APIBase = "/api/v1/"
)*/

// FrontendHandler provides the handler for the main application
func FrontendHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, fmt.Sprintf("public/index.html"))
	}
}
