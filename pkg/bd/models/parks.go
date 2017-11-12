package models


type Parks struct {
   Parkkey   string `json:"parkkey"  csv:"park.key"  db:"park.key"`
   Parkname   string `json:"parkname"  csv:"park.name"  db:"park.name"`
   Parkalias   string `json:"parkalias"  csv:"park.alias"  db:"park.alias"`
   City   string `json:"city"  csv:"city"  db:"city"`
   State   string `json:"state"  csv:"state"  db:"state"`
   Country   string `json:"country"  csv:"country"  db:"country"`
}

func (m *Parks) GetTableName() string {
  return "Parks"
}

func (m *Parks) GetFileName() string {
  return "Parks.csv"
}

func (m *Parks) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/Parks.csv"
}
