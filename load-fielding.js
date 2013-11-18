var async = require('async')
  , dataUtils = require('./data-utils') 
  , mongoUrl = process.env.MONGO_URL || 'mongodb://localhost:27017/giants'  
  , dataDir  = process.env.DATA_DIR  || '../giants'
  ;

// Creates the 'Pipeline' object with the pipeline specific functions in it.
// Returns the async.compose results
var loadFieldingStats = (function() {
  var createUpdateObjects = function( pipeObj, callback ) {
    var docs = []
      , recs = pipeObj.data[ pipeObj.prevStep ]
      ;

    for(var i=0; i < recs.length; i++ ) {
      var id = recs[i].playerID;

      delete recs[i].teamID;
      delete recs[i].lgID;

      docs.push( {query: { _id: id }, set: { $addToSet: { fieldingStats: recs[i] }}} );
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


var inputSrc = { path: dataDir + '/Fielding.csv',
              headers: 1,
          dataTypeMap: [ 'playerID', 'teamID', 'lgID', 'POS' ] };

var outputSettings = { type: 'mongodb', 
                        url: mongoUrl, 
                 collection: 'players' };

var pipeObj = { input: inputSrc,
               exitFn: null,
               output: outputSettings,
                 data: [],
             prevStep: -1 };


loadFieldingStats( pipeObj, function(err, result ) {
  if ( pipeObj.output ) {
    dataUtils.updateFns[ pipeObj.output.type ]( result.data[2], pipeObj.output, function( err, r) {
      console.log('[loadFielding] Fielding Stats loaded' );
    });
  }
});
