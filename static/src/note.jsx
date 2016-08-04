
var React = require('react');
var dates = require('./dates');

var Note = React.createClass({
    propTypes: {
        id:             React.PropTypes.number.isRequired,
        name:           React.PropTypes.string.isRequired,
        text:           React.PropTypes.string.isRequired,
        removeCallback: React.PropTypes.func.isRequired
    },

    handleRemove: function(id) {
        console.log("called Note.handleRemove");
        console.log("request to remove note with id", id);
        this.props.removeCallback();
    },

    render: function() {
        return (
            <div>
                <p>{this.props.id}</p>
                <p>{this.props.name}</p>
                <p>{this.props.text}</p>
                <p><button onClick={this.handleRemove.bind(this, this.props.id)}>delete</button></p>
            </div>
        );
    }
});

var NoteList = React.createClass({
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

    doLoadMore: function(dateFrom, dateTo, callback) {
        console.log("called NoteList.doLoadMore");
        console.log("dates", dates.yyyymmdd(dateFrom), dates.yyyymmdd(dateTo));

        if (typeof callback === 'function') {
            callback([
                {id: 1, time: "2016-08-02T13:55:32Z", name: "name", text: "text"},
                {id: 2, time: "2016-08-01T13:50:00Z", name: "eman", text: "txet"},
                {id: 3, time: "2016-07-25T12:10:30Z", name: "1234", text: "5678"}
            ]);
        }
    },

    handleLoadMore: function() {
        console.log("state", this.state);

        var dateFrom = dates.addDays(this.state.from, -this.state.days),
            dateTo = this.state.from;
        console.log("====>", dateTo, "-", this.state.days, "=", dateFrom);
        this.doLoadMore(dateFrom, dateTo, function(date, notes) {
            console.log("okok123");
            this.setState({
                from:  date,
                notes: this.state.notes.concat(notes)
            });
        }.bind(this, dateFrom));
    },

    componentDidMount: function() {
        console.log("ok");
        this.handleLoadMore();
    },

    handleRemove: function(index) {
        console.log("called NoteList.handleRemove");
        console.log("index:", index);
        console.log(index);

        var notes = this.state.notes;
        notes.splice(index, 1);

        this.setState({
            notes: notes
        });
    },

    render: function() {
        var notes = this.state.notes.map(function(note, i) {
            return (
                <Note
                    id={note.id}
                    name={note.name}
                    text={note.text}
                    removeCallback={this.handleRemove.bind(this, i)}
                />
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

exports.NoteList = NoteList;
