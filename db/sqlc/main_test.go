package db

import (
	"database/sql"
	"gomarketplace/util"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:secret@localhost:12001/gomarketplacedb?sslmode=disable"
)

var testQueries *Queries
var testConn *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load config in test file: ", err)
	}

	testConn, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	testQueries = New(testConn)

	os.Exit(m.Run())
}
