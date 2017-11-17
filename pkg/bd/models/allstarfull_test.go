package models


import (
  "testing"
  "fmt"
  "reflect"
)

func TestGetTableNameAllstarFull(t *testing.T) {
  out := AllstarFull{}
  expectedValue := "allstarfull"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNameAllstarFull(t *testing.T) {
  out := AllstarFull{}
  expectedValue := "AllstarFull.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathAllstarFull(t *testing.T) {
  out := AllstarFull{}
  expectedValue := "/Users/robertrowe/src/baseballdatabank/core/AllstarFull.csv"
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGenParseAndStoreCSVAllstarFullForError(t *testing.T) {
  out := AllstarFull{}
  actualFunc, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
  fmt.Println(reflect.TypeOf(actualFunc).Name())
  if actualErr == nil {
       t.Errorf("Calling AllstarFull.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
  }
}
