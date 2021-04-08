package main

import (
	"os"

	"github.com/spidernest-go/db/lib/sqlbuilder"
	"github.com/spidernest-go/db/postgresql"
	"github.com/spidernest-go/logger"
)

var DB sqlbuilder.Database

func connectDatabase() {
	conn := postgresql.ConnectionURL{
		Database: os.Getenv("PQ_DB"),
		Host:     os.Getenv("PQ_HOST"),
		User:     os.Getenv("PQ_USER"),
		Password: os.Getenv("PQ_PASS"),
	}

	err := *new(error)
	DB, err = postgresql.Open(conn)
	if err != nil {
		logger.Panic().Err(err).Msg("Opening the PostgreSQL connection failed.")
	}
}
