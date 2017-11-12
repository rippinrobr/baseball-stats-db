package models


type Appearances struct {
   Yearid   int `json:"yearid"  csv:"yearID"  db:"yearID"`
   Teamid   string `json:"teamid"  csv:"teamID"  db:"teamID"`
   Lgid   string `json:"lgid"  csv:"lgID"  db:"lgID"`
   Playerid   string `json:"playerid"  csv:"playerID"  db:"playerID"`
   G_All   string `json:"gAll"  csv:"G_all"  db:"G_all"`
   Gs   string `json:"gs"  csv:"GS"  db:"GS"`
   G_Batting   int `json:"gBatting"  csv:"G_batting"  db:"G_batting"`
   G_Defense   string `json:"gDefense"  csv:"G_defense"  db:"G_defense"`
   G_P   int `json:"gP"  csv:"G_p"  db:"G_p"`
   G_C   int `json:"gC"  csv:"G_c"  db:"G_c"`
   G_1B   int `json:"g1b"  csv:"G_1b"  db:"G_1b"`
   G_2B   int `json:"g2b"  csv:"G_2b"  db:"G_2b"`
   G_3B   int `json:"g3b"  csv:"G_3b"  db:"G_3b"`
   G_Ss   int `json:"gSs"  csv:"G_ss"  db:"G_ss"`
   G_Lf   int `json:"gLf"  csv:"G_lf"  db:"G_lf"`
   G_Cf   int `json:"gCf"  csv:"G_cf"  db:"G_cf"`
   G_Rf   int `json:"gRf"  csv:"G_rf"  db:"G_rf"`
   G_Of   int `json:"gOf"  csv:"G_of"  db:"G_of"`
   G_Dh   string `json:"gDh"  csv:"G_dh"  db:"G_dh"`
   G_Ph   string `json:"gPh"  csv:"G_ph"  db:"G_ph"`
   G_Pr   string `json:"gPr"  csv:"G_pr"  db:"G_pr"`
}

func (m *Appearances) GetTableName() string {
  return "Appearances"
}

func (m *Appearances) GetFileName() string {
  return "Appearances.csv"
}

func (m *Appearances) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/Appearances.csv"
}
