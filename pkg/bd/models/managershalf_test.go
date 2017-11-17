package models


import (
  "testing"
  "fmt"
  "reflect"
)

func TestGetTableNameManagersHalf(t *testing.T) {
  out := ManagersHalf{}
  expectedValue := "managershalf"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNameManagersHalf(t *testing.T) {
  out := ManagersHalf{}
  expectedValue := "ManagersHalf.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathManagersHalf(t *testing.T) {
  out := ManagersHalf{}
  expectedValue := "/Users/robertrowe/src/baseballdatabank/core/ManagersHalf.csv"
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGenParseAndStoreCSVManagersHalfForError(t *testing.T) {
  out := ManagersHalf{}
  actualFunc, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
  fmt.Println(reflect.TypeOf(actualFunc).Name())
  if actualErr == nil {
       t.Errorf("Calling ManagersHalf.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
  }
}
