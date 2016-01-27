package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"code.google.com/p/go-uuid/uuid"

	"github.com/briandowns/aion/config"
	"github.com/briandowns/aion/controllers"
	"github.com/briandowns/aion/database"
	"github.com/briandowns/aion/dispatcher"

	"github.com/codegangsta/negroni"
	"github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"github.com/gorilla/mux"
	"github.com/thoas/stats"
	"github.com/unrolled/render"
)

// CLI flags
var (
	portFlag      string
	queueHostFlag string
	setupFlag     bool
	dbUserFlag    string
	dbPassFlag    string
	dbHostFlag    string
	dbPortFlag    int
	dbNameFlag    string
	dbSetupFlag   bool
	resultWorkers int
)

var jobRegistryChan = make(chan database.Job)
var taskRegistryChan = make(chan database.Task)
var signalsChan = make(chan os.Signal, 1)

// Conf holds the current configuration
var Conf *config.Config

func init() {
	flag.StringVar(&queueHostFlag, "nsq-host", "", "NSQ server to connect to")
	flag.StringVar(&portFlag, "port", ":9898", "port to run the server")
	flag.StringVar(&dbUserFlag, "db-user", "aion", "database username")
	flag.StringVar(&dbPassFlag, "db-pass", "aion", "database password")
	flag.StringVar(&dbHostFlag, "db-host", "localhost", "database hostname")
	flag.IntVar(&dbPortFlag, "db-port", 3306, "database port")
	flag.StringVar(&dbNameFlag, "db-name", "aion", "database name")
	flag.BoolVar(&dbSetupFlag, "db-setup", false, "intial DB Configuration")
	flag.IntVar(&resultWorkers, "result-workers", 5, "number of result workers to start")
}

func main() {
	flag.Parse()

	signal.Notify(signalsChan, os.Interrupt)

	go func() {
		for sig := range signalsChan {
			log.Printf("Exiting... %v\n", sig)
			signalsChan = nil
			os.Exit(1)
		}
	}()

	if queueHostFlag == "" || dbUserFlag == "" || dbPassFlag == "" ||
		dbHostFlag == "" || dbPortFlag == 0 || dbNameFlag == "" || resultWorkers < 3 {
		flag.Usage()
		os.Exit(1)
	}

	// assign
	Conf = &config.Config{
		Database: config.DBConf{
			DBUser: dbUserFlag,
			DBPass: dbPassFlag,
			DBHost: dbHostFlag,
			DBPort: dbPortFlag,
			DBName: dbNameFlag,
		},
		QueueHost:     queueHostFlag,
		ResultWorkers: resultWorkers,
	}

	if dbSetupFlag {
		fmt.Println(Conf.Database.DBUser, Conf.Database.DBPass, Conf.Database.DBHost, Conf.Database.DBPort, Conf.Database.DBName)
		db, err := database.NewDatabase(Conf)
		if err != nil {
			log.Fatalln(err)
		}
		defer db.Conn.Close()
		db.Setup()
		os.Exit(0)
	}

	dispatch := dispatcher.NewDispatcher(Conf)
	go dispatch.Run()
	go watchForNewJobs()
	go watchForNewTasks()
	go watchForNewResults()
	go watchForNewCommands()

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

	statsMiddleware := stats.New()

	// create a router to handle the requests coming in to our endpoints
	router := mux.NewRouter()

	// Frontend Entry Point
	router.HandleFunc(controllers.FrontEnd, controllers.FrontendHandler()).Methods("GET")

	// Jobs Route
	router.HandleFunc(controllers.JobsPath, controllers.JobsRouteHandler(ren, Conf)).Methods("GET")

	// Job By ID Route
	router.HandleFunc(controllers.JobByID, controllers.JobByIDRouteHandler(ren, Conf)).Methods("GET")

	// Job Delete By ID Route
	router.HandleFunc(controllers.JobByID, controllers.JobDeleteByIDRouteHandler(ren, Conf)).Methods("DELETE")

	// New Jobs Route
	router.HandleFunc(controllers.JobsPath, controllers.NewJobRouteHandler(ren, dispatch)).Methods("POST")

	// Tasks Route
	router.HandleFunc(controllers.TasksPath, controllers.TasksRouteHandler(ren, Conf)).Methods("GET")

	// Task By ID Route
	router.HandleFunc(controllers.TaskByID, controllers.TaskByIDRouteHandler(ren, Conf)).Methods("GET")

	// Task Delete By ID Route
	router.HandleFunc(controllers.TaskByID, controllers.TaskDeleteByIDRouteHandler(ren, Conf, dispatch)).Methods("DELETE")

	// New Tasks Route
	router.HandleFunc(controllers.TasksPath, controllers.NewTaskRouteHandler(ren, dispatch)).Methods("POST")

	// Commands Route
	router.HandleFunc(controllers.CommandsPath, controllers.CommandsRouteHandler(ren, Conf)).Methods("GET")

	// Cmmand By ID Route
	router.HandleFunc(controllers.CommandByID, controllers.CommandByIDRouteHandler(ren, Conf)).Methods("GET")

	// Command Delete By ID Route
	router.HandleFunc(controllers.CommandByID, controllers.CommandDeleteByIDRouteHandler(ren, Conf)).Methods("DELETE")

	// New Commands Route
	router.HandleFunc(controllers.CommandsPath, controllers.NewCommandRouteHandler(ren, dispatch)).Methods("POST")

	// API Statistics Route
	router.HandleFunc(controllers.APIStats, controllers.AdminAionAPIServerStats(statsMiddleware)).Methods("GET")

	n.Use(statsMiddleware)
	n.UseHandler(router)
	n.Run(portFlag)
}
