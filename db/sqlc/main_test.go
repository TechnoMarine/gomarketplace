package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:secret@localhost:12001/noteitdb?sslmode=disable"
)

var testQueries *Queries
var testConn *sql.DB

func TestMain(m *testing.M) {
	var err error
	testConn, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	testQueries = New(testConn)

	os.Exit(m.Run())
}
