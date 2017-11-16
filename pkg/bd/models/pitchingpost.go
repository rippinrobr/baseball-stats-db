package models


import (
  "os"
  "log"

  "github.com/rippinrobr/baseball-databank-tools/pkg/parsers/csv"

  "github.com/rippinrobr/baseball-databank-tools/pkg/db"

)

// PitchingPost is a model that maps the CSV to a DB Table
type PitchingPost struct {
   Playerid   string `json:"playerid"  csv:"playerID"  db:"playerID"`
   Yearid   int `json:"yearid"  csv:"yearID"  db:"yearID"`
   Round   string `json:"round"  csv:"round"  db:"round"`
   Teamid   string `json:"teamid"  csv:"teamID"  db:"teamID"`
   Lgid   string `json:"lgid"  csv:"lgID"  db:"lgID"`
   W   int `json:"w"  csv:"W"  db:"W"`
   L   int `json:"l"  csv:"L"  db:"L"`
   G   int `json:"g"  csv:"G"  db:"G"`
   Gs   int `json:"gs"  csv:"GS"  db:"GS"`
   Cg   int `json:"cg"  csv:"CG"  db:"CG"`
   Sho   int `json:"sho"  csv:"SHO"  db:"SHO"`
   Sv   int `json:"sv"  csv:"SV"  db:"SV"`
   Ipouts   int `json:"ipouts"  csv:"IPouts"  db:"IPouts"`
   H   int `json:"h"  csv:"H"  db:"H"`
   Er   int `json:"er"  csv:"ER"  db:"ER"`
   Hr   int `json:"hr"  csv:"HR"  db:"HR"`
   Bb   int `json:"bb"  csv:"BB"  db:"BB"`
   So   int `json:"so"  csv:"SO"  db:"SO"`
   Baopp   string `json:"baopp"  csv:"BAOpp"  db:"BAOpp"`
   Era   string `json:"era"  csv:"ERA"  db:"ERA"`
   Ibb   string `json:"ibb"  csv:"IBB"  db:"IBB"`
   Wp   string `json:"wp"  csv:"WP"  db:"WP"`
   Hbp   string `json:"hbp"  csv:"HBP"  db:"HBP"`
   Bk   string `json:"bk"  csv:"BK"  db:"BK"`
   Bfp   string `json:"bfp"  csv:"BFP"  db:"BFP"`
   Gf   int `json:"gf"  csv:"GF"  db:"GF"`
   R   int `json:"r"  csv:"R"  db:"R"`
   Sh   string `json:"sh"  csv:"SH"  db:"SH"`
   Sf   string `json:"sf"  csv:"SF"  db:"SF"`
   Gidp   string `json:"gidp"  csv:"GIDP"  db:"GIDP"`
}

// GetTableName returns the name of the table that the data will be stored in
func (m *PitchingPost) GetTableName() string {
  return "pitchingpost"
}

// GetFileName returns the name of the source file the model was created from
func (m *PitchingPost) GetFileName() string {
  return "PitchingPost.csv"
}

// GetFilePath returns the path of the source file
func (m *PitchingPost) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/PitchingPost.csv"
}

// GenParseAndStoreCSV returns a function that will parse the source file,\n//create a slice with an object per line and store the data in the db
func (m *PitchingPost) GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) ParseAndStoreCSVFunc {
  return func() error {
    rows := make([]*PitchingPost, 0)
    numErrors := 0
    err := pfunc(f, &rows)
    if err == nil {
       numrows := len(rows)
       if numrows > 0 {
         log.Println("PitchingPost ==> Truncating")
         terr := repo.Truncate(m.GetTableName())
         if terr != nil {
            log.Println("truncate err:", terr.Error())
         }

         log.Printf("PitchingPost ==> Inserting %d Records\n", numrows)
         for _, r := range rows {
           ierr := repo.Insert(m.GetTableName(), r)
           if ierr != nil {
             log.Printf("Insert error: %s\n", ierr.Error())
             numErrors++
           }
         }
       }
       log.Printf("PitchingPost ==> %d records created\n", numrows-numErrors)
    }

    return err
   }
}
