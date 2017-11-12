package db

import "github.com/rippinrobr/baseball-databank-tools/pkg/bd/models"

// Repository is the interface used by all of the database
// objects
type Repository interface {
	CloseConn() error
	OpenConn(string) error
	Insert(models.TableObject) error
}
