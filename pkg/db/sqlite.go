package db

import (
	"errors"

	"github.com/rippinrobr/baseball-databank-tools/pkg/bd/models"

	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/sqlite"
)

func CreateSQLiteConn(dbpath string) (sqlbuilder.Database, error) {
	var cURL = sqlite.ConnectionURL{
		Database: dbpath,
	}

	return sqlite.Open(cURL)
}

type Repo struct {
	dbsess sqlbuilder.Database
}

func CreateRepo(sd sqlbuilder.Database) Repository {
	return Repo{
		dbsess: sd,
	}
}

// SQLiteRepo manages interaction with SQLite dbs
//type SQLiteRepo struct {
//dbsess sqlbuilder.Database
//}

// CloseConn closes the connection to the database
func (r Repo) CloseConn() error {
	return r.dbsess.Close()
}

// Insert adds a new record to the database
func (r Repo) Insert(o models.TableObject) error {
	return errors.New("Not implemented")
}
