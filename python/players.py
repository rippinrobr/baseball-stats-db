import utils

# this should get moved into a settings file/object 
player_files = [
        utils.base_dir+"/Master.csv", 
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
        self.batting_stats = []
        self.fielding_stats = []
        self.pitching_stats = []

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
        if stat_prop == "demographics":
            self.add_demographics(stats)

def parse():
    players= dict() 
 
    utils.process_file(player_files[0], players, "demographics", "playerID", RawPlayer)
    
    return players

players =  parse()


for p in players:
    print players[p].name_last + ", " + players[p].name_first
