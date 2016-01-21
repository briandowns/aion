# Aion

WIP

aion is a job scheduling engine that utilizes cron syntax.  All tasks are executed in their own goroutine and their results are sent into a queue for another worker to pick-up and process.

* Jobs - Jobs define tasks and potentially an outcome or expected result.
* Tasks - Tasks define an action to be executed, a potential result, and the schedule it needs to be run.

## Installation

```bash
$ go install github.com/briandowns/aion
```

## API

The aion API is broken up into a number of endpoints for managing aion.

| Method | Resource         | Description
| :----- | :-------         | :----------
| GET    | /api/v1/job      | Get a list of all jobs
| POST   | /api/v1/job      | Add a new job
| GET    | /api/v1/job/:id  | Get details for a given job ID
| GET    | /api/v1/task     | Get a list of all tasks
| POST   | /api/v1/task     | Add a new task
| GET    | /api/v1/task/:id | Get details for a given task ID

* more to come...

## Management 

aion is managed entirely through the API.  A simple web UI is also provided that interacts with the API and also shows visualizations of the data therein.

## Development

```bash
$ cd $GOPATH/src/<username>/ && git clone git@github.com:briandowns/aion.git
$ cd aion
$ make dep
```

## Contributing

* Put in an issue
* Fork and create a branch
* Submit a pull request