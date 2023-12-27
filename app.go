package main

import (
	"net/http"
	"os"

	"github.com/blacheinc/pixel/database"
	"github.com/blacheinc/pixel/database/migration"
	"github.com/blacheinc/pixel/primer"
	"github.com/blacheinc/pixel/version"
	"github.com/opensaucerer/barf"
)

func main() {

	// load environment variables
	if err := barf.Env(primer.ENV, os.Getenv("ENV_PATH")); err != nil {
		barf.Logger().Error(err.Error())
		os.Exit(1)
	}

	// configure barf
	allow := true
	if err := barf.Stark(barf.Augment{
		Port:     primer.ENV.Port,
		Logging:  &allow, // enable request logging
		Recovery: &allow, // enable panic recovery
		CORS: &barf.CORS{
			AllowedOrigins: []string{"*"},
			MaxAge:         3600,
			AllowedMethods: []string{
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodPatch,
				http.MethodDelete,
			},
		},
	}); err != nil {
		barf.Logger().Error(err.Error())
		os.Exit(1)
	}

	if err := database.NewPostgreSQLConnection(primer.ENV.PostgreSQLURI, primer.ENV.PostgreSQLConnections, primer.ENV.PostgreSQLDebug); err != nil {
		barf.Logger().Error(err.Error())
		os.Exit(1)
	}

	if err := migration.CreateTables(); err != nil {
		barf.Logger().Fatalf(`[main.main] [migration.CreateTables()] %s`, err.Error())
	}

	// if err := database.ReadFileAndExecuteQueries(primer.ENV.SQLFilePath); err != nil {
	// 	barf.Logger().Error(err.Error())
	// os.Exit(1)
	// }

	// preload v1 routes
	version.V1()

	// call upon barf to listen and serve
	if err := barf.Beck(); err != nil {
		barf.Logger().Error(err.Error())
		os.Exit(1)
	}
}
