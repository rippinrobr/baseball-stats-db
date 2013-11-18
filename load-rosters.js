var async = require('async')
  , dataUtils = require('./data-utils') 
  , MongoClient = require('mongodb').MongoClient
  , mongoUrl = process.env.MONGO_URL || 'mongodb://localhost:27017/giants'  
  ;

// Returns the async.compose results
var loadRosters = (function() {

  var getSeasonStats = function( year, stats ) {
    var match = []

    if ( stats != null ) {
      for(var i=0; i < stats.length; i++ ) {
        var s = stats[i];
        if ( s.yearID == year ) {
          match.push( s );
        };
      };
    }
    return match;
  };

  var  getPlayersOnTeam = function( pipeObj, callback ) {
    var docs = [];

    pipeObj.data[ pipeObj.prevStep ].forEach( function( rec ) {

      var obj = { nameLast: rec.nameLast, nameFirst: rec.nameFirst, bats: rec.bats, throws: rec.throws,
                  battingStats : getSeasonStats( rec.appearances.yearID, rec.battingStats ), 
                  fieldingStats: getSeasonStats( rec.appearances.yearID, rec.fieldingStats ),
                  pitchingStats: getSeasonStats( rec.appearances.yearID, rec.pitchingStats ), 
                  battingPlayoffsStats : getSeasonStats( rec.appearances.yearID, rec.battingPlayoffsStats, rec ), 
                  fieldingPlayoffsStats: getSeasonStats( rec.appearances.yearID, rec.fieldingPlayoffsStats ),
                  pitchingPlayoffsStats: getSeasonStats( rec.appearances.yearID, rec.pitchingPlayoffsStats ), 
                  appearances: rec.appearances };
      docs.push( { query: {_id: rec.appearances.yearID}, set: { $addToSet: { players: obj }} } );
    });

    pipeObj.data.push( docs );
    pipeObj.prevStep += 1;
    callback( null, pipeObj );
  };
  

  var getPlayers = function( pipeObj, callback ) {
   // do aggregate call here
    var mongoClient = new MongoClient()
      ;

    mongoClient.connect( pipeObj.input.url, function( err, db ) {   
      // unwind on appearances
      db.collection( pipeObj.input.collection )
        .aggregate([ { $project: { nameLast:1, nameFirst:1, bats:1, throws: 1, battingStats:1,
                                   fieldingStats:1, pitchingStats:1 , battingPlayoffsStats:1, 
                                   fieldingPlayoffsStats:1, pitchingPlayoffsStats:1,
                                   appearances:1, allstar:1 }}, 
                     { $unwind: '$appearances' }],
          function( err, results ) {

            pipeObj.data.push( results );
            pipeObj.prevStep += 1;
            callback( err, pipeObj );
            db.close();
          });
    });

  };

  return async.compose( 
    getPlayersOnTeam,
    getPlayers
  );

}());

var inputSrc = { type: 'mongodb', url: mongoUrl, collection: 'players' }

var outputSettings = { type: 'mongodb', url: mongoUrl, collection: 'seasons' };

var pipeObj  = { input: inputSrc, 
                 output: outputSettings,
                 data:  [], 
                 prevStep: -1 };


loadRosters( pipeObj, function(err, result ) {

  dataUtils.updateFns[ pipeObj.output.type ]( result.data[ result.prevStep ], pipeObj.output, 
    function( err, r) {
      console.log('[loadRosters] Rosters loaded' );
  });
});

