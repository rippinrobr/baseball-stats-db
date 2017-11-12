package models


type Schools struct {
   Schoolid   string `json:"schoolid"  csv:"schoolID"  db:"schoolID"`
   Name_Full   string `json:"nameFull"  csv:"name_full"  db:"name_full"`
   City   string `json:"city"  csv:"city"  db:"city"`
   State   string `json:"state"  csv:"state"  db:"state"`
   Country   string `json:"country"  csv:"country"  db:"country"`
}

func (m *Schools) GetTableName() string {
  return "Schools"
}

func (m *Schools) GetFileName() string {
  return "Schools.csv"
}

func (m *Schools) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/Schools.csv"
}
