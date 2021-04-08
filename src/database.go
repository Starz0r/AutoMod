package main

import (
	"os"

	"github.com/spidernest-go/db/lib/sqlbuilder"
	"github.com/spidernest-go/db/postgresql"
	"github.com/spidernest-go/logger"
)

var DB sqlbuilder.Database

func connectDatabase() {
	opts := make(map[string]string)
	opts["sslmode"] = "require"
	
	conn := postgresql.ConnectionURL{
		Database: os.Getenv("PQ_DB"),
		Host:     os.Getenv("PQ_HOST"),
		User:     os.Getenv("PQ_USER"),
		Password: os.Getenv("PQ_PASS"),
		Options: opts,
	}

	err := *new(error)
	DB, err = postgresql.Open(conn)
	if err != nil {
		logger.Panic().Err(err).Msg("Opening the PostgreSQL connection failed.")
	}
}
