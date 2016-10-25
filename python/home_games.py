import utils

class Park:
    def __init__(self, parkID, info):
        self.parkID = parkID
        self.info = info
        self.home_games = []

    def add_stats(self, home_games):
        self.home_games.append( home_games)
