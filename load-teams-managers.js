var async = require('async')
  , dataUtils = require('./data-utils') 
  , MongoClient = require('mongodb').MongoClient
  ;

// Returns the async.compose results
var loadSeasonManagers = (function() {
  
  var  getTeamManagers = function( pipeObj, callback ) {
    var docs = [];

    pipeObj.data[ pipeObj.prevStep ].forEach( function( rec ) {
      docs.push( { query: {_id: rec.record.yearID + "_" + rec.record.teamID}, 
                     set: { $addToSet: { managers: rec }} } );
    });

    pipeObj.data.push( docs );
    pipeObj.prevStep += 1;
    callback( null, pipeObj );
  };
  

  var getManagers = function( pipeObj, callback ) {
   // do aggregate call here
    var mongoClient = new MongoClient()
      ;

    mongoClient.connect( pipeObj.input.url, function( err, db ) {   
      // unwind on appearances
      db.collection( pipeObj.input.collection )
        .aggregate([ { $project: { nameLast:1, nameFirst:1, bats:1, record:1 }}, {$unwind: '$record'} ], 
          function( err, results ) {

            pipeObj.data.push( results );
            pipeObj.prevStep += 1;
            callback( err, pipeObj );
            db.close();
          });
    });

  };

  return async.compose( 
    getTeamManagers,
    getManagers
  );

}());

var inputSrc = { type: 'mongodb', url: dataUtils.mongoUrl, collection: 'managers' }

var outputSettings = { type: 'mongodb', url: dataUtils.mongoUrl, collection: 'teams' };

var pipeObj  = { input: inputSrc, 
                 output: outputSettings,
                 data:  [], 
                 prevStep: -1 };

loadSeasonManagers( pipeObj, function(err, result ) {
  if ( err ) console.log( err );

  dataUtils.updateFns[ pipeObj.output.type ]( result.data[ result.prevStep ], pipeObj.output, 
    function( err, r) {
      console.log('[loadSeasonManagers] SeasonManagers loaded' );
  });
});

