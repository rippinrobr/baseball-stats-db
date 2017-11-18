package models


import (
  "os"
  "log"
  "errors"

  "github.com/rippinrobr/baseball-databank-tools/pkg/parsers/csv"

  "github.com/rippinrobr/baseball-databank-tools/pkg/db"

)

// FieldingPost is a model that maps the CSV to a DB Table
type FieldingPost struct {
   Playerid   string `json:"playerid"  csv:"playerID"  db:"playerID"`
   Yearid   int `json:"yearid"  csv:"yearID"  db:"yearID"`
   Teamid   string `json:"teamid"  csv:"teamID"  db:"teamID"`
   Lgid   string `json:"lgid"  csv:"lgID"  db:"lgID"`
   Round   string `json:"round"  csv:"round"  db:"round"`
   Pos   string `json:"pos"  csv:"POS"  db:"POS"`
   G   int `json:"g"  csv:"G"  db:"G"`
   Gs   int `json:"gs"  csv:"GS"  db:"GS"`
   Innouts   int `json:"innouts"  csv:"InnOuts"  db:"InnOuts"`
   Po   int `json:"po"  csv:"PO"  db:"PO"`
   A   int `json:"a"  csv:"A"  db:"A"`
   E   int `json:"e"  csv:"E"  db:"E"`
   Dp   int `json:"dp"  csv:"DP"  db:"DP"`
   Tp   int `json:"tp"  csv:"TP"  db:"TP"`
   Pb   string `json:"pb"  csv:"PB"  db:"PB"`
   Sb   string `json:"sb"  csv:"SB"  db:"SB"`
   Cs   string `json:"cs"  csv:"CS"  db:"CS"`
}

// GetTableName returns the name of the table that the data will be stored in
func (m *FieldingPost) GetTableName() string {
  return "fieldingpost"
}

// GetFileName returns the name of the source file the model was created from
func (m *FieldingPost) GetFileName() string {
  return "FieldingPost.csv"
}

// GetFilePath returns the path of the source file
func (m *FieldingPost) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/FieldingPost.csv"
}

// GenParseAndStoreCSV returns a function that will parse the source file,\n//create a slice with an object per line and store the data in the db
func (m *FieldingPost) GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) (ParseAndStoreCSVFunc, error) {
  if f == nil {
    return func() error{return nil}, errors.New("nil File")
  }

  return func() error {
    rows := make([]*FieldingPost, 0)
    numErrors := 0
    err := pfunc(f, &rows)
    if err == nil {
       numrows := len(rows)
       if numrows > 0 {
         log.Println("FieldingPost ==> Truncating")
         terr := repo.Truncate(m.GetTableName())
         if terr != nil {
            log.Println("truncate err:", terr.Error())
         }

         log.Printf("FieldingPost ==> Inserting %d Records\n", numrows)
         for _, r := range rows {
           ierr := repo.Insert(m.GetTableName(), r)
           if ierr != nil {
             log.Printf("Insert error: %s\n", ierr.Error())
             numErrors++
           }
         }
       }
       log.Printf("FieldingPost ==> %d records created\n", numrows-numErrors)
    }

    return err
   }, nil
}
