package github

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"

	gh "github.com/google/go-github/github"
)

// rewriteTransport holds the parameters to resetup the headers for the
// needed auth token and the transport
type rewriteTransport struct {
	accessToken string
	transport   http.RoundTripper
}

// RoundTrip implements the RoundTripper interface and will be called to inject the key
// into the requet header
func (rt rewriteTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set("Authorization", fmt.Sprintf("token %s", rt.accessToken))

	// set the HTTP client to not care whether we're connecting to a server with
	// a self signed certificate
	t := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	return t.RoundTrip(r)
}

// LastCommit holds the data for the last commit made to a repo
type LastCommit struct {
	Name     string
	Date     string
	Checksum string
}

// Github holds the compoentens needed to connect to Shutterfly's Github
// enterprise installation
type Github struct {
	Conf   *config.GithubConf
	Client *gh.Client
}

// NewGithub creates a new reference to the NewSflyGHE type with
// the provided parameters
func NewGithub(conf *config.GithubConf) (*Github, error) {
	endpoint, err := url.Parse(conf.Endpoint)
	if err != nil {
		return nil, err
	}

	// create a new HTTP client to pass to the Gitub package.  We do this
	// because the Github package is meant mainly for the public Github however
	// it has the abiltiy to be used with Github Enterprise
	httpClient := &http.Client{
		Transport: rewriteTransport{
			accessToken: conf.Token,
		},
	}
	client := gh.NewClient(httpClient)
	client.BaseURL = endpoint

	ghe := &Github{
		Conf:   conf,
		Client: client,
	}

	return ghe, nil
}

// LastCommitInfo gets the data for the last commit from a repository
func (g *Github) LastCommitInfo() (*LastCommit, error) {
	commits, _, err := g.Client.Repositories.ListCommits(g.Conf.Owner, g.Conf.Repo, nil)
	if err != nil {
		return nil, err
	}

	lc := &LastCommit{
		Name:     *commits[0].Commit.Committer.Name,         // name of committer
		Date:     commits[0].Commit.Committer.Date.String(), // date of commit
		Checksum: *commits[0].SHA,                           // checksum for the commit
	}

	return lc, nil
}
