##Data Related Scripts

The scripts in this directory are used to filter the [Baseball Databank](https://github.com/chadwickbureau/baseballdatabank/tree/2012update) data by teamID 
or loading the data into MongoDb.

----------
### Running the load scripts
**0. Install node packages** - In the scripts directory run:

<pre>
npm install
</pre>

**1. Filter out the Giants stats** - After cloning the [Baseball Databank](https://github.com/chadwickbureau/baseballdatabank/tree/2012update) repository open up the <code>create-team-csv-files.sh</code> file in your favorite editor and make the following changes:

<pre>
TEAMIDS='SFN|NY1'
OUTDIR=../giants
FRANCHISEIDS='SFG|NYI'
BDBDIR=../baseballdatabank/official
</pre>

1. Set the <code>BDBDIR</code> value to the location of where you cloned the BaseballDatabank repo.  
2. Change the <code>OUTDIR</code> variable to be where you want to write out the filtered CSV files. *You will also need to set the <code>DATA_DIR</code> environment variable to the same value.  The load scripts will use this variable.*

If you wish to use a different team than the Giants, why would you, all you need to do is change the value of the <code>TEAMIDS</code> and <code>FRANCHISEIDS</code> variables.  You find the team IDs in the Teams.csv file.  The franchise IDs can be found in the TeamsFranchises.csv files. If you find more than one team or franchise ID remember to put a | in between the IDs

**2. Load MongoDB** - After prepping the <code>create-team-csv-files.sh</code>, setting the <code>DATA_DIR</code> environment variable, and running <code>create-team-csv-files.sh</code> the CSV files will be ready for loading.  The easiest way to do that is to change into the scripts directory and run:

<pre>
./load-all.sh
</pre>

<code>load-all.sh</code> runs the <code>load-managers.sh</code>, <code>load-players.sh</code> and the <code>load-seasons.sh</code> scripts.  Each of these scripts in turn call the appropriate javascript files ot load the <code>managers</code>, <code>players</code>, and <code>seasons</code> collections.

----------

### The *.sh files
**create-team-csv-files.sh**
 - Reads the [Baseball Databank](https://github.com/chadwickbureau/baseballdatabank/tree/2012update) CSV files, going through each file and parsing out the data for the team and franchise specified in the script.  The script produces CSV files that contain only data related to the specified team/franchise. These files are then used by the loading scripts.

**load-all.sh** - a wrapper script that calls <code>load-managers.sh</code>, <code>load-players.sh</code>, and <code>load-seasons.sh</code>.

**load-managers.sh** - calls all of the manager related load scripts.
  - load-managers.js    => creates the manager's document in the collectoin.  Adds the manager's record to the document. 
  - load-mgr-demog.js   => adds the manager's demographics to the document.
  - load-half-mgrs.js   => loads the manager's record for the 1981 season pre and post strike.
  - load-manager-awards => loads any awards the manager may have won into the awards array.
  

**load-players.sh** - calls all of the player related load scripts.
  - load-plyr-demog.js => creates the player's document in the players collection and adds his demographic information
  - load-allstars.js => populates the allstars array with the year(s) the player made the all star team.
  - load-appearances.js => adds records to the appearances array indicating how many games they played in a particular role.
  - load-batting.js => loads regular season batting stats
  - load-fielding.js => loads regular season fielding stats
  - load-pitching.js => loads regular season pitching stats
  - load-player-awards.js => populates the awards array with any awards the player may have won
  - load-positions.js => adds positions played by the player to the positions array.
  - load-postseason-batting.js  => loads playoff batting stats
  - load-postseason-fielding.js => loads playoff fielding stats
  - load-postseason-pitching.js => loads playoff pitching stats
  - load-salaries.js => adds seasonal salary info to the salaries array if it exists.
  - create-player-stat.indexes.js => creates the indexes on the players collection
  

**load-seasons.sh** - calls all of the season related load scripts.
  - load-seasons.js => creates the documents for each season loading in the team stats and record
  - load-seasons-managers.js => adds the managers to managers array for each season
  - load-half-season.js => adds the 1981 half season stats to the document.
  - load-rosters.js => adds a condensed player object to the roster array.
  - load-playoff-records => loads w-l records for the playoffs
