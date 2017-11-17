package models


import (
  "testing"
)

func TestGetTableNameBattingPost(t *testing.T) {
  out := BattingPost{}
  expectedValue := "battingpost"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNameBattingPost(t *testing.T) {
  out := BattingPost{}
  expectedValue := "BattingPost.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathBattingPost(t *testing.T) {
  out := BattingPost{}
  expectedValue := "/Users/robertrowe/src/baseballdatabank/core/BattingPost.csv"
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}
