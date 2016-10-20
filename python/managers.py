import csv
import sys
import utils

# this should get moved into a settings file/object 
base_dir = "/home/rob/src/baseballdatabank/core"

manager_files = [base_dir+"/Managers.csv", 
        base_dir+"/ManagersHalf.csv", 
        base_dir+"/AwardsManagers.csv", 
        base_dir+"/AwardsShareManagers.csv"]
# end of settings info

utils.hi()

managers = dict() 

class RawManager:
    """ A representation of a manager's managerial stats as read from CSV file """
    def __init__(self, playerId):
        self.playerId = playerId
        self.seasons = []
        self.awards = []

    def add_stats(self, stat_prop, stats):
        if stat_prop == "seasons":
            self.seasons.append( stats )

        if stat_prop == "awards":
            self.awards.append( stats )


def list_to_dict(headers, data):
    d = {}
    num_cols = len(headers)
    num_data_cols = len(data)
    for i in range(0, len(headers)):
        if i < num_cols and i < num_data_cols:
            if headers[i]:
                d[ headers[i] ] = data[i]
    return d


def process_file(file_path, bdb_obj_list, stat_prop, id_key):
    print "processing: " + file_path

    with open(file_path)  as csvfile:
        is_header_row = True
        csv_headers = []

        mgr_reader = csv.reader(csvfile, delimiter=',')
        counter = 0
        for mgr_season in mgr_reader:
            if is_header_row:
                csv_headers = mgr_season
                is_header_row = False
                continue

            if counter != 0:
                if counter % 20 == 0:
                    print ""
                sys.stdout.write(".")
                sys.stdout.flush()
            
            stats_dict = list_to_dict(csv_headers, mgr_season)
            id = stats_dict.pop(id_key, None) 
            if id:
                if id not in bdb_obj_list:
                    bdb_obj_list[ id] = RawManager(id)
            
                bdb_obj_list[id].add_stats(stat_prop, stats_dict)
            counter = counter + 1

    print "\nDone!"

# start the processing
process_file(manager_files[0], managers, "seasons", "playerID")
process_file(manager_files[1], managers, "seasons", "playerID")
process_file(manager_files[2], managers, "awards", "playerID")
process_file(manager_files[3], managers, "awards", "playerID")

print managers["bochybr01"].awards
