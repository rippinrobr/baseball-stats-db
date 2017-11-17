package models


import (
  "testing"
)

func TestGetTableNameManagers(t *testing.T) {
  out := Managers{}
  expectedValue := "managers"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNameManagers(t *testing.T) {
  out := Managers{}
  expectedValue := "Managers.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathManagers(t *testing.T) {
  out := Managers{}
  expectedValue := "/Users/robertrowe/src/baseballdatabank/core/Managers.csv"
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}
