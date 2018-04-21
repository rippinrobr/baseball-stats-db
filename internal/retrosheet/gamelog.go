package retrosheet

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"
)

// GameLog models a line in a retrosheet.org game log file.  For details on what each field represents
// please read http://www.retrosheet.org/gamelogs/glfields.txt
// GameOuts = length of game in outs
type GameLog struct {
	gameUniqueID             `json:",inline" bson:",inline"`
	WeekDay                  time.Weekday `json:"weekDay" bson:"weekDay" db:"week_day"`
	VisitorsTeam             string       `json:"visitorsTeam" bson:"visitorsTeam" db:"visitors_team"`
	Visitorsleague           string       `json:"visitorsleague" bson:"visitorsleague" db:"visitors_league"`
	VisitorsSeasonGameNumber int          `json:"visitorsSeasonGameNumber" bson:"visitorsSeasonGameNumber" db:"visitors_season_game_number"`
	HomeTeam                 string       `json:"homeTeam" bson:"homeTeam" db:"home_team"`
	HomeLeague               string       `json:"homeLeague" bson:"homeLeague" db:"home_league"`
	HomeSeasonGameNumber     int          `json:"homeSeasonGameNumber" bson:"homeSeasonGameNumber" db:"home_season_game_number"`
	VisitorsScore            int          `json:"visitorsScore" bson:"visitorsScore" db:"visitors_score"`
	HomeScore                int          `json:"homeScore" bson:"homeScore" db:"home_score"`
	GameOuts                 int          `json:"gameOuts" bson:"gameOuts" db:"game_outs"`
	GameType                 string       `json:"gameType" bson:"gameType" db:"game_type"`
	CompletionInformation    string       `json:"completionInformation" bson:"completionInformation" db:"completion_information"`
	ForfeitInformation       string       `json:"forfeitInformation" bson:"forfeitInformation" db:"forfeit_information"`
	ProtestInformation       string       `json:"protestInformation" bson:"protestInformation" db:"protest_information"`
	ParkID                   string       `json:"parkId" bson:"parkId" db:"park_id"`
	Attendance               int          `json:"attendance" bson:"attendance" db:"attendance"`
	GameTimeInMinutes        int          `json:"gameTimeInMinutes" bson:"gameTimeInMinutes" db:"game_time_in_minutes"`
	VisitorsLineScore        string       `json:"visitorsLineScore" bson:"visitorsLineScore" db:"visitors_line_score"`
	HomeLineScore            string       `json:"homeLineScore" bson:"homeLineScore" db:"home_line_score"`
	visitorStats             `json:",inline" bson:",inline"`
	homeStats                `json:",inline" bson:",inline"`
	umpires                  `json:",inline" bson:",inline"`
	VisitorsMgrID            string `json:"visitorsMgrId" bson:"visitorsMgrId" db:"visitors_mgr_id"`
	VisitorsMgrName          string `json:"visitorsMgrName" bson:"visitorsMgrName" db:"visitors_mgr_name"`
	HomeMgrID                string `json:"homeMgrId" bson:"homeMgrId" db:"home_mgr_id"`
	HomeMgrName              string `json:"homeMgrName" bson:"homeMgrName" db:"home_mgr_name"`
	pitching                 `json:",inline" bson:",inline"`
	GameWinningRBIBatterID   string `json:"gameWinningRBIBatterId" bson:"gameWinningRBIBatterId" db:"game_winning_rbi_batter_id"`
	GameWinningRBIBatterName string `json:"gameWinningRBIBatterName" bson:"gameWinningRBIBatterName" db:"game_winning_rbi_batter_name"`
	visitorLineup            `json:",inline" bson:",inline"`
	homeLineup               `json:",inline" bson:",inline"`
	AdditionalInformation    string `json:"additionalInformation" bson:"additionalInformation" db:"additional_information"`
	AcquisitionInformation   string `json:"acquisitionInformation" bson:"acquisitionInformation" db:"acquisition_information"`
}

// NewGameLog create a new GameLog structure takes in the season and
// a slice of strings that represents a line in the GL*.TXT file.
func NewGameLog(season int, data []string) (GameLog, error) {
	var err error
	gl := GameLog{}

	if season < 1871 || season > time.Now().Year() {
		return gl, errors.New("season must between than 1876 and the current year")
	}
	gl.Season = season

	if len(data) != 161 {
		return gl, fmt.Errorf("game log data has 161 columns data only has %d", len(data))
	}

	gl.GameDateStr = data[0]
	gl.GameMonth = monthMap[gl.GameDateStr[4:6]]
	gl.GameMonthDay, _ = strconv.Atoi(gl.GameDateStr[6:8])
	gl.GameNumber, err = convertNonIntGameNumberValues(data[1])
	if err != nil {
		log.Printf("unable to convert GameNumber value '%s'\n", data[1])
	}

	gl.WeekDay = weekDayMap[data[2]]
	gl.VisitorsTeam = data[3]
	gl.Visitorsleague = data[4]
	gl.VisitorsSeasonGameNumber, _ = strconv.Atoi(data[5])
	gl.HomeTeam = data[6]
	gl.HomeLeague = data[7]
	gl.HomeSeasonGameNumber, _ = strconv.Atoi(data[8])
	gl.VisitorsScore, _ = strconv.Atoi(data[9])
	gl.HomeScore, _ = strconv.Atoi(data[10])
	gl.GameOuts, _ = strconv.Atoi(data[11])
	gl.GameType = data[12]
	gl.CompletionInformation = data[13]
	gl.ForfeitInformation = data[14]
	gl.ProtestInformation = data[15]
	gl.ParkID = data[16]
	gl.Attendance, _ = strconv.Atoi(data[17])
	gl.GameTimeInMinutes, _ = strconv.Atoi(data[18])
	gl.VisitorsLineScore = data[19]
	gl.HomeLineScore = data[20]
	gl.VisitorsAB, _ = strconv.Atoi(data[21])
	gl.VisitorsH, _ = strconv.Atoi(data[22])
	gl.VisitorsDoubles, _ = strconv.Atoi(data[23])
	gl.VisitorsTriples, _ = strconv.Atoi(data[24])
	gl.VisitorsHR, _ = strconv.Atoi(data[25])
	gl.VisitorsRBI, _ = strconv.Atoi(data[26])
	gl.VisitorsSacHits, _ = strconv.Atoi(data[27])
	gl.VisitorsSacFlies, _ = strconv.Atoi(data[28])
	gl.VisitorsHBP, _ = strconv.Atoi(data[29])
	gl.VisitorsBB, _ = strconv.Atoi(data[30])
	gl.VisitorsIBB, _ = strconv.Atoi(data[31])
	gl.VisitorsK, _ = strconv.Atoi(data[32])
	gl.VisitorsSB, _ = strconv.Atoi(data[33])
	gl.VisitorsCS, _ = strconv.Atoi(data[34])
	gl.VisitorsGIDP, _ = strconv.Atoi(data[35])
	gl.VisitorsCatcherInterference, _ = strconv.Atoi(data[36])
	gl.VisitorsLOB, _ = strconv.Atoi(data[37])
	gl.VisitorsPitchersUsed, _ = strconv.Atoi(data[38])
	gl.VisitorsIndividualER, _ = strconv.Atoi(data[39])
	gl.VisitorsTeamER, _ = strconv.Atoi(data[40])
	gl.VisitorsWP, _ = strconv.Atoi(data[41])
	gl.VisitorsBalks, _ = strconv.Atoi(data[42])
	gl.VisitorsPO, _ = strconv.Atoi(data[43])
	gl.VisitorsA, _ = strconv.Atoi(data[44])
	gl.VisitorsE, _ = strconv.Atoi(data[45])
	gl.VisitorsPB, _ = strconv.Atoi(data[46])
	gl.VisitorsDP, _ = strconv.Atoi(data[47])
	gl.VisitorsTriplePlays, _ = strconv.Atoi(data[48])
	gl.HomeAB, _ = strconv.Atoi(data[49])
	gl.HomeH, _ = strconv.Atoi(data[50])
	gl.HomeDoubles, _ = strconv.Atoi(data[51])
	gl.HomeTriples, _ = strconv.Atoi(data[52])
	gl.HomeHR, _ = strconv.Atoi(data[53])
	gl.HomeRBI, _ = strconv.Atoi(data[54])
	gl.HomeSacHits, _ = strconv.Atoi(data[55])
	gl.HomeSacFlies, _ = strconv.Atoi(data[56])
	gl.HomeHBP, _ = strconv.Atoi(data[57])
	gl.HomeBB, _ = strconv.Atoi(data[58])
	gl.HomeIBB, _ = strconv.Atoi(data[59])
	gl.HomeK, _ = strconv.Atoi(data[60])
	gl.HomeSB, _ = strconv.Atoi(data[61])
	gl.HomeCS, _ = strconv.Atoi(data[62])
	gl.HomeGIDP, _ = strconv.Atoi(data[63])
	gl.HomeCatcherInterference, _ = strconv.Atoi(data[64])
	gl.HomeLOB, _ = strconv.Atoi(data[65])
	gl.HomePitchersUsed, _ = strconv.Atoi(data[66])
	gl.HomeIndividualER, _ = strconv.Atoi(data[67])
	gl.HomeTeamER, _ = strconv.Atoi(data[68])
	gl.HomeWP, _ = strconv.Atoi(data[69])
	gl.HomeBalks, _ = strconv.Atoi(data[70])
	gl.HomePO, _ = strconv.Atoi(data[71])
	gl.HomeA, _ = strconv.Atoi(data[72])
	gl.HomeE, _ = strconv.Atoi(data[73])
	gl.HomePB, _ = strconv.Atoi(data[74])
	gl.HomeDP, _ = strconv.Atoi(data[75])
	gl.HomeTriplePlays, _ = strconv.Atoi(data[76])
	gl.HomeUmpID = data[77]
	gl.HomeUmpName = data[78]
	gl.FirstBaseUmpID = data[79]
	gl.FirstBaseUmpName = data[80]
	gl.SecondBaseUmpID = data[81]
	gl.SecondBaseUmpName = data[82]
	gl.ThirdBaseUmpID = data[83]
	gl.ThirdBaseUmpName = data[84]
	gl.LFUmpID = data[85]
	gl.LFUmpName = data[86]
	gl.RFUmpID = data[87]
	gl.RFUmpName = data[88]
	gl.VisitorsMgrID = data[89]
	gl.VisitorsMgrName = data[90]
	gl.HomeMgrID = data[91]
	gl.HomeMgrName = data[92]
	gl.WinningPitcherID = data[93]
	gl.WinningPitcherName = data[94]
	gl.LosingPitcherID = data[95]
	gl.LosingPitcherName = data[96]
	gl.SavingPitcherID = data[97]
	gl.SavingPitcherName = data[98]
	gl.GameWinningRBIBatterID = data[99]
	gl.GameWinningRBIBatterName = data[100]
	gl.VisitorsStartingPitcherID = data[101]
	gl.VisitorsStartingPitcherName = data[102]
	gl.HomeStartingPitcherID = data[103]
	gl.HomeStartingPitcherName = data[104]
	gl.VisitorBatter1ID = data[105]
	gl.VisitorBatter1Name = data[106]
	gl.VisitorBatter1Pos, _ = strconv.Atoi(data[107])
	gl.VisitorBatter2ID = data[108]
	gl.VisitorBatter2Name = data[109]
	gl.VisitorBatter2Pos, _ = strconv.Atoi(data[110])
	gl.VisitorBatter3ID = data[111]
	gl.VisitorBatter3Name = data[112]
	gl.VisitorBatter3Pos, _ = strconv.Atoi(data[113])
	gl.VisitorBatter4ID = data[114]
	gl.VisitorBatter4Name = data[115]
	gl.VisitorBatter4Pos, _ = strconv.Atoi(data[116])
	gl.VisitorBatter5ID = data[117]
	gl.VisitorBatter5Name = data[118]
	gl.VisitorBatter5Pos, _ = strconv.Atoi(data[119])
	gl.VisitorBatter6ID = data[120]
	gl.VisitorBatter6Name = data[121]
	gl.VisitorBatter6Pos, _ = strconv.Atoi(data[122])
	gl.VisitorBatter7ID = data[123]
	gl.VisitorBatter7Name = data[124]
	gl.VisitorBatter7Pos, _ = strconv.Atoi(data[125])
	gl.VisitorBatter8ID = data[126]
	gl.VisitorBatter8Name = data[127]
	gl.VisitorBatter8Pos, _ = strconv.Atoi(data[128])
	gl.VisitorBatter9ID = data[129]
	gl.VisitorBatter9Name = data[130]
	gl.VisitorBatter9Pos, _ = strconv.Atoi(data[131])
	gl.HomeBatter1ID = data[132]
	gl.HomeBatter1Name = data[133]
	gl.HomeBatter1Pos, _ = strconv.Atoi(data[134])
	gl.HomeBatter2ID = data[135]
	gl.HomeBatter2Name = data[136]
	gl.HomeBatter2Pos, _ = strconv.Atoi(data[137])
	gl.HomeBatter3ID = data[138]
	gl.HomeBatter3Name = data[139]
	gl.HomeBatter3Pos, _ = strconv.Atoi(data[140])
	gl.HomeBatter4ID = data[141]
	gl.HomeBatter4Name = data[142]
	gl.HomeBatter4Pos, _ = strconv.Atoi(data[143])
	gl.HomeBatter5ID = data[144]
	gl.HomeBatter5Name = data[145]
	gl.HomeBatter5Pos, _ = strconv.Atoi(data[146])
	gl.HomeBatter6ID = data[147]
	gl.HomeBatter6Name = data[148]
	gl.HomeBatter6Pos, _ = strconv.Atoi(data[149])
	gl.HomeBatter7ID = data[150]
	gl.HomeBatter7Name = data[151]
	gl.HomeBatter7Pos, _ = strconv.Atoi(data[152])
	gl.HomeBatter8ID = data[153]
	gl.HomeBatter8Name = data[154]
	gl.HomeBatter8Pos, _ = strconv.Atoi(data[155])
	gl.HomeBatter9ID = data[156]
	gl.HomeBatter9Name = data[157]
	gl.HomeBatter9Pos, _ = strconv.Atoi(data[158])
	gl.AdditionalInformation = data[159]
	gl.AcquisitionInformation = data[160]

	return gl, nil
}
