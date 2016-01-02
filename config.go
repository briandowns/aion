package main

type DBConf struct {
	DBHost string
	DBUser string
	DBPass string
	DBName string
}

type Config struct {
	Database  DBConf
	QueueHost string
}
