package command

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/briandowns/aion/client/config"
	"github.com/briandowns/aion/client/utils"
	"github.com/briandowns/aion/database"

	"github.com/fatih/flags"
	"github.com/mitchellh/cli"
)

// Show holds in the passed in configuration
type Show struct {
	config *config.Configuration
}

// NewShow creates a new CommandFactory for the show subcommand
func NewShow(conf *config.Configuration) cli.CommandFactory {
	return func() (cli.Command, error) {
		return &Show{
			config: conf,
		}, nil
	}
}

// Run shows a given resource.
func (s *Show) Run(args []string) int {
	if flags.Has("help", args) || len(args) < 1 {
		fmt.Print(s.Help())
		return 1
	}

	// process the subcommand and it's options
	switch args[0] {
	case "config":
		s.showConfig()
	case "jobs":
		s.AllJobs()
	default:
		fmt.Print("ERROR: invalid option for show\n")
		return 1
	}

	return 1
}

// showConfig outputs the current running configuration
func (s *Show) showConfig() {
	fmt.Print("\naion config:\n")
	w := utils.NewTabWriter()

	v := reflect.ValueOf(*s.config)

	fmt.Fprint(w, "\n")

	// iterate through the values of the struct and write to the tabwriter
	for i := 0; i < v.NumField(); i++ {
		fmt.Fprintf(w, "%s\t%v\n", v.Type().Field(i).Name, v.Field(i).Interface())
	}

	fmt.Fprintf(w, "\n")
	w.Flush()
}

// Help provides full help inforamation for the subcommand
func (s *Show) Help() string {
	return `Usage: aion show <option> <arguments> 
  Show a resource
Options:
  jobs               Display all jobs
  tasks              Display all tasks
`
}

// Synopsis provides a brief description of the command
func (s *Show) Synopsis() string {
	return "Show an Aion resource"
}

// JobsResponse holds the response from the API
type JobsResponse struct {
	Jobs []database.Job `json:"jobs"`
}

// TasksResponse holds the response from the API
type TasksResponse struct {
	Jobs []database.Task `json:"tasks"`
}

// AllJobs gets all job entries
func (s *Show) AllJobs() {
	response, err := http.Get("http://" + s.config.Endpoint + "/api/v1/job")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	var r JobsResponse
	if err := json.NewDecoder(response.Body).Decode(&r); err != nil {
		fmt.Println(err)
	}

	w := utils.NewTabWriter()

	fmt.Fprintf(w, "\nName\tDescription\tTasks")
	fmt.Fprintf(w, "\n----------\t----------\t----------\n")

	for _, i := range r.Jobs {
		fmt.Fprintf(w, "%s\t%s\t%s\n", i.Name, i.Desc, i.Tasks)
	}

	fmt.Fprintf(w, "\n")
	w.Flush()
}
