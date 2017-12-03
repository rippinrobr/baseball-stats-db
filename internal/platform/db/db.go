package db

import (
	"errors"
	"fmt"

	db "upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mongo"
)

const (
	// DatabaseNotSupportedErrMsg is the string used to indicate the provided DB is not supported
	DatabaseNotSupportedErrMsg = "Database not supported"

	// DBSQLite is a string that indicates sqlite is supported
	DBSQLite string = "sqlite"

	// DBPostgres is a string that indicates postgres is supported
	DBPostgres string = "postgres"

	// DBMongoDB is a string that indicates MongoDB is a supported database
	DBMongoDB string = "mongodb"

	// DBMySQL is a string that indicates mysql is supported
	DBMySQL string = "mysql"

	// DefaultSQLiteDBName will be used when user indicates they
	// want to create a SQLite db but do not provide a path/name
	DefaultSQLiteDBName string = "baseball_databank.sqlite3"
)

// CreateDBConnFunc is the type of function that all supported
// databases must support
type CreateDBConnFunc func(Options) (sqlbuilder.Database, error)

// InsertFunc  Inserts the data into the table specified in the
// first parameter
type InsertFunc func(string, interface{}) error

//ErrorDBNotSupported is used to inform the caller that the
// database type the provided is not a supporeted database
var ErrorDBNotSupported = errors.New(DatabaseNotSupportedErrMsg)

// CreateConnection is the public facing function that is responsible for
// providing a database connection to the outside world.  It acts as a
// proxy fo the db package using the `dbtype` to determine which database
// connection to return. If somehow the dbtype is value we do not support
// a DB Not Supported error will be returned
func CreateConnection(opts Options) (sqlbuilder.Database, error) {
	connFuncs := map[string]CreateDBConnFunc{
		DBSQLite:   createSQLiteConn,
		DBPostgres: createPostgesqlConn,
		DBMySQL:    createMySQLConn,
	}

	if IsSupportedDB(opts.Type) {
		if f, ok := connFuncs[opts.Type]; ok {
			return f(opts)
		}
	}

	return nil, ErrorDBNotSupported
}

// CreateNoSQLConnection is the public facing function that is responsible for
// providing a nosql database connection to the outside world.  It acts as a
// proxy fo the db package using the `dbtype` to determine which database
// connection to return. If somehow the dbtype is value we do not support
// a DB Not Supported error will be returned
func CreateNoSQLConnection(opts Options) (db.Database, error) {
	if IsSupportedDB(opts.Type) {
		var settings = mongo.ConnectionURL{
			Database: opts.Name,
			Host:     opts.Host,
			User:     opts.User,
			Password: opts.Pass,
		}

		return mongo.Open(settings)
	}

	return nil, ErrorDBNotSupported
}

// IsSupportedDB checks to see if the database name given is
// one of the dbs supported by Baseball Databank Tools
func IsSupportedDB(s string) bool {
	for _, d := range []string{DBSQLite, DBPostgres, DBMySQL, DBMongoDB} {
		if s == d {
			return true
		}
	}
	return false
}

// Options are used to connect to the database file or server depending
// on the type of database being used
type Options struct {
	Host    string
	Name    string
	Pass    string
	Path    string
	Port    int
	Type    string
	User    string
	Verbose bool
}

func (o Options) String() string {
	return fmt.Sprintf("Options: Host='%s', Name='%s', Pass='%s', Path='%s', Port='%d', Type='%s', User='%s', Verbose='%t'\n",
		o.Host, o.Name, o.Pass, o.Path, o.Port, o.Type, o.User, o.Verbose)
}
