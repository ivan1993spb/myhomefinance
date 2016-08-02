
var React = require('react');
var dates = require('./dates');

exports.noteList = React.createClass({
    propTypes: {
        from: React.PropTypes.object.isRequired,
        days: React.PropTypes.number.isRequired
    },

    getInitialState: function() {
        return {
            from:  this.props.from,
            days:  this.props.days,
            notes: []
        };
    },

    doLoadMore: function(a, b, callback) {
        console.log("doLoadMore");
        if (typeof callback === 'function') {
            callback([
                {id: 1, time: "2016-08-02T13:55:32Z", name: "name", text: "text"},
                {id: 2, time: "2016-08-01T13:50:00Z", name: "eman", text: "txet"}
            ]);
        }
    },

    handleLoadMore: function() {
        console.log("state22", this.state);
        this.doLoadMore(this.state.from, dates.addDays(this.state.from, this.state.days), function(notes) {
            console.log("okok123");
            this.setState({
                from:  dates.addDays(this.state.from, this.state.days),
                notes: this.state.notes.concat(notes)
            });
        }.bind(this));
    },

    componentDidMount: function() {
        console.log("ok");
        this.handleLoadMore();
    },

    handleRemove: function(i, id) {
        console.log(i, id)
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
                <button onClick={this.handleLoadMore.bind(this)}>load more</button>
            </div>
        );
    }
});
