import utils

home_game_files = [
        utils.base_dir + "/HomeGames.csv",
]

class HomeGame:
    def __init__(self, id):
        self.id = id
        self.data =  {}

    def __str__(self):
        return "[id: " + self.id + "]"

    def add_stats(self, prop, stats):
        print "add_stats"
        print stats
        self.data = stats
        print self.data


def parse():
    games = {} 

    utils.process_file_with_composite_key( home_game_files[0], games, ["year.key", "team.key", "park.key"], HomeGame)
  
    return games

