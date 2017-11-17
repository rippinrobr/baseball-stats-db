package models


import (
  "testing"
)

func TestGetTableNameParks(t *testing.T) {
  out := Parks{}
  expectedValue := "parks"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNameParks(t *testing.T) {
  out := Parks{}
  expectedValue := "Parks.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathParks(t *testing.T) {
  out := Parks{}
  expectedValue := "/Users/robertrowe/src/baseballdatabank/core/Parks.csv"
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}
