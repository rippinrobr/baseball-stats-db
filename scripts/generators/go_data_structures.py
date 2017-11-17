import inflection
import sys
from core_data_structures import *

def create_go_datastructure(args, headers, data_types):
    original_output_stream = sys.stdout
    test_file = ""
    file_name = get_file_name(args.input).lower().replace(".csv", "")
        
    if args.output_dir != None:
        output_file = os.path.join(args.output_dir, file_name+".go")
        print "Writing to the file", output_file
        sys.stdout = open(output_file, "w+")

    print_code_file(args, headers, data_types, "TableObject")
    sys.stdout = original_output_stream

    if args.output_dir != None:
        test_file = os.path.join(args.output_dir, file_name +"_test.go")
        print "Writing to the file", test_file
        sys.stdout = open(test_file, "w+")

    print_test_file(args, headers, data_types, "TableObject")
    sys.stdout = original_output_stream
    
def get_package_name(args):
    package = "// DON'T FORGET TO ADD YOUR PACKAGE LINE HERE\n\n"
    if args.package:
        if args.output_dir != None and args.output_dir != "":
            parts = args.output_dir.split("/")
            package = "package " + parts[-1]  + "\n\n"

    return package

def print_code_file(args, headers, data_types, interface_name):
    data_structure_name, table_name = get_data_struct_name(args)
    index = 0
    types = {
        "float" : "float64"
    }

    print get_package_name(args)
    print "import (\n  \"os\"\n  \"log\"\n"
    print "  \"github.com/rippinrobr/baseball-databank-tools/pkg/parsers/csv\"\n"
    print "  \"github.com/rippinrobr/baseball-databank-tools/pkg/db\"\n"
    print ")\n"
    print "// "+data_structure_name+" is a model that maps the CSV to a DB Table"
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
            db_tag =  "db:\""+raw_col.replace(".","")+"\""  

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
    print ""
    print_parse_csv_func(data_structure_name)

def print_get_table_name_func(struct_name, table_name):
    print "// GetTableName returns the name of the table that the data will be stored in"
    print "func (m *"+struct_name+") GetTableName() string {\n  return \""+table_name+"\"\n}" 

def print_get_file_name_func(struct_name, file_path):
    file_name = get_file_name(file_path)
    print "// GetFileName returns the name of the source file the model was created from"
    print "func (m *"+struct_name+") GetFileName() string {\n  return \""+file_name+"\"\n}" 

def print_get_file_path_func(struct_name, file_path):
    print "// GetFilePath returns the path of the source file"
    print "func (m *"+struct_name+") GetFilePath() string {\n  return \""+file_path+"\"\n}" 

def print_parse_csv_func(struct_name):
    print "// GenParseAndStoreCSV returns a function that will parse the source file,\\n//create a slice with an object per line and store the data in the db"
    print "func (m *"+struct_name+") GenParseAndStoreCSV(f *os.File, repo db.Repository, pfunc csv.ParserFunc) ParseAndStoreCSVFunc {"
    print "  return func() error {"
    print "    rows := make([]*"+struct_name+", 0)"
    print "    numErrors := 0"
    print "    err := pfunc(f, &rows)"
    print "    if err == nil {"
    print "       numrows := len(rows)"
    print "       if numrows > 0 {"
    print "         log.Println(\""+struct_name+" ==> Truncating\")"
    print "         terr := repo.Truncate(m.GetTableName())"
    print "         if terr != nil {"
    print "            log.Println(\"truncate err:\", terr.Error())"
    print "         }"
    print ""
    print "         log.Printf(\""+struct_name+" ==> Inserting %d Records\\n\", numrows)"
    print "         for _, r := range rows {"
    print "           ierr := repo.Insert(m.GetTableName(), r)"
    print "           if ierr != nil {"
    print "             log.Printf(\"Insert error: %s\\n\", ierr.Error())"
    print "             numErrors++"
    print "           }"
    print "         }"
    print "       }"
    print "       log.Printf(\""+struct_name+" ==> %d records created\\n\", numrows-numErrors)"
    print "    }"
    print ""
    print "    return err"
    print "   }"
    print "}"

def print_test_file(args, headers, data_types, interface_name):
    data_structure_name, table_name = get_data_struct_name(args)
    index = 0

    print get_package_name(args)

    print "import ("
    print "  \"testing\""
    print "  \"github.com/rippinrobr/baseball-databank-tools/pkg/db\"\n"
    print "  \"github.com/rippinrobr/baseball-databank-tools/pkg/parsers\"\n"
    print ")\n"

    print_get_table_name_test_func(data_structure_name, table_name)
    print ""
    print_get_file_name_test_func(data_structure_name, get_file_name(args.input))
    print ""
    print_get_file_path_test_func(data_structure_name, args.input)

def print_get_table_name_test_func(struct_name, table_name):
    print "func TestGetTableName"+struct_name+"(t *testing.T) {"
    print "  out := "+struct_name+"{}"
    print "  expectedValue := \""+table_name+"\""
    print "  actualValue := out.GetTableName()\n"
    print "  if actualValue != expectedValue {"
    print "    t.Errorf(\"actualValue (%s) != expectedValue (%s)\\n\", actualValue, expectedValue)"
    print "  }"
    print "}"

def print_get_file_name_test_func(struct_name, file_name):
    print "func TestGetFileName"+struct_name+"(t *testing.T) {"
    print "  out := "+struct_name+"{}"
    print "  expectedValue := \""+file_name+"\""
    print "  actualValue := out.GetFileName()\n"
    print "  if actualValue != expectedValue {"
    print "    t.Errorf(\"actualValue (%s) != expectedValue (%s)\\n\", actualValue, expectedValue)"
    print "  }"
    print "}"

def print_get_file_path_test_func(struct_name, file_path):
    print "func TestGetFilePath"+struct_name+"(t *testing.T) {"
    print "  out := "+struct_name+"{}"
    print "  expectedValue := \""+file_path+"\""
    print "  actualValue := out.GetFilePath()\n"
    print "  if actualValue != expectedValue {"
    print "    t.Errorf(\"actualValue (%s) != expectedValue (%s)\\n\", actualValue, expectedValue)"
    print "  }"
    print "}"

def print_parse_csv_test_func(struct_name):
    print "func TestGenParseAndStoreCSV"+struct_name+"(t *testing.T) {"
    print "  out := "+struct_name+"{}"
    print "  r := db.Repo{}"
    print "  if out.TestGenParseAndStoreCSV(nil, ) != nil {"
    print "    t.Errorf(\"actualValue (%s) != expectedValue (%s)\\n\", actualValue, expectedValue)"
    print "  }"
    print "}"