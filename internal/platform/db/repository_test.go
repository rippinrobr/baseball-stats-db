package db

import (
	"fmt"
	"testing"
)

func TestRepoCloseConn(t *testing.T) {
	d := &mockDatabase{}

	r := CreateRepo(d)

	r.CloseConn()
	if !d.WasCloseCalled {
		t.Error("db.Close was not called")
	}
}

func TestRepoInsert(t *testing.T) {
	d := &mockDatabase{
		TestCollection: &mockCollection{},
	}
	colName := "mytesttable"
	r := CreateRepo(d)

	r.Insert(colName, "ddd")
	if d.CollectionNamePassed != colName {
		t.Error("The collection name passed to the collection call did not match the one passed to Insert")
	}

	fmt.Printf("d.TestCollection %+v\n", d.TestCollection)
	if !d.TestCollection.WasInsertCalled {
		t.Error("The Collection.Insert was not called")
	}

}

func TestRepoTruncate(t *testing.T) {
	d := &mockDatabase{
		TestCollection: &mockCollection{},
	}
	colName := "mytesttable"
	r := CreateRepo(d)

	r.Truncate(colName)
	if d.CollectionNamePassed != colName {
		t.Error("The collection name passed to the collection call did not match the one passed to Truncate")
	}

	if !d.TestCollection.WasTruncateCalled {
		t.Error("The Collection.Truncate was not called")
	}

}
