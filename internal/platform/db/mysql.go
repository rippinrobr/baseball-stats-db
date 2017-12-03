package db

import (
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
)

// createMySQLConn opens a connection to the given MySQL database
// and returns a `sqlbuilder.Database` object upon success and returns
// an `error` on failure
func createMySQLConn(opts Options) (sqlbuilder.Database, error) {
	var cURL = mysql.ConnectionURL{
		Database: opts.Name,
		Host:     opts.Host,
		User:     opts.User,
		Password: opts.Pass,
		Options:  map[string]string{},
	}

	return mysql.Open(cURL)
}
