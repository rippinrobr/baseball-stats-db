package databank


import (
  "testing"
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
  expectedValue := "/mytestpath/FieldingPost.csv"
  out.SetInputDirectory("/mytestpath/")
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestSetInputDirectoryFieldingPost(t *testing.T) {
  out := FieldingPost{}
  testPath := "/mytestpath/"
  expectedValue := testPath + "FieldingPost.csv"

  out.SetInputDirectory(testPath)
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGenParseAndStoreCSVFieldingPostForError(t *testing.T) {
  out := FieldingPost{}
  _, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
  if actualErr == nil {
       t.Errorf("Calling FieldingPost.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
  }
}
