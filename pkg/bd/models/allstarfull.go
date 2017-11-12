package models


type AllstarFull struct {
   Playerid   string `json:"playerid"  csv:"playerID"  db:"playerID"`
   Yearid   int `json:"yearid"  csv:"yearID"  db:"yearID"`
   Gamenum   int `json:"gamenum"  csv:"gameNum"  db:"gameNum"`
   Gameid   string `json:"gameid"  csv:"gameID"  db:"gameID"`
   Teamid   string `json:"teamid"  csv:"teamID"  db:"teamID"`
   Lgid   string `json:"lgid"  csv:"lgID"  db:"lgID"`
   Gp   string `json:"gp"  csv:"GP"  db:"GP"`
   Startingpos   string `json:"startingpos"  csv:"startingPos"  db:"startingPos"`
}

func (m *AllstarFull) GetTableName() string {
  return "AllstarFull"
}

func (m *AllstarFull) GetFileName() string {
  return "AllstarFull.csv"
}

func (m *AllstarFull) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/AllstarFull.csv"
}
