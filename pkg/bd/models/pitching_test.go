package models


import (
  "testing"
  "fmt"
  "reflect"
)

func TestGetTableNamePitching(t *testing.T) {
  out := Pitching{}
  expectedValue := "pitching"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNamePitching(t *testing.T) {
  out := Pitching{}
  expectedValue := "Pitching.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathPitching(t *testing.T) {
  out := Pitching{}
  expectedValue := "/Users/robertrowe/src/baseballdatabank/core/Pitching.csv"
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGenParseAndStoreCSVPitchingForError(t *testing.T) {
  out := Pitching{}
  actualFunc, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
  fmt.Println(reflect.TypeOf(actualFunc).Name())
  if actualErr == nil {
       t.Errorf("Calling Pitching.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
  }
}
