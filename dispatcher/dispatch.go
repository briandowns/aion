package dispatcher

import (
	"log"
	"os/exec"

	"github.com/briandowns/aion/config"
	"github.com/briandowns/aion/database"

	"github.com/robfig/cron"
)

// Sender is an interface for sending data to NSQ
type Sender interface {
	Send(db *database.Database) error
}

// Dispatcher holds the values that comprise the Aion dispatcher
type Dispatcher struct {
	Conf           *config.Config
	cron           *cron.Cron
	ResultChan     chan database.Result
	JobProcChan    chan database.Job
	TaskProcChan   chan database.Task
	RemoveTaskChan chan database.Task
	SenderChan     chan Sender
}

// NewDispatcher creates a new refence of type Dispatcher
func NewDispatcher(conf *config.Config) *Dispatcher {
	return &Dispatcher{
		Conf:           conf,
		cron:           cron.New(),
		ResultChan:     make(chan database.Result),
		SenderChan:     make(chan Sender),
		JobProcChan:    make(chan database.Job),
		TaskProcChan:   make(chan database.Task),
		RemoveTaskChan: make(chan database.Task),
	}
}

// generateTaskFunc generates a function suitable for use in the scheduler
func (d *Dispatcher) taskFuncFactory(task *database.Task) func() {
	fn := func() {
		var r database.Result
		r.TaskID = task.ID

		out, err := exec.Command(task.CMD, task.Args).Output()
		if err != nil {
			log.Println(err)
			r.Error = err.Error()
		}

		db, err := database.NewDatabase(d.Conf)
		if err != nil {
			log.Fatalln(err)
		}
		defer db.Conn.Close()

		r.Result = out

		r.Send(db)
	}
	task.Func = fn
	return fn
}

// AddExistingTasks adds active tasks from the database to the scheduler
func (d *Dispatcher) AddExistingTasks() {
	db, err := database.NewDatabase(d.Conf)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Conn.Close()

	tasks := db.GetTasks()
	for _, task := range tasks {
		d.cron.AddFunc(task.Schedule, d.taskFuncFactory(&task))
	}
}

// Run starts the dispatcher
func (d *Dispatcher) Run() error {
	db, err := database.NewDatabase(d.Conf)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Conn.Close()
	log.Println("Starting scheduler...")
	d.cron.Start()
	defer d.cron.Stop()
	d.AddExistingTasks()

	for {
		select {
		case data := <-d.SenderChan:
			if err := data.Send(db); err != nil {
				log.Println(err)
			}

		// listen for new tasks and add them to the scheduler
		case job := <-d.JobProcChan:
			//

		// listen for new tasks and add them to the scheduler
		case task := <-d.TaskProcChan:
			d.cron.AddFunc(task.Schedule, d.taskFuncFactory(&task))

		// remove tasks from the scheduler
		case task := <-d.RemoveTaskChan:
			entries := d.cron.Entries()
			//a = append(a[:i], a[i+1:]...)
			log.Println(task)
			for i, t := range entries {
				log.Printf("%d: %+v\n", i, t)
			}
		}
	}
}

// ResultWorkers starts the result workers
func (d *Dispatcher) ResultWorkers() {
}
