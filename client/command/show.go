package command

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/briandowns/aion/client/config"
	"github.com/briandowns/aion/client/utils"
	"github.com/fatih/flags"
	"github.com/mitchellh/cli"
)

// Show holds in the passed in configuration
type Show struct {
	config *config.Configuration
}

// NewShow creates a new instance of Delete
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
		s.GetAll(args[0])
	case "tasks":
		s.GetAll(args[0])
	default:
		fmt.Print("ERROR: invalid option for show\n\n")
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

// Getter
type Getter interface {
	GetAll() ([]Resource, error)
}

// Resource
type Resource interface {
	Getter
}

// GetAll gets all entries for a given resource
func (s *Show) GetAll(resource string) ([]Resource, error) {
	response, err := http.Get(s.config.Endpoint + "/api/v1/" + resource)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var r Resource
	if err := json.NewDecoder(response.Body).Decode(&r); err != nil {
		return nil, err
	}

	fmt.Println(r)
	/*w := utils.NewTabWriter()

	fmt.Fprintf(w, "\nHealth\tStatus\tName\tShards\tReplicas\tDocuments\tSize")
	fmt.Fprintf(w, "\n----------\t----------\t----------\t----------\t----------\t----------\t----------\n")

	for _, i := range r {
		fmt.Fprintf(w, "%s\t%s\t%s\t%d\t%d\t%d\t%d\n",
			i.Health, i.Status, i.Name, i.Shards, i.Replicas, i.Docs.Deleted, i.Store.Size)
	}

	fmt.Fprintf(w, "\n")
	w.Flush()*/

	return nil, nil
}
