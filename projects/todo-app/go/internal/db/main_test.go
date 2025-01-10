package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/litmus-zhang/90min-app-todo/internal/config"
)

var testDB *sql.DB
var testQueries *Queries

func TestMain(m *testing.M) {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("cannot setup a new config:", err)
	}
	testDB, err = sql.Open(cfg.DbDriver, cfg.DbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
