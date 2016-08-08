
var React = require('react');

var HistoryRecord = React.createClass({
    displayName: "HistoryRecord",

    propTypes: {
        guid:    React.PropTypes.string.isRequired,
        time:    React.PropTypes.string.isRequired,
        name:    React.PropTypes.string.isRequired,
        amount:  React.PropTypes.number.isRequired,
        balance: React.PropTypes.number.isRequired
    },

    render: function() {
        return (
            <div>
                <p>{this.props.guid}</p>
                <p>{this.props.time}</p>
                <p>{this.props.name.trim()}</p>
                <p>{this.props.amount}</p>
                <p>{this.props.balance}</p>
            </div>
        );
    }
});

exports.HistoryRecord = HistoryRecord;