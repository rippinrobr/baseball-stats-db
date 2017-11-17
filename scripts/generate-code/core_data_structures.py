import os

def col_name_cleaner(col_name):
    if col_name[0].isdigit():
        if col_name == "2B":
            return "doubles"
        if col_name == "3B":
            return "triples"
    
    return col_name.replace(".", "").replace("_","")

def convert_to_lang_specific_type(typeMap, type):
    if type in typeMap:
        return typeMap[type]

    return type

def get_data_struct_name(args):
   table_name = os.path.splitext(get_file_name(args.input))[0] 
   if args.name != None:
        return args.name, table_name.lower()
   else:
        return table_name, table_name.lower()

def get_file_name(file_path):
    return os.path.basename(file_path)