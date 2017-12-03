package db

import (
	db "upper.io/db.v3"
	"upper.io/db.v3/mongo"
)

// createMySQLConn opens a connection to the given MySQL database
// and returns a `sqlbuilder.Database` object upon success and returns
// an `error` on failure
func createMongoDBConn(opts Options) (db.Database, error) {
	var cURL = mongo.ConnectionURL{
		Database: opts.Name,
		Host:     opts.Host,
		User:     opts.User,
		Password: opts.Pass,
		Options:  map[string]string{},
	}

	return mongo.Open(cURL)
}
