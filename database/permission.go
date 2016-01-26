package database

// GetPermissions gets all permissions from the database
func (d *Database) GetPermissions() []Permission {
	var data []Permission
	d.Conn.Find(&data)
	return data
}

// GetPermissionByID gets the user for the given ID
func (d *Database) GetPermissionByID(id int) []Permission {
	var data []Permission
	d.Conn.Where("id = ?", id).Find(&data)
	return data
}

// DeletePermission deletes a task
func (d *Database) DeletePermission(id int) {
	d.Conn.Delete(&Permission{ID: id})
}
