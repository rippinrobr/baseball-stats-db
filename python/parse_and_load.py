import utils

table_schemas = {}
raw_all_stars, schema = utils.new_process_file(utils.BASE_DIR + "/AllStarFull.csv")
table_schemas[schema['table_name']] = schema

raw_appearances, schema = utils.new_process_file(utils.BASE_DIR + "/Appearances.csv")
table_schemas[schema['table_name']] = schema

raw_awards_managers, schema = utils.new_process_file(utils.BASE_DIR + "/AwardsManagers.csv")
table_schemas[schema['table_name']] = schema

raw_awards_players, schema = utils.new_process_file(utils.BASE_DIR + "/AwardsPlayers.csv")
table_schemas[schema['table_name']] = schema

raw_awards_share_managers, schema = utils.new_process_file(utils.BASE_DIR + "/AwardsShareManagers.csv")
table_schemas[schema['table_name']] = schema

raw_awards_share_players, schema = utils.new_process_file(utils.BASE_DIR + "/AwardsSharePlayers.csv")
table_schemas[schema['table_name']] = schema

raw_batting, schema = utils.new_process_file(utils.BASE_DIR + "/Batting.csv")
table_schemas[schema['table_name']] = schema

raw_batting_post, schema = utils.new_process_file(utils.BASE_DIR + "/BattingPost.csv")
table_schemas[schema['table_name']] = schema

raw_college_playing, schema = utils.new_process_file(utils.BASE_DIR + "/CollegePlaying.csv")
table_schemas[schema['table_name']] = schema

raw_fielding, schema = utils.new_process_file(utils.BASE_DIR + "/Fielding.csv")
table_schemas[schema['table_name']] = schema

raw_fielding_of, schema = utils.new_process_file(utils.BASE_DIR + "/FieldingOF.csv")
table_schemas[schema['table_name']] = schema

raw_fielding_post, schema = utils.new_process_file(utils.BASE_DIR + "/FieldingPost.csv")
table_schemas[schema['table_name']] = schema

raw_hof, schema = utils.new_process_file(utils.BASE_DIR + "/HallOfFame.csv")
table_schemas[schema['table_name']] = schema

raw_home_games, schema = utils.new_process_file(utils.BASE_DIR + "/HomeGames.csv")
table_schemas[schema['table_name']] = schema

raw_managers, schema = utils.new_process_file(utils.BASE_DIR + "/Managers.csv")
table_schemas[schema['table_name']] = schema

raw_managers_half, schema = utils.new_process_file(utils.BASE_DIR + "/ManagersHalf.csv")
table_schemas[schema['table_name']] = schema

raw_master, schema = utils.new_process_file(utils.BASE_DIR + "/Master.csv")
table_schemas[schema['table_name']] = schema

raw_parks, schema = utils.new_process_file(utils.BASE_DIR + "/Parks.csv")
table_schemas[schema['table_name']] = schema

raw_pitching, schema = utils.new_process_file(utils.BASE_DIR + "/Pitching.csv")
table_schemas[schema['table_name']] = schema

raw_pitching_post, schema = utils.new_process_file(utils.BASE_DIR + "/PitchingPost.csv")
table_schemas[schema['table_name']] = schema

raw_salaries, schema = utils.new_process_file(utils.BASE_DIR + "/Salaries.csv")
table_schemas[schema['table_name']] = schema

raw_schools, schema = utils.new_process_file(utils.BASE_DIR + "/Schools.csv")
table_schemas[schema['table_name']] = schema

raw_series_post, schema = utils.new_process_file(utils.BASE_DIR + "/SeriesPost.csv")
table_schemas[schema['table_name']] = schema

raw_teams, schema = utils.new_process_file(utils.BASE_DIR + "/Teams.csv")
table_schemas[schema['table_name']] = schema

raw_team_franchises, schema = utils.new_process_file(utils.BASE_DIR + "/TeamsFranchises.csv")
table_schemas[schema['table_name']] = schema

raw_teams_half, schema = utils.new_process_file(utils.BASE_DIR + "/TeamsHalf.csv")
table_schemas[schema['table_name']] = schema

utils.write_sql_schema(table_schemas, utils.SQLITE)