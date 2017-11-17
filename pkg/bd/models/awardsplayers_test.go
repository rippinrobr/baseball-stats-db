package models


import (
  "testing"
  "fmt"
  "reflect"
)

func TestGetTableNameAwardsPlayers(t *testing.T) {
  out := AwardsPlayers{}
  expectedValue := "awardsplayers"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNameAwardsPlayers(t *testing.T) {
  out := AwardsPlayers{}
  expectedValue := "AwardsPlayers.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathAwardsPlayers(t *testing.T) {
  out := AwardsPlayers{}
  expectedValue := "/Users/robertrowe/src/baseballdatabank/core/AwardsPlayers.csv"
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGenParseAndStoreCSVAwardsPlayersForError(t *testing.T) {
  out := AwardsPlayers{}
  actualFunc, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
  fmt.Println(reflect.TypeOf(actualFunc).Name())
  if actualErr == nil {
       t.Errorf("Calling AwardsPlayers.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
  }
}
