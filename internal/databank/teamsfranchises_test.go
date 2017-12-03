package databank


import (
  "testing"
)

func TestGetTableNameTeamsFranchises(t *testing.T) {
  out := TeamsFranchises{}
  expectedValue := "teamsfranchises"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNameTeamsFranchises(t *testing.T) {
  out := TeamsFranchises{}
  expectedValue := "TeamsFranchises.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathTeamsFranchises(t *testing.T) {
  out := TeamsFranchises{}
  expectedValue := "/mytestpath/TeamsFranchises.csv"
  out.SetInputDirectory("/mytestpath/")
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestSetInputDirectoryTeamsFranchises(t *testing.T) {
  out := TeamsFranchises{}
  testPath := "/mytestpath/"
  expectedValue := testPath + "TeamsFranchises.csv"

  out.SetInputDirectory(testPath)
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGenParseAndStoreCSVTeamsFranchisesForError(t *testing.T) {
  out := TeamsFranchises{}
  _, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
  if actualErr == nil {
       t.Errorf("Calling TeamsFranchises.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
  }
}
