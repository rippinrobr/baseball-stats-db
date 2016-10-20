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


def process_file(file_path, stat_prop, id_key):
    
    with open(file_path)  as csvfile:
        is_header_row = True
        csv_headers = []

        mgr_reader = csv.reader(csvfile, delimiter=',')
        for mgr_season in mgr_reader:
            if is_header_row:
                csv_headers = mgr_season
                is_header_row = False
                continue

            stats_dict = list_to_dict(csv_headers, mgr_season)
            id = stats_dict.pop(id_key, None) 
            if id:
                if id not in managers:
                    managers[ id] = RawManager(id)
            
                managers[id].add_stats(stat_prop, stats_dict)

# start the processing
process_file(manager_files[0], "seasons", "playerID")
process_file(manager_files[1], "seasons", "playerID")
process_file(manager_files[2], "awards", "playerID")
process_file(manager_files[3], "awards", "playerID")

print managers["bochybr01"].awards
