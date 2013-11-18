var async = require('async')
  , dataUtils = require('./data-utils') 
  ;

dataUtils.createIndex( dataUtils.mongoUrl, 'players', {'battingStats.yearID':1});
dataUtils.createIndex( dataUtils.mongoUrl, 'players', {'fieldingStats.yearID':1});
dataUtils.createIndex( dataUtils.mongoUrl, 'players', {'pitchingStats.yearID':1});
dataUtils.createIndex( dataUtils.mongoUrl, 'players', 
  {'battingPlayoffStats.yearID':1, 'battingPlayoffStats.round':1});
dataUtils.createIndex( dataUtils.mongoUrl, 'players', 
  {'fieldingPlayoffStats.yearID':1, 'fieldingPlayoffStats.round':1});
dataUtils.createIndex( dataUtils.mongoUrl, 'players', 
  {'pitchingPlayoffStats.yearID':1, 'pitchingPlayoffStats.round':1});
