import csv;

# this should get moved into a settings file/object 
base_dir = "/home/rob/src/baseballdatabank/core"

manager_files = [base_dir+"/Managers.csv", 
        base_dir+"/ManagersHalf.csv", 
        base_dir+"/AwardsManagers.csv", 
        base_dir+"/AwardsShareManagers.csv"]
# end of settings info

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

def process_manager_files_with_id_in_first_col(file_path, stat_prop):
    
    with open(file_path)  as csvfile:
        is_header_row = True
        csv_headers = []

        mgr_reader = csv.reader(csvfile, delimiter=',')
        for mgr_season in mgr_reader:
            if is_header_row:
                csv_headers = mgr_season[1:]
                is_header_row = False
                continue

            stats_dict = list_to_dict(csv_headers, mgr_season[1:])
            if mgr_season[0]:
                if mgr_season[0] not in managers:
                    managers[mgr_season[0]] = RawManager(mgr_season[0])
            
                managers[mgr_season[0]].add_stats(stat_prop, stats_dict)

# start the processing
process_manager_files_with_id_in_first_col(manager_files[0], "seasons")
process_manager_files_with_id_in_first_col(manager_files[1], "seasons")
process_manager_files_with_id_in_first_col(manager_files[2], "awards")

print( managers["bochybr01"].seasons )
print( managers["bochybr01"].awards )
#for k in managers:
#    print( managers[k].seasons)

