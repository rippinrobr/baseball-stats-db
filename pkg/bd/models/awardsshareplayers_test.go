package models


import (
  "testing"
)

func TestGetTableNameAwardsSharePlayers(t *testing.T) {
  out := AwardsSharePlayers{}
  expectedValue := "awardsshareplayers"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNameAwardsSharePlayers(t *testing.T) {
  out := AwardsSharePlayers{}
  expectedValue := "AwardsSharePlayers.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathAwardsSharePlayers(t *testing.T) {
  out := AwardsSharePlayers{}
  expectedValue := "/Users/robertrowe/src/baseballdatabank/core/AwardsSharePlayers.csv"
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGenParseAndStoreCSVAwardsSharePlayersForError(t *testing.T) {
  out := AwardsSharePlayers{}
  _, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
  if actualErr == nil {
       t.Errorf("Calling AwardsSharePlayers.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
  }
}
