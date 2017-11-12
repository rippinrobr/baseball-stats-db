package models


type Managers struct {
   Playerid   string `json:"playerid"  csv:"playerID"  db:"playerID"`
   Yearid   int `json:"yearid"  csv:"yearID"  db:"yearID"`
   Teamid   string `json:"teamid"  csv:"teamID"  db:"teamID"`
   Lgid   string `json:"lgid"  csv:"lgID"  db:"lgID"`
   Inseason   int `json:"inseason"  csv:"inseason"  db:"inseason"`
   G   int `json:"g"  csv:"G"  db:"G"`
   W   int `json:"w"  csv:"W"  db:"W"`
   L   int `json:"l"  csv:"L"  db:"L"`
   Rank   string `json:"rank"  csv:"rank"  db:"rank"`
   Plyrmgr   string `json:"plyrmgr"  csv:"plyrMgr"  db:"plyrMgr"`
}

func (m *Managers) GetTableName() string {
  return "Managers"
}

func (m *Managers) GetFileName() string {
  return "Managers.csv"
}

func (m *Managers) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/Managers.csv"
}
