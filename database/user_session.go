package database

// DeleteResult deletes a task
func (d *Database) DeleteResult(id int) {
	d.Conn.Delete(&Result{ID: id})
}

// GetUserSessions gets all user sessions from the database
func (d *Database) GetUserSessions() []UserSession {
	var data []UserSession
	d.Conn.Find(&data)
	return data
}

// GetUserSessionByID gets the user session for the given ID
func (d *Database) GetUserSessionByID(sessionKey string) []UserSession {
	var data []UserSession
	d.Conn.Where("sessionKey = ?", sessionKey).Find(&data)
	return data
}

// DeleteUserSession deletes a user session
func (d *Database) DeleteUserSession(sessionKey string) {
	d.Conn.Delete(&UserSession{SessionKey: sessionKey})
}
