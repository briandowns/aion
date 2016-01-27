# vim:ft=make:

GOCMD = go
GOBUILD = $(GOCMD) build
GOGET = $(GOCMD) get
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
	$(GOGET) -u -v github.com/pborman/uuid
	$(GOGET) -u -v github.com/codegangsta/negroni
	$(GOGET) -u -v github.com/goincremental/negroni-sessions
	$(GOGET) -u -v github.com/gorilla/mux
	$(GOGET) -u -v github.com/unrolled/render
	$(GOGET) -u -v github.com/go-sql-driver/mysql
	$(GOGET) -u -v github.com/jinzhu/gorm
	$(GOGET) -u -v github.com/hashicorp/consul/api
	$(GOGET) -u -v github.com/cloudfoundry/gosigar
	$(GOGET) -u -v github.com/StalkR/goircbot/lib/disk
	$(GOGET) -u -v github.com/gorilla/websocket
	$(GOGET) -u -v github.com/google/go-github/github
	$(GOGET) -u -v github.com/robfig/cron
	$(GOGET) -u -v github.com/gorhill/cronexpr
	$(GOGET) -u -v github.com/jinzhu/gorm
	$(GOGET) -u -v github.com/nsqio/go-nsq
	$(GOGET) -u -v github.com/mikespook/gorbac
	$(GOGET) -u -v github.com/emicklei/forest
	$(GOGET) -u -v github.com/StephanDollberg/go-json-rest-middleware-jwt

# build all and place the binary
install: clean get
	$(GOINSTALL) -v

# install dependancies
dep: get

# remove aion artifact(s)
clean:
	$(GOCLEAN) -n -i -x
	rm -f $(GOPATH)/bin/aiond
	rm -rf bin/aiond
	rm -f client/aion

# build the backend
build: clean build-server build-client

build-server:
	$(GOBUILD) -v -race -o aiond
	
build-client: 
	cd ./client; $(GOBUILD) -v -race -o aion

# run the backend unit tests
test:
	$(GOTEST) -v -cover github.com/briandowns/aion

# check for git
git:
	$(if $(shell git), , $(error $(GIT_ERROR)))

# check for mercurial
hg:
	$(if $(shell hg), , $(error $(HG_ERROR)))