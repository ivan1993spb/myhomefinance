
var React = require('react');

var Note = React.createClass({
    propTypes: {
        id:           React.PropTypes.number.isRequired,
        time:         React.PropTypes.string.isRequired,
        name:         React.PropTypes.string.isRequired,
        text:         React.PropTypes.string.isRequired,
        handleEdit:   React.PropTypes.func.isRequired,
        handleRemove: React.PropTypes.func.isRequired
    },

    handleEdit: function() {
        console.log("called Note.handleEdit");
        this.props.handleEdit(this.props.id);
    },

    handleRemove: function() {
        console.log("called Note.handleRemove");
        this.props.handleRemove(this.props.id);
    },

    render: function() {
        return (
            <div>
                <p>{this.props.id}</p>
                <p>{this.props.time}</p>
                <p>{this.props.name.trim()}</p>
                <p>{this.props.text.trim()}</p>
                <p><button onClick={this.handleEdit}>edit</button></p>
                <p><button onClick={this.handleRemove}>delete</button></p>
            </div>
        );
    }
});

exports.Note = Note;
