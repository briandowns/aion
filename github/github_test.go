package github

import (
	"path/filepath"
	"reflect"
	"testing"

	"gh.internal.shutterfly.com/shutterfly/slam/config"
)

var prodConfig = "../prod-conf.json"

// setUp is used to setup what's needed to do a Github test
func setUp() (*Github, error) {
	file, err := filepath.Abs(prodConfig)
	if err != nil {
		return nil, err
	}

	conf, err := config.Load(file)
	if err != nil {
		return nil, err
	}

	github, err := NewGithub(&conf.Github)
	if err != nil {
		return nil, err
	}

	return github, nil
}

// TestNewGithub verifies the NewGithub function return a pointer to an Github object
func TestNewGithub(t *testing.T) {
	t.Parallel()

	file, err := filepath.Abs(prodConfig)
	if err != nil {
		t.Log(err)
	}

	conf, err := config.Load(file)
	if err != nil {
		t.Log(err)
	}

	github, err := NewGithub(&conf.Github)
	if err != nil {
		t.Error(err)
	}

	t.Log(reflect.TypeOf(github).String())

	// make sure we get the right data type returened
	if reflect.TypeOf(github).String() != "*github.Github" {
		t.Error("Incorrect data type pointer returned from NewGithub function")
	}
}

// TestLastCommitInfo verifies that the expected data from the API is returned
func TestLastCommitInfo(t *testing.T) {
	t.Parallel()

	sflyGithub, err := setUp()
	if err != nil {
		t.Error(err)
	}

	lc, err := sflyGithub.LastCommitInfo()
	if err != nil {
		t.Error(err)
	}

	// make sure we get the right data type returened
	if reflect.TypeOf(lc).String() != "*github.LastCommit" {
		t.Error("Incorrect data type pointer returned from LastCommitInfo function")
	}
}
