package types

type Env struct {
	// Port for the server to listen on
	Port string `barfenv:"key=PORT;required=true"`
	// Database connection string
	PostgreSQLURI string `barfenv:"key=POSTGRESQL_URI;required=true"`
	// Number of connections to the database
	PostgreSQLConnections int32 `barfenv:"key=POSTGRESQL_CONNECTIONS;required=true"`
	// Enables verbose logging of database queries
	PostgreSQLDebug bool `barfenv:"key=POSTGRESQL_DEBUG;required=true"`
	// Path to SQL file containing queries to be executed on startup
	SQLFilePath string `barfenv:"key=SQL_FILE_PATH;required=true"`
	// Secret for generating JWT signatures
	JWTSecret string `barfenv:"key=JWT_SECRET;required=true"`
	// Name of the app instance
	AppName string `barfenv:"key=APP_NAME;required=true"`
}
