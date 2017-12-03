package databank


import (
  "testing"
)

func TestGetTableNameManagers(t *testing.T) {
  out := Managers{}
  expectedValue := "managers"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNameManagers(t *testing.T) {
  out := Managers{}
  expectedValue := "Managers.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathManagers(t *testing.T) {
  out := Managers{}
  expectedValue := "/mytestpath/Managers.csv"
  out.SetInputDirectory("/mytestpath/")
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestSetInputDirectoryManagers(t *testing.T) {
  out := Managers{}
  testPath := "/mytestpath/"
  expectedValue := testPath + "Managers.csv"

  out.SetInputDirectory(testPath)
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGenParseAndStoreCSVManagersForError(t *testing.T) {
  out := Managers{}
  _, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
  if actualErr == nil {
       t.Errorf("Calling Managers.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
  }
}
