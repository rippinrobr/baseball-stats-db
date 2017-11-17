package models


import (
  "testing"
  "fmt"
  "reflect"
)

func TestGetTableNameBatting(t *testing.T) {
  out := Batting{}
  expectedValue := "batting"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNameBatting(t *testing.T) {
  out := Batting{}
  expectedValue := "Batting.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathBatting(t *testing.T) {
  out := Batting{}
  expectedValue := "/Users/robertrowe/src/baseballdatabank/core/Batting.csv"
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGenParseAndStoreCSVBattingForError(t *testing.T) {
  out := Batting{}
  actualFunc, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
  fmt.Println(reflect.TypeOf(actualFunc).Name())
  if actualErr == nil {
       t.Errorf("Calling Batting.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
  }
}
