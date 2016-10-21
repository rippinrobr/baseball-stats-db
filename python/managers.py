import utils

# this should get moved into a settings file/object 
manager_files = [
        utils.base_dir+"/Managers.csv", 
        utils.base_dir+"/ManagersHalf.csv", 
        utils.base_dir+"/AwardsManagers.csv", 
        utils.base_dir+"/AwardsShareManagers.csv"
]
# end of settings info

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


def parse():
    managers = dict() 
    
    utils.process_file(manager_files[0], managers, "seasons", "playerID", RawManager)
    utils.process_file(manager_files[1], managers, "seasons", "playerID", RawManager)
    utils.process_file(manager_files[2], managers, "awards", "playerID", RawManager)
    utils.process_file(manager_files[3], managers, "awards", "playerID", RawManager)

    #print managers["bochybr01"].awards
    return managers

