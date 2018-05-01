var express = require('express');
var router = express.Router();

/* GET Healthcheck */
router.get('/', function(req, res) {
    res.json({
        message: 'healthcheck',
        status: "healthy"
    });
});

module.exports = router;