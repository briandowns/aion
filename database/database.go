package database

import (
	"fmt"

	// _ "github.com/go-sql-driver/mysql" is to load the MySQL driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/briandowns/aion/config"
)

// DataAccess is an interface
type DataAccess interface {
	GetAll(d *Database)
	GetByID(d *Database, id int)
	Delete(d *Database, id int)
}

// Database holds db conf and a connection
type Database struct {
	Conf *config.Config
	Conn *gorm.DB
}

// NewDatabase creates a new Database object
func NewDatabase(conf *config.Config) (*Database, error) {
	d := &Database{
		Conf: conf,
	}
	if err := d.connect(); err != nil {
		return nil, err
	}
	return d, nil
}

// Connect will provide the caller with a db connection
func (d *Database) connect() error {
	db, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=%s&charset=utf8&parseTime=True&loc=Local",
			d.Conf.Database.DBUser, d.Conf.Database.DBPass, d.Conf.Database.DBHost, d.Conf.Database.DBPort, d.Conf.Database.DBName, "60s"))
	if err != nil {
		return err
	}
	db.LogMode(true)
	d.Conn = &db
	return nil
}
