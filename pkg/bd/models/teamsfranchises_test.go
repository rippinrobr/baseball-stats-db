package models

import (
	"testing"
)

func TestGetTableNameTeamsFranchises(t *testing.T) {
	out := TeamsFranchises{}
	expectedValue := "teamsfranchises"
	actualValue := out.GetTableName()

	if actualValue != expectedValue {
		t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
	}
}

func TestGetFileNameTeamsFranchises(t *testing.T) {
	out := TeamsFranchises{}
	expectedValue := "TeamsFranchises.csv"
	actualValue := out.GetFileName()

	if actualValue != expectedValue {
		t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
	}
}

func TestGetFilePathTeamsFranchises(t *testing.T) {
	out := TeamsFranchises{}
	expectedValue := "/Users/robertrowe/src/baseballdatabank/core/TeamsFranchises.csv"
	actualValue := out.GetFilePath()

	if actualValue != expectedValue {
		t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
	}
}

func TestGenParseAndStoreCSVTeamsFranchisesForError(t *testing.T) {
	out := TeamsFranchises{}
	_, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
	if actualErr == nil {
		t.Errorf("Calling TeamsFranchises.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
	}
}
