package command

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

// Version holds the current version
type Version struct {
	version string
}

// NewVersion creates a new CommandFactory for the version subcommand
func NewVersion(version string) cli.CommandFactory {
	return func() (cli.Command, error) {
		return &Version{
			version: version,
		}, nil
	}
}

// Run shows the current version
func (v *Version) Run(args []string) int {
	fmt.Fprintln(os.Stderr, v.version)
	return 0
}

// Help displays the below string
func (v *Version) Help() string {
	return "Prints the aion version"
}

// Synopsis provides a brief description of the command
func (v *Version) Synopsis() string {
	return "Prints the aion version"
}
