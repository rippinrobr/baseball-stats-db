var async = require('async')
  , csv = require('csv')
  , http = require('https')
  , MongoClient = require('mongodb').MongoClient
  , Server = require('mongodb').Server
  , lookupFns = {}
  , queryFns = {}
  , storageFns = {}
  , updateFns = {}
  , mongoUrl = process.env.MONGO_URL || 'mongodb://localhost:27017/bdb'  
  , dataDir  = process.env.DATA_DIR  || '../giants'
  , baseGithubUrl = process.env.BASE_GITHUB_URL || 'https://raw.github.com/chadwickbureau/baseballdatabank/2012update/official'
  ;

var convertDataTypes = function( fields, prop, val, floatCols ) {
  if ( fields && fields.indexOf( prop ) > -1 ) {
    return val;
  } else {
    if ( floatCols  && floatCols.indexOf( prop ) > -1 ) {
      return ( val ) ? parseFloat( val ) : null; 
    } else {
      return ( val ) ? parseInt( val ) : null;
    }
  }
};

exports.createIndex = function( mongoUrl, collName, index ) {
  mongoClient = new MongoClient();

  mongoClient.connect( mongoUrl, function( err, db ) {
    db.collection( collName ).ensureIndex(index, function(err, results) {
     console.log( '[createIndex]', 'created index on ', collName,  results );
     db.close();
    });
  });
};

exports.emptyCollection = function( mongoUrl, collName, cb ) {
  mongoClient = new MongoClient();

  mongoClient.connect( mongoUrl, function( err, db ) {
    db.collection( collName ).remove({}, function(err, numRemoved) {
     cb( err, numRemoved );
     db.close();
    });
  });
};

var addCsvSourceToPipelineData = function( csvSrcObj, pipeObj, cb) {
  csvSrcObj.to.array( function( data, count ) {
     pipeObj.data.push({ cols: data[0], data : data.slice(1), count: count });
     pipeObj.prevStep += 1;
     cb( null, pipeObj );
   });
};

exports.readCsv = function(pipeObj, callback) { 
  addCsvSourceToPipelineData( csv().from( pipeObj.input.path ), pipeObj, callback );
};

exports.readRemoteCsv = function( pipeObj, callback) {
  http.get( pipeObj.input.path, function( res ) {
    var body = ""
      ;
     res.setEncoding('utf8');
     res.on('data',  function( chunk ) {
       body += chunk
     });

     res.on('end', function() { 
       addCsvSourceToPipelineData( csv().from.string( body ), pipeObj, callback );
     });

  });
};

exports.createObjects = function( pipeObj, callback) {
  var objects = []
    , cols    = pipeObj.data[ pipeObj.prevStep ].cols
    , data    = pipeObj.data[ pipeObj.prevStep ].data
    ;

  for(var i=0 ; i < data.length; i++ ) {
    var obj = {}
    for(var j=0; j < cols.length; j++ ) {
      var propName = cols[j];

      obj[ propName ] = convertDataTypes( pipeObj.input.dataTypeMap, propName, data[i][j], pipeObj.input.floatColMap ); 
    }
    objects.push( obj );
  };

  pipeObj.prevStep += 1;
  pipeObj.data.push( objects );
  callback( null, pipeObj );
};

exports.loadData = function( pipeObj, pipeLine, savingFn, cb) {
  pipeLine( pipeObj, function(err, result ) {
    console.log( err );
 
    if( pipeObj.output ) {
      storageFns[ pipeObj.output.type ]( result.data[2], pipeObj.output, function( err, r) {
        console.log( pipeObj.input.path, 'Loaded!');
      });
    }
  });
};

exports.loadInput = function( pipeObj, callback) {
  queryFns[ pipeObj.input.type ]( pipeObj, callback );  
};

queryFns['csv'] = function( pipeObj, cb ) { 
 readCsv( pipeObj, cb ); 
}

queryFns['remoteCsv'] = function( pipeObj, cb ) { 
  this.readCsv( pipeObj, cb ); 
}

lookupFns['mongodb'] = function( settings, cb ) {
  var mongoClient = new MongoClient()
    ;

  mongoClient.connect( settings.url, function( err, db ) {
    if (err) throw err;

    db.collection( settings.collection ).find( settings.query, settings.projection ).toArray(function( err, result ) {
      var cols = []
        ;

      for(var i in result[0] ) { cols.push( i ); }
      cb( err, { cols: cols, data : result , count: result.length });
      db.close();
    });
  });
};

storageFns['mongodb'] = function( data, settings, cb ) {
  mongoClient = new MongoClient();

  mongoClient.connect( settings.url, function( err, db ) {
    if (err) throw err;

    db.collection( settings.collection ).insert( data, function(err, result ) {
      cb( err, result );
      db.close();
    });
  });
};

updateFns['mongodb'] = function( data, settings, cb ) {
  mongoClient = new MongoClient();

  mongoClient.connect( settings.url, function( err, db ) {
    if (err) throw err;
    
    async.map( data, function( item,cb ) {
      db.collection( settings.collection )
        .update( item.query, item.set, function( err, result ) {
          cb( err, result );
    });
    }, 
    function(err, results) {
      db.close();   
      cb( err, results );
    });
  });
};

exports.mongoUrl      = mongoUrl;
exports.dataDir       = dataDir;
exports.baseGithubUrl = baseGithubUrl;

exports.lookupFns  = lookupFns;
exports.queryFns   = queryFns;
exports.storageFns = storageFns;
exports.updateFns  = updateFns;
