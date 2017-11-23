package models

import (
	"testing"
)

func TestGetTableNameSeriesPost(t *testing.T) {
	out := SeriesPost{}
	expectedValue := "seriespost"
	actualValue := out.GetTableName()

	if actualValue != expectedValue {
		t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
	}
}

func TestGetFileNameSeriesPost(t *testing.T) {
	out := SeriesPost{}
	expectedValue := "SeriesPost.csv"
	actualValue := out.GetFileName()

	if actualValue != expectedValue {
		t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
	}
}

func TestGetFilePathSeriesPost(t *testing.T) {
	out := SeriesPost{}
	expectedValue := "/Users/robertrowe/src/baseballdatabank/core/SeriesPost.csv"
	actualValue := out.GetFilePath()

	if actualValue != expectedValue {
		t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
	}
}

func TestGenParseAndStoreCSVSeriesPostForError(t *testing.T) {
	out := SeriesPost{}
	_, actualErr := out.GenParseAndStoreCSV(nil, &RepositoryMock{}, ParserTestingFunc)
	if actualErr == nil {
		t.Errorf("Calling SeriesPost.GenParseAndStoreCSV with a nil file pointer should have returned an error\n")
	}
}
