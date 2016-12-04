Data Relations
 - data store
 	- File/primary key
    - AllStarFull.csv   		  playerID, gameID
    - Appearances.csv   		  playerID, teamID, yearID
    - AwardsManager.csv 		  playerID, awardID, yearID
    - AwardsPlayers.csv 			playerID, awardID, yearID
    - AwardsShareManagers.csv playerID, awardID, yearID
    - AwardsSharePlayers.csv 	playerID, awardID, yearID
	- Batting.csv 						playerID, yearID, teamID, stint
	- BattingPost.csv 				playerID, teamID, yearID, round
	- CollegePlaying.csv 		  playerID, schoolID, yearID
	- Fielding.csv            playerID, teamID, yearID, POS
	- FieldingPost.csv        playerID, teamID, yearID, POS, round
	- FieldingOF.csv				  playerID, yearID, stint
	- HallOfFame.csv          playerID, category
	- HomeGames.csv           year, teamID, park
	- Managers.csv            playerID, yearID, teamID

 - Update readme
  - fixed issue with HallOfFame.csv's yearid, changed it to yearID

 	- include column name changes for HomeGames columns
		- year.key to yearID
		- league.key to lgID
		- team.key to teamID
		- park.key to parkID
		- span.first to firstGame
		- span.last to lastGame

	- include column name changes for Parks
		- park.key to parkID
		- park.name to name
		- park.alias to alias
