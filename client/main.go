package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/cli"

	"github.com/briandowns/aion/client/command"
	"github.com/briandowns/aion/client/config"
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
	conf, err := config.Load()
	if err != nil {
		log.Fatalln(err)
	}

	c := &cli.CLI{
		Name:     aionName,
		Version:  aionVersion,
		Args:     os.Args[1:],
		HelpFunc: cli.BasicHelpFunc(aionName),
		Commands: map[string]cli.CommandFactory{
			"show":    command.NewShow(conf),
			"version": command.NewVersion(aionVersion),
		},
	}

	retval, err := c.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
		return 1, err
	}

	return retval, nil
}
