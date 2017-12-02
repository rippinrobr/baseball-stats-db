package databank


import (
  "testing"
)

func TestGetTableNameCollegePlaying(t *testing.T) {
  out := CollegePlaying{}
  expectedValue := "collegeplaying"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNameCollegePlaying(t *testing.T) {
  out := CollegePlaying{}
  expectedValue := "CollegePlaying.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathCollegePlaying(t *testing.T) {
  out := CollegePlaying{}
  expectedValue := "/mytestpath/CollegePlaying.csv"
  out.SetInputDirectory("/mytestpath/")
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestSetInputDirectoryCollegePlaying(t *testing.T) {
  out := CollegePlaying{}
  testPath := "/mytestpath/"
  expectedValue := testPath + "CollegePlaying.csv"

  out.SetInputDirectory(testPath)
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGenParseAndStoreCSVCollegePlayingForError(t *testing.T) {
  out := CollegePlaying{}
  _, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
  if actualErr == nil {
       t.Errorf("Calling CollegePlaying.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
  }
}
