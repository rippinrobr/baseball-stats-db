package models


import (
  "testing"
)

func TestGetTableNameHomeGames(t *testing.T) {
  out := HomeGames{}
  expectedValue := "homegames"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNameHomeGames(t *testing.T) {
  out := HomeGames{}
  expectedValue := "HomeGames.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathHomeGames(t *testing.T) {
  out := HomeGames{}
  expectedValue := "/Users/robertrowe/src/baseballdatabank/core/HomeGames.csv"
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}
