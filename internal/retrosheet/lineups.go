package retrosheet

type visitorLineup struct {
	VisitorBatter1ID   string `json:"visitorBatter1Id" bson:"visitorBatter1Id" db:"visitor_batter1_id"`
	VisitorBatter1Name string `json:"visitorBatter1Name" bson:"visitorBatter1Name" db:"visitor_batter1_name"`
	VisitorBatter1Pos  int    `json:"visitorBatter1Pos" bson:"visitorBatter1Pos" db:"visitor_batter1_pos"`
	VisitorBatter2ID   string `json:"visitorBatter2Id" bson:"visitorBatter2Id" db:"visitor_batter2_id"`
	VisitorBatter2Name string `json:"visitorBatter2Name" bson:"visitorBatter2Name" db:"visitor_batter2_name"`
	VisitorBatter2Pos  int    `json:"visitorBatter2Pos" bson:"visitorBatter2Pos" db:"visitor_batter2_pos"`
	VisitorBatter3ID   string `json:"visitorBatter3Id" bson:"visitorBatter3Id" db:"visitor_batter3_id"`
	VisitorBatter3Name string `json:"visitorBatter3Name" bson:"visitorBatter3Name" db:"visitor_batter3_name"`
	VisitorBatter3Pos  int    `json:"visitorBatter3Pos" bson:"visitorBatter3Pos" db:"visitor_batter3_pos"`
	VisitorBatter4ID   string `json:"visitorBatter4Id" bson:"visitorBatter4Id" db:"visitor_batter4_id"`
	VisitorBatter4Name string `json:"visitorBatter4Name" bson:"visitorBatter4Name" db:"visitor_batter4_name"`
	VisitorBatter4Pos  int    `json:"visitorBatter4Pos" bson:"visitorBatter4Pos" db:"visitor_batter4_pos"`
	VisitorBatter5ID   string `json:"visitorBatter5Id" bson:"visitorBatter5Id" db:"visitor_batter5_id"`
	VisitorBatter5Name string `json:"visitorBatter5Name" bson:"visitorBatter5Name" db:"visitor_batter5_name"`
	VisitorBatter5Pos  int    `json:"visitorBatter5Pos" bson:"visitorBatter5Pos" db:"visitor_batter5_pos"`
	VisitorBatter6ID   string `json:"visitorBatter6Id" bson:"visitorBatter6Id" db:"visitor_batter6_id"`
	VisitorBatter6Name string `json:"visitorBatter6Name" bson:"visitorBatter6Name" db:"visitor_batter6_name"`
	VisitorBatter6Pos  int    `json:"visitorBatter6Pos" bson:"visitorBatter6Pos" db:"visitor_batter6_pos"`
	VisitorBatter7ID   string `json:"VisitorBatter7Id" bson:"VisitorBatter7Id" db:"visitor_batter7_id"`
	VisitorBatter7Name string `json:"visitorBatter7Name" bson:"visitorBatter7Name" db:"visitor_batter7_name"`
	VisitorBatter7Pos  int    `json:"visitorBatter7Pos" bson:"visitorBatter7Pos" db:"visitor_batter7_pos"`
	VisitorBatter8ID   string `json:"VisitorBatter8Id" bson:"VisitorBatter8Id" db:"visitor_batter8_id"`
	VisitorBatter8Name string `json:"visitorBatter8Name" bson:"visitorBatter8Name" db:"visitor_batter8_name"`
	VisitorBatter8Pos  int    `json:"visitorBatter8Pos" bson:"visitorBatter8Pos" db:"visitor_batter8_pos"`
	VisitorBatter9ID   string `json:"VisitorBatter9Id" bson:"VisitorBatter9Id" db:"visitor_batter9_id"`
	VisitorBatter9Name string `json:"visitorBatter9Name" bson:"visitorBatter9Name" db:"visitor_batter9_name"`
	VisitorBatter9Pos  int    `json:"visitorBatter9Pos" bson:"visitorBatter9Pos" db:"visitor_batter9_pos"`
}

type homeLineup struct {
	HomeBatter1ID   string `json:"homeBatter1Id" bson:"homeBatter1Id" db:"home_batter1_id"`
	HomeBatter1Name string `json:"homeBatter1Name" bson:"homeBatter1Name" db:"home_batter1_name"`
	HomeBatter1Pos  int    `json:"homeBatter1Pos" bson:"homeBatter1Pos" db:"home_batter1_pos"`
	HomeBatter2ID   string `json:"homeBatter2Id" bson:"homeBatter2Id" db:"home_batter2_id"`
	HomeBatter2Name string `json:"homeBatter2Name" bson:"homeBatter2Name" db:"home_batter2_pos"`
	HomeBatter2Pos  int    `json:"homeBatter2Pos" bson:"homeBatter2Pos" db:"home_batter2_pos"`
	HomeBatter3ID   string `json:"homeBatter3Id" bson:"homeBatter3Id" db:"home_batter3_id"`
	HomeBatter3Name string `json:"homeBatter3Name" bson:"homeBatter3Name" db:"home_batter3_name"`
	HomeBatter3Pos  int    `json:"homeBatter3Pos" bson:"homeBatter3Pos" db:"home_batter3_pos"`
	HomeBatter4ID   string `json:"homeBatter4Id" bson:"homeBatter4Id" db:"home_batter4_id"`
	HomeBatter4Name string `json:"homeBatter4Name" bson:"homeBatter4Name" db:"home_batter4_name"`
	HomeBatter4Pos  int    `json:"homeBatter4Pos" bson:"homeBatter4Pos" db:"home_batter4_pos"`
	HomeBatter5ID   string `json:"homeBatter5Id" bson:"homeBatter5Id" db:"home_batter5_id"`
	HomeBatter5Name string `json:"homeBatter5Name" bson:"homeBatter5Name" db:"home_batter5_name"`
	HomeBatter5Pos  int    `json:"homeBatter5Pos" bson:"homeBatter5Pos" db:"home_batter5_pos"`
	HomeBatter6ID   string `json:"homeBatter6Id" bson:"homeBatter6Id" db:"home_batter6_id"`
	HomeBatter6Name string `json:"homeBatter6Name" bson:"homeBatter6Name" db:"home_batter6_name"`
	HomeBatter6Pos  int    `json:"homeBatter6Pos" bson:"homeBatter6Pos" db:"home_batter6_pos"`
	HomeBatter7ID   string `json:"homeBatter7Id" bson:"homeBatter7Id" db:"home_batter7_id"`
	HomeBatter7Name string `json:"homeBatter7Name" bson:"homeBatter7Name" db:"home_batter7_name"`
	HomeBatter7Pos  int    `json:"homeBatter7Pos" bson:"homeBatter7Pos" db:"home_batter7_pos"`
	HomeBatter8ID   string `json:"homeBatter8Id" bson:"homeBatter8Id" db:"home_batter8_id"`
	HomeBatter8Name string `json:"homeBatter8Name" bson:"homeBatter8Name" db:"home_batter8_name"`
	HomeBatter8Pos  int    `json:"homeBatter8Pos" bson:"homeBatter8Pos" db:"home_batter8_pos"`
	HomeBatter9ID   string `json:"homeBatter9Id" bson:"homeBatter9Id" db:"home_batter9_id"`
	HomeBatter9Name string `json:"homeBatter9Name" bson:"homeBatter9Name" db:"home_batter9_name"`
	HomeBatter9Pos  int    `json:"homeBatter9Pos" bson:"homeBatter9Pos" db:"home_batter9_pos"`
}
