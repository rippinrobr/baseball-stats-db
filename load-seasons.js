var async = require('async')
  , dataUtils = require('./data-utils') 
  , mongoUrl = process.env.MONGO_URL || 'mongodb://localhost:27017/giants'  
  , dataDir  = process.env.DATA_DIR  || '../giants'
  ;

// Creates the 'Pipeline' object with the pipeline specific functions in it.
// Returns the async.compose results
var loadSeasons = (function() {
    var createSeasonObjects = function( pipeObj, callback ) {
    var docs = []
      , recs = pipeObj.data[ pipeObj.prevStep ]
      ;

    for(var i=0; i < recs.length; i++ ) {
      recs[i]._id      = recs[i].yearID;
      recs[i].managers = [];
      recs[i].players  = []

      delete recs[i].teamIDlahman45;

      docs.push( recs[i] );
    }

    pipeObj.data.push( docs );
    pipeObj.prevStep += 1;
    callback( null, pipeObj );
  };


  return  async.compose(
     createSeasonObjects, 
     dataUtils.createObjects,
     dataUtils.readCsv
  );

}());


var inputSrc = { path: dataDir + '/Teams.csv',
              headers: 1,
          dataTypeMap: [ 'lgID', 'teamID', 'franchID', 'divID', 'name', 'park', 'teamIDBR', 'teamIDretro', 'DivWin', 'WCWin', 'LgWin', 'WSWin'],
          floatColMap: [ 'ERA' ]  };

var outputSettings = { type: 'mongodb', 
                        url: mongoUrl, 
                 collection: 'seasons' };

var pipeObj = { input: inputSrc,
               exitFn: null,
               output: outputSettings,
                 data: [],
             prevStep: -1 };


loadSeasons( pipeObj, function(err, result ) {

  if ( pipeObj.output ) {
    dataUtils.emptyCollection( pipeObj.output.url, 'seasons', function(err, numDeleted) { 
      console.log('[loadSeasons]', 'Deleted', numDeleted, 'seasons documents.');
      dataUtils.storageFns[ pipeObj.output.type ]( result.data[2], pipeObj.output, function( err, r) {
        console.log('[loadSeasons]', 'Seasons Loaded!');
        dataUtils.createIndex( mongoUrl, 'seasons', { 'managers._id':1});
        dataUtils.createIndex( mongoUrl, 'seasons', { 'players._id':1});
      });
    });
  }
});
