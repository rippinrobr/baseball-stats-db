package db

import upperdb "upper.io/db.v3"

// Logger MOCK
type mockLogger struct {
}

func (s mockLogger) Log(q *upperdb.QueryStatus) {
}
