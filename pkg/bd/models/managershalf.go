package models


type ManagersHalf struct {
   Playerid   string `json:"playerid"  csv:"playerID"  db:"playerID"`
   Yearid   int `json:"yearid"  csv:"yearID"  db:"yearID"`
   Teamid   string `json:"teamid"  csv:"teamID"  db:"teamID"`
   Lgid   string `json:"lgid"  csv:"lgID"  db:"lgID"`
   Inseason   int `json:"inseason"  csv:"inseason"  db:"inseason"`
   Half   int `json:"half"  csv:"half"  db:"half"`
   G   int `json:"g"  csv:"G"  db:"G"`
   W   int `json:"w"  csv:"W"  db:"W"`
   L   int `json:"l"  csv:"L"  db:"L"`
   Rank   int `json:"rank"  csv:"rank"  db:"rank"`
}

func (m *ManagersHalf) GetTableName() string {
  return "ManagersHalf"
}

func (m *ManagersHalf) GetFileName() string {
  return "ManagersHalf.csv"
}

func (m *ManagersHalf) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/ManagersHalf.csv"
}
