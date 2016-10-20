import csv;

# this should get moved into a settings file/object 
base_dir = "/home/rob/src/baseballdatabank/core"

manager_files = [base_dir+"/Managers.csv", 
        base_dir+"/ManagersHalf.csv", 
        base_dir+"/AwardsManagers.csv", 
        base_dir+"/AwardsShareManagers.csv"]
# end of settings info

class RawManager:
    """ A representation of a manager's managerial stats as read from CSV file """
    def __init__(self, playerId, season):
        self.playerId = playerId
        self.seasons = [season]

def list_to_dict(headers, data):
    d = {}
    num_cols = len(headers)
    num_data_cols = len(data)
    for i in range(0, len(headers)-1):
        if i < num_cols and i < num_data_cols:
            if headers[i]:
                print i, "=>", headers[i]
                d[ headers[i] ] = data[i]
                print headers[i], " => ",  data[i]

    return d

managers = dict() 
is_header_row = True
csv_headers = []
with open(manager_files[0])  as csvfile:

    mgr_reader = csv.reader(csvfile, delimiter=',')
    for mgr_season in mgr_reader:
        if is_header_row:
            csv_headers = mgr_season[1:]
            is_header_row = False
            continue

        season_dict = list_to_dict(csv_headers, mgr_season[1:])
        if mgr_season[0] not in managers:
            if mgr_season[0]:
                managers[mgr_season[0]] = RawManager(mgr_season[0], season_dict)
            else:
                managers[mgr_season[0]].seasons.append(season_dict)

#print("hello")
#for k in managers:
#    print( managers[k].seasons)

