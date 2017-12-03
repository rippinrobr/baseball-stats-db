package databank


import (
  "testing"
)

func TestGetTableNameSeriesPost(t *testing.T) {
  out := SeriesPost{}
  expectedValue := "seriespost"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNameSeriesPost(t *testing.T) {
  out := SeriesPost{}
  expectedValue := "SeriesPost.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathSeriesPost(t *testing.T) {
  out := SeriesPost{}
  expectedValue := "/mytestpath/SeriesPost.csv"
  out.SetInputDirectory("/mytestpath/")
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestSetInputDirectorySeriesPost(t *testing.T) {
  out := SeriesPost{}
  testPath := "/mytestpath/"
  expectedValue := testPath + "SeriesPost.csv"

  out.SetInputDirectory(testPath)
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGenParseAndStoreCSVSeriesPostForError(t *testing.T) {
  out := SeriesPost{}
  _, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
  if actualErr == nil {
       t.Errorf("Calling SeriesPost.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
  }
}
