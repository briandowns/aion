package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"

	"code.google.com/p/go-uuid/uuid"

	"github.com/codegangsta/negroni"
	"github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// CLI flags
var portFlag string
var queueHostFlag string
var setupFlag bool

var signalsChan = make(chan os.Signal, 1)

func init() {
	flag.StringVar(&queueHostFlag, "h", "", "NSQ server to connect to")
	flag.StringVar(&portFlag, "p", ":9898", "port to run server on in :8888 format. Default 8888")
}

func main() {
	flag.Parse()

	signal.Notify(signalsChan, os.Interrupt)

	// launch a go routine to listen for an operating system signals
	go func() {
		for sig := range signalsChan {
			log.Printf("Received ctrl^c.  Exiting... %v\n", sig)
			signalsChan = nil
			os.Exit(1)
		}
	}()

	// run database setup if -s is provided at the CLI
	if setupFlag {
		// do some setup. ie create queues and topics
		os.Exit(0)
	}

	// setup the renderer for returning our JSON
	ren := render.New(render.Options{})

	store := cookiestore.New([]byte(uuid.NewUUID().String()))

	// initialize the web framework
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		negroni.NewStatic(http.Dir("public")),
	)

	n.Use(sessions.Sessions("session", store))

	// create a router to handle the requests coming in to our endpoints
	router := mux.NewRouter()

	// Frontend Entry Point
	router.HandleFunc(frontEnd, FrontendHandler()).Methods("GET")

	// Jobs Route
	router.HandleFunc(JobsPath, JobsRouteHandler(ren)).Methods("GET")

	// New Jobs Route
	router.HandleFunc(JobsPath, NewJobsRouteHandler(ren)).Methods("POST")

	// Tasks Route
	router.HandleFunc(TasksPath, TasksRouteHandler(ren)).Methods("GET")

	// New Tasks Route
	router.HandleFunc(TasksPath, NewTasksRouteHandler(ren)).Methods("POST")

	n.UseHandler(router)
	n.Run(portFlag)
}
