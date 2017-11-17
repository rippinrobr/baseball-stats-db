package models


import (
  "testing"
)

func TestGetTableNameTeamsHalf(t *testing.T) {
  out := TeamsHalf{}
  expectedValue := "teamshalf"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNameTeamsHalf(t *testing.T) {
  out := TeamsHalf{}
  expectedValue := "TeamsHalf.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathTeamsHalf(t *testing.T) {
  out := TeamsHalf{}
  expectedValue := "/Users/robertrowe/src/baseballdatabank/core/TeamsHalf.csv"
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}
