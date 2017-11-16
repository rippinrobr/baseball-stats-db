package models


import (
  "os"
  "log"

  "github.com/rippinrobr/baseball-databank-tools/pkg/parsers/csv"

  "github.com/rippinrobr/baseball-databank-tools/pkg/db"

)

// FieldingOFsplit is a model that maps the CSV to a DB Table
type FieldingOFsplit struct {
   Playerid   string `json:"playerid"  csv:"playerID"  db:"playerID"`
   Yearid   int `json:"yearid"  csv:"yearID"  db:"yearID"`
   Stint   int `json:"stint"  csv:"stint"  db:"stint"`
   Teamid   string `json:"teamid"  csv:"teamID"  db:"teamID"`
   Lgid   string `json:"lgid"  csv:"lgID"  db:"lgID"`
   Pos   string `json:"pos"  csv:"POS"  db:"POS"`
   G   int `json:"g"  csv:"G"  db:"G"`
   Gs   string `json:"gs"  csv:"GS"  db:"GS"`
   Innouts   string `json:"innouts"  csv:"InnOuts"  db:"InnOuts"`
   Po   string `json:"po"  csv:"PO"  db:"PO"`
   A   string `json:"a"  csv:"A"  db:"A"`
   E   string `json:"e"  csv:"E"  db:"E"`
   Dp   string `json:"dp"  csv:"DP"  db:"DP"`
   Pb   string `json:"pb"  csv:"PB"  db:"PB"`
   Wp   string `json:"wp"  csv:"WP"  db:"WP"`
   Sb   string `json:"sb"  csv:"SB"  db:"SB"`
   Cs   string `json:"cs"  csv:"CS"  db:"CS"`
   Zr   string `json:"zr"  csv:"ZR"  db:"ZR"`
}

// GetTableName returns the name of the table that the data will be stored in
func (m *FieldingOFsplit) GetTableName() string {
  return "fieldingofsplit"
}

// GetFileName returns the name of the source file the model was created from
func (m *FieldingOFsplit) GetFileName() string {
  return "FieldingOFsplit.csv"
}

// GetFilePath returns the path of the source file
func (m *FieldingOFsplit) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/FieldingOFsplit.csv"
}

// GenParseAndStoreCSV returns a function that will parse the source file,\n//create a slice with an object per line and store the data in the db
func (m *FieldingOFsplit) GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) ParseAndStoreCSVFunc {
  return func() error {
    rows := make([]*FieldingOFsplit, 0)
    numErrors := 0
    err := pfunc(f, &rows)
    if err == nil {
       numrows := len(rows)
       if numrows > 0 {
         log.Println("FieldingOFsplit ==> Truncating")
         terr := repo.Truncate(m.GetTableName())
         if terr != nil {
            log.Println("truncate err:", terr.Error())
         }

         log.Printf("FieldingOFsplit ==> Inserting %d Records\n", numrows)
         for _, r := range rows {
           ierr := repo.Insert(m.GetTableName(), r)
           if ierr != nil {
             log.Printf("Insert error: %s\n", ierr.Error())
             numErrors++
           }
         }
       }
       log.Printf("FieldingOFsplit ==> %d records created\n", numrows-numErrors)
    }

    return err
   }
}
