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
	dbSource = "postgres://zeqnqsup:9Epyab6mf2Rq2NcQXq0iyRhtR6P9zZuf@cornelius.db.elephantsql.com/zeqnqsup"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cant connect to db", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
