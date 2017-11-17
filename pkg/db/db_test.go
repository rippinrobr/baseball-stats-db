package db

import (
	"testing"
)
  
func TestCreateConnection(t *testing.T) {
	_, err := CreateConnection("FredDB", "")
	if err == nil {
	  t.Errorf("CreateConnection should have thrown error since FredDB isn't a supported database.\n")
	}
}

func TestCreateConnectionWithEmptydbconn(t *testing.T) {
	_ , err := CreateConnection("SQLITE", "")
	if err == nil {
		t.Error("CreateConnection should have thrown an error since it received an empty string for dbconn")
	}
}