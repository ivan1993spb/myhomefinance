
var React = require('react'),
    Note = require("./Note").Note,
    dates = require("./dates"),
    client = require("./client");

var NoteList = React.createClass({
    displayName: "NoteList",

    propTypes: {
        dateTo:       React.PropTypes.object,
        loadDays:     React.PropTypes.number,
        handleEdit:   React.PropTypes.func.isRequired,
        handleRemove: React.PropTypes.func.isRequired
    },

    getDefaultProps: function() {
        return {
            dateTo:   new Date(),
            loadDays: 20
        };
    },

    getInitialState: function() {
        return {
            dateTo:  this.props.dateTo,
            notes:   [],
            loading: true
        };
    },

    doLoadMore: function(from, to, callback) {
        console.log("called NoteList.doLoadMore");
        console.log("dates 1", dates.yyyymmdd(from), dates.yyyymmdd(to));
        console.log("dates 2", from, to);

        if (typeof callback === 'function') {
            client.getNotesByDateRange(from, to, function() {
                callback([
                    {id: 1, time: "2016-08-02T13:55:32Z", name: "name", text: "text"},
                    {id: 2, time: "2016-08-01T13:50:00Z", name: "eman", text: "txet"},
                    {id: 3, time: "2016-07-25T12:10:30Z", name: "1234", text: "5678"}
                ]);
            });
        }
    },

    handleLoadMore: function() {
        var from = dates.addDays(this.state.dateTo, -this.props.loadDays),
            to   = this.state.dateTo;

        this.setState({
            loading: true
        });

        this.doLoadMore(from, to, function(date, notes) {
            console.log("load more: setting component state");
            this.setState({
                from:    date,
                notes:   this.state.notes.concat(notes),
                loading: false
            });
        }.bind(this, dates.addDays(from, -1)));
    },

    componentDidMount: function() {
        console.log("ok");
        this.handleLoadMore();
    },

    handleRemove: function(index) {
        console.log("called NoteList.handleRemove");
        console.log("index:", index);
        console.log(index);

        // TODO if (this.props.handleRemove(id)) {
        var notes = this.state.notes;
        notes.splice(index, 1);

        this.setState({
            notes: notes
        });
        // }
    },

    handleEdit: function(note, handleUpdate) {
        console.log("edit", note, handleUpdate);
        this.setState({
            noteToEdit: note
        });
        // TODO this.props.handleEdit(id)
    },

    render: function() {
        console.log("render note list");
        var notes = this.state.notes.map(function(note, index) {
            return (
                <Note
                    id={note.id}
                    key={index}
                    time={note.time}
                    name={note.name}
                    text={note.text}
                    handleRemove={this.handleRemove.bind(this, index)}
                    handleEdit={this.handleEdit.bind(this, note)}
                />
            );
        }.bind(this));

        console.log(this.state.noteToEdit);

        return (
            <div>
                <h1>Notes list {this.state.page}</h1>
                <div>
                    {notes}
                </div>
                {
                    this.state.loading ?
                    "loading" :
                    <button onClick={this.handleLoadMore}>load more</button>
                }

            </div>
        );
    }
});

exports.NoteList = NoteList;
