package models


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

func (m *TeamsHalf) GetTableName() string {
  return "TeamsHalf"
}

func (m *TeamsHalf) GetFileName() string {
  return "TeamsHalf.csv"
}

func (m *TeamsHalf) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/TeamsHalf.csv"
}
