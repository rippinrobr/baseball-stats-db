package db

import (
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/postgresql"
)

// CreatePostgresConn opens a connection to the given PostgeSQL database
// and returns a `sqlbuilder.Database` object upon success and returns
// an `error` on failure
func createPostgesqlConn(opts Options) (sqlbuilder.Database, error) {
	var cURL = postgresql.ConnectionURL{
		Database: opts.Name,
		Host:     opts.Host,
		User:     opts.User,
		Password: opts.Pass,
		Options: map[string]string{
			"sslmode": "disable",
		},
	}

	return postgresql.Open(cURL)
}
