package models


import (
  "testing"
  "fmt"
  "reflect"
)

func TestGetTableNameCollegePlaying(t *testing.T) {
  out := CollegePlaying{}
  expectedValue := "collegeplaying"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNameCollegePlaying(t *testing.T) {
  out := CollegePlaying{}
  expectedValue := "CollegePlaying.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathCollegePlaying(t *testing.T) {
  out := CollegePlaying{}
  expectedValue := "/Users/robertrowe/src/baseballdatabank/core/CollegePlaying.csv"
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGenParseAndStoreCSVCollegePlayingForError(t *testing.T) {
  out := CollegePlaying{}
  actualFunc, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
  fmt.Println(reflect.TypeOf(actualFunc).Name())
  if actualErr == nil {
       t.Errorf("Calling CollegePlaying.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
  }
}
