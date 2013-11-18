var async = require('async')
  , dataUtils = require('./data-utils') 
  , mongoUrl = process.env.MONGO_URL || 'mongodb://localhost:27017/giants'  
  ;

dataUtils.createIndex( mongoUrl, 'players', {'battingStats.yearID':1});
dataUtils.createIndex( mongoUrl, 'players', {'fieldingStats.yearID':1});
dataUtils.createIndex( mongoUrl, 'players', {'pitchingStats.yearID':1});
dataUtils.createIndex( mongoUrl, 'players', {'battingPlayoffStats.yearID':1, 'battingPlayoffStats.round':1});
dataUtils.createIndex( mongoUrl, 'players', {'fieldingPlayoffStats.yearID':1, 'fieldingPlayoffStats.round':1});
dataUtils.createIndex( mongoUrl, 'players', {'pitchingPlayoffStats.yearID':1, 'pitchingPlayoffStats.round':1});
