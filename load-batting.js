var async = require('async')
  , dataUtils = require('./data-utils') 
  ;

// Creates the 'Pipeline' object with the pipeline specific functions in it.
// Returns the async.compose results
var loadBattingStats = (function() {
  var createUpdateObjects = function( pipeObj, callback ) {
    var docs = []
      , recs = pipeObj.data[ pipeObj.prevStep ]
      ;

    for(var i=0; i < recs.length; i++ ) {
      var id = recs[i].playerID;

      delete recs[i].teamID;
      delete recs[i].lgID;

      docs.push( {query: { _id: id }, set: { $addToSet: { battingStats: recs[i] }}} );
    }

    pipeObj.data.push( docs );
    pipeObj.prevStep += 1;
    callback( null, pipeObj );
  };

  return async.compose(
    createUpdateObjects,
    dataUtils.createObjects,
    dataUtils.readRemoteCsv
  );

}());


var inputSrc = { path: dataUtils.baseGithubUrl + '/Batting.csv',
              headers: 1,
          dataTypeMap: [ 'playerID', 'teamID', 'lgID'] };

var outputSettings = { type: 'mongodb', 
                        url: dataUtils.mongoUrl, 
                 collection: 'players' };

var pipeObj = { input: inputSrc,
               exitFn: null,
               output: outputSettings,
                 data: [],
             prevStep: -1 };


loadBattingStats( pipeObj, function(err, result ) {
  if ( pipeObj.output ) {
    dataUtils.updateFns[ pipeObj.output.type ]( result.data[2], pipeObj.output, function( err, r) {
      console.log('[loadBatting] Batting Stats loaded' );
    });
  }
});
