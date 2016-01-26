package database

// AddCommand adds a new command record to the database
func (d *Database) AddCommand(c Command) {
	d.Conn.NewRecord(c)
	d.Conn.Create(&c)
	d.Conn.NewRecord(c)
}

// GetCommands gets all commands from the database
func (d *Database) GetCommands() []Command {
	var data []Command
	d.Conn.Find(&data)
	return data
}

// GetCommandByID gets the command for the given ID
func (d *Database) GetCommandByID(id int) []Command {
	var data []Command
	d.Conn.Where("id = ?", id).Find(&data)
	return data
}

// DeleteCommand deletes a command
func (d *Database) DeleteCommand(id int) {
	d.Conn.Delete(&Command{ID: id})
}
