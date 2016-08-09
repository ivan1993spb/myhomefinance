
var urlPathAPI = '/api';

var $ = require("jquery");
var dates = require('./dates');

exports.createNote = function(time, name, text, success, error) {
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
            if (typeof success === 'function') {
                success();
            }
        }
    });
};

exports.deleteNote = function(id, success, error) {
    // TODO implement request to delete note
    console.log("delete note id", id);
    if (typeof success === 'function') {
        success();
    }
};

exports.getNotesByDateRange = function(from, to, success, error) {
    $.ajax({
        mathod:  'GET',
        url:     urlPathAPI + '/notes/range',
        data:    {
            from: dates.yyyymmdd(from),
            to:   dates.yyyymmdd(to)
        },
        success: function(data, status, xhr) {
            console.log(data);
            if (typeof success === 'function') {
                success(data);
            }
        }
    });
};

exports.getHistoryRecordsByDateRange = function(from, to, success, error) {
    $.ajax({
        mathod:  'GET',
        url:     urlPathAPI + '/history/range',
        data:    {
            from: dates.yyyymmdd(from),
            to:   dates.yyyymmdd(to)
        },
        success: function(data, status, xhr) {
            console.log(data);
            if (typeof success === 'function') {
                success(data);
            }
        }
    });
};
