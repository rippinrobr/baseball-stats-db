package databank

import (
	"context"
	"time"
	upperdb "upper.io/db.v3"
)

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
