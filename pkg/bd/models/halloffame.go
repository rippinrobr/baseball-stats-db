package models


type HallOfFame struct {
   Playerid   string `json:"playerid"  csv:"playerID"  db:"playerID"`
   Yearid   int `json:"yearid"  csv:"yearid"  db:"yearid"`
   Votedby   string `json:"votedby"  csv:"votedBy"  db:"votedBy"`
   Ballots   string `json:"ballots"  csv:"ballots"  db:"ballots"`
   Needed   string `json:"needed"  csv:"needed"  db:"needed"`
   Votes   string `json:"votes"  csv:"votes"  db:"votes"`
   Inducted   string `json:"inducted"  csv:"inducted"  db:"inducted"`
   Category   string `json:"category"  csv:"category"  db:"category"`
   Needed_Note   string `json:"neededNote"  csv:"needed_note"  db:"needed_note"`
}

func (m *HallOfFame) GetTableName() string {
  return "HallOfFame"
}

func (m *HallOfFame) GetFileName() string {
  return "HallOfFame.csv"
}

func (m *HallOfFame) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/HallOfFame.csv"
}
