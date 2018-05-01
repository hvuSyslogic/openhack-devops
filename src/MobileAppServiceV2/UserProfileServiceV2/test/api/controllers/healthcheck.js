var should = require('should');
var request = require('supertest');
var server = require('../../../app');

describe('controllers', function() {

  describe('healthcheck', function() {

    describe('GET /', function() {

      it('should return a default json', function(done) {

        request(server)
          .get('/')
          .set('Accept', 'application/json')
          .expect('Content-Type', /json/)
          .expect(200)
          .end(function(err, res) {
            should.not.exist(err);

            res.body.should.eql('{\
                "message": "healthcheck",\
                "status": "healthy"\
            }');

            done();
          });
      });

    });

  });

});
