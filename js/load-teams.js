var async = require('async')
  , dataUtils = require('./data-utils') 
  ;

// Creates the 'Pipeline' object with the pipeline specific functions in it.
// Returns the async.compose results
var loadTeams = (function() {
    var createSeasonObjects = function( pipeObj, callback ) {
    var docs = []
      , recs = pipeObj.data[ pipeObj.prevStep ]
      ;
 
    console.log( recs.length );
    for(var i=0; i < recs.length; i++ ) {
      recs[i]._id      = recs[i].yearID + "_" + recs[i].teamID;
      recs[i].managers = [];

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
     dataUtils.readRemoteCsv
  );

}());


var inputSrc = { path: dataUtils.baseGithubUrl + '/Teams.csv',
              headers: 1,
          dataTypeMap: [ 'lgID', 'teamID', 'franchID', 'divID', 'name', 'park', 'teamIDBR', 'teamIDretro', 'DivWin', 'WCWin', 'LgWin', 'WSWin'],
          floatColMap: [ 'ERA' ]  };

var outputSettings = { type: 'mongodb', 
                        url: dataUtils.mongoUrl, 
                 collection: 'teams' };

var pipeObj = { input: inputSrc,
               exitFn: null,
               output: outputSettings,
                 data: [],
             prevStep: -1 };


loadTeams( pipeObj, function(err, result ) {

  if ( pipeObj.output ) {
    dataUtils.emptyCollection( pipeObj.output.url, 'teams', function(err, numDeleted) { 
      console.log('[loadTeams]', 'Deleted', numDeleted, 'teams documents.');
      dataUtils.storageFns[ pipeObj.output.type ]( result.data[2], pipeObj.output, function( err, r) {
        console.log('[loadTeams]', 'Teams Loaded!');
        dataUtils.createIndex( dataUtils.mongoUrl, 'teams', { 'managers._id':1});
        dataUtils.createIndex( dataUtils.mongoUrl, 'teams', { 'players._id':1});
      });
    });
  }
});
