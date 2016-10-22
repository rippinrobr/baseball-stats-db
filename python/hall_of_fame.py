import utils

hall_of_fame_files = [
            utils.base_dir + "/HallOfFame.csv"
        ]

class RawHallOfFame:
    def __init__(self, playerId):
        self.playerID = playerId
        self.inductions = []

    def add_stats(self, stat_prop, data):
        self.inductions.append(data)


def parse():
    hoFamers = dict() 
    
    utils.process_file(hall_of_fame_files[0], hoFamers, "", "playerID", RawHallOfFame)
