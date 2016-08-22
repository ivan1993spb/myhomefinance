
var React = require('react'),
    dates = require("./dates"),
    client = require("./client"),
    ChartistGraph = require('react-chartist');

var loadDaysLimit = 70;

var Graphs = React.createClass({
    displayName: "Graphs",

    propTypes: {
        dateStart: React.PropTypes.object,
        loadDays:  React.PropTypes.number
    },

    getDefaultProps: function() {
        var currentDate = dates.addDays(new Date(), 1);
        return {
            dateStart: new Date(currentDate.getFullYear(), currentDate.getMonth(), currentDate.getDate(), 0, 0, 0, 0),
            loadDays:  10
        };
    },

    getInitialState: function() {
        return {
            dateTo:   this.props.dateStart,
            loadDays: this.props.loadDays < loadDaysLimit ? this.props.loadDays : loadDaysLimit,
            balances: [],
            inflow:   [],
            outflow:  [],
            loading:  true
        };
    },

    handleLoadMore: function() {
        var from = dates.addDays(this.state.dateTo, -this.state.loadDays),
            to   = this.state.dateTo;

        this.setState({
            loading: true,
            dateTo:  from
        });

        client.getHistoryRecordsByDateRange(from, to, function(historyRecords) {
            var newState = {
                loading: false
            };

            var balances = [],
                inflow = [],
                outflow = [];

            historyRecords.forEach(function(historyRecord) {
                var amount = Math.abs(historyRecord.amount);
                if (historyRecord.amount > 0) {
                    balances.push(historyRecord.balance-amount);
                    inflow.push(historyRecord.amount);
                    outflow.push(0);
                } else {
                    balances.push(historyRecord.balance);
                    inflow.push(0);
                    outflow.push(amount);
                }
            });

            if (balances.length) {
                balances.reverse();
                newState.balances = balances;
            }

            if (inflow.length) {
                inflow.reverse();
                newState.inflow = inflow;
            }

            if (outflow.length) {
                outflow.reverse();
                newState.outflow = outflow;
            }

            console.log("newState", newState);

            this.setState(newState);
        }.bind(this));
    },

    componentDidMount: function() {
        this.handleLoadMore();
    },

    render: function() {
        console.log(this.state);

        var labels = new Array(this.state.balances.length);

        labels.fill("", 0, this.state.balances.length);
        labels[0] = this.state.dateTo.toDateString();
        labels[labels.length-1] = this.state.dateTo.toDateString();
        console.log(labels);
        var data = {
            labels: labels,
            series: [
                this.state.balances,
                this.state.inflow,
                this.state.outflow
            ]
        };

        var options = {
            axisX: {
                labelInterpolationFnc: function(value, index) {
                    return value;
                }
            },
            stackBars: true
        };

        var type = 'Bar';

        var style = {
            height: '500px'
        };

        return (
            <div>
                <p>Between <i>{this.state.dateTo.toDateString()}</i> and {this.props.dateStart.toDateString()}</p>
                <div><button type="button">load more</button></div>
                <ChartistGraph data={data} options={options} type={type} style={style}/>
            </div>
        );
    }
});

exports.Graphs = Graphs;
