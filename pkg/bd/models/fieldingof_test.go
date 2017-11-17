package models


import (
  "testing"
)

func TestGetTableNameFieldingOF(t *testing.T) {
  out := FieldingOF{}
  expectedValue := "fieldingof"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNameFieldingOF(t *testing.T) {
  out := FieldingOF{}
  expectedValue := "FieldingOF.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathFieldingOF(t *testing.T) {
  out := FieldingOF{}
  expectedValue := "/Users/robertrowe/src/baseballdatabank/core/FieldingOF.csv"
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}
