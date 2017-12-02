package databank

import (
	upperdb "upper.io/db.v3"
)

type mockResult struct{}

// String satisfies fmt.Stringer and returns a SELECT statement for
// the result.
func (s mockResult) String() string {
	return "mock"
}

// Limit defines the maximum number of results in this set. It only has
// effect on `One()`, `All()` and `Next()`. A negative limit cancels any
// previous limit settings.
func (s mockResult) Limit(int) upperdb.Result {
	return mockResult{}
}

// Offset ignores the first *n* results. It only has effect on `One()`,
// `All()` and `Next()`. A negative offset cancels any previous offset
// settings.
func (s mockResult) Offset(int) upperdb.Result {
	return mockResult{}
}

// OrderBy receives field names that define the order in which elements will be
// returned in a query, field names may be prefixed with a minus sign (-)
// indicating descending order, ascending order will be used otherwise.
func (s mockResult) OrderBy(...interface{}) upperdb.Result {
	return mockResult{}
}

// Select defines specific columns to be returned from the elements of the
// set.
func (s mockResult) Select(...interface{}) upperdb.Result {
	return mockResult{}
}

// Where discards all the previously set filtering constraints (if any) and
// sets new ones. Commonly used when the conditions of the result depend on
// external parameters that are yet to be evaluated:
//
//   res := col.Find()
//
//   if ... {
//     res.Where(...)
//   } else {
//     res.Where(...)
//   }
func (s mockResult) Where(...interface{}) upperdb.Result {
	return mockResult{}
}

// And adds more filtering conditions on top of the existing constraints.
//
//   res := col.Find(...).And(...)
func (s mockResult) And(...interface{}) upperdb.Result {
	return mockResult{}
}

// Group is used to group results that have the same value in the same column
// or columns.
func (s mockResult) Group(...interface{}) upperdb.Result {
	return mockResult{}
}

// Delete deletes all items within the result set. `Offset()` and `Limit()` are
// not honoured by `Delete()`.
func (s mockResult) Delete() error {
	return nil
}

// Update modifies all items within the result set. `Offset()` and `Limit()`
// are not honoured by `Update()`.
func (s mockResult) Update(interface{}) error {
	return nil
}

// Count returns the number of items that match the set conditions. `Offset()`
// and `Limit()` are not honoured by `Count()`
func (s mockResult) Count() (uint64, error) {
	return 0, nil
}

// Exists returns true if at least one item on the collection exists. False
// otherwise.
func (s mockResult) Exists() (bool, error) {
	return true, nil
}

// Next fetches the next result within the result set and dumps it into the
// given pointer to struct or pointer to map. You must call
// `Close()` after finishing using `Next()`.
func (s mockResult) Next(ptrToStruct interface{}) bool {
	return true
}

// Err returns the last error that has happened with the result set, nil
// otherwise.
func (s mockResult) Err() error {
	return nil
}

// One fetches the first result within the result set and dumps it into the
// given pointer to struct or pointer to map. The result set is automatically
// closed after picking the element, so there is no need to call Close()
// after using One().
func (s mockResult) One(ptrToStruct interface{}) error {
	return nil
}

// All fetches all results within the result set and dumps them into the
// given pointer to slice of maps or structs.  The result set is
// automatically closed, so there is no need to call Close() after
// using All().
func (s mockResult) All(sliceOfStructs interface{}) error {
	return nil
}

// Paginate splits the results of the query into pages containing pageSize
// items.  When using pagination previous settings for Limit and Offset are
// ignored. Page numbering starts at 1.
//
// Use Page() to define the specific page to get results from.
//
// Example:
//
//   r = q.Paginate(12)
//
// You can provide constraints an order settings when using pagination:
//
// Example:
//
//   res := q.Where(conds).OrderBy("-id").Paginate(12)
//   err := res.Page(4).All(&items)
func (s mockResult) Paginate(pageSize uint) upperdb.Result {
	return mockResult{}
}

// Page makes the result set return results only from the page identified by
// pageNumber. Page numbering starts from 0.
//
// Example:
//
//   r = q.Paginate(12).Page(4)
func (s mockResult) Page(pageNumber uint) upperdb.Result {
	return mockResult{}
}

// Cursor defines the column that is going to be taken as basis for
// cursor-based pagination.
//
// Example:
//
//   a = q.Paginate(10).Cursor("id")
//   b = q.Paginate(12).Cursor("-id")
//
// You can set "" as cursorColumn to disable cursors.
func (s mockResult) Cursor(cursorColumn string) upperdb.Result {
	return mockResult{}
}

// NextPage returns the next results page according to the cursor. It expects
// a cursorValue, which is the value the cursor column had on the last item
// of the current result set (lower bound).
//
// Example:
//
//   cursor = q.Paginate(12).Cursor("id")
//   res = cursor.NextPage(items[len(items)-1].ID)
//
// Note that NextPage requires a cursor, any column with an absolute order
// (given two values one always precedes the other) can be a cursor.
//
// You can define the pagination order and add constraints to your result:
//
//	 cursor = q.Where(...).OrderBy("id").Paginate(10).Cursor("id")
//   res = cursor.NextPage(lowerBound)
func (s mockResult) NextPage(cursorValue interface{}) upperdb.Result {
	return mockResult{}
}

// PrevPage returns the previous results page according to the cursor. It
// expects a cursorValue, which is the value the cursor column had on the
// fist item of the current result set (upper bound).
//
// Example:
//
//   current = current.PrevPage(items[0].ID)
//
// Note that PrevPage requires a cursor, any column with an absolute order
// (given two values one always precedes the other) can be a cursor.
//
// You can define the pagination order and add constraints to your result:
//
//   cursor = q.Where(...).OrderBy("id").Paginate(10).Cursor("id")
//   res = cursor.PrevPage(upperBound)
func (s mockResult) PrevPage(cursorValue interface{}) upperdb.Result {
	return mockResult{}
}

// TotalPages returns the total number of pages the result could produce.  If
// no pagination has been set this value equals 1.
func (s mockResult) TotalPages() (uint, error) {
	return 0, nil
}

// TotalEntries returns the total number of entries in the query.
func (s mockResult) TotalEntries() (uint64, error) {
	return 0, nil
}

// Close closes the result set and frees all locked resources.
func (s mockResult) Close() error {
	return nil
}
