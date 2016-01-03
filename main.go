package main

import (
	"flag"
	"fmt"
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
var (
	portFlag      string
	queueHostFlag string
	setupFlag     bool
	dbUserFlag    string
	dbPassFlag    string
	dbHostFlag    string
	dbNameFlag    string
	dbSetupFlag   bool
)

var jobRegistryChan = make(chan Job)
var taskRegistryChan = make(chan Task)
var signalsChan = make(chan os.Signal, 1)

func init() {
	flag.StringVar(&queueHostFlag, "nsq-host", "", "NSQ server to connect to")
	flag.StringVar(&portFlag, "port", ":9898", "port to run the server")
	flag.StringVar(&dbUserFlag, "db-user", "aion", "database username")
	flag.StringVar(&dbPassFlag, "db-pass", "aion", "database password")
	flag.StringVar(&dbHostFlag, "db-host", "localhost", "database hostname")
	flag.StringVar(&dbNameFlag, "db-name", "aion", "database name")
	flag.BoolVar(&dbSetupFlag, "db-setup", false, "intial DB configuration")
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

	if queueHostFlag == "" || dbUserFlag == "" || dbPassFlag == "" || dbHostFlag == "" || dbNameFlag == "" {
		flag.Usage()
		os.Exit(1)
	}

	conf := &Config{
		Database: DBConf{
			DBUser: dbUserFlag,
			DBPass: dbPassFlag,
			DBHost: dbHostFlag,
			DBName: dbNameFlag,
		},
		QueueHost: queueHostFlag,
	}

	if dbSetupFlag {
		fmt.Println(conf.Database.DBUser, conf.Database.DBPass, conf.Database.DBHost, conf.Database.DBName)
		db, err := NewDatabase(conf)
		if err != nil {
			log.Fatalln(err)
		}
		defer db.Conn.Close()
		db.Setup()
		os.Exit(0)
	}

	dispatcher := NewDispatcher(conf)
	go dispatcher.Run()

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
	router.HandleFunc(JobsPath, NewJobsRouteHandler(ren, dispatcher)).Methods("POST")

	// Tasks Route
	router.HandleFunc(TasksPath, TasksRouteHandler(ren, conf)).Methods("GET")

	// New Tasks Route
	router.HandleFunc(TasksPath, NewTasksRouteHandler(ren, dispatcher)).Methods("POST")

	n.UseHandler(router)
	n.Run(portFlag)
}
