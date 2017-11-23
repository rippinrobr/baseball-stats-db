package models

import (
	"testing"
)

func TestGetTableNameHallOfFame(t *testing.T) {
	out := HallOfFame{}
	expectedValue := "halloffame"
	actualValue := out.GetTableName()

	if actualValue != expectedValue {
		t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
	}
}

func TestGetFileNameHallOfFame(t *testing.T) {
	out := HallOfFame{}
	expectedValue := "HallOfFame.csv"
	actualValue := out.GetFileName()

	if actualValue != expectedValue {
		t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
	}
}

func TestGetFilePathHallOfFame(t *testing.T) {
	out := HallOfFame{}
	expectedValue := "/Users/robertrowe/src/baseballdatabank/core/HallOfFame.csv"
	actualValue := out.GetFilePath()

	if actualValue != expectedValue {
		t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
	}
}

func TestGenParseAndStoreCSVHallOfFameForError(t *testing.T) {
	out := HallOfFame{}
	_, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
	if actualErr == nil {
		t.Errorf("Calling HallOfFame.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
	}
}
