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
| DELETE | /api/v1/job/:id  | Delete a job
| GET    | /api/v1/task     | Get a list of all tasks
| POST   | /api/v1/task     | Add a new task
| GET    | /api/v1/task/:id | Get details for a given task ID
| DELETE | /api/v1/task/:id | Delete a task

* more to come...

## Management 

aion is managed entirely through the API.  A simple web UI is also provided that interacts with the API and also shows visualizations of the data therein.

## Development

```bash
$ cd $GOPATH/src/<username>/ && git clone git@github.com:briandowns/aion.git
$ cd aion
$ make dep
```

## Statistics

You can get statistics from Aion simply by going to the `/api/v1/admin/api/stats` endpoint.  This will yield the following results.  There's intent to ahve these and other statistics gathered and stored either in the MySQL database or another datastore of the user's choosing.

```javascript
{
	pid: 50415,
	uptime: "46.357446503s",
	uptime_sec: 46.357446503,
	time: "2016-01-23 14:56:49.683847989 -0700 MST",
	unixtime: 1453586209,
	status_code_count: { },
	total_status_code_count: {
	200: 6
},
	count: 0,
	total_count: 6,
	total_response_time: "1.591559ms",
	total_response_time_sec: 0.0015915590000000002,
	average_response_time: "265.259Âµs",
	average_response_time_sec: 0.000265259
}
```

## Contributing

* Put in an issue
* Fork and create a branch
* Submit a pull request