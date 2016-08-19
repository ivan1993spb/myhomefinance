
var React = require('react');

var NoteSmall = React.createClass({
    displayName: "NoteSmall",

    propTypes: {
        id:   React.PropTypes.number.isRequired,
        time: React.PropTypes.object.isRequired,
        name: React.PropTypes.string.isRequired,
        text: React.PropTypes.string.isRequired
    },

    render: function() {
        return (
            <div style={{'backgroundColor': '#ff0'}}>
                <p>{this.props.time.toDateString()} {this.props.time.toTimeString()}</p>
                <p>{this.props.name}</p>
                <p>{this.props.text}</p>
            </div>
        );
    }
});

exports.NoteSmall = NoteSmall;
