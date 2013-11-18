var async = require('async')
  , dataUtils = require('./data-utils') 
  ;

// Creates the 'Pipeline' object with the pipeline specific functions in it.
// Returns the async.compose results
var loadAllStars = (function() {

  var createUpdateObjects = function( pipeObj, callback ) {
    var docs = []
      , recs = pipeObj.data[ pipeObj.prevStep ]
      ;

    for(var i=0; i < recs.length; i++ ) {
      var id = recs[i].playerID;

      docs.push( {query: {_id: id}, set: { $addToSet: { allstar: recs[i].yearID}}} );
    }

    pipeObj.data.push( docs );
    pipeObj.prevStep += 1;
    callback( null, pipeObj );
  };


  return  async.compose(
     createUpdateObjects,
     dataUtils.createObjects,
     dataUtils.readRemoteCsv
    );

}());

var inputSrc = { path: dataUtils.baseGithubUrl + '/AllstarFull.csv',
              headers: 1,
          dataTypeMap: ['playerID', 'gameID', 'teamID'] };

var outputSettings = { type: 'mongodb', 
                        url: dataUtils.mongoUrl, 
                 collection: 'players' };

var pipeObj = { input: inputSrc,
               exitFn: null,
               output: outputSettings,
                 data: [],
             prevStep: -1 };


loadAllStars( pipeObj, function(err, result ) {
  if (err) return console.log( err );

  if( pipeObj.output ) {
    dataUtils.updateFns[ pipeObj.output.type ]( result.data[ result.prevStep ], pipeObj.output, function( err, r) {
     console.log('[loadAllStars] AllStars loaded' );
  });
}});

