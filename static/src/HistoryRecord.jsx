
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
                <h2>{this.props.name.trim()} {this.props.guid} - {this.props.time.toDateString()} {this.props.time.toTimeString()}</h2>
                <p className={this.props.amount > 0 ? "inflow" : "outflow"}>{this.props.amount}</p>
                <p className="balance">Balance: {this.props.balance}</p>
            </div>
        );
    }
});

exports.HistoryRecord = HistoryRecord;
