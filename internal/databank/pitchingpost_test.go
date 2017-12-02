package databank


import (
  "testing"
)

func TestGetTableNamePitchingPost(t *testing.T) {
  out := PitchingPost{}
  expectedValue := "pitchingpost"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNamePitchingPost(t *testing.T) {
  out := PitchingPost{}
  expectedValue := "PitchingPost.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathPitchingPost(t *testing.T) {
  out := PitchingPost{}
  expectedValue := "/mytestpath/PitchingPost.csv"
  out.SetInputDirectory("/mytestpath/")
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestSetInputDirectoryPitchingPost(t *testing.T) {
  out := PitchingPost{}
  testPath := "/mytestpath/"
  expectedValue := testPath + "PitchingPost.csv"

  out.SetInputDirectory(testPath)
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGenParseAndStoreCSVPitchingPostForError(t *testing.T) {
  out := PitchingPost{}
  _, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
  if actualErr == nil {
       t.Errorf("Calling PitchingPost.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
  }
}
