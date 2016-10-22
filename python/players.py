import utils

# this should get moved into a settings file/object 
player_files = [
        utils.base_dir+"/Master.csv", 
        utils.base_dir+"/AllstarFull.csv", 
        utils.base_dir+"/Appearances.csv", 
        utils.base_dir+"/Batting.csv", 
        utils.base_dir+"/Fielding.csv", 
        utils.base_dir+"/Pitching.csv", 
        utils.base_dir+"/BattingPost.csv", 
        utils.base_dir+"/FieldingPost.csv", 
        utils.base_dir+"/PitchingPost.csv", 
        utils.base_dir+"/FieldingOF.csv", 
        utils.base_dir+"/AwardsPlayers.csv", 
        utils.base_dir+"/Salaries.csv", 
        utils.base_dir+"/CollegePlaying.csv",
        utils.base_dir+"/HallOfFame.csv",
]

class RawPlayer:
    " A representation of a players's stats and demographics as read from CSV file """
    def __init__(self, playerId):
        self.playerID = playerId
        self.birthCountry = ""
        self.birthState = ""
        self.birthCity = ""
        self.deathCountry = ""
        self.deathState = ""
        self.deathCity = ""
        self.nameFirst = ""
        self.nameLast = ""
        self.nameGiven = ""
        self.bats = ""
        self.throws = ""
        self.debut = ""
        self.finalGame = ""
        self.college = ""
        self.lahman40ID = ""
        self.lahman45ID = ""
        self.retroID = ""
        self.holtzID = ""
        self.bbrefID = ""
        
        self.allstarAppearances = []
        self.appearances =  []
        self.awards = []
        self.battingStats = []
        self.college = []
        self.fieldingStats = []
        self.fieldingOfStats = []
        self.hallOfFame = []
        self.pitchingStats = []
        self.playoffBattingStats = []
        self.playoffFieldingStats = []
        self.playoffPitchingStats = []
        self.salaries = []

    def add_allstarAppearances(self, appearance):
        self.allstarAppearances.append( appearance )

    def add_appearances(self, appearance):
        self.appearances.append( appearance )

    def add_awards(self, awards):
        self.awards.append( awards )

    def add_battingStats(self, stats):
        self.battingStats.append( stats )

    def add_college(self, stats):
        self.college.append( stats )

    def add_demographics(self, demog):
        self.birthCountry =  demog["birthCountry"]
        self.birthState = demog["birthState"]
        self.birthCity = demog["birthCity"]
        self.deathCountry = demog["deathCountry"]
        self.deathState = demog["deathCountry"]
        self.deathCity = demog["deathCountry"]
        self.nameFirst = demog["nameFirst"]
        self.nameLast = demog["nameLast"]
        self.nameGiven = demog["nameGiven"]
        self.weight = demog["weight"]
        self.height = demog["height"]
        self.bats = demog["bats"]
        self.throws = demog["throws"] 
        self.debut = demog["debut"]
        self.finalGame = demog["finalGame"]
        self.retroID = demog["retroID"]
        self.bbrefID = demog["bbrefID"]

    def add_fieldingStats(self, stats):
        self.fieldingStats.append( stats )

    def add_fieldingOfStats(self, stats):
        self.fieldingOfStats.append( stats )

    def add_hallOfFame(self, stats):
        self.hallOfFame.append(stats)

    def add_pitchingStats(self, stats):
        self.pitchingStats.append( stats )

    def add_playoffBattingStats(self, stats):
        self.playoffBattingStats.append( stats )

    def add_playoffFieldingStats(self, stats):
        self.playoffFieldingStats.append( stats )

    def add_playoffPitchingStats(self, stats):
        self.playoffPitchingStats.append( stats )

    def add_salaries(self, stats):
        self.salaries.append(stats)

    def add_stats(self, stat_prop, stats):
        if stat_prop == "allstars":
            self.add_allstarAppearances(stats)

        if stat_prop == "appearances":
            self.add_appearances(stats)

        if stat_prop == "awards":
            self.add_awards(stats)

        if stat_prop == "demographics":
            self.add_demographics(stats)

        if stat_prop == "batting":
            self.add_battingStats(stats)

        if stat_prop == "college":
            self.add_college(stats)

        if stat_prop == "fielding":
            self.add_fieldingStats(stats)

        if stat_prop == "fielding-OF":
            self.add_fieldingOfStats(stats)

        if stat_prop == "hallOfFame":
            self.add_hallOfFame(stats)

        if stat_prop == "pitching":
            self.add_pitchingStats(stats)

        if stat_prop == "playoffBatting":
            self.add_playoffBattingStats(stats)

        if stat_prop == "playoffFielding":
            self.add_playoffFieldingStats(stats)

        if stat_prop == "playoffPitching":
            self.add_playoffPitchingStats(stats)

        if stat_prop == "salaries":
            self.add_salaries(stats)

def parse():
    players= dict() 
 
    utils.process_file(player_files[0], players, "demographics", "playerID", RawPlayer)
    utils.process_file(player_files[1], players, "allstars", "playerID", RawPlayer)
    utils.process_file(player_files[2], players, "appearances", "playerID", RawPlayer)
    utils.process_file(player_files[3], players, "batting", "playerID", RawPlayer)
    utils.process_file(player_files[4], players, "fielding", "playerID", RawPlayer)
    utils.process_file(player_files[5], players, "pitching", "playerID", RawPlayer)
    utils.process_file(player_files[6], players, "playoffPitching", "playerID", RawPlayer)
    utils.process_file(player_files[7], players, "playoffFielding", "playerID", RawPlayer)
    utils.process_file(player_files[8], players, "playoffPitching", "playerID", RawPlayer)
    utils.process_file(player_files[9], players, "fielding-OF", "playerID", RawPlayer)
    utils.process_file(player_files[10], players, "awards", "playerID", RawPlayer)
    utils.process_file(player_files[11], players, "salaries", "playerID", RawPlayer)
    utils.process_file(player_files[12], players, "college", "playerID", RawPlayer)
    utils.process_file(player_files[13], players, "hallOfFame", "playerID", RawPlayer)

    return players
