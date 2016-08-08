
var React = require('react'),
    HistoryRecord = require("./HistoryRecord").HistoryRecord,
    dates = require("./dates"),
    client = require("./client");

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
            loading:        true
        };
    },

    increaseLoadDays: function(loadDays) {
        if (loadDays >= loadDaysLimit) {
            return loadDays;
        }
        return loadDays * 2;
    },

    doLoadMore: function(from, to, callback) {
        console.log("called History.doLoadMore");
        console.log("dates 1", dates.yyyymmdd(from), dates.yyyymmdd(to));
        console.log("dates 2", from, to);

        if (typeof callback === 'function') {
            client.getHistoryRecordsByDateRange(from, to, function(data) {
                callback(data);
            });
        }
    },

    handleLoadMore: function() {
        var from = dates.addDays(this.state.dateTo, -this.state.loadDays),
            to   = this.state.dateTo;

        this.setState({
            loading: true,
            dateTo:  from
        });

        this.doLoadMore(from, to, function(historyRecords) {
            console.log("load more: setting component state" );
            var newState = {
                loading: false
            };

            if (historyRecords.length) {
                newState.historyRecords = this.state.historyRecords.concat(historyRecords);
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
        console.log("ok");
        this.handleLoadMore();
    },

    render: function() {
        console.log("render history");
        var historyRecords = this.state.historyRecords.map(function(historyRecord, index) {
            return (
                <HistoryRecord
                    guid={historyRecord.guid}
                    key={index}
                    time={historyRecord.time}
                    name={historyRecord.name}
                    amount={historyRecord.amount}
                    balance={historyRecord.balance}
                />
            );
        }.bind(this));

        return (
            <div>
                <h2>History</h2>
                <p>Between <i>{this.state.dateTo.toDateString()}</i> and {this.props.dateStart.toDateString()}</p>
                <hr />
                <div>
                    {historyRecords.length > 0 ? historyRecords : (this.state.loading ? "loading" : "empty")}
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
