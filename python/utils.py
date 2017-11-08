import csv
import sys
import os

SQLITE = "sqlite"
PKs = {
    "Master" : ["playerID"],
    "Managers" : ["playerID", "yearID", "teamID"]
}

class Parser(object):

    def __list_to_dict(self, headers, data, schema):
        p_data = {}

        num_cols = len(headers)
        num_data_cols = len(data)
        for i in range(0, len(headers)):
            if i < num_cols and i < num_data_cols:
                if headers[i]:
                    value = data[i]
                    curr_header = headers[i]
                    p_data[curr_header] = data[i]
                    if schema[curr_header]['length'] < len(value):
                        schema[curr_header]['length'] = len(value)

                    if schema[curr_header]['data_type'] != "string":
                        if value != "":
                            try:
                                int(value)
                                schema[curr_header]['data_type'] = "int"
                            except ValueError:
                                try:
                                    float(value)
                                    schema[curr_header]['data_type'] = "float"
                                except ValueError:
                                    schema[curr_header]['data_type'] = "string"

        return p_data

    def parse_file(self, file_path):
        bdb_obj_list = []
        schema_info = {}

        #print "processing: " + file_path
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

                row = self.__list_to_dict(csv_headers, record, schema_info)
                bdb_obj_list.append(row)

        return bdb_obj_list, schema_info

class SchemaGenerator(object):
    def write_sql_schema(self, tables, db_type):
        print "from peewee import *"
        print "from playhouse.postgres_ext import *"
        print "import logging"
        print "import sys"
        print ""

        print "# Print the connection info here for %s" % db_type

        self.__create_table(tables["Master"])
        self.__create_table(tables["Managers"])
        for table_name in tables:
            if table_name not in ["Master", "Managers"]:
                self.__create_table(tables[table_name])

    def __create_table(self, table):
        print "class %s(BaseModel):" % table["table_name"]
        for col in table:
            if col != "table_name":
                data_type = self.__col_data_type(table[col], self.__is_pk(col, table["table_name"]))
                print "    %s = %s" % (col, data_type)
        print ""

    def __col_data_type(self, column, is_primary_key):
        null_str = "null='true'"
        if is_primary_key:
            null_str = "null='false'"

        if column["data_type"] == "string":
            return "CharField(%s)" % null_str

        if column["data_type"] == "int":
            return "IntegerField(%s)" % null_str

        if column["data_type"] == "float":
            return "FloatField(%s)" % null_str

    def __is_pk(self, column, table_name):
        try:
            return column in PKs[table_name]
        except:
            return False
