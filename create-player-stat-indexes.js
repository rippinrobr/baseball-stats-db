var async = require('async')
  , dataUtils = require('./data-utils') 
  ;

dataUtils.createIndex( dataUtils.mongoUrl, 'players', {'battingStats.yearID':1});
dataUtils.createIndex( dataUtils.mongoUrl, 'players', {'fieldingStats.yearID':1});
dataUtils.createIndex( dataUtils.mongoUrl, 'players', {'pitchingStats.yearID':1});
dataUtils.createIndex( dataUtils.mongoUrl, 'players', {'battingStats.season':1});
dataUtils.createIndex( dataUtils.mongoUrl, 'players', {'fieldingStats.season':1});
dataUtils.createIndex( dataUtils.mongoUrl, 'players', {'pitchingStats.season':1});
dataUtils.createIndex( dataUtils.mongoUrl, 'players', {'battingPlayoffsStats.season':1});
dataUtils.createIndex( dataUtils.mongoUrl, 'players', {'fieldingPlayoffsStats.season':1});
dataUtils.createIndex( dataUtils.mongoUrl, 'players', {'pitchingPlayoffsStats.season':1});

dataUtils.createIndex( dataUtils.mongoUrl, 'players', 
  {'battingPlayoffStats.yearID':1, 'battingPlayoffStats.round':1});
dataUtils.createIndex( dataUtils.mongoUrl, 'players', 
  {'fieldingPlayoffStats.yearID':1, 'fieldingPlayoffStats.round':1});
dataUtils.createIndex( dataUtils.mongoUrl, 'players', 
  {'pitchingPlayoffStats.yearID':1, 'pitchingPlayoffStats.round':1});
