package models


import (
  "testing"
)

func TestGetTableNameFieldingOFsplit(t *testing.T) {
  out := FieldingOFsplit{}
  expectedValue := "fieldingofsplit"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNameFieldingOFsplit(t *testing.T) {
  out := FieldingOFsplit{}
  expectedValue := "FieldingOFsplit.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathFieldingOFsplit(t *testing.T) {
  out := FieldingOFsplit{}
  expectedValue := "/Users/robertrowe/src/baseballdatabank/core/FieldingOFsplit.csv"
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}
