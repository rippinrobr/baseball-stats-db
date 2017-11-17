package models


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
  expectedValue := "/Users/robertrowe/src/baseballdatabank/core/PitchingPost.csv"
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}
