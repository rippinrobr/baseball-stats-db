var async = require('async')
  , dataUtils = require('./data-utils') 
  , mongoUrl = process.env.MONGO_URL || 'mongodb://localhost:27017/giants'  
  , dataDir  = process.env.DATA_DIR  || '../giants'
  ;

// Creates the 'Pipeline' object with the pipeline specific functions in it.
// Returns the async.compose results
var loadPlyrDemog = (function() {
    var createPlayerObjects = function( pipeObj, callback ) {
    var docs = []
      , recs = pipeObj.data[ pipeObj.prevStep ]
      ;

    for(var i=0; i < recs.length; i++ ) {
      if ( recs[i] != null && recs[i].playerID != '' && recs[i].playerID != null ) {
        recs[i]._id       = recs[i].playerID;
        recs[i].allstar   = [];
        recs[i].positions = []

        recs[i].battingStats  = [];
        recs[i].fieldingStats = [];
        recs[i].pitchingStats = [];

        recs[i].pitchingPlayoffsStats = [];
        recs[i].battingPlayoffsStats  = [];
        recs[i].fieldingPlayoffsStats = [];

        delete recs[i].managerID;
        delete recs[i].lahman40ID;
        delete recs[i].lahman45ID;
        delete recs[i].holtzID;

        docs.push( recs[i] );
      }
    }

    pipeObj.data.push( docs );
    pipeObj.prevStep += 1;
    callback( null, pipeObj );
  };


  return  async.compose(
     createPlayerObjects, 
     dataUtils.createObjects,
     dataUtils.readRemoteCsv
  );

}());


var inputSrc = { path: dataUtils.baseGithubUrl + '/Master.csv',
              headers: 1,
          dataTypeMap: [ 'playerID','managerID','hofID','birthCountry','birthState','birthCity','deathCountry',
                         'deathState','deathCity','nameFirst','nameLast','nameNote','nameGiven','nameNick','bats',
                         'throws','debut','finalGame','college','lahman40ID','lahman45ID','retroID','holtzID',
                         'bbrefID'] };

var outputSettings = { type: 'mongodb', 
                        url: dataUtils.mongoUrl, 
                 collection: 'players' };

var pipeObj = { input: inputSrc,
               exitFn: null,
               output: outputSettings,
                 data: [],
             prevStep: -1 };


loadPlyrDemog( pipeObj, function(err, result ) {

  if ( pipeObj.output ) {

   dataUtils.emptyCollection( pipeObj.output.url, 'players', function(err, numDeleted) {
     console.log('[loadPlyrDemog]', 'Deleted', numDeleted, 'player documents.');
     dataUtils.storageFns[ pipeObj.output.type ]( result.data[2], pipeObj.output, function( err, r) {
      if( err ) return console.log( err );
       console.log('[loadPlyrDemog]', 'Player Demographics Loaded!');
       dataUtils.createIndex( mongoUrl, 'players', {nameLast:1});
     });
   });
  }
});
