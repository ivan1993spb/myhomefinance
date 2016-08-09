
var React = require('react'),
    Note = require("./Note").Note,
    Overlay = require('./Overlay').Overlay,
    dates = require("./dates"),
    client = require("./client");

var loadDaysLimit = 70;

var NoteList = React.createClass({
    displayName: "NoteList",

    propTypes: {
        dateStart: React.PropTypes.object,
        loadDays:  React.PropTypes.number
    },

    getDefaultProps: function() {
        var currentDate = dates.addDays(new Date(), 1);
        return {
            dateStart: new Date(currentDate.getFullYear(), currentDate.getMonth(), currentDate.getDate(), 0, 0, 0, 0),
            loadDays:  1
        };
    },

    getInitialState: function() {
        return {
            dateTo:   this.props.dateStart,
            loadDays: this.props.loadDays < loadDaysLimit ? this.props.loadDays : loadDaysLimit,
            notes:    [],
            loading:  true,

            noteIndex:    -1,    // Index of note to remove or edit
            deleteDialog: false,
            formDialog:   false
        };
    },

    increaseLoadDays: function(loadDays) {
        if (loadDays >= loadDaysLimit) {
            return loadDays;
        }
        return loadDays * 2;
    },

    doLoadMore: function(from, to, callback) {
        console.log("called NoteList.doLoadMore");
        console.log("dates 1", dates.yyyymmdd(from), dates.yyyymmdd(to));
        console.log("dates 2", from, to);

        if (typeof callback === 'function') {
            client.getNotesByDateRange(from, to, function(notes) {
                callback(notes);
            });
        }
    },

    handleLoadMore: function() {
        var from = dates.addDays(this.state.dateTo, -this.state.loadDays),
            to   = this.state.dateTo;

        this.setState({
            loading: true,
            dateTo:  from
        });

        this.doLoadMore(from, to, function(notes) {
            console.log("load more: setting component state" );
            var newState = {
                loading: false
            };

            if (notes.length) {
                newState.notes = this.state.notes.concat(notes);
                // return initial value
                if (this.state.loadDays != this.props.loadDays) {
                    newState.loadDays = this.props.loadDays;
                }
            } else {
                newState.loadDays = this.increaseLoadDays(this.state.loadDays);
            }

            this.setState(newState);
        }.bind(this));
    },

    componentDidMount: function() {
        console.log("ok");
        this.handleLoadMore();
    },

    handleEdit: function(index, id) {
        console.log("index:", index);
        console.log("id:", id);

        this.setState({
            noteIndex:  index,
            formDialog: true
        });
    },

    handleRemove: function(index, id) {
        console.log("index:", index);
        console.log("id:", id);

        this.setState({
            noteIndex:    index,
            deleteDialog: true
        });
    },

    doDelete: function(index) {
        client.deleteNote(this.state.notes[index].id, function() {
            var notes = this.state.notes;
            notes.splice(index, 1);
            this.setState({
                notes:        notes,
                noteIndex:    -1,
                deleteDialog: false
            });
        }.bind(this));
    },

    render: function() {
        console.log("render note list", this.state.notes);
        var notes = this.state.notes.map(function(note, index) {
            return (
                <Note
                    id={note.id}
                    key={index}
                    time={note.time}
                    name={note.name}
                    text={note.text}
                    handleEdit={this.handleEdit.bind(this, index)}
                    handleRemove={this.handleRemove.bind(this, index)}
                />
            );
        }.bind(this));

        return (
            <div>

                {/* DELETE DIALOG */}
                {this.state.deleteDialog ? <Overlay topic="Delete" close={function(){this.setState({
                    deleteDialog: false
                });}.bind(this)} >
                    <h3>{this.state.notes[this.state.noteIndex].name}</h3>
                    <b>{this.state.notes[this.state.noteIndex].time}</b>
                    <p>{this.state.notes[this.state.noteIndex].text}</p>
                    <button onClick={this.doDelete.bind(this, this.state.noteIndex)}>delete</button>
                </Overlay> : null}

                {/* FORM DIALOG */}
                {this.state.formDialog ? <Overlay topic="Form" close={function(){this.setState({
                    formDialog: false
                });}.bind(this)} >
                    <input value={this.state.noteIndex > -1 ? this.state.notes[this.state.noteIndex].name : null} />
                    <input value={this.state.noteIndex > -1 ? this.state.notes[this.state.noteIndex].time : null} />
                    <textarea value={this.state.noteIndex > -1 ? this.state.notes[this.state.noteIndex].text : null} />
                    <br />
                    <button>save</button>
                </Overlay> : null}

                <h2>Notes list</h2>
                <p>Between <i>{this.state.dateTo.toDateString()}</i> and {this.props.dateStart.toDateString()}</p>
                <hr />
                <div>
                    {notes.length > 0 ? notes : (this.state.loading ? "loading" : "empty")}
                </div>
                <hr />
                <p>Between <i>{this.state.dateTo.toDateString()}</i> and {this.props.dateStart.toDateString()}</p>

                {this.state.loading ? "loading..." : "loaded: "+this.state.notes.length}

                <button
                    onClick={this.handleLoadMore}
                    disabled={this.state.loading}>load more: {this.state.loadDays} days</button>

            </div>
        );
    }
});

exports.NoteList = NoteList;
