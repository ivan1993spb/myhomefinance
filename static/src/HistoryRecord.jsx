
var React = require('react');

var HistoryRecord = React.createClass({
    displayName: "HistoryRecord",

    propTypes: {
        guid:    React.PropTypes.string.isRequired,
        time:    React.PropTypes.object.isRequired,
        name:    React.PropTypes.string.isRequired,
        amount:  React.PropTypes.number.isRequired,
        balance: React.PropTypes.number.isRequired
    },

    render: function() {
        return (
            <div className="history-record">
                <p>{this.props.guid}</p>
                <p>{this.props.time.toDateString()} {this.props.time.toTimeString()}</p>
                <p>{this.props.name.trim()}</p>
                <p className={this.props.amount > 0 ? "inflow" : "outflow"}>{this.props.amount}</p>
                <p className="balance">{this.props.balance}</p>
            </div>
        );
    }
});

exports.HistoryRecord = HistoryRecord;
