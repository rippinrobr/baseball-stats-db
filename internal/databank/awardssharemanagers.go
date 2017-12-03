package databank


import (
  "os"
  "log"
  "errors"
  "path/filepath"

  "github.com/rippinrobr/baseball-stats-db/internal/platform/parsers/csv"

  "github.com/rippinrobr/baseball-stats-db/internal/platform/db"

)

// AwardsShareManagers is a model that maps the CSV to a DB Table
type AwardsShareManagers struct {
   Awardid   string `json:"awardID"  csv:"awardID"  db:"awardID,omitempty"  bson:"awardID"`
   Yearid   int `json:"yearID"  csv:"yearID"  db:"yearID"  bson:"yearID"`
   Lgid   string `json:"lgID"  csv:"lgID"  db:"lgID,omitempty"  bson:"lgID"`
   Playerid   string `json:"playerID"  csv:"playerID"  db:"playerID,omitempty"  bson:"playerID"`
   Pointswon   int `json:"pointsWon"  csv:"pointsWon"  db:"pointsWon"  bson:"pointsWon"`
   Pointsmax   int `json:"pointsMax"  csv:"pointsMax"  db:"pointsMax"  bson:"pointsMax"`
   Votesfirst   int `json:"votesFirst"  csv:"votesFirst"  db:"votesFirst"  bson:"votesFirst"`
  inputDir string
}

// GetTableName returns the name of the table that the data will be stored in
func (m *AwardsShareManagers) GetTableName() string {
  return "awardssharemanagers"
}

// GetFileName returns the name of the source file the model was created from
func (m *AwardsShareManagers) GetFileName() string {
  return "AwardsShareManagers.csv"
}

// GetFilePath returns the path of the source file
func (m *AwardsShareManagers) GetFilePath() string {
  return filepath.Join(m.inputDir, "AwardsShareManagers.csv")
}

// SetInputDirectory sets the input directory's path so it can be used to create the full path to the file
func (m *AwardsShareManagers) SetInputDirectory(inputDir string) {
  m.inputDir=inputDir
}

// GenParseAndStoreCSV returns a function that will parse the source file,\n//create a slice with an object per line and store the data in the db
func (m *AwardsShareManagers) GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) (ParseAndStoreCSVFunc, error) {
  if f == nil {
    return func() error{return nil}, errors.New("nil File")
  }

  return func() error {
    rows := make([]*AwardsShareManagers, 0)
    numErrors := 0
    err := pfunc(f, &rows)
    if err == nil {
       numrows := len(rows)
       if numrows > 0 {
         log.Println("AwardsShareManagers ==> Truncating")
         terr := repo.Truncate(m.GetTableName())
         if terr != nil {
            log.Println("truncate err:", terr.Error())
         }

         log.Printf("AwardsShareManagers ==> Inserting %d Records\n", numrows)
         for _, r := range rows {
           ierr := repo.Insert(m.GetTableName(), r)
           if ierr != nil {
             log.Printf("Insert error: %s\n", ierr.Error())
             numErrors++
           }
         }
       }
       log.Printf("AwardsShareManagers ==> %d records created\n", numrows-numErrors)
    }

    return err
   }, nil
}
