package models


import (
  "testing"
  "fmt"
  "reflect"
)

func TestGetTableNameTeams(t *testing.T) {
  out := Teams{}
  expectedValue := "teams"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNameTeams(t *testing.T) {
  out := Teams{}
  expectedValue := "Teams.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathTeams(t *testing.T) {
  out := Teams{}
  expectedValue := "/Users/robertrowe/src/baseballdatabank/core/Teams.csv"
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGenParseAndStoreCSVTeamsForError(t *testing.T) {
  out := Teams{}
  actualFunc, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
  fmt.Println(reflect.TypeOf(actualFunc).Name())
  if actualErr == nil {
       t.Errorf("Calling Teams.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
  }
}
