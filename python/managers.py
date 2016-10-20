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
    for i in range(0, len(headers)-1):
        if i < num_cols and i < num_data_cols:
            if headers[i]:
                d[ headers[i] ] = data[i]
    return d

with open(manager_files[0])  as csvfile:
    is_header_row = True
    csv_headers = []

    mgr_reader = csv.reader(csvfile, delimiter=',')
    for mgr_season in mgr_reader:
        if is_header_row:
            csv_headers = mgr_season[1:]
            is_header_row = False
            continue

        season_dict = list_to_dict(csv_headers, mgr_season[1:])
        if mgr_season[0]:
            if mgr_season[0] not in managers:
                managers[mgr_season[0]] = RawManager(mgr_season[0]) #, season_dict)
            
            managers[mgr_season[0]].add_stats("seasons", season_dict)

print( managers["bochybr01"].seasons )
#for k in managers:
#    print( managers[k].seasons)

