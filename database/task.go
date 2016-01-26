package database

// AddTask adds a new task record to the database
func (d *Database) AddTask(t Task) {
	d.Conn.NewRecord(t)
	d.Conn.Create(&t)
	d.Conn.NewRecord(t)
}

// GetTasks gets all tasks from the database
func (d *Database) GetTasks() []Task {
	var data []Task
	d.Conn.Find(&data)
	return data
}

// GetTaskByID gets the task for the given ID
func (d *Database) GetTaskByID(id int) []Task {
	var data []Task
	d.Conn.Where("id = ?", id).Find(&data)
	return data
}

// DeleteTask deletes a task
func (d *Database) DeleteTask(id int) {
	d.Conn.Delete(&Task{ID: id})
}
