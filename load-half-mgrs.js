var async = require('async')
  , dataUtils = require('./data-utils') 
  , mongoUrl = process.env.MONGO_URL || 'mongodb://localhost:27017/giants'  
  , dataDir  = process.env.DATA_DIR  || '../giants'
  ;

// Creates the 'Pipeline' object with the pipeline specific functions in it.
// Returns the async.compose results
var loadHalfManagers = (function() {

  var createUpdateObjects = function( pipeObj, callback ) {
    var docs = []
      , recs = pipeObj.data[ pipeObj.prevStep ]
      ;

    for(var i=0; i < recs.length; i++ ) {
      var id = recs[i].managerID;
      delete recs[i].managerID;

      docs.push( {query: {_id: id}, set: { $addToSet: { record : recs[i] }}} );
    }

    pipeObj.data.push( docs );
    pipeObj.prevStep += 1;
    callback( null, pipeObj );
  };

  return async.compose( 
      createUpdateObjects,
      dataUtils.createObjects, 
      dataUtils.readCsv );  
}());

// Configuration for the pipeline
var inputSrc = { type: 'csv', path: dataDir + '/ManagersHalf.csv', headers:1, dataTypeMap: ['teamID', 'lgID', 'managerID', 'plyrMgr'] };
var outputSettings = { type: 'mongodb', url: mongoUrl, collection: 'managers' };
var pipeObj  = { input: inputSrc, 
                 exitFn: null, // this can/will be used to 'chain' these pipelines together.
                 output: outputSettings,
                 data:  [], 
                 prevStep: -1 };

loadHalfManagers( pipeObj, function( err, result ) {
  if (err) return console.log( err );

  if( pipeObj.output ) {
    dataUtils.updateFns[ pipeObj.output.type ]( result.data[2], pipeObj.output, function( err, r) {
     console.log('[loadHalfManagers] ManagersHalf loaded' );
  });
  }
});

