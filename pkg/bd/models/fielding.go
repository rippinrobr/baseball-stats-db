package models


import (
  "os"
  "log"
  "errors"

  "github.com/rippinrobr/baseball-databank-tools/pkg/parsers/csv"

  "github.com/rippinrobr/baseball-databank-tools/pkg/db"

)

// Fielding is a model that maps the CSV to a DB Table
type Fielding struct {
   Playerid   string `json:"playerid"  csv:"playerID"  db:"playerID"`
   Yearid   int `json:"yearid"  csv:"yearID"  db:"yearID"`
   Stint   int `json:"stint"  csv:"stint"  db:"stint"`
   Teamid   string `json:"teamid"  csv:"teamID"  db:"teamID"`
   Lgid   string `json:"lgid"  csv:"lgID"  db:"lgID"`
   Pos   string `json:"pos"  csv:"POS"  db:"POS"`
   G   int `json:"g"  csv:"G"  db:"G"`
   Gs   string `json:"gs"  csv:"GS"  db:"GS"`
   Innouts   string `json:"innouts"  csv:"InnOuts"  db:"InnOuts"`
   Po   int `json:"po"  csv:"PO"  db:"PO"`
   A   int `json:"a"  csv:"A"  db:"A"`
   E   string `json:"e"  csv:"E"  db:"E"`
   Dp   int `json:"dp"  csv:"DP"  db:"DP"`
   Pb   string `json:"pb"  csv:"PB"  db:"PB"`
   Wp   string `json:"wp"  csv:"WP"  db:"WP"`
   Sb   string `json:"sb"  csv:"SB"  db:"SB"`
   Cs   string `json:"cs"  csv:"CS"  db:"CS"`
   Zr   string `json:"zr"  csv:"ZR"  db:"ZR"`
}

// GetTableName returns the name of the table that the data will be stored in
func (m *Fielding) GetTableName() string {
  return "fielding"
}

// GetFileName returns the name of the source file the model was created from
func (m *Fielding) GetFileName() string {
  return "Fielding.csv"
}

// GetFilePath returns the path of the source file
func (m *Fielding) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/Fielding.csv"
}

// GenParseAndStoreCSV returns a function that will parse the source file,\n//create a slice with an object per line and store the data in the db
func (m *Fielding) GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) (ParseAndStoreCSVFunc, error) {
  if f == nil {
    return func() error{return nil}, errors.New("nil File")
  }

  return func() error {
    rows := make([]*Fielding, 0)
    numErrors := 0
    err := pfunc(f, &rows)
    if err == nil {
       numrows := len(rows)
       if numrows > 0 {
         log.Println("Fielding ==> Truncating")
         terr := repo.Truncate(m.GetTableName())
         if terr != nil {
            log.Println("truncate err:", terr.Error())
         }

         log.Printf("Fielding ==> Inserting %d Records\n", numrows)
         for _, r := range rows {
           ierr := repo.Insert(m.GetTableName(), r)
           if ierr != nil {
             log.Printf("Insert error: %s\n", ierr.Error())
             numErrors++
           }
         }
       }
       log.Printf("Fielding ==> %d records created\n", numrows-numErrors)
    }

    return err
   }, nil
}
