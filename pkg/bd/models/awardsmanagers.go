package models


type AwardsManagers struct {
   Playerid   string `json:"playerid"  csv:"playerID"  db:"playerID"`
   Awardid   string `json:"awardid"  csv:"awardID"  db:"awardID"`
   Yearid   int `json:"yearid"  csv:"yearID"  db:"yearID"`
   Lgid   string `json:"lgid"  csv:"lgID"  db:"lgID"`
   Tie   string `json:"tie"  csv:"tie"  db:"tie"`
   Notes   string `json:"notes"  csv:"notes"  db:"notes"`
}

func (m *AwardsManagers) GetTableName() string {
  return "AwardsManagers"
}

func (m *AwardsManagers) GetFileName() string {
  return "AwardsManagers.csv"
}

func (m *AwardsManagers) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/AwardsManagers.csv"
}
