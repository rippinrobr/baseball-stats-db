package models


import (
  "os"
  "log"
  "errors"

  "github.com/rippinrobr/baseball-databank-db/pkg/parsers/csv"

  "github.com/rippinrobr/baseball-databank-db/pkg/db"

)

// Teams is a model that maps the CSV to a DB Table
type Teams struct {
   Yearid   int `json:"yearID"  csv:"yearID"  db:"yearID"`
   Lgid   string `json:"lgID"  csv:"lgID"  db:"lgID,omitempty"`
   Teamid   string `json:"teamID"  csv:"teamID"  db:"teamID,omitempty"`
   Franchid   string `json:"franchID"  csv:"franchID"  db:"franchID,omitempty"`
   Divid   string `json:"divID"  csv:"divID"  db:"divID,omitempty"`
   Rank   int `json:"rank"  csv:"Rank"  db:"Rank"`
   G   int `json:"g"  csv:"G"  db:"G"`
   Ghome   int `json:"ghome"  csv:"Ghome"  db:"Ghome"`
   W   int `json:"w"  csv:"W"  db:"W"`
   L   int `json:"l"  csv:"L"  db:"L"`
   Divwin   string `json:"divWin"  csv:"DivWin"  db:"DivWin,omitempty"`
   Wcwin   string `json:"wCWin"  csv:"WCWin"  db:"WCWin,omitempty"`
   Lgwin   string `json:"lgWin"  csv:"LgWin"  db:"LgWin,omitempty"`
   Wswin   string `json:"wSWin"  csv:"WSWin"  db:"WSWin,omitempty"`
   R   int `json:"r"  csv:"R"  db:"R"`
   Ab   int `json:"aB"  csv:"AB"  db:"AB"`
   H   int `json:"h"  csv:"H"  db:"H"`
   Doubles   int `json:"doubles"  csv:"2B"  db:"doubles"`
   Triples   int `json:"triples"  csv:"3B"  db:"triples"`
   Hr   int `json:"hR"  csv:"HR"  db:"HR"`
   Bb   int `json:"bB"  csv:"BB"  db:"BB"`
   So   int `json:"sO"  csv:"SO"  db:"SO"`
   Sb   int `json:"sB"  csv:"SB"  db:"SB"`
   Cs   int `json:"cS"  csv:"CS"  db:"CS"`
   Hbp   int `json:"hBP"  csv:"HBP"  db:"HBP"`
   Sf   int `json:"sF"  csv:"SF"  db:"SF"`
   Ra   int `json:"rA"  csv:"RA"  db:"RA"`
   Er   int `json:"eR"  csv:"ER"  db:"ER"`
   Era   float64 `json:"eRA"  csv:"ERA"  db:"ERA"`
   Cg   int `json:"cG"  csv:"CG"  db:"CG"`
   Sho   int `json:"sHO"  csv:"SHO"  db:"SHO"`
   Sv   int `json:"sV"  csv:"SV"  db:"SV"`
   Ipouts   int `json:"iPouts"  csv:"IPouts"  db:"IPouts"`
   Ha   int `json:"hA"  csv:"HA"  db:"HA"`
   Hra   int `json:"hRA"  csv:"HRA"  db:"HRA"`
   Bba   int `json:"bBA"  csv:"BBA"  db:"BBA"`
   Soa   int `json:"sOA"  csv:"SOA"  db:"SOA"`
   E   int `json:"e"  csv:"E"  db:"E"`
   Dp   int `json:"dP"  csv:"DP"  db:"DP"`
   Fp   float64 `json:"fP"  csv:"FP"  db:"FP"`
   Name   string `json:"name"  csv:"name"  db:"name,omitempty"`
   Park   string `json:"park"  csv:"park"  db:"park,omitempty"`
   Attendance   int `json:"attendance"  csv:"attendance"  db:"attendance"`
   Bpf   int `json:"bPF"  csv:"BPF"  db:"BPF"`
   Ppf   int `json:"pPF"  csv:"PPF"  db:"PPF"`
   Teamidbr   string `json:"teamIDBR"  csv:"teamIDBR"  db:"teamIDBR,omitempty"`
   Teamidlahman45   string `json:"teamIDlahman45"  csv:"teamIDlahman45"  db:"teamIDlahman45,omitempty"`
   Teamidretro   string `json:"teamIDretro"  csv:"teamIDretro"  db:"teamIDretro,omitempty"`
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
func (m *Teams) GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) (ParseAndStoreCSVFunc, error) {
  if f == nil {
    return func() error{return nil}, errors.New("nil File")
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
