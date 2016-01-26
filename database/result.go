package database

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
