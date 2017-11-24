package models


import (
  "os"
  "log"
  "errors"
  "path/filepath"

  "github.com/rippinrobr/baseball-databank-db/pkg/parsers/csv"

  "github.com/rippinrobr/baseball-databank-db/pkg/db"

)

// PitchingPost is a model that maps the CSV to a DB Table
type PitchingPost struct {
   Playerid   string `json:"playerID"  csv:"playerID"  db:"playerID,omitempty"`
   Yearid   int `json:"yearID"  csv:"yearID"  db:"yearID"`
   Round   string `json:"round"  csv:"round"  db:"round,omitempty"`
   Teamid   string `json:"teamID"  csv:"teamID"  db:"teamID,omitempty"`
   Lgid   string `json:"lgID"  csv:"lgID"  db:"lgID,omitempty"`
   W   int `json:"w"  csv:"W"  db:"W"`
   L   int `json:"l"  csv:"L"  db:"L"`
   G   int `json:"g"  csv:"G"  db:"G"`
   Gs   int `json:"gS"  csv:"GS"  db:"GS"`
   Cg   int `json:"cG"  csv:"CG"  db:"CG"`
   Sho   int `json:"sHO"  csv:"SHO"  db:"SHO"`
   Sv   int `json:"sV"  csv:"SV"  db:"SV"`
   Ipouts   int `json:"iPouts"  csv:"IPouts"  db:"IPouts"`
   H   int `json:"h"  csv:"H"  db:"H"`
   Er   int `json:"eR"  csv:"ER"  db:"ER"`
   Hr   int `json:"hR"  csv:"HR"  db:"HR"`
   Bb   int `json:"bB"  csv:"BB"  db:"BB"`
   So   int `json:"sO"  csv:"SO"  db:"SO"`
   Baopp   float64 `json:"bAOpp"  csv:"BAOpp"  db:"BAOpp"`
   Era   float64 `json:"eRA"  csv:"ERA"  db:"ERA"`
   Ibb   int `json:"iBB"  csv:"IBB"  db:"IBB"`
   Wp   int `json:"wP"  csv:"WP"  db:"WP"`
   Hbp   int `json:"hBP"  csv:"HBP"  db:"HBP"`
   Bk   int `json:"bK"  csv:"BK"  db:"BK"`
   Bfp   int `json:"bFP"  csv:"BFP"  db:"BFP"`
   Gf   int `json:"gF"  csv:"GF"  db:"GF"`
   R   int `json:"r"  csv:"R"  db:"R"`
   Sh   int `json:"sH"  csv:"SH"  db:"SH"`
   Sf   int `json:"sF"  csv:"SF"  db:"SF"`
   Gidp   int `json:"gIDP"  csv:"GIDP"  db:"GIDP"`
  inputDir string
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
  return filepath.Join(m.inputDir, "PitchingPost.csv")
}

// SetInputDirectory sets the input directory's path so it can be used to create the full path to the file
func (m *PitchingPost) SetInputDirectory(inputDir string) {
  m.inputDir=inputDir
}

// GenParseAndStoreCSV returns a function that will parse the source file,\n//create a slice with an object per line and store the data in the db
func (m *PitchingPost) GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) (ParseAndStoreCSVFunc, error) {
  if f == nil {
    return func() error{return nil}, errors.New("nil File")
  }

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
   }, nil
}
