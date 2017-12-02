package databank


import (
  "testing"
)

func TestGetTableNameFieldingOFsplit(t *testing.T) {
  out := FieldingOFsplit{}
  expectedValue := "fieldingofsplit"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNameFieldingOFsplit(t *testing.T) {
  out := FieldingOFsplit{}
  expectedValue := "FieldingOFsplit.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathFieldingOFsplit(t *testing.T) {
  out := FieldingOFsplit{}
  expectedValue := "/mytestpath/FieldingOFsplit.csv"
  out.SetInputDirectory("/mytestpath/")
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestSetInputDirectoryFieldingOFsplit(t *testing.T) {
  out := FieldingOFsplit{}
  testPath := "/mytestpath/"
  expectedValue := testPath + "FieldingOFsplit.csv"

  out.SetInputDirectory(testPath)
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGenParseAndStoreCSVFieldingOFsplitForError(t *testing.T) {
  out := FieldingOFsplit{}
  _, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
  if actualErr == nil {
       t.Errorf("Calling FieldingOFsplit.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
  }
}
