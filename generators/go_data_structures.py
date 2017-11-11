import inflection
import sys
from core_data_structures import *

def create_go_datastructure(args, headers, data_types):
    types = {
        "float" : "float64"
    }

    index = 0
    
    original_output_stream = sys.stdout
    
    if args.output_dir != None:
        file_name = get_file_name(args.input).lower()
        output_file = os.path.join(args.output_dir, file_name)
        print "Writing to the file", output_file
        sys.stdout = open(output_file, "w+")

    data_structure_name, table_name = get_data_struct_name(args)
    print "DON'T FORGET TO ADD YOUR PACKAGE LINE HERE\n\n "
    print "type", data_structure_name, "struct {"
    for raw_col in headers:
        json_tag = ""
        csv_tag = ""
        db_tag = ""
        tags = ""

        col = col_name_cleaner(raw_col).lower()
        if args.json:
            json_tag = "json:\"" +inflection.camelize(col,False) +"\""

        if args.csv:
            csv_tag =  "csv:\""+raw_col+"\""  

        if args.db:
            db_tag =  "db:\""+raw_col+"\""  

        if json_tag != "" or csv_tag != "" or db_tag != "":
            if json_tag != "" and csv_tag == "" and db_tag == "":
                tags = "`"+json_tag+"`"
            elif csv_tag != "" and json_tag == "" and db_tag == "":
                tags = "`"+csv_tag+"`"
            elif db_tag != "" and json_tag == "" and csv_tag == "":
                tags = "`"+csv_tag+"`"
            else:
                tags = "`"+json_tag+"  "+csv_tag+"  "+db_tag+"`"
                
        
        print "  ", col.title(), " ", convert_to_lang_specific_type(types, data_types[index]), tags
        index += 1

    print "}\n"
    print_get_table_name_func(data_structure_name, table_name)
    print ""
    print_get_file_name_func(data_structure_name, args.input)
    print ""
    print_get_file_path_func(data_structure_name, args.input)

    sys.stdout = original_output_stream

def print_get_table_name_func(struct_name, table_name):
    print "func (m *"+struct_name+") GetTableName() string {\n  return \""+table_name+"\"\n}" 

def print_get_file_name_func(struct_name, file_path):
    file_name = get_file_name(file_path)
    print "func (m *"+struct_name+") GetFileName() string {\n  return \""+file_name+"\"\n}" 

def print_get_file_path_func(struct_name, file_path):
    print "func (m *"+struct_name+") GetFilePath() string {\n  return \""+file_path+"\"\n}" 