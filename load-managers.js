var async = require('async')
  , dataUtils = require('./data-utils') 
  ;

// Creates the 'Pipeline' object with the pipeline specific functions in it.
// Returns the async.compose results
var loadManagers = (function() {

  var createMgrDocuments = function( pipeObj, callback ) {
    var docs = []
      , input = pipeObj.data[ pipeObj.prevStep ]
      , mgrs = {}
    ;
  
    for(var i=0; i < input.length; i++ ) {
      var id = input[i].managerID;
      
      delete input[i].managerID;

      if (mgrs[ id ] == null ) {
        mgrs[ id ] = { record: [] }
      }
      mgrs[ id ].record.push( input[i] );
    };

    for(var i in mgrs ) { 
      mgrs[i]['_id'] = i; 
      docs.push( mgrs[ i ] ); 
    }
    pipeObj.data.push( docs ); 
    pipeObj.prevStep += 1;
    callback( null, pipeObj );
  };


  var loadManagers =  async.compose( createMgrDocuments, 
      dataUtils.createObjects,
      dataUtils.readRemoteCsv );  

  return loadManagers;
}());

// Pipeline Config
var inputSrc = { path: dataUtils.baseGithubUrl + '/Managers.csv',
                 type: 'remoteCsv',
              headers: 1, 
          dataTypeMap: ['teamID', 'lgID', 'managerID', 'plyrMgr']};

var outputSettings = { type: 'mongodb', url: dataUtils.mongoUrl, collection: 'managers' };

var pipeObj  = { input: inputSrc, 
                 exitFn: null, // this can/will be used to 'chain' these pipelines together.
                 output: outputSettings,
                 data:  [], 
                 prevStep: -1 };

loadManagers( pipeObj, function( err, result ) {

    if ( pipeObj.output ) {
      var data = pipeObj.data[ pipeObj.prevStep ]
        ;

    dataUtils.emptyCollection( pipeObj.output.url, 'managers', function(err, numDeleted) { 
     console.log('[loadManagers]', 'Deleted', numDeleted, 'manager documents.');
     dataUtils.storageFns[ pipeObj.output.type ]( data, pipeObj.output, function(err, res) {
        if( err) throw err;
        console.log('[loadManagers] Managers loaded' );
        dataUtils.createIndex( dataUtils.mongoUrl, 'managers', { nameLast:1 });
        dataUtils.createIndex( dataUtils.mongoUrl, 'managers', { 'record.yearID':1});
      });
     });
    }

    if ( result.exitFn ) {
      console.log( 'should call the exitFn' );
    } 
  });
