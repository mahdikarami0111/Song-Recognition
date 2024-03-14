package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/mahdikarami0111/cloud-computing-1/api"
	db "github.com/mahdikarami0111/cloud-computing-1/db/sqlc"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://zeqnqsup:9Epyab6mf2Rq2NcQXq0iyRhtR6P9zZuf@cornelius.db.elephantsql.com/zeqnqsup"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cant connect to db", err)
	}
	queries := db.New(conn)
	server := api.NewServer(queries)
	go api.RunServiceTwo("amqps://rigqkizo:CBNChsj9lZoMSSzHXKB84-0glFjLZsT8@hawk.rmq.cloudamqp.com/rigqkizo", queries)
	go api.RunServiceThree(queries)
	err = server.Start("0.0.0.0:8080")
	if err != nil {
		log.Fatal("cant start server", err)
	}

}
