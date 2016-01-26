package database

// AddJob adds a new job record to the database
func (d *Database) AddJob(j Job) {
	d.Conn.NewRecord(j)
	d.Conn.Create(&j)
	d.Conn.NewRecord(j)
}

// GetJobs gets all jobs from the database
func (d *Database) GetJobs() []Job {
	var data []Job
	d.Conn.Find(&data)
	return data
}

// GetJobByID gets the job for the given ID
func (d *Database) GetJobByID(id int) []Job {
	var data []Job
	d.Conn.Where("id = ?", id).Find(&data)
	return data
}

// DeleteJob deletes a task
func (d *Database) DeleteJob(id int) {
	d.Conn.Delete(&Job{ID: id})
}
