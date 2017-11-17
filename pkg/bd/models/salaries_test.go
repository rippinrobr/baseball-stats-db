package models


import (
  "testing"
  "fmt"
  "reflect"
)

func TestGetTableNameSalaries(t *testing.T) {
  out := Salaries{}
  expectedValue := "salaries"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNameSalaries(t *testing.T) {
  out := Salaries{}
  expectedValue := "Salaries.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathSalaries(t *testing.T) {
  out := Salaries{}
  expectedValue := "/Users/robertrowe/src/baseballdatabank/core/Salaries.csv"
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGenParseAndStoreCSVSalariesForError(t *testing.T) {
  out := Salaries{}
  actualFunc, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
  fmt.Println(reflect.TypeOf(actualFunc).Name())
  if actualErr == nil {
       t.Errorf("Calling Salaries.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
  }
}
