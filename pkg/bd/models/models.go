package models

// TableObject is an interface all database related
// models must implement
type TableObject interface {
	GetTableName() string
	GetFileName() string
	GetFilePath() string
}
