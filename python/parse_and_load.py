import bdb_schema
import os
import utils

BASE_DIR = os.getenv("BDB_BASE_DIR", "/home/rob/src/baseballdatabank/core")
FILES = ["/AllstarFull.csv", "/Appearances.csv", "/AwardsManagers.csv", "/AwardsPlayers.csv",
         "/AwardsShareManagers.csv", "/AwardsSharePlayers.csv", "/Batting.csv", "/BattingPost.csv",
         "/CollegePlaying.csv", "/Fielding.csv", "/FieldingOF.csv", "/FieldingPost.csv",
         "/HallOfFame.csv", "/HomeGames.csv", "/Managers.csv", "/ManagersHalf.csv", "/Master.csv",
         "/Parks.csv", "/Pitching.csv", "/PitchingPost.csv", "/Salaries.csv", "/Schools.csv",
         "/SeriesPost.csv", "/Teams.csv", "/TeamsFranchises.csv", "/TeamsHalf.csv"]

tables = [
    'Master', 'TeamsFranchises', 'AllstarFull', 'Appearances', 'AwardsManagers', 
    'AwardsPlayers', 'AwardsShareManagers', 'AwardsSharePlayers', 'Batting', 
    'BattingPost', 'CollegePlaying', 'Fielding', 'FieldingOF', 'FieldingPost', 
    'HallOfFame', 'HomeGames', 'Managers', 'ManagersHalf', 'Parks', 'Pitching', 
    'PitchingPost', 'Salaries', 'Schools', 'SeriesPost', 'Teams', 'TeamsHalf'
]

TABLE_DATA = {}
TABLE_SCHEMAS = {}

PARSER = utils.Parser()
#SCHEMA_GENERATOR = utils.SchemaGenerator()

for csv_file in FILES:
    data, schema = PARSER.parse_file(BASE_DIR + csv_file)
    #print schema['table_name']
    #TABLE_SCHEMAS[schema['table_name']] = schema
    #print schema['table_name'], " ", schema
    #TABLE_DATA[schema['table_name']] = data

# bdb_schema.Create_Tables()
