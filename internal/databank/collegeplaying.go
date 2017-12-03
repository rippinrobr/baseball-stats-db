package databank


import (
  "os"
  "log"
  "errors"
  "path/filepath"

  "github.com/rippinrobr/baseball-stats-db/internal/platform/parsers/csv"

  "github.com/rippinrobr/baseball-stats-db/internal/platform/db"

)

// CollegePlaying is a model that maps the CSV to a DB Table
type CollegePlaying struct {
   Playerid   string `json:"playerID"  csv:"playerID"  db:"playerID,omitempty"  bson:"playerID"`
   Schoolid   string `json:"schoolID"  csv:"schoolID"  db:"schoolID,omitempty"  bson:"schoolID"`
   Yearid   int `json:"yearID"  csv:"yearID"  db:"yearID"  bson:"yearID"`
  inputDir string
}

// GetTableName returns the name of the table that the data will be stored in
func (m *CollegePlaying) GetTableName() string {
  return "collegeplaying"
}

// GetFileName returns the name of the source file the model was created from
func (m *CollegePlaying) GetFileName() string {
  return "CollegePlaying.csv"
}

// GetFilePath returns the path of the source file
func (m *CollegePlaying) GetFilePath() string {
  return filepath.Join(m.inputDir, "CollegePlaying.csv")
}

// SetInputDirectory sets the input directory's path so it can be used to create the full path to the file
func (m *CollegePlaying) SetInputDirectory(inputDir string) {
  m.inputDir=inputDir
}

// GenParseAndStoreCSV returns a function that will parse the source file,\n//create a slice with an object per line and store the data in the db
func (m *CollegePlaying) GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) (ParseAndStoreCSVFunc, error) {
  if f == nil {
    return func() error{return nil}, errors.New("nil File")
  }

  return func() error {
    rows := make([]*CollegePlaying, 0)
    numErrors := 0
    err := pfunc(f, &rows)
    if err == nil {
       numrows := len(rows)
       if numrows > 0 {
         log.Println("CollegePlaying ==> Truncating")
         terr := repo.Truncate(m.GetTableName())
         if terr != nil {
            log.Println("truncate err:", terr.Error())
         }

         log.Printf("CollegePlaying ==> Inserting %d Records\n", numrows)
         for _, r := range rows {
           ierr := repo.Insert(m.GetTableName(), r)
           if ierr != nil {
             log.Printf("Insert error: %s\n", ierr.Error())
             numErrors++
           }
         }
       }
       log.Printf("CollegePlaying ==> %d records created\n", numrows-numErrors)
    }

    return err
   }, nil
}
