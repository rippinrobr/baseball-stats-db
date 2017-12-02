import os

player_id_tables = [
    "People", "AllstarFull", "Appearances", "AwardsManagers", "AwardsPlayers",
    "AwardsShareManagers", "AwardsSharePlayers", "Batting", "BattingPost", 
    "CollegePlaying", "Fielding", "FieldingOF", "FieldingPost", "FieldingOFsplit",
    "HallOfFame", "Managers", "ManagersHalf", "Pitching", "PitchingPost", "Salaries"
]

park_key_tables = ["Parks"]
school_id_tables = ["Schools"]
team_franchises_tables = [ "TeamsFranchises"]
team_id_tables = [ "Teams", "TeamsHalf"]
team_key_tables = [ "HomeGames"]

def change_col_name_for_easier_sql(col_name):
    if col_name == "2B":
        return "doubles"
    if col_name == "3B":
        return "triples"
    if col_name == "yearid":
        return "yearID"
    
    return col_name

def col_name_cleaner(col_name):
    col_name = change_col_name_for_easier_sql(col_name)
    return col_name.replace(".", "").replace("_","")

def convert_to_lang_specific_type(typeMap, type):
    if type in typeMap:
        return typeMap[type]

    if type == "":
        return "string"

    return type

def get_data_struct_name(args):
   table_name = os.path.splitext(get_file_name(args.input))[0] 
   if args.name != None:
        return args.name, table_name.lower()
   else:
        return table_name, table_name.lower()

def get_file_name(file_path):
    return os.path.basename(file_path)

def mongo_key_checks(struct_name, col_name):
    if struct_name in player_id_tables and col_name == "playerID":
        return "_id"

    if struct_name in team_id_tables and col_name == "teamID":
        return "_id"
    
    if struct_name in school_id_tables and col_name == "schoolID":
        return "_id"
    
    if struct_name in team_franchises_tables and col_name == "franchID":
        return "_id"

    if struct_name in team_key_tables and col_name == "teamkey":
        return "_id"

    if struct_name in park_key_tables and col_name == "parkkey":
        return "_id"

    return col_name