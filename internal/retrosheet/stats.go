package retrosheet

// SacHits - may include sac flies prior to 1954
// SacFlies - are what you think they are from 1954 on
// PO may not equal 3 times the number of innings pitched.  Prior to 1931 no putout was given when
// a runner was hit by a batted ball.
type visitorStats struct {
	VisitorsAB                  int `json:"visitorsAB" bson:"visitorsAB" db:"visitors_ab"`
	VisitorsH                   int `json:"visitorsH" bson:"visitorsH" db:"visitors_h"`
	VisitorsDoubles             int `json:"visitorsDoubles" bson:"visitorsDoubles" db:"visitors_doubles"`
	VisitorsTriples             int `json:"visitorsTriples" bson:"visitorsTriples" db:"visitors_triples"`
	VisitorsHR                  int `json:"visitorsHR" bson:"visitorsHR" db:"visitors_hr"`
	VisitorsRBI                 int `json:"visitorsRBI" bson:"visitorsRBI" db:"visitors_rbi"`
	VisitorsSacHits             int `json:"visitorsSacHits" bson:"visitorsSacHits" db:"visitors_sac_hits"`
	VisitorsSacFlies            int `json:"visitorsSacFlies" bson:"visitorsSacFlies" db:"visitors_sac_flies"`
	VisitorsHBP                 int `json:"visitorsHBP" bson:"visitorsHBP" db:"visitors_hbp"`
	VisitorsBB                  int `json:"visitorsBB" bson:"visitorsBB" db:"visitors_bb"`
	VisitorsIBB                 int `json:"visitorsIBB" bson:"visitorsIBB" db:"visitors_ibb"`
	VisitorsK                   int `json:"visitorsK" bson:"visitorsK" db:"visitors_k"`
	VisitorsSB                  int `json:"visitorsSB" bson:"visitorsSB" db:"visitors_sb"`
	VisitorsCS                  int `json:"visitorsCS" bson:"visitorsCS" db:"visitors_cs"`
	VisitorsGIDP                int `json:"visitorsGIDP" bson:"visitorsGIDP" db:"visitors_gidp"`
	VisitorsCatcherInterference int `json:"visitorsCatcherInterference" bson:"catcherInterference" db:"visitors_catcher_interference"`
	VisitorsLOB                 int `json:"visitorsLOB" bson:"visitorsLOB" db:"visitors_lob"`
	VisitorsPitchersUsed        int `json:"visitorsPitchersUsed" bson:"visitorsPitchersUsed" db:"visitors_pitchers_used"`
	VisitorsIndividualER        int `json:"visitorsIndividualER" bson:"visitorsIndividualER" db:"visitors_individual_er"`
	VisitorsTeamER              int `json:"visitorsTeamER" bson:"visitorsTeamER" db:"visitors_team_er"`
	VisitorsWP                  int `json:"visitorsWP" bson:"visitorsWP" db:"visitors_wp"`
	VisitorsBalks               int `json:"visitorsBalks" bson:"visitorsBalks" db:"visitors_balks"`
	VisitorsPO                  int `json:"visitorsPO" bson:"visitorsPO" db:"visitors_po"`
	VisitorsA                   int `json:"visitorsA" bson:"visitorsA" db:"visitors_a"`
	VisitorsE                   int `json:"visitorsE" bson:"visitorsE" db:"visitors_e"`
	VisitorsPB                  int `json:"visitorsPB" bson:"visitorsPB" db:"visitors_pb"`
	VisitorsDP                  int `json:"visitorsDP" bson:"visitorsDP" db:"visitors_dp"`
	VisitorsTriplePlays         int `json:"visitorsTriplePlays" bson:"visitorsTriplePlays" db:"visitors_triple_plays"`
}

// SacHits - may include sac flies prior to 1954
// SacFlies - are what you think they are from 1954 on
// PO may not equal 3 times the number of innings pitched.  Prior to 1931 no putout was given when
// a runner was hit by a batted ball.
type homeStats struct {
	HomeAB                  int `json:"homeAB" bson:"homeAB" db:"home_ab"`
	HomeH                   int `json:"homeH" bson:"homeH" db:"home_h"`
	HomeDoubles             int `json:"homeDoubles" bson:"homeDoubles" db:"home_doubles"`
	HomeTriples             int `json:"homeTriples" bson:"homeTriples" db:"home_triples"`
	HomeHR                  int `json:"homeHR" bson:"homeHR" db:"home_hr"`
	HomeRBI                 int `json:"homeRBI" bson:"homeRBI" db:"home_rbi"`
	HomeSacHits             int `json:"homeSacHits" bson:"homeSacHits" db:"home_sac_hits"`
	HomeSacFlies            int `json:"homeSacFlies" bson:"homeSacFlies" db:"home_sac_flies"`
	HomeHBP                 int `json:"homeHBP" bson:"homeHBP" db:"home_hbp"`
	HomeBB                  int `json:"homeBB" bson:"homeBB" db:"home_bb"`
	HomeIBB                 int `json:"homeIBB" bson:"homeIBB" db:"home_ibb"`
	HomeK                   int `json:"homeK" bson:"homeK" db:"home_k"`
	HomeSB                  int `json:"homeSB" bson:"homeSB" db:"home_sb"`
	HomeCS                  int `json:"homeCS" bson:"homeCS" db:"home_cs"`
	HomeGIDP                int `json:"homeGIDP" bson:"homeGIDP" db:"home_gidp"`
	HomeCatcherInterference int `json:"homeCatcherInterference" bson:"homeCatcherInterference" db:"home_catcher_interference"`
	HomeLOB                 int `json:"homeLOB" bson:"homeLOB" db:"home_lob"`
	HomePitchersUsed        int `json:"homePitchersUsed" bson:"homePitchersUsed" db:"home_pitchers_used"`
	HomeIndividualER        int `json:"homeIndividualER" bson:"homeIndividualER" db:"home_individual_er"`
	HomeTeamER              int `json:"homeTeamER" bson:"homeTeamER" db:"home_team_er"`
	HomeWP                  int `json:"homeWP" bson:"homeWP" db:"home_wp"`
	HomeBalks               int `json:"homeBalks" bson:"homeBalks" db:"home_balks"`
	HomePO                  int `json:"homePO" bson:"homePO" db:"home_po"`
	HomeA                   int `json:"homeA" bson:"homeA" db:"home_a"`
	HomeE                   int `json:"homeE" bson:"homeE" db:"home_e"`
	HomePB                  int `json:"homePB" bson:"homePB" db:"home_pb"`
	HomeDP                  int `json:"homeDP" bson:"homeDP" db:"home_dp"`
	HomeTriplePlays         int `json:"homeTriplePlays" bson:"homeTriplePlays" db:"home_triple_plays"`
}
