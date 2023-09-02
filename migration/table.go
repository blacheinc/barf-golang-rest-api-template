package database

import (
	"context"

	"github.com/blacheinc/pixel/database"
	userRepository "github.com/blacheinc/pixel/repository/v1/user"
	"github.com/opensaucerer/barf"
)

var Table = []interface{}{

	&userRepository.User{},
}

// CreateTables creates tables that do not already exist. Although we have connections to other DBs configure.Save should only handle migration for configure.Save DB.
func CreateTables() error {
	for _, m := range Table {
		_, err := database.PostgreSQLDB.NewCreateTable().
			IfNotExists().
			Model(m).Exec(context.TODO())
		if err != nil {
			barf.Logger().Warnf("failed to create %v table", m)
			return err
		}
	}
	return nil
}

// migrate effects any database schema migration
func Migrate() error {
	return nil
}
