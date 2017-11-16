package models


import (
  "os"
  "log"

  "github.com/rippinrobr/baseball-databank-tools/pkg/parsers/csv"

  "github.com/rippinrobr/baseball-databank-tools/pkg/db"

)

// Teams is a model that maps the CSV to a DB Table
type Teams struct {
   Yearid   int `json:"yearid"  csv:"yearID"  db:"yearID"`
   Lgid   string `json:"lgid"  csv:"lgID"  db:"lgID"`
   Teamid   string `json:"teamid"  csv:"teamID"  db:"teamID"`
   Franchid   string `json:"franchid"  csv:"franchID"  db:"franchID"`
   Divid   string `json:"divid"  csv:"divID"  db:"divID"`
   Rank   int `json:"rank"  csv:"Rank"  db:"Rank"`
   G   int `json:"g"  csv:"G"  db:"G"`
   Ghome   string `json:"ghome"  csv:"Ghome"  db:"Ghome"`
   W   int `json:"w"  csv:"W"  db:"W"`
   L   int `json:"l"  csv:"L"  db:"L"`
   Divwin   string `json:"divwin"  csv:"DivWin"  db:"DivWin"`
   Wcwin   string `json:"wcwin"  csv:"WCWin"  db:"WCWin"`
   Lgwin   string `json:"lgwin"  csv:"LgWin"  db:"LgWin"`
   Wswin   string `json:"wswin"  csv:"WSWin"  db:"WSWin"`
   R   int `json:"r"  csv:"R"  db:"R"`
   Ab   int `json:"ab"  csv:"AB"  db:"AB"`
   H   int `json:"h"  csv:"H"  db:"H"`
   Doubles   int `json:"doubles"  csv:"2B"  db:"2B"`
   Triples   int `json:"triples"  csv:"3B"  db:"3B"`
   Hr   int `json:"hr"  csv:"HR"  db:"HR"`
   Bb   int `json:"bb"  csv:"BB"  db:"BB"`
   So   string `json:"so"  csv:"SO"  db:"SO"`
   Sb   string `json:"sb"  csv:"SB"  db:"SB"`
   Cs   string `json:"cs"  csv:"CS"  db:"CS"`
   Hbp   string `json:"hbp"  csv:"HBP"  db:"HBP"`
   Sf   string `json:"sf"  csv:"SF"  db:"SF"`
   Ra   int `json:"ra"  csv:"RA"  db:"RA"`
   Er   int `json:"er"  csv:"ER"  db:"ER"`
   Era   string `json:"era"  csv:"ERA"  db:"ERA"`
   Cg   int `json:"cg"  csv:"CG"  db:"CG"`
   Sho   int `json:"sho"  csv:"SHO"  db:"SHO"`
   Sv   int `json:"sv"  csv:"SV"  db:"SV"`
   Ipouts   int `json:"ipouts"  csv:"IPouts"  db:"IPouts"`
   Ha   int `json:"ha"  csv:"HA"  db:"HA"`
   Hra   int `json:"hra"  csv:"HRA"  db:"HRA"`
   Bba   int `json:"bba"  csv:"BBA"  db:"BBA"`
   Soa   int `json:"soa"  csv:"SOA"  db:"SOA"`
   E   int `json:"e"  csv:"E"  db:"E"`
   Dp   string `json:"dp"  csv:"DP"  db:"DP"`
   Fp   string `json:"fp"  csv:"FP"  db:"FP"`
   Name   string `json:"name"  csv:"name"  db:"name"`
   Park   string `json:"park"  csv:"park"  db:"park"`
   Attendance   string `json:"attendance"  csv:"attendance"  db:"attendance"`
   Bpf   int `json:"bpf"  csv:"BPF"  db:"BPF"`
   Ppf   int `json:"ppf"  csv:"PPF"  db:"PPF"`
   Teamidbr   string `json:"teamidbr"  csv:"teamIDBR"  db:"teamIDBR"`
   Teamidlahman45   string `json:"teamidlahman45"  csv:"teamIDlahman45"  db:"teamIDlahman45"`
   Teamidretro   string `json:"teamidretro"  csv:"teamIDretro"  db:"teamIDretro"`
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
  return "/Users/robertrowe/src/baseballdatabank/core/Teams.csv"
}

// GenParseAndStoreCSV returns a function that will parse the source file,\n//create a slice with an object per line and store the data in the db
func (m *Teams) GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) ParseAndStoreCSVFunc {
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
   }
}
