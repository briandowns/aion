package database

import "log"

// Setup ...sets up the database
func (d *Database) Setup() {
	log.Println("Aion database setup starting...")
	d.Conn.CreateTable(&Job{})
	d.Conn.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Job{})

	d.Conn.CreateTable(&Task{})
	d.Conn.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Task{})

	d.Conn.CreateTable(&User{})
	d.Conn.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})
	d.Conn.Model(&User{}).AddForeignKey("permission_id", "permissions(id)", "RESTRICT", "RESTRICT")

	d.Conn.CreateTable(&UserSession{})
	d.Conn.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&UserSession{})
	d.Conn.Model(&UserSession{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")

	d.Conn.CreateTable(&Permission{})
	d.Conn.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Permission{})

	d.Conn.CreateTable(&Result{})
	d.Conn.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Result{})
	d.Conn.Model(&Result{}).AddIndex("idx_start_end", "start_time", "end_time")

	d.Conn.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&Job{}, &Task{}, &User{}, &UserSession{}, &Permission{}, &Result{})
	log.Println("Complete!")
}
