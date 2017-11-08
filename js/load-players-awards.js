var async = require('async')
  , dataUtils = require('./data-utils') 
  ;

// Creates the 'Pipeline' object with the pipeline specific functions in it.
// Returns the async.compose results
var loadPlayerAwards = (function() {
  
  var  getAwardWinningPlayers = function( pipeObj, callback ) {
    var currStep = pipeObj.prevStep + 1;

    if ( pipeObj.lookups[ currStep ] != {} )  {
      var lookup = pipeObj.lookups[ currStep ]
        , lookupObj = {}
        , docs = []
        ;

      dataUtils.lookupFns[ lookup.type ]( pipeObj.lookups[ currStep ], 
        function( err, lookupData ) {
          lookupData.data.forEach( function( ld ) { 
            lookupObj[ ld._id ] = [];
            ld.battingStats.forEach( function( r ) { 
              lookupObj[ ld._id ].push( r.yearID ); 
            });
          });

          pipeObj.data[ pipeObj.prevStep ].forEach( function( rec ) {
            if ( lookupObj[ rec.playerID ] && lookupObj[ rec.playerID ].indexOf( rec.yearID ) > -1 ) {
              var id = rec.playerID;

              docs.push( { query: {_id: id}, set: { $addToSet: { awards: rec }} } );
            } 
          });

        pipeObj.data.push( docs );
        pipeObj.prevStep = currStep;
        callback( null, pipeObj );
      });

    }
  };
  return async.compose( 
    getAwardWinningPlayers,
    dataUtils.createObjects,
    dataUtils.readRemoteCsv
  );

}());

var inputSrc = { path: dataUtils.baseGithubUrl + '/AwardsPlayers.csv',
              headers: 1,
                 type: 'csv',
          dataTypeMap: [ 'playerID', 'awardID', 'lgID', 'tie', 'notes' ],
          floatColMap: null }

var outputSettings = { type: 'mongodb', url: dataUtils.mongoUrl, collection: 'players' };

var pipeObj  = { input: inputSrc, 
                 lookups: [ {}, {}, 
                           {step: 2, type: 'mongodb', 
                             url: dataUtils.mongoUrl, 
                            collection: 'players', query: {}, 
                            projection: {_id:1, 'battingStats':1} } ],
                 output: outputSettings,
                 data:  [], 
                 prevStep: -1 };


loadPlayerAwards( pipeObj, function(err, result ) {
  if(err) return  console.log( err );
  
  dataUtils.updateFns[ pipeObj.output.type ]( result.data[2], pipeObj.output, 
    function( err, r) {
      console.log('[loadPlayersAwards] Players Awards loaded' );
  });
});
