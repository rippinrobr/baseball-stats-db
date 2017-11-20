from sql_schema import *
import sys

MYSQL = "mysql"
POSTGRES = "postgres"
SQLITE = "sqlite"

def define_parameters(parser):
    parser.add_argument("--dbtype", choices=[SQLITE,POSTGRES,MYSQL], help="the database type you'd like to generate the schema for", type=str, required=True)
    parser.add_argument("--dbhost", help="host of the database server", default="localhost", type=str)
    parser.add_argument("--dbname", help="Name of the database where the tables are to be added. REQUIRED if not sqlite", type=str)
    parser.add_argument("--dbpath", help="SQLITE ONLY - the path for the newly created database", type=str)
    parser.add_argument("--dbpass", help="The password for the user given in the --dbuser option, ignored for SQLite", type=str)
    parser.add_argument("--dbport", help="The port the database servier is listeing on, ignored for SQLite, defaults to appropriate value for server type if not provided", type=int)
    parser.add_argument("--dbuser", help="username to use when creating the database, ignorred for SQLite databases, REQUIRED for others.", type=str)

def main():
    param_parser = argparse.ArgumentParser(description="Generates a DB schema based on the Baseball Databank csv files.")    
    define_parameters(param_parser)
    args = param_parser.parse_args()

    if args.dbname == None and args.dbtype.tolower() != SQLITE:
        print "ERROR: If --dbtype is not sqlite then --dbname is required"
        param_parser.print_help()
        sys.exit(1)

    bdb_db = ""
    
    sqlite_db_name = "baseball_databank.sqlite3"
    if args.dbpath != None and args.dbpath != "":
        if args.dbtype == SQLITE:
            sqlite_db_name = args.dbpath 
        else:
            print "The argument --dbpath is only needed for Sqlite databases, ignoring."

    if args.dbtype == SQLITE:
        bdb_db = SqliteDatabase(sqlite_db_name)
    else:
        if args.dbuser == None:
            param_parser.print_help() 
            sys.exit(1)
        
        if args.dbtype == POSTGRES:
            port = 5432
            if port > 0:
                port = args.dbport
            bdb_db = PostgresqlDatabase('baseballdatabank', host=args.dbhost, user=args.dbuser, port=port)

        if args.dbtype == MYSQL:  #>r#(-t:Hu4&w
            port = 3306
            password = ""
            if args.dbport > 0:
                port = args.dbport
            
            if args.dbpass != None:
                password = args.dbpass
            bdb_db = MySQLDatabase('baseballdatabank',  host=args.dbhost, port=port, user=args.dbuser, password=args.dbpass)
       
    initialize_db_and_connect(bdb_db)
    
if __name__ == "__main__":
    main()