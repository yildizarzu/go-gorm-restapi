package test

import (
	"log"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestDBConnectionTrue(t *testing.T) {
	var DNS = "host=postgres user=arzu password=12345678 dbname=testpostgres port=5432"
	DBConnection(DNS)

}

func TestDBConnectionFalse(t *testing.T) {
	var DNS = "host=failedpostgres user=arzu password=12345678 dbname=failedpostgres port=5432"
	DBConnection(DNS)

}

var DB *gorm.DB

func DBConnection(DSN string) {
	var err error

	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("DB Connection")
	}
}

var DSN = "host=postgres user=arzu password=12345678 dbname=postgres port=5432"
var err error

func Connection() {
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("DB Connection")
	}
}

//dene
