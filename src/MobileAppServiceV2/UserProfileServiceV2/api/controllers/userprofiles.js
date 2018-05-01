var express = require('express');
var router = express.Router();
var TYPES = require('tedious').TYPES;
var queries = require('./queries');

/* GET user profiles */
router.get('/', function (req, res) {
    req.sql(queries.SELECT_USER_PROFILES)
        .into(res);
});

/* GET single profile. */
router.get('/:user_profile_id', function (req, res) {
    req.sql(queries.SELECT_USER_PROFILE_BY_ID)
        .param('user_profile_id', req.params.user_profile_id, TYPES.NVarChar)
        .into(res, '{}');
});

/* POST create profile. */
router.post('/', function (req, res) {
    var qry = queries.INSERT_USER_PROFILE;
    req.sql(queries.INSERT_USER_PROFILE)
        .param('UserProfileJson', req.body, TYPES.NVarChar)
        .exec(res);
});

/* PUT update profile. */
router.put('/:id', function (req, res) {
    req.sql("EXEC UpdateProductFromJson @id, @json")
        .param('json', req.body, TYPES.NVarChar)
        .param('id', req.params.id, TYPES.Int)
        .exec(res);
});

/* DELETE delete profile. */
router.delete('/:id', function (req, res) {
    req.sql("DELETE Product WHERE ProductId = @id")
        .param('id', req.params.id, TYPES.Int)
        .exec(res);
});

module.exports = router;