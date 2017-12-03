package databank


import (
  "testing"
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
  expectedValue := "/mytestpath/AllstarFull.csv"
  out.SetInputDirectory("/mytestpath/")
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestSetInputDirectoryAllstarFull(t *testing.T) {
  out := AllstarFull{}
  testPath := "/mytestpath/"
  expectedValue := testPath + "AllstarFull.csv"

  out.SetInputDirectory(testPath)
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGenParseAndStoreCSVAllstarFullForError(t *testing.T) {
  out := AllstarFull{}
  _, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
  if actualErr == nil {
       t.Errorf("Calling AllstarFull.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
  }
}
