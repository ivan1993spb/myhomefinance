
var React = require('react'),
    HistoryRecord = require("./HistoryRecord").HistoryRecord,
    dates = require("./dates"),
    client = require("./client");

var loadDaysLimit = 70;

var History = React.createClass({
    displayName: "History",

    propTypes: {
        dateFrom: React.PropTypes.object,
        loadDays: React.PropTypes.number
    },

    getDefaultProps: function() {
        var currentDate = dates.addDays(new Date(), 1);
        return {
            dateFrom: new Date(currentDate.getFullYear(), currentDate.getMonth(), currentDate.getDate(), 0, 0, 0, 0),
            loadDays: 1
        };
    },

    getInitialState: function() {
        return {
            dateTo:         this.props.dateFrom,
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
            // TODO use client.getHistoryRecordsByDateRange(from, to, callback)

            // test:
            callback([{
                guid:    "f536509e-bbc8-4da8-9028-32fab440e2fc",
                time:    "2016-08-08T10:59:15+03:00",
                name:    "name 1",
                amount:  545.56,
                balance: 596.22
            }, {
                guid:    "972fe78d-4571-437e-aa05-108dfd02633d",
                time:    "2016-08-07T10:59:15+03:00",
                name:    "name 2",
                amount:  -45.55,
                balance: 659
            }, {
                guid:    "673e9b4b-6960-4f2c-8e05-47101c4a536c",
                time:    "2016-08-06T10:59:15+03:00",
                name:    "name 3",
                amount:  33.25,
                balance: 5465.05
            }]);
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
                <p>Between {this.props.dateFrom.toDateString()} and {this.state.dateTo.toDateString()}</p>
                <hr />
                <div>
                    {historyRecords.length > 0 ? historyRecords : (this.state.loading ? "loading" : "empty")}
                </div>
                <hr />
                <p>Between {this.props.dateFrom.toDateString()} and {this.state.dateTo.toDateString()}</p>

                {this.state.loading ? "loading..." : "loaded: "+this.state.historyRecords.length}

                <button
                    onClick={this.handleLoadMore}
                    disabled={this.state.loading}>load more: {this.state.loadDays} days</button>

            </div>
        );
    }
});

exports.History = History;
