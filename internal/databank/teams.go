package databank

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/rippinrobr/baseball-stats-db/internal/platform/parsers/csv"

	"github.com/rippinrobr/baseball-stats-db/internal/platform/db"
)

// Teams is a model that maps the CSV to a DB Table
type Teams struct {
	Yearid         int     `json:"yearID"  csv:"yearID"  db:"yearID"  bson:"yearID"`
	Lgid           string  `json:"lgID"  csv:"lgID"  db:"lgID,omitempty"  bson:"lgID"`
	Teamid         string  `json:"teamID"  csv:"teamID"  db:"teamID,omitempty"  bson:"teamID"`
	Franchid       string  `json:"franchID"  csv:"franchID"  db:"franchID,omitempty"  bson:"franchID"`
	Divid          string  `json:"divID"  csv:"divID"  db:"divID,omitempty"  bson:"divID"`
	Rank           int     `json:"rank"  csv:"Rank"  db:"Rank"  bson:"Rank"`
	G              int     `json:"g"  csv:"G"  db:"G"  bson:"G"`
	Ghome          int     `json:"ghome"  csv:"Ghome"  db:"Ghome"  bson:"Ghome"`
	W              int     `json:"w"  csv:"W"  db:"W"  bson:"W"`
	L              int     `json:"l"  csv:"L"  db:"L"  bson:"L"`
	Divwin         string  `json:"divWin"  csv:"DivWin"  db:"DivWin,omitempty"  bson:"DivWin"`
	Wcwin          string  `json:"wCWin"  csv:"WCWin"  db:"WCWin,omitempty"  bson:"WCWin"`
	Lgwin          string  `json:"lgWin"  csv:"LgWin"  db:"LgWin,omitempty"  bson:"LgWin"`
	Wswin          string  `json:"wSWin"  csv:"WSWin"  db:"WSWin,omitempty"  bson:"WSWin"`
	R              int     `json:"r"  csv:"R"  db:"R"  bson:"R"`
	Ab             int     `json:"aB"  csv:"AB"  db:"AB"  bson:"AB"`
	H              int     `json:"h"  csv:"H"  db:"H"  bson:"H"`
	Doubles        int     `json:"doubles"  csv:"2B"  db:"doubles"  bson:"doubles"`
	Triples        int     `json:"triples"  csv:"3B"  db:"triples"  bson:"triples"`
	Hr             int     `json:"hR"  csv:"HR"  db:"HR"  bson:"HR"`
	Bb             float32 `json:"bB"  csv:"BB"  db:"BB"  bson:"BB"`
	So             int     `json:"sO"  csv:"SO"  db:"SO"  bson:"SO"`
	Sb             float32 `json:"sB"  csv:"SB"  db:"SB"  bson:"SB"`
	Cs             float32 `json:"cS"  csv:"CS"  db:"CS"  bson:"CS"`
	Hbp            float32 `json:"hBP"  csv:"HBP"  db:"HBP"  bson:"HBP"`
	Sf             int     `json:"sF"  csv:"SF"  db:"SF"  bson:"SF"`
	Ra             int     `json:"rA"  csv:"RA"  db:"RA"  bson:"RA"`
	Er             int     `json:"eR"  csv:"ER"  db:"ER"  bson:"ER"`
	Era            float64 `json:"eRA"  csv:"ERA"  db:"ERA"  bson:"ERA"`
	Cg             int     `json:"cG"  csv:"CG"  db:"CG"  bson:"CG"`
	Sho            int     `json:"sHO"  csv:"SHO"  db:"SHO"  bson:"SHO"`
	Sv             int     `json:"sV"  csv:"SV"  db:"SV"  bson:"SV"`
	Ipouts         int     `json:"iPouts"  csv:"IPouts"  db:"IPouts"  bson:"IPouts"`
	Ha             int     `json:"hA"  csv:"HA"  db:"HA"  bson:"HA"`
	Hra            int     `json:"hRA"  csv:"HRA"  db:"HRA"  bson:"HRA"`
	Bba            int     `json:"bBA"  csv:"BBA"  db:"BBA"  bson:"BBA"`
	Soa            int     `json:"sOA"  csv:"SOA"  db:"SOA"  bson:"SOA"`
	E              int     `json:"e"  csv:"E"  db:"E"  bson:"E"`
	Dp             int     `json:"dP"  csv:"DP"  db:"DP"  bson:"DP"`
	Fp             float64 `json:"fP"  csv:"FP"  db:"FP"  bson:"FP"`
	Name           string  `json:"name"  csv:"name"  db:"name,omitempty"  bson:"name"`
	Park           string  `json:"park"  csv:"park"  db:"park,omitempty"  bson:"park"`
	Attendance     int     `json:"attendance"  csv:"attendance"  db:"attendance"  bson:"attendance"`
	Bpf            int     `json:"bPF"  csv:"BPF"  db:"BPF"  bson:"BPF"`
	Ppf            int     `json:"pPF"  csv:"PPF"  db:"PPF"  bson:"PPF"`
	Teamidbr       string  `json:"teamIDBR"  csv:"teamIDBR"  db:"teamIDBR,omitempty"  bson:"teamIDBR"`
	Teamidlahman45 string  `json:"teamIDlahman45"  csv:"teamIDlahman45"  db:"teamIDlahman45,omitempty"  bson:"teamIDlahman45"`
	Teamidretro    string  `json:"teamIDretro"  csv:"teamIDretro"  db:"teamIDretro,omitempty"  bson:"teamIDretro"`
	inputDir       string
}

// GetTableName returns the name of the table that the data will be stored in
func (m *Teams) GetTableName() string {
	return "teams"
}

// GetFileName returns the name of the source file the model was created from
func (m *Teams) GetFileName() string {
	return "Teams.csv"
}

// GetFilePath returns the path of the source file
func (m *Teams) GetFilePath() string {
	return filepath.Join(m.inputDir, "Teams.csv")
}

// SetInputDirectory sets the input directory's path so it can be used to create the full path to the file
func (m *Teams) SetInputDirectory(inputDir string) {
	m.inputDir = inputDir
}

// GenParseAndStoreCSV returns a function that will parse the source file,\n//create a slice with an object per line and store the data in the db
func (m *Teams) GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) (ParseAndStoreCSVFunc, error) {
	if f == nil {
		return func() error { return nil }, errors.New("nil File")
	}

	return func() error {
		rows := make([]*Teams, 0)
		numErrors := 0
		err := pfunc(f, &rows)
		if err == nil {
			numrows := len(rows)
			if numrows > 0 {
				log.Println("Teams ==> Truncating")
				terr := repo.Truncate(m.GetTableName())
				if terr != nil {
					log.Println("truncate err:", terr.Error())
				}

				log.Printf("Teams ==> Inserting %d Records\n", numrows)
				for _, r := range rows {
					ierr := repo.Insert(m.GetTableName(), r)
					if ierr != nil {
						log.Printf("Insert error: %s\n", ierr.Error())
						numErrors++
					}
				}
			}
			log.Printf("Teams ==> %d records created\n", numrows-numErrors)
		}

		return err
	}, nil
}
