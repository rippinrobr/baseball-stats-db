package databank

import (
	"testing"
)

func TestGetTableObjects(t *testing.T) {
	objs := GetTableObjects()

	if len(objs) != 27 {
		t.Errorf("There were only %d TableObjects returned, should be 27", len(objs))
	}
}
