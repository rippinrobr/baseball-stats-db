package db

import (
	"fmt"
	upperdb "upper.io/db.v3"
)

type mockCollection struct {
	WasInsertCalled   bool
	WasTruncateCalled bool
}

func (s *mockCollection) Insert(interface{}) (interface{}, error) {
	s.WasInsertCalled = true
	fmt.Println("Insert", s.WasInsertCalled)
	return nil, nil
}

// InsertReturning is like Insert() but it updates the passed pointer to map
// or struct with the newly inserted element (and with automatic fields, like
// IDs, timestamps, etc). This is all done atomically within a transaction.
// If the database does not support transactions this method returns
// db.ErrUnsupported.
func (s mockCollection) InsertReturning(interface{}) error {
	return nil
}

// UpdateReturning takes a pointer to map or struct and tries to update the
// given item on the collection based on the item's primary keys. Once the
// element is updated, UpdateReturning will query the element that was just
// updated. If the database does not support transactions this method returns
// db.ErrUnsupported
func (s mockCollection) UpdateReturning(interface{}) error {
	return nil
}

// Exists returns true if the collection exists, false otherwise.
func (s mockCollection) Exists() bool {
	return true
}

// Find defines a new result set with elements from the collection.
func (s mockCollection) Find(...interface{}) upperdb.Result {
	return mockResult{}
}

// Truncate removes all elements on the collection and resets the
// collection's IDs.
func (s *mockCollection) Truncate() error {
	s.WasTruncateCalled = true
	return nil
}

// Name returns the name of the collection.
func (s mockCollection) Name() string {
	return "mockCollection"
}
