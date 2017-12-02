package databank


import (
  "testing"
)

func TestGetTableNameSalaries(t *testing.T) {
  out := Salaries{}
  expectedValue := "salaries"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNameSalaries(t *testing.T) {
  out := Salaries{}
  expectedValue := "Salaries.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathSalaries(t *testing.T) {
  out := Salaries{}
  expectedValue := "/mytestpath/Salaries.csv"
  out.SetInputDirectory("/mytestpath/")
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestSetInputDirectorySalaries(t *testing.T) {
  out := Salaries{}
  testPath := "/mytestpath/"
  expectedValue := testPath + "Salaries.csv"

  out.SetInputDirectory(testPath)
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGenParseAndStoreCSVSalariesForError(t *testing.T) {
  out := Salaries{}
  _, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
  if actualErr == nil {
       t.Errorf("Calling Salaries.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
  }
}
