package db

import (
	"errors"
	"strings"

	"upper.io/db.v3/lib/sqlbuilder"
)

const (
	// DBSQLite is a string that indicates sqlite is supported
	DBSQLite string = "sqlite"
	// DefaultSQLiteDBName will be used when user indicates they
	// want to create a SQLite db but do not provide a path/name
	DefaultSQLiteDBName string = "baseball_databank.sqlite3"
)

// InsertFunc  Inserts the data into the table specified in the
// first parameter
type InsertFunc func(string, interface{}) error

//ErrorDBNotSupported is used to inform the caller that the
// database type the provided is not a supporeted database
var ErrorDBNotSupported = errors.New("Database not supported")

// CreateConnection is the public facing function that is responsible for
// providing a database connection to the outside world.  It acts as a
// proxy fo the db package using the `dbtype` to determine which database
// connection to return. If somehow the dbtype is value we do not support
// a DB Not Supported error will be returned
func CreateConnection(dbtype, dbconn string) (sqlbuilder.Database, error) {
	if strings.ToLower(dbtype) != DBSQLite {
		return nil, ErrorDBNotSupported
	}

	return createSQLiteConn(dbconn)
}

// IsSupportedDB checks to see if the database name given is
// one of the dbs supported by Baseball Databank Tools
func IsSupportedDB(s string) bool {

	if strings.ToLower(s) == DBSQLite {
		return true
	}

	return false
}
