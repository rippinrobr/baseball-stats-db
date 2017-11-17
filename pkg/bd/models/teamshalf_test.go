package models


import (
  "testing"
  "fmt"
  "reflect"
)

func TestGetTableNameTeamsHalf(t *testing.T) {
  out := TeamsHalf{}
  expectedValue := "teamshalf"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNameTeamsHalf(t *testing.T) {
  out := TeamsHalf{}
  expectedValue := "TeamsHalf.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathTeamsHalf(t *testing.T) {
  out := TeamsHalf{}
  expectedValue := "/Users/robertrowe/src/baseballdatabank/core/TeamsHalf.csv"
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGenParseAndStoreCSVTeamsHalfForError(t *testing.T) {
  out := TeamsHalf{}
  actualFunc, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
  fmt.Println(reflect.TypeOf(actualFunc).Name())
  if actualErr == nil {
       t.Errorf("Calling TeamsHalf.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
  }
}
