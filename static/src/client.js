
var urlPathAPI = '/api';

var $ = require("jquery");
var dates = require('./dates');

exports.createNote = function(time, name, text, callback) {
    var data = {
        name: name
    };

    if (!time) {
        data.time = new Date();
    }

    if (text) {
        data.text = text;
    }

    $.ajax({
        mathod:  'POST',
        url:     urlPathAPI + '/note',
        data:    data,
        success: function(data, status, xhr) {
            console.log(data);
            if (typeof callback === 'function') {
                callback();
            }
        }
    });
};

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
                callback(data);
            }
        }
    });
};

exports.getHistoryRecordsByDateRange = function(from, to, callback) {
    // TODO implement ajax request
};
