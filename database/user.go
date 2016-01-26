package database

// GetUsers gets all users from the database
func (d *Database) GetUsers() []User {
	var data []User
	d.Conn.Find(&data)
	return data
}

// GetUserByID gets the user for the given ID
func (d *Database) GetUserByID(id int) []User {
	var data []User
	d.Conn.Where("id = ?", id).Find(&data)
	return data
}

// DeleteUser deletes a task
func (d *Database) DeleteUser(id int) {
	d.Conn.Delete(&User{ID: id})
}
