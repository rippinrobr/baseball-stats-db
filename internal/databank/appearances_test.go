package databank


import (
  "testing"
)

func TestGetTableNameAppearances(t *testing.T) {
  out := Appearances{}
  expectedValue := "appearances"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNameAppearances(t *testing.T) {
  out := Appearances{}
  expectedValue := "Appearances.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathAppearances(t *testing.T) {
  out := Appearances{}
  expectedValue := "/mytestpath/Appearances.csv"
  out.SetInputDirectory("/mytestpath/")
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestSetInputDirectoryAppearances(t *testing.T) {
  out := Appearances{}
  testPath := "/mytestpath/"
  expectedValue := testPath + "Appearances.csv"

  out.SetInputDirectory(testPath)
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGenParseAndStoreCSVAppearancesForError(t *testing.T) {
  out := Appearances{}
  _, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
  if actualErr == nil {
       t.Errorf("Calling Appearances.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
  }
}
