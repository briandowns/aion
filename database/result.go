package database

// AddResult adds a new result record to the database
func (d *Database) AddResult(r Result) {
	d.Conn.NewRecord(r)
	d.Conn.Create(&r)
	d.Conn.NewRecord(r)
}

// GetResults gets all permissions from the database
func (d *Database) GetResults() []Permission {
	var data []Permission
	d.Conn.Find(&data)
	return data
}

// GetResultByID gets the user for the given ID
func (d *Database) GetResultByID(id int) []Permission {
	var data []Permission
	d.Conn.Where("id = ?", id).Find(&data)
	return data
}

// DeleteResult deletes a task
func (d *Database) DeleteResult(id int) {
	d.Conn.Delete(&Result{ID: id})
}
