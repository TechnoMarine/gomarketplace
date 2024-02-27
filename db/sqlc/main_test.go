package db

import (
	"database/sql"
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
	var err error
	testConn, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	testQueries = New(testConn)

	os.Exit(m.Run())
}
