from sql_schema import *

SQLITE = "sqlite"
def define_parameters(parser):
    parser.add_argument("--dbtype", choices=[SQLITE], help="the database type you'd like to generate the schema for", type=str, default="sqlite" )
    parser.add_argument("--dbpath", help="SQLITE ONLY - the path for the newly created database", type=str)

def main():
    param_parser = argparse.ArgumentParser(description="Generates a DB schema based on the Baseball Databank csv files.")    
    define_parameters(param_parser)
    args = param_parser.parse_args()

    bdb_db = ""
    sqlite_db_name = "baseball_databank.sqlite3"
    if args.dbpath != None and args.dbpath != "":
        if args.dbtype == SQLITE:
            sqlite_db_name = args.dbpath 
        else:
            print "The argument --dbpath is only needed for Sqlite databases, ignoring."

    if args.dbtype == SQLITE:
        bdb_db = SqliteDatabase(sqlite_db_name)

    initialize_db_and_connect(bdb_db)
    
if __name__ == "__main__":
    main()