package models


type FieldingOF struct {
   Playerid   string `json:"playerid"  csv:"playerID"  db:"playerID"`
   Yearid   int `json:"yearid"  csv:"yearID"  db:"yearID"`
   Stint   int `json:"stint"  csv:"stint"  db:"stint"`
   Glf   string `json:"glf"  csv:"Glf"  db:"Glf"`
   Gcf   string `json:"gcf"  csv:"Gcf"  db:"Gcf"`
   Grf   string `json:"grf"  csv:"Grf"  db:"Grf"`
}

func (m *FieldingOF) GetTableName() string {
  return "FieldingOF"
}

func (m *FieldingOF) GetFileName() string {
  return "FieldingOF.csv"
}

func (m *FieldingOF) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/FieldingOF.csv"
}
