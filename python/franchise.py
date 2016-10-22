class RawFranchise:
    def __init__(self, franchId):
        self.franchID = franchId 
        # this will have a key of the team name and 
        # each season will be an array of seasons
        self.half_seasons = [] 
        self.seasons =  dict()
        self.playoffs =  dict()
        self.teams = dict() 

    def add_half_season(self, teamId, half):
        self.half_seasons.append( half)

    def add_team(self, teamId, name):
        self.teams[ teamId ] = name 
   
    def add_season(self, teamId, season):
        if teamId not in self.seasons:
            self.seasons[ teamId ] = []

        self.seasons[ teamId ].append( season )
   
    def add_playoffs(self, teamId, season):
        if teamId not in self.playoffs:
            self.playoffs[ teamId ] = []

        self.playoffs[ teamId ].append( season )
