
var urlPathAPI = '/api';

var $ = require("jquery");
var dates = require('./dates');

function Note(id, time, name, text) {
    this.id = id;
    if (typeof time == 'object') {
        this.time = time;
    } else {
        this.time = new Date(time);
    }
    this.name = name;
    this.text = text;
};

function HistoryRecord(guid, time, name, amount, balance) {
    this.guid = guid;
    if (typeof time == 'object') {
        this.time = time;
    } else {
        this.time = new Date(time);
    }
    this.name = name;
    this.amount = amount;
    this.balance = balance;
};

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
                success(data.map(function(rawNote) {
                    with (rawNote) {
                        return new Note(id, time, name, text);
                    }
                }));
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
                success(data.map(function(rawHistoryRecord) {
                    with (rawHistoryRecord) {
                        return new HistoryRecord(guid, time, name, amount, balance);
                    }
                }));
            }
        }
    });
};
