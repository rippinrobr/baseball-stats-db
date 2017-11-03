import csv
import argparse

TYPE_STRING = "string"
TYPE_INT = "int"
TYPE_FLOAT = "float"

def define_parameters(parser):
    """Creates the support command line options"""
    parser.add_argument("--input", help="the path to the input CSV file",type=str, required=True )
    parser.add_argument("--verbose", help="more output during the parsing and creation of the datastructures", action="store_true")

def get_data_type(col_val):
    """Deterimes if the current value is numeric or a string"""
    dtype = TYPE_STRING


    if col_val.replace('.', '',1).replace ('-', '',1).replace(',', '', -1).isdigit():
        if '.' in col_val:
            dtype = TYPE_FLOAT
        else:
            dtype = TYPE_INT

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
                    if col != "" and data_types[index] != TYPE_STRING:
                        data_types[index] = get_data_type(col)
                    index += 1

            else:
                headers = line   
                for col in line:
                    data_types.append("")
                
                have_columns = True 
                print "HEADERS: ", headers
                print "DATATYPES: ", data_types

    return headers, data_types

def main():
    parser = argparse.ArgumentParser(description="Test the SMS notification Service.")
    define_parameters(parser)
    args = parser.parse_args()

    headers, data_types = parse_file(args)
    print "------------------------------------------------------------------"

    index = 0
    for field in headers:
        print field, " ", data_types[index]
        index += 1



if __name__ == "__main__":
    main()