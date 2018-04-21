package retrosheet

type pitching struct {
	WinningPitcherID            string `json:"winningPitcherId" bson:"winningPitcherId" db:"winning_pitcher_id"`
	WinningPitcherName          string `json:"winningPitcherName" bson:"winningPitcherName" db:"winning_pitcher_name"`
	LosingPitcherID             string `json:"losingPitcherId" bson:"losingPitcherId" db:"losing_pitcher_id"`
	LosingPitcherName           string `json:"losingPitcherName" bson:"losingPitcherName" db:"losing_pitcher_name"`
	SavingPitcherID             string `json:"savingPitcherId" bson:"savingPitcherId" db:"saving_pitcher_id"`
	SavingPitcherName           string `json:"savingPitcherName" bson:"savingPitcherName" db:"saving_pitcher_name"`
	VisitorsStartingPitcherID   string `json:"visitorsStartingPitcherId" bson:"visitorsStartingPitcherId" db:"visitors_starting_pitcher_id"`
	VisitorsStartingPitcherName string `json:"visitorsStartingPitcherName" bson:"visitorsStartingPitcherName" db:"visitors_starting_pitcher_name"`
	HomeStartingPitcherID       string `json:"homeStartingPitcherId" bson:"homeStartingPitcherId" db:"home_starting_pitcher_id"`
	HomeStartingPitcherName     string `json:"homeStartingPitcherName" bson:"homeStartingPitcherName" db:"home_starting_pitcher_name"`
}
