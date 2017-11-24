package db

import (
	"testing"
)

func TestCreateConnection(t *testing.T) {
	_, err := CreateConnection(Options{Type: "FredDB"})
	if err == nil {
		t.Errorf("CreateConnection should have thrown error since FredDB isn't a supported database.\n")
	}
}

func TestCreateConnectionWithEmptydbconn(t *testing.T) {

	_, err := CreateConnection(Options{Type: DBSQLite})
	if err == nil {
		t.Error("CreateConnection should have thrown an error since it received an empty string for dbconn")
	}
}

func TestIsSupportedDBWithSqlite(t *testing.T) {
	if !IsSupportedDB(DBSQLite) {
		t.Error("IsSupportedDB should have returned true for DBSQLite but it returned false")
	}
}

func TestIsSupportedDBWithPostgress(t *testing.T) {
	if !IsSupportedDB(DBPostgres) {
		t.Error("IsSupportedDB should have returned true for DBPostgres but it returned false")
	}
}

func TestIsSupportedDBWithMySQL(t *testing.T) {
	if !IsSupportedDB(DBMySQL) {
		t.Error("IsSupportedDB should have returned true for DBMySQL but it returned false")
	}
}

func TestIsSupportedDBWithUnsupportedType(t *testing.T) {
	if IsSupportedDB("FredDB") {
		t.Error("IsSupportedDB should have returned true for DBPostgres but it returned false")
	}
}
