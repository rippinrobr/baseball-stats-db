package db

import (
	"context"
	"database/sql"
	"errors"
	upperdb "upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)

type mockDatabase struct {
	WasCloseCalled       bool
	CollectionNamePassed string
	TestCollection       *mockCollection
	mockSettings
}

func (s *mockDatabase) Driver() interface{} {
	return errors.New("Driver() Not Implemented")
}

func (s *mockDatabase) Open(tc upperdb.ConnectionURL) error {
	return errors.New("Open(ConnectionURL) Not Implemented")
}

func (s *mockDatabase) Ping() error {
	return errors.New("Ping() not implemented")
}

func (s *mockDatabase) Close() error {
	s.WasCloseCalled = true
	return nil
}

func (s *mockDatabase) Collection(c string) upperdb.Collection {
	s.CollectionNamePassed = c
	return s.TestCollection
}

func (s *mockDatabase) Collections() ([]string, error) {
	return nil, errors.New("Collections() not implemented")
}

func (s *mockDatabase) Name() string {
	return "mockDatabase"
}

func (s *mockDatabase) ConnectionURL() upperdb.ConnectionURL {
	return mockConnectionURL{}
}

func (s *mockDatabase) ClearCache() {}

func (s *mockDatabase) DeleteFrom(st string) sqlbuilder.Deleter {
	return nil
}

func (s *mockDatabase) Exec(interface{}, ...interface{}) (sql.Result, error) {
	return nil, nil
}

func (s *mockDatabase) ExecContext(context.Context, interface{}, ...interface{}) (sql.Result, error) {
	return nil, nil
}

func (s *mockDatabase) InsertInto(st string) sqlbuilder.Inserter {
	return nil
}

func (s *mockDatabase) Iterator(interface{}, ...interface{}) sqlbuilder.Iterator {
	return nil
}

func (s *mockDatabase) IteratorContext(context.Context, interface{}, ...interface{}) sqlbuilder.Iterator {
	return nil
}

func (s *mockDatabase) NewTx(context.Context) (sqlbuilder.Tx, error) {
	return nil, nil
}

func (s *mockDatabase) Prepare(i interface{}) (*sql.Stmt, error) {
	return nil, nil
}

func (s *mockDatabase) PrepareContext(c context.Context, i interface{}) (*sql.Stmt, error) {
	return nil, nil
}

func (s *mockDatabase) Query(interface{}, ...interface{}) (*sql.Rows, error) {
	return nil, nil
}

func (s *mockDatabase) QueryContext(context.Context, interface{}, ...interface{}) (*sql.Rows, error) {
	return nil, nil
}

func (s *mockDatabase) QueryRow(interface{}, ...interface{}) (*sql.Row, error) {
	return nil, nil
}

func (s *mockDatabase) QueryRowContext(context.Context, interface{}, ...interface{}) (*sql.Row, error) {
	return nil, nil
}

func (s *mockDatabase) Select(...interface{}) sqlbuilder.Selector {
	return nil
}

func (s *mockDatabase) SelectFrom(...interface{}) sqlbuilder.Selector {
	return nil
}

func (s *mockDatabase) SetTxOptions(sql.TxOptions) {}

func (s *mockDatabase) Tx(context.Context, func(sqlbuilder.Tx) error) error {
	return nil
}

func (s *mockDatabase) TxOptions() *sql.TxOptions {
	return nil
}

func (s *mockDatabase) Update(string) sqlbuilder.Updater {
	return nil
}

func (s *mockDatabase) WithContext(context.Context) sqlbuilder.Database {
	return s
}
