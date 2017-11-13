package models

// TableObject is an interface all database related
// models must implement
type TableObject interface {
	GetTableName() string
	GetFileName() string
	GetFilePath() string
}

// GetTableObjects returns an array of pointers to
// the TableObjects for each table in the
// Baseball Databank Database
func GetTableObjects() []TableObject {
	return []TableObject{
		&AllstarFull{},
		&Appearances{},
		&AwardsManagers{},
	}
}
