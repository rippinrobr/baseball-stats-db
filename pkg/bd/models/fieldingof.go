package models


import (
  "os"
  "log"

  "github.com/rippinrobr/baseball-databank-tools/pkg/parsers/csv"

  "github.com/rippinrobr/baseball-databank-tools/pkg/db"

)

// FieldingOF is a model that maps the CSV to a DB Table
type FieldingOF struct {
   Playerid   string `json:"playerid"  csv:"playerID"  db:"playerID"`
   Yearid   int `json:"yearid"  csv:"yearID"  db:"yearID"`
   Stint   int `json:"stint"  csv:"stint"  db:"stint"`
   Glf   string `json:"glf"  csv:"Glf"  db:"Glf"`
   Gcf   string `json:"gcf"  csv:"Gcf"  db:"Gcf"`
   Grf   string `json:"grf"  csv:"Grf"  db:"Grf"`
}

// GetTableName returns the name of the table that the data will be stored in
func (m *FieldingOF) GetTableName() string {
  return "fieldingof"
}

// GetFileName returns the name of the source file the model was created from
func (m *FieldingOF) GetFileName() string {
  return "FieldingOF.csv"
}

// GetFilePath returns the path of the source file
func (m *FieldingOF) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/FieldingOF.csv"
}

// GenParseAndStoreCSV returns a function that will parse the source file,\n//create a slice with an object per line and store the data in the db
func (m *FieldingOF) GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) ParseAndStoreCSVFunc {
  return func() error {
    rows := make([]*FieldingOF, 0)
    numErrors := 0
    err := pfunc(f, &rows)
    if err == nil {
       numrows := len(rows)
       if numrows > 0 {
         log.Println("FieldingOF ==> Truncating")
         terr := repo.Truncate(m.GetTableName())
         if terr != nil {
            log.Println("truncate err:", terr.Error())
         }

         log.Printf("FieldingOF ==> Inserting %d Records\n", numrows)
         for _, r := range rows {
           ierr := repo.Insert(m.GetTableName(), r)
           if ierr != nil {
             log.Printf("Insert error: %s\n", ierr.Error())
             numErrors++
           }
         }
       }
       log.Printf("FieldingOF ==> %d records created\n", numrows-numErrors)
    }

    return err
   }
}
