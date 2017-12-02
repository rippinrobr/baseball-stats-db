package databank


import (
  "testing"
)

func TestGetTableNameTeamsHalf(t *testing.T) {
  out := TeamsHalf{}
  expectedValue := "teamshalf"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNameTeamsHalf(t *testing.T) {
  out := TeamsHalf{}
  expectedValue := "TeamsHalf.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathTeamsHalf(t *testing.T) {
  out := TeamsHalf{}
  expectedValue := "/mytestpath/TeamsHalf.csv"
  out.SetInputDirectory("/mytestpath/")
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestSetInputDirectoryTeamsHalf(t *testing.T) {
  out := TeamsHalf{}
  testPath := "/mytestpath/"
  expectedValue := testPath + "TeamsHalf.csv"

  out.SetInputDirectory(testPath)
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGenParseAndStoreCSVTeamsHalfForError(t *testing.T) {
  out := TeamsHalf{}
  _, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
  if actualErr == nil {
       t.Errorf("Calling TeamsHalf.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
  }
}
