package models


import (
  "testing"
  "fmt"
  "reflect"
)

func TestGetTableNameAwardsShareManagers(t *testing.T) {
  out := AwardsShareManagers{}
  expectedValue := "awardssharemanagers"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNameAwardsShareManagers(t *testing.T) {
  out := AwardsShareManagers{}
  expectedValue := "AwardsShareManagers.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathAwardsShareManagers(t *testing.T) {
  out := AwardsShareManagers{}
  expectedValue := "/Users/robertrowe/src/baseballdatabank/core/AwardsShareManagers.csv"
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGenParseAndStoreCSVAwardsShareManagersForError(t *testing.T) {
  out := AwardsShareManagers{}
  actualFunc, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
  fmt.Println(reflect.TypeOf(actualFunc).Name())
  if actualErr == nil {
       t.Errorf("Calling AwardsShareManagers.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
  }
}
