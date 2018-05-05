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

// LookupTableObject returns the TableObject that relates to
// the file being processed
func LookupTableObject(fileName string) TableObject {
	switch fileName {
	case "People.csv":
		return &People{}
	case "AllstarFull.csv":
		return &AllstarFull{}
	case "Appearances.csv":
		return &Appearances{}
	case "AwardsManagers.csv":
		return &AwardsManagers{}
	case "AwardsPlayers.csv":
		return &AwardsPlayers{}
	case "AwardsShareManagers.csv":
		return &AwardsShareManagers{}
	case "AwardsSharePlayers.csv":
		return &AwardsSharePlayers{}
	case "Batting.csv":
		return &Batting{}
	case "BattingPost.csv":
		return &BattingPost{}
	case "CollegePlaying.csv":
		return &CollegePlaying{}
	case "Fielding.csv":
		return &Fielding{}
	case "FieldingOF.csv":
		return &FieldingOF{}
	case "FieldingPost.csv":
		return &FieldingPost{}
	case "FieldingOFsplit.csv":
		return &FieldingOFsplit{}
	case "HallOfFame.csv":
		return &HallOfFame{}
	case "HomeGames.csv":
		return &HomeGames{}
	case "Managers.csv":
		return &Managers{}
	case "ManagersHalf.csv":
		return &ManagersHalf{}
	case "Parks.csv":
		return &Parks{}
	case "Pitching.csv":
		return &Pitching{}
	case "PitchingPost.csv":
		return &PitchingPost{}
	case "Salaries.csv":
		return &Salaries{}
	case "Schools.csv":
		return &Schools{}
	case "SeriesPost.csv":
		return &SeriesPost{}
	case "Teams.csv":
		return &Teams{}
	case "TeamsFranchises.csv":
		return &TeamsFranchises{}
	case "TeamsHalf.csv":
		return &TeamsHalf{}
	}

	return nil //&People{}
}
