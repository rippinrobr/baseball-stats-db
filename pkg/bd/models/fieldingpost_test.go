package models


import (
  "testing"
  "fmt"
  "reflect"
)

func TestGetTableNameFieldingPost(t *testing.T) {
  out := FieldingPost{}
  expectedValue := "fieldingpost"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNameFieldingPost(t *testing.T) {
  out := FieldingPost{}
  expectedValue := "FieldingPost.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathFieldingPost(t *testing.T) {
  out := FieldingPost{}
  expectedValue := "/Users/robertrowe/src/baseballdatabank/core/FieldingPost.csv"
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGenParseAndStoreCSVFieldingPostForError(t *testing.T) {
  out := FieldingPost{}
  actualFunc, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
  fmt.Println(reflect.TypeOf(actualFunc).Name())
  if actualErr == nil {
       t.Errorf("Calling FieldingPost.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
  }
}
