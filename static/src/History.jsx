
var React = require('react'),
    HistoryRecord = require("./HistoryRecord").HistoryRecord,
    NoteSmall = require("./NoteSmall").NoteSmall,
    dates = require("./dates"),
    client = require("./client"),
    parallel = require('async/parallel');

var loadDaysLimit = 70;

var History = React.createClass({
    displayName: "History",

    propTypes: {
        dateStart: React.PropTypes.object,
        loadDays:  React.PropTypes.number
    },

    getDefaultProps: function() {
        var currentDate = dates.addDays(new Date(), 1);
        return {
            dateStart: new Date(currentDate.getFullYear(), currentDate.getMonth(), currentDate.getDate(), 0, 0, 0, 0),
            loadDays:  1
        };
    },

    getInitialState: function() {
        return {
            dateTo:         this.props.dateStart,
            loadDays:       this.props.loadDays < loadDaysLimit ? this.props.loadDays : loadDaysLimit,
            historyRecords: [],
            notes:          [],
            loading:        true
        };
    },

    increaseLoadDays: function(loadDays) {
        if (loadDays >= loadDaysLimit) {
            return loadDays;
        }
        return loadDays * 2;
    },

    handleLoadMore: function() {
        var from = dates.addDays(this.state.dateTo, -this.state.loadDays),
            to   = this.state.dateTo;

        this.setState({
            loading: true,
            dateTo:  from
        });

        parallel({
            notes: function(callback) {
                client.getNotesByDateRange(from, to, function(notes) {
                    callback(null, notes);
                });
            },
            historyRecords: function(callback) {
                client.getHistoryRecordsByDateRange(from, to, function(historyRecords) {
                    callback(null, historyRecords);
                });
            }
        }, function(err, results) {
            var newState = {
                loading: false
            };

            if (results.historyRecords.length) {
                newState.historyRecords = this.state.historyRecords.concat(results.historyRecords);
            }

            if (results.notes.length) {
                newState.notes = this.state.notes.concat(results.notes);
            }

            if (results.historyRecords.length || results.notes.length) {
                // return initial value
                if (this.state.loadDays != this.props.loadDays) {
                    newState.loadDays = this.props.loadDays;
                }
            } else {
                newState.loadDays = this.increaseLoadDays(this.state.loadDays);
            }

            this.setState(newState);
        }.bind(this));
    },

    componentDidMount: function() {
        this.handleLoadMore();
    },

    renderHistoryRecord: function(historyRecord, key) {
        return (
            <HistoryRecord
                guid={historyRecord.guid}
                time={historyRecord.time}
                name={historyRecord.name}
                amount={historyRecord.amount}
                balance={historyRecord.balance}
                key={key}
            />
        );
    },

    renderNote: function(note, key) {
        return (
            <NoteSmall
                id={note.id}
                time={note.time}
                name={note.name}
                text={note.text}
                key={key}
            />
        );
    },

    renderHostoryRecordsAndNotes: function() {
        var i = 0,
            j = 0;

        var note, historyRecord;

        var renderedHostoryRecordsAndNotes = [];
        var key = 0;

        while (i < this.state.historyRecords.length && j < this.state.notes.length) {
            historyRecord = this.state.historyRecords[i];
            note = this.state.notes[j];

            if (historyRecord.time > note.time) {
                renderedHostoryRecordsAndNotes[key] = this.renderHistoryRecord(historyRecord, key);
                i++;
            } else {
                renderedHostoryRecordsAndNotes[key] = this.renderNote(note, key);
                j++;
            }

            key++;
        }

        if (i < this.state.historyRecords.length) {
            this.state.historyRecords.slice(i).forEach(function(historyRecord) {
                renderedHostoryRecordsAndNotes[key] = this.renderHistoryRecord(historyRecord, key++);
            }.bind(this));
        } else if (j < this.state.notes.length) {
            this.state.notes.slice(j).forEach(function(note) {
                renderedHostoryRecordsAndNotes[key] = this.renderNote(note, key++);
            }.bind(this));
        }

        return renderedHostoryRecordsAndNotes;
    },

    render: function() {
        var historyRecordsAndNotes = this.renderHostoryRecordsAndNotes();

        return (
            <div>
                <h2>History</h2>
                <p>Between <i>{this.state.dateTo.toDateString()}</i> and {this.props.dateStart.toDateString()}</p>
                <hr />
                <div>
                    {historyRecordsAndNotes.length > 0 ? historyRecordsAndNotes : (this.state.loading ? "loading" : "empty")}
                </div>
                <hr />
                <p>Between <i>{this.state.dateTo.toDateString()}</i> and {this.props.dateStart.toDateString()}</p>

                {this.state.loading ? "loading..." : "loaded: "+this.state.historyRecords.length}

                <button
                    onClick={this.handleLoadMore}
                    disabled={this.state.loading}>load more: {this.state.loadDays} days</button>

            </div>
        );
    }
});

exports.History = History;
