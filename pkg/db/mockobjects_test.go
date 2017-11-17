package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
	upperdb "upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)

// ConnectionURL MOCK
type mockConnectionURL struct{}

func (s mockConnectionURL) String() string {
	return "TestConnectinURL"
}

// Logger MOCK
type mockLogger struct {
}

func (s mockLogger) Log(q *upperdb.QueryStatus) {
}

//////////////////// mockSettings ///////////////////////

type mockSettings struct{}

func (s mockSettings) SetLogging(b bool) {

}

// LoggingEnabled returns true if logging is enabled, false otherwise.
func (s mockSettings) LoggingEnabled() bool {
	return false
}

// SetLogger defines which logger to use.
func (s mockSettings) SetLogger(l upperdb.Logger) {

}

// Returns the currently configured logger.
func (s mockSettings) Logger() upperdb.Logger {
	return mockLogger{}
}

// SetPreparedStatementCache enables or disables the prepared statement
// cache.
func (s mockSettings) SetPreparedStatementCache(bool) {}

// PreparedStatementCacheEnabled returns true if the prepared statement cache
// is enabled, false otherwise.
func (s mockSettings) PreparedStatementCacheEnabled() bool {
	return false
}

// SetConnMaxLifetime sets the default maximum amount of time a connection
// may be reused.
func (s mockSettings) SetConnMaxLifetime(d time.Duration) {
}

// ConnMaxLifetime returns the default maximum amount of time a connection
// may be reused.
func (s mockSettings) ConnMaxLifetime() time.Duration {
	return time.Second
}

// SetMaxIdleConns sets the default maximum number of connections in the idle
// connection pool.
func (s mockSettings) SetMaxIdleConns(i int) {
}

// MaxIdleConns returns the default maximum number of connections in the idle
// connection pool.
func (s mockSettings) MaxIdleConns() int {
	return -1
}

// SetMaxOpenConns sets the default maximum number of open connections to the
// database.
func (s mockSettings) SetMaxOpenConns(i int) {
}

// MaxOpenConns returns the default maximum number of open connections to the
// database.
func (s mockSettings) MaxOpenConns() int {
	return -1
}

func (s mockSettings) Context() context.Context {
	return context.Background()
}

type mockDatabase struct {
	WasCloseCalled bool
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
	fmt.Println("in mockDatabase.Close()")
	s.WasCloseCalled = true
	return nil
}

func (s *mockDatabase) Collection(c string) upperdb.Collection {
	return mockCollection{}
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
