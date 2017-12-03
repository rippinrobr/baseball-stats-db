package db

import (
	db "upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)

// Repository is the interface used by all of the database
// objects
type Repository interface {
	CloseConn() error
	Insert(string, interface{}) error
	Truncate(string) error
}

// Repo is the struct that manages the inserting of data
// and closing of db connections across the supported
// databases
type Repo struct {
	dbsess sqlbuilder.Database
}

// CloseConn closes the connection to the database
func (r Repo) CloseConn() error {
	return r.dbsess.Close()
}

// Insert takes a slice of and adds it to the given table
func (r Repo) Insert(tableName string, data interface{}) error {
	_, err := r.dbsess.Collection(tableName).Insert(data)
	return err
}

// Truncate takes the name of the table to truncate
// and returns any errors that occurr
func (r Repo) Truncate(tableName string) error {
	return r.dbsess.Collection(tableName).Truncate()
}

// CreateRepo creates a Repositoryobject that is
// connected to the database.
func CreateRepo(sd sqlbuilder.Database) Repository {
	return &Repo{
		dbsess: sd,
	}
}

// NoSQLRepo handles interaction between the database and
// the caller
type NoSQLRepo struct {
	dbsess db.Database
}

// CloseConn closes the connection to the database
func (r NoSQLRepo) CloseConn() error {
	return r.dbsess.Close()
}

// Insert takes a slice of and adds it to the given table
func (r NoSQLRepo) Insert(tableName string, data interface{}) error {
	_, err := r.dbsess.Collection(tableName).Insert(data)
	return err
}

// Truncate takes the name of the table to truncate
// and returns any errors that occurr
func (r NoSQLRepo) Truncate(tableName string) error {
	return r.dbsess.Collection(tableName).Truncate()
}

// CreateNoSQLRepo creates a Repository object that is
// connected to the database.
func CreateNoSQLRepo(sd db.Database) Repository {
	return &NoSQLRepo{
		dbsess: sd,
	}
}
