package databank


import (
  "os"
  "log"
  "errors"
  "path/filepath"

  "github.com/rippinrobr/baseball-stats-db/internal/platform/parsers/csv"

  "github.com/rippinrobr/baseball-stats-db/internal/platform/db"

)

// TeamsFranchises is a model that maps the CSV to a DB Table
type TeamsFranchises struct {
   Franchid   string `json:"franchID"  csv:"franchID"  db:"franchID,omitempty"  bson:"franchID"`
   Franchname   string `json:"franchName"  csv:"franchName"  db:"franchName,omitempty"  bson:"franchName"`
   Active   string `json:"active"  csv:"active"  db:"active,omitempty"  bson:"active"`
   Naassoc   string `json:"nAassoc"  csv:"NAassoc"  db:"NAassoc,omitempty"  bson:"NAassoc"`
  inputDir string
}

// GetTableName returns the name of the table that the data will be stored in
func (m *TeamsFranchises) GetTableName() string {
  return "teamsfranchises"
}

// GetFileName returns the name of the source file the model was created from
func (m *TeamsFranchises) GetFileName() string {
  return "TeamsFranchises.csv"
}

// GetFilePath returns the path of the source file
func (m *TeamsFranchises) GetFilePath() string {
  return filepath.Join(m.inputDir, "TeamsFranchises.csv")
}

// SetInputDirectory sets the input directory's path so it can be used to create the full path to the file
func (m *TeamsFranchises) SetInputDirectory(inputDir string) {
  m.inputDir=inputDir
}

// GenParseAndStoreCSV returns a function that will parse the source file,\n//create a slice with an object per line and store the data in the db
func (m *TeamsFranchises) GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) (ParseAndStoreCSVFunc, error) {
  if f == nil {
    return func() error{return nil}, errors.New("nil File")
  }

  return func() error {
    rows := make([]*TeamsFranchises, 0)
    numErrors := 0
    err := pfunc(f, &rows)
    if err == nil {
       numrows := len(rows)
       if numrows > 0 {
         log.Println("TeamsFranchises ==> Truncating")
         terr := repo.Truncate(m.GetTableName())
         if terr != nil {
            log.Println("truncate err:", terr.Error())
         }

         log.Printf("TeamsFranchises ==> Inserting %d Records\n", numrows)
         for _, r := range rows {
           ierr := repo.Insert(m.GetTableName(), r)
           if ierr != nil {
             log.Printf("Insert error: %s\n", ierr.Error())
             numErrors++
           }
         }
       }
       log.Printf("TeamsFranchises ==> %d records created\n", numrows-numErrors)
    }

    return err
   }, nil
}
