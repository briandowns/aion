package main

// DBConf holds the given values for the database
type DBConf struct {
	DBHost string
	DBUser string
	DBPass string
	DBName string
}

// Config holds the running config values
type Config struct {
	Database  DBConf
	QueueHost string
}
