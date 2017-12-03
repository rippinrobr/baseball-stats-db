package databank


import (
  "testing"
)

func TestGetTableNameParks(t *testing.T) {
  out := Parks{}
  expectedValue := "parks"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNameParks(t *testing.T) {
  out := Parks{}
  expectedValue := "Parks.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathParks(t *testing.T) {
  out := Parks{}
  expectedValue := "/mytestpath/Parks.csv"
  out.SetInputDirectory("/mytestpath/")
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestSetInputDirectoryParks(t *testing.T) {
  out := Parks{}
  testPath := "/mytestpath/"
  expectedValue := testPath + "Parks.csv"

  out.SetInputDirectory(testPath)
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGenParseAndStoreCSVParksForError(t *testing.T) {
  out := Parks{}
  _, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
  if actualErr == nil {
       t.Errorf("Calling Parks.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
  }
}
