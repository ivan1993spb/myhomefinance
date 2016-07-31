
var React = require('react');

exports.noteList = React.createClass({
    getInitialState: function() {
        return {
            from:  0,
            days: 20,
            notes: [
                {id: 1, name: "first", text: "text first text"},
                {id: 2, name: "second", text: "text second text"},
                {id: 3, name: "third", text: "text third text"},
                {id: 4, name: "fourth", text: "text fourth text"}
            ]
        };
    },

    handleRemove: function() {},

    handleLoadMore: function() {
        this.setState({
        });
    },

    render: function() {
        var notes = this.state.notes.map(function(note, i) {
            return (
                <div>
                    <p>{note.id}</p>
                    <p>{note.name}</p>
                    <p>{note.text}</p>
                    <p><button onClick={this.handleRemove.bind(this, i, note.id)}>delete</button></p>
                </div>
            );
        }.bind(this));

        return (
            <div>
                <h1>Notes list {this.state.page}</h1>
                <div>
                    {notes}
                </div>
                <button onClick={this.handleLoadMore.bind(this)}>load</button>
            </div>
        );
    }
});
