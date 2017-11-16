package models


import (
  "os"
  "log"

  "github.com/rippinrobr/baseball-databank-tools/pkg/parsers/csv"

  "github.com/rippinrobr/baseball-databank-tools/pkg/db"

)

// TeamsHalf is a model that maps the CSV to a DB Table
type TeamsHalf struct {
   Yearid   int `json:"yearid"  csv:"yearID"  db:"yearID"`
   Lgid   string `json:"lgid"  csv:"lgID"  db:"lgID"`
   Teamid   string `json:"teamid"  csv:"teamID"  db:"teamID"`
   Half   int `json:"half"  csv:"Half"  db:"Half"`
   Divid   string `json:"divid"  csv:"divID"  db:"divID"`
   Divwin   string `json:"divwin"  csv:"DivWin"  db:"DivWin"`
   Rank   int `json:"rank"  csv:"Rank"  db:"Rank"`
   G   int `json:"g"  csv:"G"  db:"G"`
   W   int `json:"w"  csv:"W"  db:"W"`
   L   int `json:"l"  csv:"L"  db:"L"`
}

// GetTableName returns the name of the table that the data will be stored in
func (m *TeamsHalf) GetTableName() string {
  return "teamshalf"
}

// GetFileName returns the name of the source file the model was created from
func (m *TeamsHalf) GetFileName() string {
  return "TeamsHalf.csv"
}

// GetFilePath returns the path of the source file
func (m *TeamsHalf) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/TeamsHalf.csv"
}

// GenParseAndStoreCSV returns a function that will parse the source file,\n//create a slice with an object per line and store the data in the db
func (m *TeamsHalf) GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) ParseAndStoreCSVFunc {
  return func() error {
    rows := make([]*TeamsHalf, 0)
    numErrors := 0
    err := pfunc(f, &rows)
    if err == nil {
       numrows := len(rows)
       if numrows > 0 {
         log.Println("TeamsHalf ==> Truncating")
         terr := repo.Truncate(m.GetTableName())
         if terr != nil {
            log.Println("truncate err:", terr.Error())
         }

         log.Printf("TeamsHalf ==> Inserting %d Records\n", numrows)
         for _, r := range rows {
           ierr := repo.Insert(m.GetTableName(), r)
           if ierr != nil {
             log.Printf("Insert error: %s\n", ierr.Error())
             numErrors++
           }
         }
       }
       log.Printf("TeamsHalf ==> %d records created\n", numrows-numErrors)
    }

    return err
   }
}
