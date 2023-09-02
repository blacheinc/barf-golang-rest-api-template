package test

import (
	"log"
	"os"

	"github.com/blacheinc/pixel/database"
	"github.com/blacheinc/pixel/primer"
	"github.com/opensaucerer/barf"
)

// Setup prepares the application for testing
// Setup could be made to use a separate environment file pointing to a test database
// such that the tables are dropped and recreated for each test
func Setup() {
	// load environment variables
	if err := barf.Env(primer.ENV, os.Getenv("ENV_PATH")); err != nil {
		log.Fatal(err)
	}

	database.NewPostgreSQLConnection(primer.ENV.PostgreSQLURI, primer.ENV.PostgreSQLConnections, primer.ENV.PostgreSQLDebug)

	database.ReadFileAndExecuteQueries(primer.ENV.SQLFilePath)
}

// Teardown cleans up the application after testing
// drop all tables in the database
func Teardown() {
	// database.DropAllTables()
}
