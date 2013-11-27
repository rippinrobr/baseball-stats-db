##Data Related Scripts

The scripts in this directory are used to load the [Baseball Databank](https://github.com/chadwickbureau/baseballdatabank/tree/2012update) into MongoDb database.

----------
### Running the load scripts
**0. Install node packages** - In the scripts directory run:

<pre>
npm install
</pre>

**1. Setting the environment variables** - Before running the load scripts you may need to set an the environment variable `MONGO_URL` to indicate which server to connect to and which database to use.  If you choose not to set the variable the scripts assume that there is MongoDB server running on the same machine as the scripts.  The default database name is `bdb`.  That is the only setting that you will need to worry about.  The data is fetched from github and parsed.

**2. Loading the Data** - The easiest way to do that is to change into the scripts directory and run:

<pre>
./load-all.sh
</pre>

<code>load-all.sh</code> runs the <code>load-managers.sh</code>, <code>load-players.sh</code> and the <code>load-seasons.sh</code> scripts. Each of these scripts in turn call the appropriate javascript files ot load the <code>managers</code>, <code>players</code>, and <code>seasons</code> collections.

----------

### The *.sh files
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
  - load-postseason-batting.js  => loads playoff batting stats
  - load-postseason-fielding.js => loads playoff fielding stats
  - load-postseason-pitching.js => loads playoff pitching stats
  - load-salaries.js => adds seasonal salary info to the salaries array if it exists.
  - create-player-stat.indexes.js => creates the indexes on the players collection
  

**load-teams.sh** - calls all of the team related load scripts.
  - load-teams.js => creates the documents for each team loading in the team stats and record
  - load-teams-managers.js => adds the managers to managers array for each season
  - load-half-season.js => adds the 1981 half season stats to the document.
  - load-playoff-records => loads w-l records for the playoffs
