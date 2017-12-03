package databank

import (
	"os"

	"github.com/rippinrobr/baseball-stats-db/internal/platform/db"
	"github.com/rippinrobr/baseball-stats-db/internal/platform/parsers/csv"
)

// ParseAndStoreCSVFunc are functions created by the models
// to parse and store the data from the CSV files.  These
// are needed because the gocsv library didn't like the
// []interface{} I was passing in.
type ParseAndStoreCSVFunc func() error

// TableObject is an interface all database related
// models must implement
type TableObject interface {
	GetTableName() string
	GetFileName() string
	GetFilePath() string
	GenParseAndStoreCSV(*os.File, db.Repository, csv.ParserFunc) (ParseAndStoreCSVFunc, error)
	SetInputDirectory(string)
}

// GetTableObjects returns an array of pointers to
// the TableObjects for each table in the
// Baseball Databank Database
func GetTableObjects() []TableObject {
	return []TableObject{
		&People{},
		&AllstarFull{},
		&Appearances{},
		&AwardsManagers{},
		&AwardsPlayers{},
		&AwardsShareManagers{},
		&AwardsSharePlayers{},
		&Batting{},
		&BattingPost{},
		&CollegePlaying{},
		&Fielding{},
		&FieldingOF{},
		&FieldingPost{},
		&FieldingOFsplit{},
		&HallOfFame{},
		&HomeGames{},
		&Managers{},
		&ManagersHalf{},
		&Parks{},
		&Pitching{},
		&PitchingPost{},
		&Salaries{},
		&Schools{},
		&SeriesPost{},
		&Teams{},
		&TeamsFranchises{},
		&TeamsHalf{},
	}
}
