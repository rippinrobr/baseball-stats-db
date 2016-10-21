import utils

# this should get moved into a settings file/object 
player_files = [
        utils.base_dir+"/Master.csv", 
        utils.base_dir+"/AllstarFull.csv", 
        utils.base_dir+"/Appearances.csv", 
        utils.base_dir+"/Batting.csv", 
        utils.base_dir+"/Fielding.csv", 
        utils.base_dir+"/Pitching.csv", 
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
        self.batting_stats = []
        self.fielding_stats = []
        self.pitching_stats = []
        self.playoff_batting_stats = []
        self.playoff_fielding_stats = []
        self.playoff_pitching_stats = []

    def add_allstar_appearances(self, appearance):
        self.allstar_appearances.append( appearance )

    def add_appearances(self, appearance):
        self.appearances.append( appearance )

    def add_batting_stats(self, stats):
        self.batting_stats.append( stats )

    def add_fielding_stats(self, stats):
        self.fielding_stats.append( stats )

    def add_pitching_stats(self, stats):
        self.pitching_stats.append( stats )

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

    def add_stats(self, stat_prop, stats):
        if stat_prop == "allstars":
            self.add_allstar_appearances(stats)

        if stat_prop == "appearances":
            self.add_appearances(stats)

        if stat_prop == "demographics":
            self.add_demographics(stats)

        if stat_prop == "batting":
            self.add_batting_stats(stats)

        if stat_prop == "fielding":
            self.add_fielding_stats(stats)

        if stat_prop == "pitching":
            self.add_pitching_stats(stats)

def parse():
    players= dict() 
 
    utils.process_file(player_files[0], players, "demographics", "playerID", RawPlayer)
    utils.process_file(player_files[1], players, "allstars", "playerID", RawPlayer)
    utils.process_file(player_files[2], players, "appearances", "playerID", RawPlayer)
    utils.process_file(player_files[3], players, "batting", "playerID", RawPlayer)
    utils.process_file(player_files[4], players, "fielding", "playerID", RawPlayer)
    utils.process_file(player_files[5], players, "pitching", "playerID", RawPlayer)
    
    return players

players =  parse()


#for p in players:
    #print players[p].appearances
