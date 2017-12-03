package databank


import (
  "testing"
)

func TestGetTableNamePeople(t *testing.T) {
  out := People{}
  expectedValue := "people"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNamePeople(t *testing.T) {
  out := People{}
  expectedValue := "People.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathPeople(t *testing.T) {
  out := People{}
  expectedValue := "/mytestpath/People.csv"
  out.SetInputDirectory("/mytestpath/")
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestSetInputDirectoryPeople(t *testing.T) {
  out := People{}
  testPath := "/mytestpath/"
  expectedValue := testPath + "People.csv"

  out.SetInputDirectory(testPath)
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGenParseAndStoreCSVPeopleForError(t *testing.T) {
  out := People{}
  _, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
  if actualErr == nil {
       t.Errorf("Calling People.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
  }
}
