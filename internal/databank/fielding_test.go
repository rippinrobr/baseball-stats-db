package databank


import (
  "testing"
)

func TestGetTableNameFielding(t *testing.T) {
  out := Fielding{}
  expectedValue := "fielding"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNameFielding(t *testing.T) {
  out := Fielding{}
  expectedValue := "Fielding.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathFielding(t *testing.T) {
  out := Fielding{}
  expectedValue := "/mytestpath/Fielding.csv"
  out.SetInputDirectory("/mytestpath/")
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestSetInputDirectoryFielding(t *testing.T) {
  out := Fielding{}
  testPath := "/mytestpath/"
  expectedValue := testPath + "Fielding.csv"

  out.SetInputDirectory(testPath)
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGenParseAndStoreCSVFieldingForError(t *testing.T) {
  out := Fielding{}
  _, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
  if actualErr == nil {
       t.Errorf("Calling Fielding.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
  }
}
