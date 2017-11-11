import inflection
import sys
from core_data_structures import *

def create_haskell_datastructure(args, headers, data_types):
    comma = ""
    index = 0
    original_output_stream = sys.stdout
    if args.output_dir != None:
        file_name = get_file_name(args.input).lower().replace(".csv", ".hs")
        output_file = os.path.join(args.output_dir, file_name)
        print "Writing to the file", output_file
        sys.stdout = open(output_file, "w+")

    name = get_data_struct_name(args) 
    print "data", name, "=", name , "{"
    for raw_col in headers:
        if index > 0:
            comma=","
        
        col = col_name_cleaner(raw_col).lower()
        print "  ", comma, col[:1].lower() + col[1:], "::", data_types[index].title()
        index += 1
    print "} deriving (Show)"

    sys.stdout = original_output_stream
