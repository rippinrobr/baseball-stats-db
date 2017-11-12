package db

import (
	"errors"

	"github.com/rippinrobr/baseball-databank-tools/pkg/bd/models"

	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/sqlite"
)

// SQLiteRepo manages interaction with SQLite dbs
type SQLiteRepo struct {
	dbsess sqlbuilder.Database
}

// CloseConn closes the connection to the database
func (s *SQLiteRepo) CloseConn() error {
	return s.dbsess.Close()
}

// OpenConn opens a connection to the given database or
// returns an error
func (s *SQLiteRepo) OpenConn(dbPath string) error {
	var cURL = sqlite.ConnectionURL{
		Database: dbPath,
	}

	sess, err := sqlite.Open(cURL)
	if err != nil {
		return err
	}
	s.dbsess = sess

	return nil
}

// Insert adds a new record to the database
func (s *SQLiteRepo) Insert(o models.TableObject) error {
	return errors.New("Not implemented")
}
