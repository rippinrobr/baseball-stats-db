import csv
import sys
import os

BASE_DIR = os.getenv("BDB_BASE_DIR", "/home/rob/src/baseballdatabank/core")
SQLITE = "sqlite"
PKs = {
    "Master" : ["playerID"],
    "Managers" : ["playerID", "yearID", "teamID"]
}

def list_to_dict(headers, data, schema):
    d = {}
    
    num_cols = len(headers)
    num_data_cols = len(data)
    for i in range(0, len(headers)):
        if i < num_cols and i < num_data_cols:
            if headers[i]:
                value = data[i]
                curr_header = headers[i]
                d[curr_header] = data[i]
                if schema[curr_header]['length'] < len(value):
                    schema[curr_header]['length'] = len(value)

                if schema[curr_header]['data_type'] != "string":
                    try:
                        int(value)
                        schema[curr_header]['data_type'] = "int"
                    except ValueError:
                        try:
                            float(value)
                            schema[curr_header]['data_type'] = "float"
                        except ValueError:
                            schema[curr_header]['data_type'] = "string"

    return d

def new_process_file(file_path):
    bdb_obj_list = []
    schema_info = {}

    print "processing: " + file_path
    file_name = os.path.splitext(file_path)[0]
    last_slash_index = file_name.rfind('/')
    schema_info['table_name'] = file_name[last_slash_index+1:]
                    
    with open(file_path)  as csvfile:
        is_header_row = True
        csv_headers = []

        reader = csv.reader(csvfile, delimiter=',')
        for record in reader:
            if is_header_row:
                csv_headers = record
                for header in csv_headers:
                    schema_info[header] = {'data_type': "", 'length': 0}

                is_header_row = False
                continue

            row= list_to_dict(csv_headers, record, schema_info)
            bdb_obj_list.append(row)
            
    print "\nDone!"
    return bdb_obj_list, schema_info

def write_sql_schema(tables, db_type):
    print "from peewee import *"
    print "from playhouse.postgres_ext import *"
    print "import logging"
    print "import sys"
    print ""

    # Print the connection info here 

    __create_table(tables["Master"], db_type)
    __create_table(tables["Managers"], db_type)
    
def __create_table(table, db_type):
    print "class %s(BaseModel):" % table["table_name"]
    for col in table:
        if col != "table_name":
            print "    %s = %s" % (col, __col_data_type(table[col], __is_pk(col, table["table_name"])))
    print ");"
    print ""

def __col_data_type(column, is_primary_key):
    null_str = "null='true'"
    if is_primary_key:
        null_str = "null='false'"

    if column["data_type"] == "string":
        return "CharField(%s)" % null_str
    
    if column["data_type"] == "int":
        return "IntegerField(%s)" % null_str

def __is_pk(column, table_name):
    return column in PKs[table_name] 