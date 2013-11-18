var async = require('async')
  , dataUtils = require('./data-utils') 
  , mongoUrl = 'mongodb://localhost:27017/giants'  
  , dataDir  = process.env.DATA_DIR  || '../giants'
  ;

// Creates the 'Pipeline' object with the pipeline specific functions in it.
// Returns the async.compose results
var loadPlayoffRecords = (function() {
  var createUpdateObjects = function( pipeObj, callback ) {
    var docs = []
      , recs = pipeObj.data[ pipeObj.prevStep ]
      ;

    for(var i=0; i < recs.length; i++ ) {
      var id = recs[i].yearID;
      
      docs.push( {query: { _id: id }, set: { $addToSet: { playoffs: recs[i] }}} );
    }

    pipeObj.data.push( docs );
    pipeObj.prevStep += 1;
    callback( null, pipeObj );
  };

  return async.compose(
    createUpdateObjects,
    dataUtils.createObjects,
    dataUtils.readCsv
  );

}());


var inputSrc = { path: dataDir + '/SeriesPost.csv',
              headers: 1,
          dataTypeMap: [ 'round', 'teamIDwinner', 'lgIDwinner', 'teamIDloser', 'lgIDloser' ] };

var outputSettings = { type: 'mongodb', 
                        url: mongoUrl, 
                 collection: 'seasons' };

var pipeObj = { input: inputSrc,
               exitFn: null,
               output: outputSettings,
                 data: [],
             prevStep: -1 };


loadPlayoffRecords( pipeObj, function(err, result ) {
  if ( pipeObj.output ) {
    dataUtils.updateFns[ pipeObj.output.type ]( result.data[2], pipeObj.output, function( err, r) {
      console.log('[loadPlayoffRecords] PlayoffRecords loaded' );
    });
  }
});
