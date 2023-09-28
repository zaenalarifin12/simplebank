package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/zaenalarifin12/simplebank/utils"
	"log"
	"os"
	"testing"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("can't load config", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect db: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())

}
