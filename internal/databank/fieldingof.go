package databank


import (
  "os"
  "log"
  "errors"
  "path/filepath"

  "github.com/rippinrobr/baseball-stats-db/internal/platform/parsers/csv"

  "github.com/rippinrobr/baseball-stats-db/internal/platform/db"

)

// FieldingOF is a model that maps the CSV to a DB Table
type FieldingOF struct {
   Playerid   string `json:"playerID"  csv:"playerID"  db:"playerID,omitempty"  bson:"playerID"`
   Yearid   int `json:"yearID"  csv:"yearID"  db:"yearID"  bson:"yearID"`
   Stint   int `json:"stint"  csv:"stint"  db:"stint"  bson:"stint"`
   Glf   int `json:"glf"  csv:"Glf"  db:"Glf"  bson:"Glf"`
   Gcf   int `json:"gcf"  csv:"Gcf"  db:"Gcf"  bson:"Gcf"`
   Grf   int `json:"grf"  csv:"Grf"  db:"Grf"  bson:"Grf"`
  inputDir string
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
  return filepath.Join(m.inputDir, "FieldingOF.csv")
}

// SetInputDirectory sets the input directory's path so it can be used to create the full path to the file
func (m *FieldingOF) SetInputDirectory(inputDir string) {
  m.inputDir=inputDir
}

// GenParseAndStoreCSV returns a function that will parse the source file,\n//create a slice with an object per line and store the data in the db
func (m *FieldingOF) GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) (ParseAndStoreCSVFunc, error) {
  if f == nil {
    return func() error{return nil}, errors.New("nil File")
  }

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
   }, nil
}
