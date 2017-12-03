package databank


import (
  "testing"
)

func TestGetTableNameAwardsPlayers(t *testing.T) {
  out := AwardsPlayers{}
  expectedValue := "awardsplayers"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNameAwardsPlayers(t *testing.T) {
  out := AwardsPlayers{}
  expectedValue := "AwardsPlayers.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathAwardsPlayers(t *testing.T) {
  out := AwardsPlayers{}
  expectedValue := "/mytestpath/AwardsPlayers.csv"
  out.SetInputDirectory("/mytestpath/")
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestSetInputDirectoryAwardsPlayers(t *testing.T) {
  out := AwardsPlayers{}
  testPath := "/mytestpath/"
  expectedValue := testPath + "AwardsPlayers.csv"

  out.SetInputDirectory(testPath)
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGenParseAndStoreCSVAwardsPlayersForError(t *testing.T) {
  out := AwardsPlayers{}
  _, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
  if actualErr == nil {
       t.Errorf("Calling AwardsPlayers.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
  }
}
