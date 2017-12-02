package databank


import (
  "testing"
)

func TestGetTableNameAwardsShareManagers(t *testing.T) {
  out := AwardsShareManagers{}
  expectedValue := "awardssharemanagers"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNameAwardsShareManagers(t *testing.T) {
  out := AwardsShareManagers{}
  expectedValue := "AwardsShareManagers.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathAwardsShareManagers(t *testing.T) {
  out := AwardsShareManagers{}
  expectedValue := "/mytestpath/AwardsShareManagers.csv"
  out.SetInputDirectory("/mytestpath/")
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestSetInputDirectoryAwardsShareManagers(t *testing.T) {
  out := AwardsShareManagers{}
  testPath := "/mytestpath/"
  expectedValue := testPath + "AwardsShareManagers.csv"

  out.SetInputDirectory(testPath)
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGenParseAndStoreCSVAwardsShareManagersForError(t *testing.T) {
  out := AwardsShareManagers{}
  _, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
  if actualErr == nil {
       t.Errorf("Calling AwardsShareManagers.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
  }
}
