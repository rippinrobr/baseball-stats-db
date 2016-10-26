import csv
import sys
import os

base_dir = os.getenv("BDB_BASE_DIR", "/home/rob/src/baseballdatabank/core")

def list_to_dict(headers, data):
    d = {}
    num_cols = len(headers)
    num_data_cols = len(data)
    for i in range(0, len(headers)):
        if i < num_cols and i < num_data_cols:
            if headers[i]:
                d[ headers[i] ] = data[i]
    return d

def process_file(file_path, bdb_obj_list, stat_prop, id_key, entity):
    print "processing: " + file_path

    with open(file_path)  as csvfile:
        is_header_row = True
        csv_headers = []

        reader = csv.reader(csvfile, delimiter=',')
        counter = 0
        for record in reader:
            if is_header_row:
                csv_headers = record
                is_header_row = False
                continue

            if counter != 0:
                if counter % 60 == 0:
                    print ""
                sys.stdout.write(".")
                sys.stdout.flush()
            
            stats_dict = list_to_dict(csv_headers, record)
            id = stats_dict.pop(id_key, None) 
            if id:
                if id not in bdb_obj_list:
                    bdb_obj_list[ id] = entity(id)
            
                bdb_obj_list[id].add_stats(stat_prop, stats_dict)
            counter = counter + 1

    print "\nDone!"

def process_file_with_composite_key(file_path, bdb_obj_list, key_field_list, entity):
    print "processing: " + file_path

    with open(file_path)  as csvfile:
        is_header_row = True
        csv_headers = []

        reader = csv.reader(csvfile, delimiter=',')
        counter = 0
        for record in reader:
            if is_header_row:
                csv_headers = record
                is_header_row = False
                continue

            if counter != 0:
                if counter % 60 == 0:
                    print ""
                sys.stdout.write(".")
                sys.stdout.flush()
            
            stats_dict = list_to_dict(csv_headers, record)
            id = ""
            for i,f in enumerate(key_field_list):
                if i == 0:
                    id = stats_dict.pop(f)
                else:
                    id = id + "_" + stats_dict.pop(f)

            if id:
                if id not in bdb_obj_list:
                    bdb_obj_list[id] = entity(id)
            
                bdb_obj_list[id].add_stats("", stats_dict)
            counter = counter + 1

    print "Done!"
