package models


import (
  "testing"
  "fmt"
  "reflect"
)

func TestGetTableNamePitchingPost(t *testing.T) {
  out := PitchingPost{}
  expectedValue := "pitchingpost"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNamePitchingPost(t *testing.T) {
  out := PitchingPost{}
  expectedValue := "PitchingPost.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathPitchingPost(t *testing.T) {
  out := PitchingPost{}
  expectedValue := "/Users/robertrowe/src/baseballdatabank/core/PitchingPost.csv"
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGenParseAndStoreCSVPitchingPostForError(t *testing.T) {
  out := PitchingPost{}
  actualFunc, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
  fmt.Println(reflect.TypeOf(actualFunc).Name())
  if actualErr == nil {
       t.Errorf("Calling PitchingPost.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
  }
}
