package main

import (
	"context"
	"database/sql"
	_ "embed"
	"updates/db"

	_ "github.com/glebarez/go-sqlite"
)

var (
	dbCtx     context.Context
	dbQueries *db.Queries

	//go:embed sqlc/schema.sql
	dbSchema string
)

func init() {
	dbCtx = context.Background()

	sqldb, err := sql.Open("sqlite", config.DbPath)
	if err != nil {
		panic(err)
	}

	// create tables
	if _, err = sqldb.ExecContext(dbCtx, dbSchema); err != nil {
		panic(err)
	}

	dbQueries = db.New(sqldb)
}
