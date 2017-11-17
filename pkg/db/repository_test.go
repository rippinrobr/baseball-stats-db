package db

import (
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
