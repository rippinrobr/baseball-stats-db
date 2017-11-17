package models


import (
  "testing"
)

func TestGetTableNameHallOfFame(t *testing.T) {
  out := HallOfFame{}
  expectedValue := "halloffame"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNameHallOfFame(t *testing.T) {
  out := HallOfFame{}
  expectedValue := "HallOfFame.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathHallOfFame(t *testing.T) {
  out := HallOfFame{}
  expectedValue := "/Users/robertrowe/src/baseballdatabank/core/HallOfFame.csv"
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}
