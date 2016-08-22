
var React = require('react');

function formatMoney(money) {
    money = Math.round(money*100) / 100;
    money = (""+money).indexOf(".") > -1 ? money + "00" : money + ".00";
    var dec = money.indexOf(".");
    return dec == money.length-3 || dec == 0 ? money : money.substring(0,dec+3);
}

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
        var transactionClass = this.props.amount > 0 ? "inflow" : "outflow",
            transactionText = "";

        if (this.props.amount > 0) {
            transactionText += "+";
        } else if (this.props.amount < 0) {
            transactionText += "-";
        }

        transactionText += formatMoney(Math.abs(this.props.amount));

        return (
            <div className="history-record">
                <h2>{this.props.name.trim()} {this.props.guid} - {this.props.time.toDateString()} {this.props.time.toTimeString()}</h2>
                <p>
                    <span className="balance">Balance: {this.props.balance}</span>
                    <span className={transactionClass}>{transactionText}</span>
                </p>
            </div>
        );
    }
});

exports.HistoryRecord = HistoryRecord;
