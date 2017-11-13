package db

import (
	"github.com/rippinrobr/baseball-databank-tools/pkg/bd/models"
	"upper.io/db.v3/lib/sqlbuilder"
)

// Repository is the interface used by all of the database
// objects
type Repository interface {
	CloseConn() error
	Insert(models.TableObject) error
}

type Inserter interface {
	Insert(sqlbuilder.Database, models.TableObject) error
}
