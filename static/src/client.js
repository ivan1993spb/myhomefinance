
var urlPathAPI = '/api';

var $ = require("jquery");
var dates = require('./dates');

exports.getNotesByDateRange = function(from, to, callback) {
    $.ajax({
        mathod:  'GET',
        url:     urlPathAPI + '/notes/range',
        data:    {
            from: dates.yyyymmdd(from),
            to:   dates.yyyymmdd(to)
        },
        success: function(data, status, xhr) {
            console.log(data);
            if (typeof callback === 'function') {
                callback();
            }
        }
    });
};
