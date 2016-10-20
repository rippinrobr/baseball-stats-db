var async = require('async')
  , dataUtils = require('./data-utils') 
  , MongoClient = require('mongodb').MongoClient
  , mongoUrl = process.env.MONGO_URL || 'mongodb://localhost:27017/giants'  
  ;

// Returns the async.compose results
var loadManagersPlayoffRecord = (function() {
  
  var  getPlayoffManagers = function( pipeObj, callback ) {
    var docs = [];

    pipeObj.data[ pipeObj.prevStep ].forEach( function( rec ) {
        var numManagers = rec.managers.length
          , mgrId = ''
          ;

        if ( numManagers == 1 ) {
          mgrId = rec.managers[0]._id;
        } else {
          rec.managers.forEach( function( mgr ) {
            if ( mgr.record.inseason == numManagers ) mgrId = mgr._id;
          });
        }      

        if ( mgrId != '' ) {
          rec.playoffs.forEach( function( r ) { 
            docs.push( { query: {_id: mgrId}, set: { $addToSet: { playoffs: r }} } );
          });
        } 
    });

    pipeObj.data.push( docs );
    pipeObj.prevStep += 1;
    callback( null, pipeObj );
  };

  var getManagers = function( pipeObj, callback ) {
    var mongoClient = new MongoClient()
      ;

    mongoClient.connect( pipeObj.input.url, function( err, db ) {   
      db.collection( pipeObj.input.collection ).find({playoffs: {$exists:true} }, {managers:1, playoffs:1}).toArray(
        function( err, results ) {
          if (err) cb( err, pipeObj );
             
          pipeObj.data.push( results );
          pipeObj.prevStep += 1;
          callback( err, pipeObj );
          db.close();
         }); 
    });

  };

  return async.compose( 
    getPlayoffManagers,
    getManagers
  );

}());

var inputSrc = { type: 'mongodb', url: mongoUrl, collection: 'seasons' }

var outputSettings = { type: 'mongodb', url: mongoUrl, collection: 'managers' };

var pipeObj  = { input: inputSrc, 
                 output: outputSettings,
                 data:  [], 
                 prevStep: -1 };

loadManagersPlayoffRecord( pipeObj, function(err, result ) {
  if ( err ) console.log( err );

  dataUtils.updateFns[ pipeObj.output.type ]( result.data[ result.prevStep ], pipeObj.output, 
    function( err, r) {
      console.log('[loadManagersPlayoffRecord] Manager Playoff Records loaded' );
  });
});
