import csv
import sys

import franchise
import team
import utils

team_files = [
        utils.base_dir + "/Teams.csv",
        utils.base_dir + "/TeamsHalf.csv", # 1981 strike
        utils.base_dir + "/SeriesPost.csv",
]

# file is Teams.csv
class RawSeason:
    """ A representation of a team's performance for a season. """
    def __init__(self, yearId, teamId):
        self.ID = yearId + "_" + teamId
        self.yearID = yearId
        self.teamID = teamId
        self.lgID = ""
        self.franchID = ""
        self.name = ""
        self.park = ""
        self.teamIDBR = ""
        self.teamIDlahman45 = ""
        self.teamIDretro = ""

        self.playoff_stats = [] 
        self.stats = dict()

    def add_playoffs(self, stats):
        self.playoff_stats.append( stats )

    def add_stats(self, stats):
        self.lgID = stats.pop("yearID",None)
        self.franchID = stats.pop("franchID",None)
        self.name = stats.pop("name", None)
        self.park = stats.pop("park", None)
        self.teamIDBR = stats.pop("teamIDBR", None)
        self.teamIDlahman45 = stats.pop("teamIDlahman45", None)
        self.teamIDretro = stats.pop("teamIDretro", None)

        self.stats = stats
    print "Done!"


def process_season_file(file_path, seasons_list, teams_list, franchise_list, stat_prop):
    print "processing: " + file_path

    with open(file_path)  as csvfile:
        is_header_row = True
        csv_headers = []

        reader = csv.reader(csvfile, delimiter=',')
        counter = 0
        for season in reader:
            if is_header_row:
                csv_headers = season
                is_header_row = False
                continue

            if counter != 0:
                if counter % 40 == 0:
                    print ""
                sys.stdout.write(".")
                sys.stdout.flush()
           
            stats_dict = utils.list_to_dict(csv_headers, season)
            # set up the RawSeason object and load it in the 
            # list
            yearId  = stats_dict.pop("yearID", None) 
            teamId  = stats_dict.pop("teamID", None) 
            seasonKey = yearId + "_"  + teamId
            if seasonKey:
                curr_season = RawSeason(yearId, teamId)
                curr_season.add_stats(stats_dict)
                seasons_list[seasonKey] = curr_season

                if curr_season.teamID not in teams_list:
                    teams_list[ curr_season.teamID ] = team.RawTeam(teamId)
                    
                if curr_season.franchID not in franchise_list:
                    franchise_list[ curr_season.franchID ] = franchise.RawFranchise( curr_season.franchID )
                
                if curr_season.name not in franchise_list[curr_season.franchID].teams:
                    franchise_list[ curr_season.franchID ].add_team( teamId, curr_season.name)

                if stat_prop == "half_season":
                    teams_list[ curr_season.teamID ].add_half_season( curr_season )
                    franchise_list[ curr_season.franchID ].add_half_season( teamId, curr_season )

                if stat_prop == "season":
                    teams_list[ curr_season.teamID ].add_season( curr_season )
                    franchise_list[ curr_season.franchID ].add_season( teamId, curr_season )

            counter = counter + 1
    print "Done!"

def process_playoff_file(file_path, seasons_list, teams_list):
    print "processing: " + file_path

    with open(file_path)  as csvfile:
        is_header_row = True
        csv_headers = []

        reader = csv.reader(csvfile, delimiter=',')
        counter = 0
        for season in reader:
            if is_header_row:
                csv_headers = season
                is_header_row = False
                continue

            if counter != 0:
                if counter % 40 == 0:
                    print ""
                sys.stdout.write(".")
                sys.stdout.flush()
           
            stats_dict = utils.list_to_dict(csv_headers, season)
            # set up the RawSeason object and load it in the 
            # list
            yearId  = stats_dict.pop("yearID", None) 
            winningTeamId  = stats_dict.pop("teamIDwinner", None) 
            #
            # THIS NEEDS TO BE FUNC THAT TAKES the yearID and teamId
            # if STL it needs to see what season it is to determine 
            # the appropriate teamId
            if winningTeamId == "NYP":
                winningTeamId = "NY4"

            if winningTeamId == "CHC":
                winningTeamId = "CHN"

            if winningTeamId == "STL":
                winningTeamId = "SL2"

            losingTeamId  = stats_dict.pop("teamIDloser", None) 
            if losingTeamId == "NYP":
                losingTeamId = "NY4"

            if losingTeamId == "CHC":
                losingTeamId = "CHN"

            if losingTeamId == "STL":
                losingTeamId = "SL2"

            winnerKey= yearId + "_"  + winningTeamId
            losingKey= yearId + "_"  + losingTeamId
            if winningTeamId and losingTeamId:
                seasons_list[winnerKey].add_playoffs( stats_dict )
                teams_list[ winningTeamId ].add_playoffs( stats_dict )

                seasons_list[losingKey].add_playoffs( stats_dict )
                teams_list[ losingTeamId ].add_playoffs( stats_dict )

            counter = counter + 1
    print "Done!"

# Run here
franchises = dict()
seasons = dict()
teams = dict()

process_season_file(team_files[0], seasons, teams, franchises, "season")
process_season_file(team_files[1], seasons, teams, franchises, "half_season")
process_playoff_file(team_files[2], seasons, teams)
