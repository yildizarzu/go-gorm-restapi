package test

import (
	"testing"

	"github.com/yildizarzu/go-gorm-restapi/db"
)

func TestDBConnectionSuccess(t *testing.T) {
	db.DBConnection("host=postgres user=arzu password=12345678 dbname=testpostgres port=5432")
}

func TestDBConnectionFailed(t *testing.T) {
	err := db.DBConnection("host=postgres user=failed password=failed dbname=testpostgres port=5432")

	if err == nil {
		t.Errorf("FAILED: " + err.Error())
	}
}
