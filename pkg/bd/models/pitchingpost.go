package models


type PitchingPost struct {
   Playerid   string `json:"playerid"  csv:"playerID"  db:"playerID"`
   Yearid   int `json:"yearid"  csv:"yearID"  db:"yearID"`
   Round   string `json:"round"  csv:"round"  db:"round"`
   Teamid   string `json:"teamid"  csv:"teamID"  db:"teamID"`
   Lgid   string `json:"lgid"  csv:"lgID"  db:"lgID"`
   W   int `json:"w"  csv:"W"  db:"W"`
   L   int `json:"l"  csv:"L"  db:"L"`
   G   int `json:"g"  csv:"G"  db:"G"`
   Gs   int `json:"gs"  csv:"GS"  db:"GS"`
   Cg   int `json:"cg"  csv:"CG"  db:"CG"`
   Sho   int `json:"sho"  csv:"SHO"  db:"SHO"`
   Sv   int `json:"sv"  csv:"SV"  db:"SV"`
   Ipouts   int `json:"ipouts"  csv:"IPouts"  db:"IPouts"`
   H   int `json:"h"  csv:"H"  db:"H"`
   Er   int `json:"er"  csv:"ER"  db:"ER"`
   Hr   int `json:"hr"  csv:"HR"  db:"HR"`
   Bb   int `json:"bb"  csv:"BB"  db:"BB"`
   So   int `json:"so"  csv:"SO"  db:"SO"`
   Baopp   string `json:"baopp"  csv:"BAOpp"  db:"BAOpp"`
   Era   string `json:"era"  csv:"ERA"  db:"ERA"`
   Ibb   string `json:"ibb"  csv:"IBB"  db:"IBB"`
   Wp   string `json:"wp"  csv:"WP"  db:"WP"`
   Hbp   string `json:"hbp"  csv:"HBP"  db:"HBP"`
   Bk   string `json:"bk"  csv:"BK"  db:"BK"`
   Bfp   string `json:"bfp"  csv:"BFP"  db:"BFP"`
   Gf   int `json:"gf"  csv:"GF"  db:"GF"`
   R   int `json:"r"  csv:"R"  db:"R"`
   Sh   string `json:"sh"  csv:"SH"  db:"SH"`
   Sf   string `json:"sf"  csv:"SF"  db:"SF"`
   Gidp   string `json:"gidp"  csv:"GIDP"  db:"GIDP"`
}

func (m *PitchingPost) GetTableName() string {
  return "PitchingPost"
}

func (m *PitchingPost) GetFileName() string {
  return "PitchingPost.csv"
}

func (m *PitchingPost) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/PitchingPost.csv"
}
