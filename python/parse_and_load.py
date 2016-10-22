import hall_of_fame
import managers
import players
import season

print "### PROCESSING MANAGER DATA ###"
raw_managers = managers.parse()

print "### PROCESSING PLAYER DATA ###"
raw_players = players.parse()

print "### PROCESSING SEASON  DATA ###"
raw_seasons = players.parse()

print "### PROCESSING SEASON  DATA ###"
raw_hof = hall_of_fame.parse()
