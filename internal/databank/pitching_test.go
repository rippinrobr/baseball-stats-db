package databank


import (
  "testing"
)

func TestGetTableNamePitching(t *testing.T) {
  out := Pitching{}
  expectedValue := "pitching"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNamePitching(t *testing.T) {
  out := Pitching{}
  expectedValue := "Pitching.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathPitching(t *testing.T) {
  out := Pitching{}
  expectedValue := "/mytestpath/Pitching.csv"
  out.SetInputDirectory("/mytestpath/")
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestSetInputDirectoryPitching(t *testing.T) {
  out := Pitching{}
  testPath := "/mytestpath/"
  expectedValue := testPath + "Pitching.csv"

  out.SetInputDirectory(testPath)
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGenParseAndStoreCSVPitchingForError(t *testing.T) {
  out := Pitching{}
  _, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
  if actualErr == nil {
       t.Errorf("Calling Pitching.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
  }
}
