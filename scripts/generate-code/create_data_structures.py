import argparse
import csv
import inflection
import glob
import os, errno, sys
import re
from go_data_structures import create_go_datastructure

LANG_GO = "go"

TYPE_STRING = "string"
TYPE_INT = "int"
TYPE_FLOAT = "float"


def define_parameters(parser):
    """Creates the support command line options"""
    parser.add_argument("--bson", help="If the language supports add bson tags or bson representation to the structure", action="store_true")
    parser.add_argument("--csv", help="If the language supports add CSV tags or CSV representation to the structure", action="store_true")
    parser.add_argument("--db", help="adds db tags to Go structs." , action="store_true")
    parser.add_argument("--input", default="", help="the path to the input CSV file",type=str )
    parser.add_argument("--input-dir", default="", help="the path to the directory that contains the CSV files to parse",type=str )
    parser.add_argument("--json", help="If the language supports add JSON tags or JSON representation to the structure", action="store_true")
    parser.add_argument("--package", help="sets the package to the correct name for Go structs", action="store_true")
    parser.add_argument("--language", choices=[LANG_GO], default=LANG_GO, help="create language specific data structures", required=True, type=str)
    parser.add_argument("--name", help="name of the data structure being created", type=str)
    parser.add_argument("--output-dir", help="the directory where the generated file should be written.  If not provided file will be written to stdout")
    parser.add_argument("--verbose", help="more output during the parsing and creation of the data structures", action="store_true")

def get_data_type(col_val):
    """Determines if the current value is numeric or a string"""
    dtype = ""

    original_col_val = col_val
    digits_only = col_val.replace('-', '',1).replace(',', '', -1).replace(".", "")
    if digits_only.isdigit():
        try:
            int(original_col_val)
            dtype = TYPE_INT
        except ValueError:
            dtype = TYPE_FLOAT
    
    return dtype

def parse_file(args):
    """Parses the CSV file and tracks the datatypes of each column

    this will be a map with the column name and the data type of the col
    all cols will default to a empty string so that I will know if a column has already
    been declared a string, once a string always string.  If after all lines
    have been read and a column still has an empty string it will be set as 
    a string.  Currently the values for data types are string, int or float"""

    data_types = []
    headers = []

    with open(args.input, "r") as csvfile:
        reader = csv.reader(csvfile)
        have_columns = False

        for line in reader:
            if have_columns:
                index = 0
                for col in line:
                    if col != "": 
                        if data_types[index] != TYPE_STRING and data_types[index] != TYPE_FLOAT:
                            data_types[index] = get_data_type(col)
                    # else:
                    #     data_types[index] = TYPE_STRING
                    index += 1

            else:
                headers = line   
                for col in line:
                    data_types.append("")
                have_columns = True 

    return headers, data_types




def create_output_directory(dir_path):
    try:
        os.makedirs(dir_path)
    except OSError as e:
        if e.errno != errno.EEXIST:
            raise

def main():
    parser = argparse.ArgumentParser(description="Generate language specific data structures that model each of the Baseball Databank CSV files.")
    define_parameters(parser)
    args = parser.parse_args()

    if args.input == "" and args.input_dir == "":
        print "ERROR: Please provide one of the following arguments --input or --input-dir\n"
        parser.print_help()
        sys.exit(1)

    LANGUAGE_FUNCS = {
        LANG_GO: create_go_datastructure
    }

    if args.output_dir != None and args.output_dir != "":
        create_output_directory(args.output_dir)

    if args.input_dir != None and args.input_dir != "":    
       process_directory(args, LANGUAGE_FUNCS[args.language])
    else:
       generate_struct(args, LANGUAGE_FUNCS[args.language])

def process_directory(args, gen_func):
    success = False
    if os.path.exists(args.input_dir):
        if os.path.isdir(args.input_dir):
            process_files(args,  gen_func)
            success = True
        else:
            print "\nERROR: The value of --input-dir '"+args.input_dir+" is not a directory. Please provide a directory\n"
            parser.print_help()
    else:
        print "\nERROR: The directory, '"+args.input_dir+"' does not exist. Please provide a valid directory path for --input-dir\n"
    
    return success 
        
def process_files(args, gen_func):
    for file in glob.glob(os.path.join(args.input_dir, "*.csv")):
        args.input = file
        generate_struct(args,gen_func) 

def generate_struct(args, gen_func):
    headers, data_types = parse_file(args)
    if args.verbose:
        print "Data types\n------------------------------------------------------------------"
        index = 0
        for field in headers:
            print field, " ", data_types[index]
            index += 1

    gen_func(args, headers, data_types)

if __name__ == "__main__":
    main()
