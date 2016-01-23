package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

const (
	aionVersion = "0.1"
	aionName    = "aion"
)

// Commands is the mapping of all the available wasteband commands.
var Commands map[string]cli.CommandFactory

func main() {
	if retval, err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(retval)
	}
}

func run() (int, error) {
	conf, err := config.Load(configFile)
	if err != nil {
		log.Fatalln(err)
	}

	if !conf.HasEndpoint() {
		fmt.Printf("ERROR: no Elasticsearch endpoint set.\n")
		return 1, nil
	}

	c := &cli.CLI{
		Name:     wastebandName,
		Version:  wastebandVersion,
		Args:     os.Args[1:],
		HelpFunc: cli.BasicHelpFunc(wastebandName),
		Commands: map[string]cli.CommandFactory{
			"show":     command.NewShow(conf),
			"version":  command.NewVersion(wastebandVersion),
		},
	}

	retval, err := c.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
		return 1, err
	}

	return retval, nil
}
