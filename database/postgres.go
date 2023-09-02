package database

import (
	"database/sql"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

func NewPostgreSQLConnection(uri string, connections int32, debug bool) error {

	pgDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(uri)))

	pgDB.SetMaxOpenConns(int(connections))

	PostgreSQLDB := bun.NewDB(pgDB, pgdialect.New())

	if debug {
		PostgreSQLDB.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	}

	return PostgreSQLDB.Ping()
}

func ReadFileAndExecuteQueries(path string) error {
	queries, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	_, err = PostgreSQLDB.Exec(string(queries))
	if err != nil {
		return err
	}
	return nil
}
