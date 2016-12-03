import os
import utils

BASE_DIR = os.getenv("BDB_BASE_DIR", "/home/rob/src/baseballdatabank/core")
FILES = ["/AllstarFull.csv", "/Appearances.csv", "/AwardsManagers.csv", "/AwardsPlayers.csv",
         "/AwardsShareManagers.csv", "/AwardsSharePlayers.csv", "/Batting.csv", "/BattingPost.csv",
         "/CollegePlaying.csv", "/Fielding.csv", "/FieldingOF.csv", "/FieldingPost.csv",
         "/HallOfFame.csv", "/HomeGames.csv", "/Managers.csv", "/ManagersHalf.csv", "/Master.csv",
         "/Parks.csv", "/Pitching.csv", "/PitchingPost.csv", "/Salaries.csv", "/Schools.csv",
         "/SeriesPost.csv", "/Teams.csv", "/TeamsFranchises.csv", "/TeamsHalf.csv"]

PARSER = utils.Parser()
SCHEMA_GENERATOR = utils.SchemaGenerator()

TABLE_DATA = {}
TABLE_SCHEMAS = {}

for csv_file in FILES:
    data, schema = PARSER.parse_file(BASE_DIR + csv_file)
    TABLE_SCHEMAS[schema['table_name']] = schema
    print schema['table_name'], " ", schema
    TABLE_DATA[schema['table_name']] = data

SCHEMA_GENERATOR.write_sql_schema(TABLE_SCHEMAS, utils.SQLITE)
