package models


import (
  "testing"
  "fmt"
  "reflect"
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
  actualFunc, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
  fmt.Println(reflect.TypeOf(actualFunc).Name())
  if actualErr == nil {
       t.Errorf("Calling AwardsSharePlayers.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
  }
}
