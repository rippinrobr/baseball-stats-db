package models


import (
  "testing"
  "fmt"
  "reflect"
)

func TestGetTableNameAwardsManagers(t *testing.T) {
  out := AwardsManagers{}
  expectedValue := "awardsmanagers"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNameAwardsManagers(t *testing.T) {
  out := AwardsManagers{}
  expectedValue := "AwardsManagers.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathAwardsManagers(t *testing.T) {
  out := AwardsManagers{}
  expectedValue := "/Users/robertrowe/src/baseballdatabank/core/AwardsManagers.csv"
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGenParseAndStoreCSVAwardsManagersForError(t *testing.T) {
  out := AwardsManagers{}
  actualFunc, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
  fmt.Println(reflect.TypeOf(actualFunc).Name())
  if actualErr == nil {
       t.Errorf("Calling AwardsManagers.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
  }
}
