package models


import (
  "testing"
)

func TestGetTableNameFielding(t *testing.T) {
  out := Fielding{}
  expectedValue := "fielding"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNameFielding(t *testing.T) {
  out := Fielding{}
  expectedValue := "Fielding.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathFielding(t *testing.T) {
  out := Fielding{}
  expectedValue := "/Users/robertrowe/src/baseballdatabank/core/Fielding.csv"
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}
