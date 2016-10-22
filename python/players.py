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
]

class RawPlayer:
    " A representation of a players's stats and demographics as read from CSV file """
    def __init__(self, playerId):
        self.playerID = playerId
        self.birth_country = ""
        self.birth_state = ""
        self.birth_city = ""
        self.death_country = ""
        self.death_state = ""
        self.death_city = ""
        self.name_first = ""
        self.name_last = ""
        self.name_note = ""
        self.name_given = ""
        self.name_nick = ""
        self.bats = ""
        self.throws = ""
        self.debut = ""
        self.final_game = ""
        self.college = ""
        self.lahman40ID = ""
        self.lahman45ID = ""
        self.retroID = ""
        self.holtzID = ""
        self.bbrefID = ""
        
        self.allstar_appearances = []
        self.appearances =  []
        self.awards = []
        self.batting_stats = []
        self.fielding_stats = []
        self.fielding_of_stats = []
        self.pitching_stats = []
        self.playoff_batting_stats = []
        self.playoff_fielding_stats = []
        self.playoff_pitching_stats = []
        self.salaries = []

    def add_allstar_appearances(self, appearance):
        self.allstar_appearances.append( appearance )

    def add_appearances(self, appearance):
        self.appearances.append( appearance )

    def add_awards(self, awards):
        self.awards.append( awards )

    def add_batting_stats(self, stats):
        self.batting_stats.append( stats )

    def add_demographics(self, demog):
        self.birth_country =  demog["birthCountry"]
        self.birth_state = demog["birthState"]
        self.birth_city = demog["birthCity"]
        self.death_country = demog["deathCountry"]
        self.death_state = demog["deathCountry"]
        self.death_city = demog["deathCountry"]
        self.name_first = demog["nameFirst"]
        self.name_last = demog["nameLast"]
        self.name_given = demog["nameGiven"]
        self.weight = demog["weight"]
        self.height = demog["height"]
        self.bats = demog["bats"]
        self.throws = demog["throws"] 
        self.debut = demog["debut"]
        self.final_game = demog["finalGame"]
        self.retroID = demog["retroID"]
        self.bbrefID = demog["bbrefID"]

    def add_fielding_stats(self, stats):
        self.fielding_stats.append( stats )

    def add_fielding_of_stats(self, stats):
        self.fielding_of_stats.append( stats )

    def add_pitching_stats(self, stats):
        self.pitching_stats.append( stats )

    def add_playoff_batting_stats(self, stats):
        self.playoff_batting_stats.append( stats )

    def add_playoff_fielding_stats(self, stats):
        self.playoff_fielding_stats.append( stats )

    def add_playoff_pitching_stats(self, stats):
        self.playoff_pitching_stats.append( stats )

    def add_salaries(self, stats):
        self.salaries.append(stats)

    def add_stats(self, stat_prop, stats):
        if stat_prop == "allstars":
            self.add_allstar_appearances(stats)

        if stat_prop == "appearances":
            self.add_appearances(stats)

        if stat_prop == "awards":
            self.add_awards(stats)

        if stat_prop == "demographics":
            self.add_demographics(stats)

        if stat_prop == "batting":
            self.add_batting_stats(stats)

        if stat_prop == "fielding":
            self.add_fielding_stats(stats)

        if stat_prop == "fielding-OF":
            self.add_fielding_of_stats(stats)

        if stat_prop == "pitching":
            self.add_pitching_stats(stats)

        if stat_prop == "playoff_batting":
            self.add_playoff_batting_stats(stats)

        if stat_prop == "playoff_fielding":
            self.add_playoff_fielding_stats(stats)

        if stat_prop == "playoff_pitching":
            self.add_playoff_pitching_stats(stats)

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
    utils.process_file(player_files[6], players, "playoff_pitching", "playerID", RawPlayer)
    utils.process_file(player_files[7], players, "playoff_fielding", "playerID", RawPlayer)
    utils.process_file(player_files[8], players, "playoff_pitching", "playerID", RawPlayer)
    utils.process_file(player_files[9], players, "fielding-OF", "playerID", RawPlayer)
    utils.process_file(player_files[10], players, "awards", "playerID", RawPlayer)
    utils.process_file(player_files[11], players, "salaries", "playerID", RawPlayer)

    return players

players =  parse()


#for p in players:
    #print players[p].appearances
