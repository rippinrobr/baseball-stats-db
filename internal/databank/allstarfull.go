package databank


import (
  "os"
  "log"
  "errors"
  "path/filepath"

  "github.com/rippinrobr/baseball-stats-db/internal/platform/parsers/csv"

  "github.com/rippinrobr/baseball-stats-db/internal/platform/db"

)

// AllstarFull is a model that maps the CSV to a DB Table
type AllstarFull struct {
   Playerid   string `json:"playerID"  csv:"playerID"  db:"playerID,omitempty"  bson:"playerID"`
   Yearid   int `json:"yearID"  csv:"yearID"  db:"yearID"  bson:"yearID"`
   Gamenum   int `json:"gameNum"  csv:"gameNum"  db:"gameNum"  bson:"gameNum"`
   Gameid   string `json:"gameID"  csv:"gameID"  db:"gameID,omitempty"  bson:"gameID"`
   Teamid   string `json:"teamID"  csv:"teamID"  db:"teamID,omitempty"  bson:"teamID"`
   Lgid   string `json:"lgID"  csv:"lgID"  db:"lgID,omitempty"  bson:"lgID"`
   Gp   int `json:"gP"  csv:"GP"  db:"GP"  bson:"GP"`
   Startingpos   int `json:"startingPos"  csv:"startingPos"  db:"startingPos"  bson:"startingPos"`
  inputDir string
}

// GetTableName returns the name of the table that the data will be stored in
func (m *AllstarFull) GetTableName() string {
  return "allstarfull"
}

// GetFileName returns the name of the source file the model was created from
func (m *AllstarFull) GetFileName() string {
  return "AllstarFull.csv"
}

// GetFilePath returns the path of the source file
func (m *AllstarFull) GetFilePath() string {
  return filepath.Join(m.inputDir, "AllstarFull.csv")
}

// SetInputDirectory sets the input directory's path so it can be used to create the full path to the file
func (m *AllstarFull) SetInputDirectory(inputDir string) {
  m.inputDir=inputDir
}

// GenParseAndStoreCSV returns a function that will parse the source file,\n//create a slice with an object per line and store the data in the db
func (m *AllstarFull) GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) (ParseAndStoreCSVFunc, error) {
  if f == nil {
    return func() error{return nil}, errors.New("nil File")
  }

  return func() error {
    rows := make([]*AllstarFull, 0)
    numErrors := 0
    err := pfunc(f, &rows)
    if err == nil {
       numrows := len(rows)
       if numrows > 0 {
         log.Println("AllstarFull ==> Truncating")
         terr := repo.Truncate(m.GetTableName())
         if terr != nil {
            log.Println("truncate err:", terr.Error())
         }

         log.Printf("AllstarFull ==> Inserting %d Records\n", numrows)
         for _, r := range rows {
           ierr := repo.Insert(m.GetTableName(), r)
           if ierr != nil {
             log.Printf("Insert error: %s\n", ierr.Error())
             numErrors++
           }
         }
       }
       log.Printf("AllstarFull ==> %d records created\n", numrows-numErrors)
    }

    return err
   }, nil
}
