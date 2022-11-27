package db

import (
	"errors"

	"github.com/yildizarzu/go-gorm-restapi/models"
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

func CreateTestDB() {
	DB, err = gorm.Open(postgres.Open("host=postgres user=arzu password=12345678 dbname=testpostgres port=5432"), &gorm.Config{})
	if err != nil {
		DbDisconnect()
		DBConnection("host=postgres user=arzu password=12345678 port=5432")
		DB.Exec("CREATE DATABASE testpostgres")
		DB.Exec("DROP TABLE IF EXISTS tickets")
		DB.AutoMigrate(&models.Ticket{})
	} else {
		return
	}

}
