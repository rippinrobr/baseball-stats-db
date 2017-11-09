from sql_schema import *

SQLITE = "sqlite"
def define_parameters(parser):
    parser.add_argument("--dbtype", choices=[SQLITE], help="the database type you'd like to generate the schema for", type=str, default="sqlite" )

def main():
    param_parser = argparse.ArgumentParser(description="Generates a DB schema based on the Baseball Databank csv files.")    
    define_parameters(param_parser)
    args = param_parser.parse_args()

    bdb_db = ""
    if args.dbtype == SQLITE:
        bdb_db = SqliteDatabase('db/bdb_v0.0.1.sqlite')

    initialize_db_and_connect(bdb_db)
    
if __name__ == "__main__":
    main()