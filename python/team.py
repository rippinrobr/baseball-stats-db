
class RawTeam:
    def __init__(self, teamId):
        self.teamID = teamId 
        
        self.half_seasons = []
        self.names = []
        self.playoffs = []
        self.seasons = []

    def add_half_season(self, half):
        self.half_seasons.append(half)

    def add_name(self, name):
        self.names.append(name)

    def add_season(self, season):
        self.seasons.append(season)

    def add_playoffs(self, playoffs):
        self.playoffs.append(playoffs)
