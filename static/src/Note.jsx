
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

    getInitialState: function() {
        return {
            time: this.props.time,
            name: this.props.name,
            text: this.props.text,
        };
    },

    handleRemove: function() {
        console.log("called Note.handleRemove");
        console.log("request to remove note with id", id);
        this.props.handleRemove(this.props.id);
    },

    handleEdit: function() {
        console.log("called Note.handleEdit");
        this.props.handleEdit(this.props.id);
    },

    render: function() {
        return (
            <div>
                <p>{this.props.id}</p>
                <p>{this.state.time}</p>
                <p>{this.state.name}</p>
                <p>{this.state.text}</p>
                <p><button onClick={this.handleRemove}>delete</button></p>
                <p><button onClick={this.handleEdit}>edit</button></p>
            </div>
        );
    }
});

exports.Note = Note;
