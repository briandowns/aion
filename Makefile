# vim:ft=make:

GOCMD = go
GOBUILD = $(GOCMD) build
GOGET = $(GOCMD) get -v
GOCLEAN = $(GOCMD) clean
GOINSTALL = $(GOCMD) install
GOTEST = $(GOCMD) test

.PHONY: all

define GIT_ERROR

FATAL: Git (git) is required to download aion dependencies.
endef

define HG_ERROR

FATAL: Mercurial (hg) is required to download aion dependencies.
endef

# default target
all : build

# download 3rd party dependencies
get: git hg
	$(GOGET) code.google.com/p/go-uuid/uuid
	$(GOGET) github.com/codegangsta/negroni
	$(GOGET) github.com/goincremental/negroni-sessions
	$(GOGET) github.com/gorilla/mux
	$(GOGET) github.com/unrolled/render
	$(GOGET) github.com/go-sql-driver/mysql
	$(GOGET) github.com/jinzhu/gorm
	$(GOGET) github.com/hashicorp/consul/api
	$(GOGET) github.com/cloudfoundry/gosigar
	$(GOGET) github.com/StalkR/goircbot/lib/disk
	$(GOGET) github.com/gorilla/websocket
	$(GOGET) github.com/google/go-github/github
	$(GOGET) github.com/robfig/cron
	$(GOGET) github.com/gorhill/cronexpr
	$(GOGET) github.com/jinzhu/gorm

# build all and place the binary
install: clean get
	$(GOINSTALL) -v

# remove aion artifact(s)
clean:
	$(GOCLEAN) -n -i -x
	rm -f $(GOPATH)/bin/aion
	rm -rf bin/aion

# build the backend
build: clean
	$(GOBUILD) -v -race -o aion

# run the backend unit tests
test:
	$(GOTEST) -v -cover github.com/briandowns/aion

# check for git
git:
	$(if $(shell git), , $(error $(GIT_ERROR)))

# check for mercurial
hg:
	$(if $(shell hg), , $(error $(HG_ERROR)))