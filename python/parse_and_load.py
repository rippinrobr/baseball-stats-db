import utils
import players
import season

raw_all_stars = utils.new_process_file(utils.base_dir + "/AllStarFull.csv")
raw_appearances = utils.new_process_file(utils.base_dir + "/Appearances.csv")
raw_awards_managers = utils.new_process_file(utils.base_dir + "/AwardsManagers.csv")
raw_awards_players = utils.new_process_file(utils.base_dir + "/AwardsPlayers.csv")
raw_awards_share_managers = utils.new_process_file(utils.base_dir + "/AwardsShareManagers.csv")
raw_awards_share_players = utils.new_process_file(utils.base_dir + "/AwardsSharePlayers.csv")
print "raw_awards_share_players[0]: ", raw_awards_share_players[0]
raw_batting = utils.new_process_file(utils.base_dir + "/Batting.csv")
raw_batting_post = utils.new_process_file(utils.base_dir + "/BattingPost.csv")
raw_college_playing = utils.new_process_file(utils.base_dir + "/CollegePlaying.csv")
raw_fielding = utils.new_process_file(utils.base_dir + "/Fielding.csv")
raw_fielding_of = utils.new_process_file(utils.base_dir + "/FieldingOF.csv")
raw_fielding_post = utils.new_process_file(utils.base_dir + "/FieldingPost.csv")
raw_hof = utils.new_process_file(utils.base_dir + "/HallOfFame.csv")
raw_home_games = utils.new_process_file(utils.base_dir + "/HomeGames.csv")
raw_managers = utils.new_process_file(utils.base_dir + "/Managers.csv")
raw_managers_half = utils.new_process_file(utils.base_dir + "/ManagersHalf.csv")
raw_master = utils.new_process_file(utils.base_dir + "/Master.csv")
raw_parks = utils.new_process_file(utils.base_dir + "/Parks.csv")
raw_pitching = utils.new_process_file(utils.base_dir + "/Pitching.csv")
raw_pitching_post = utils.new_process_file(utils.base_dir + "/PitchingPost.csv")
raw_salaries = utils.new_process_file(utils.base_dir + "/Salaries.csv")
raw_schools = utils.new_process_file(utils.base_dir + "/Schools.csv")
print "raw_schools[0]: ", raw_schools[0]
raw_series_post = utils.new_process_file(utils.base_dir + "/SeriesPost.csv")
raw_teams = utils.new_process_file(utils.base_dir + "/Teams.csv")
raw_teams_half = utils.new_process_file(utils.base_dir + "/TeamsHalf.csv")
raw_team_franchises = utils.new_process_file(utils.base_dir + "/TeamsFranchises.csv")
print "raw_team_franchises[0]: ", raw_team_franchises[0]

