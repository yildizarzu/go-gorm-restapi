package db

import (
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=postgres user=arzu password=12345678 dbname=postgres port=5432"
var DB *gorm.DB
var err error

func DBConnection(Dsn string) error {
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		return errors.New("connection failed")
	} else {
		return nil
	}
}

func DbDisconnect() {
	db, _ := DB.DB()
	db.Close()
}
