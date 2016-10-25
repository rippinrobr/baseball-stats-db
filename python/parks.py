import utils

park_files = [
        utils.base_dir + "/Parks.csv",
        ]

class Park:
    def __init__(self, parkID):
        self.parkID = parkID
        self.info = {}
        self.home_games = []

    def add_stats(self, stats, stat_prop):

        if stat_prop == "home_games":
            self.home_games.append( stats )

        if stat_prop == "info":
            self.info = stats

def parse():
    parks = dict()

    utils.process_file(park_files[0], parks, "info","park.key", Park),

    #print parks
    return parks

parse()
