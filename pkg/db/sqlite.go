package db

import (
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/sqlite"
)

// CreateSQLiteConn opens a connection to the given SQLite database
// and returns a `sqlbuilder.Database` object upon success and returns
// an `error` on failure
func createSQLiteConn(dbpath string) (sqlbuilder.Database, error) {
	var cURL = sqlite.ConnectionURL{
		Database: dbpath,
	}

	return sqlite.Open(cURL)
}
