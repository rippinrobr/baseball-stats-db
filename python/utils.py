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

def new_process_file(file_path):
    bdb_obj_list = []

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
            
            bdb_obj_list.append( list_to_dict(csv_headers, record) )
            counter = counter + 1

    print "\nDone!"
    return bdb_obj_list

#def process_file(file_path, bdb_obj_list, stat_prop, id_key, entity):
#    print "processing: " + file_path
#
#    with open(file_path)  as csvfile:
#        is_header_row = True
#        csv_headers = []
#
#        mgr_reader = csv.reader(csvfile, delimiter=',')
#        counter = 0
#        for mgr_season in mgr_reader:
#            if is_header_row:
#                csv_headers = mgr_season
#                is_header_row = False
#                continue
#
#            if counter != 0:
#                if counter % 60 == 0:
#                    print ""
#                sys.stdout.write(".")
#                sys.stdout.flush()
#            
#            stats_dict = list_to_dict(csv_headers, mgr_season)
#            id = stats_dict.pop(id_key, None) 
#            if id:
#                if id not in bdb_obj_list:
#                    bdb_obj_list[ id] = entity(id)
#            
#                bdb_obj_list[id].add_stats(stat_prop, stats_dict)
#            counter = counter + 1
#
#    print "\nDone!"
