'use strict';

var SwaggerExpress = require('swagger-express-mw');
var bodyParser = require('body-parser');
var express = require('express');
var tediousExpress = require('express4-tedious');
var sqlConfig = require('./config/sqlConfig');

var app = express();
app.use(express.static('wwwroot'));

//Enable swagger
var swaggerConfig = {
  appRoot: __dirname
};

//
SwaggerExpress.create(swaggerConfig, function (err, swaggerExpress) {
  if (err) {
    throw err;
  }

  swaggerExpress.register(app);

  // configure app to use bodyParser() which will let us get the data from a POST easily
  app.use(bodyParser.text({ type: 'application/json' }))

  //setup db connection info
  app.use(function (req, res, next) {
    req.sql = tediousExpress(sqlConfig);
    next();
  });

  //setup routes for user profile and healthcheck
  app.use('/api/userprofiles', require('./api/controllers/userprofiles'));
  app.use('/api/healthcheck', require('./api/controllers/healthcheck'));

  // catch 404 and forward to error handler
  app.use(function (req, res, next) {
    var err = new Error('Not Found' + req.originalUrl);
    err.status = 404;
    next(err);
  });

  //set service port to 8080
  var port = process.env.PORT || 8080;

  //start handling request
  var server = app.listen(port, function() {
    console.log('Express server listening on port ' + port);
});

  module.exports = app; // for testing
});