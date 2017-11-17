package models


import (
  "testing"
  "fmt"
  "reflect"
)

func TestGetTableNamePeople(t *testing.T) {
  out := People{}
  expectedValue := "people"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNamePeople(t *testing.T) {
  out := People{}
  expectedValue := "People.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathPeople(t *testing.T) {
  out := People{}
  expectedValue := "/Users/robertrowe/src/baseballdatabank/core/People.csv"
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGenParseAndStoreCSVPeopleForError(t *testing.T) {
  out := People{}
  actualFunc, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
  fmt.Println(reflect.TypeOf(actualFunc).Name())
  if actualErr == nil {
       t.Errorf("Calling People.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
  }
}
