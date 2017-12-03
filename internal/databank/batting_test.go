package databank


import (
  "testing"
)

func TestGetTableNameBatting(t *testing.T) {
  out := Batting{}
  expectedValue := "batting"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNameBatting(t *testing.T) {
  out := Batting{}
  expectedValue := "Batting.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathBatting(t *testing.T) {
  out := Batting{}
  expectedValue := "/mytestpath/Batting.csv"
  out.SetInputDirectory("/mytestpath/")
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestSetInputDirectoryBatting(t *testing.T) {
  out := Batting{}
  testPath := "/mytestpath/"
  expectedValue := testPath + "Batting.csv"

  out.SetInputDirectory(testPath)
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGenParseAndStoreCSVBattingForError(t *testing.T) {
  out := Batting{}
  _, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
  if actualErr == nil {
       t.Errorf("Calling Batting.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
  }
}
