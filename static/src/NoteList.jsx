
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

    handleCloseDialog: function() {
        var newState = {};
        if (this.state.formDialog) {
            newState['formDialog'] = false;
        }
        if (this.state.deleteDialog) {
            newState['deleteDialog'] = false;
        }
        this.setState(newState);
    },

    handleLoadMore: function() {
        var from = dates.addDays(this.state.dateTo, -this.state.loadDays),
            to   = this.state.dateTo;

        this.setState({
            loading: true,
            dateTo:  from
        });

        client.getNotesByDateRange(from, to, function(notes) {
            var newState = {
                loading: false
            };

            if (notes.length) {
                newState.notes = this.state.notes.concat(notes);
                // return to initial value if was loaded notes
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

    renderEditDialog: function() {
        if (!this.state.formDialog) {
            return null;
        }

        var name = '',
            time = '',
            text = '';

        with (this.state) {
            if (noteIndex in notes) {
                name = notes[noteIndex].name;
                time = notes[noteIndex].time;
                text = notes[noteIndex].text;
            }
        }

        return (
            <Overlay topic="Form" close={this.handleCloseDialog}>
                <input value={name} />
                <input value={time} />
                <textarea value={text} />
                <br />
                <button>save</button>
            </Overlay>
        );
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

    renderDeleteDialog: function() {
        if (!this.state.deleteDialog || !(this.state.noteIndex in this.state.notes)) {
            return null
        }

        var note = this.state.notes[this.state.noteIndex];

        return (
            <Overlay topic="Delete" close={this.handleCloseDialog}>
                <h3>{note.name}</h3>
                <b>{note.time.toDateString()}</b>
                <p>{note.text}</p>
                <button onClick={this.doDelete.bind(this, this.state.noteIndex)}>delete</button>
            </Overlay>
        );
    },

    render: function() {
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

                {this.renderDeleteDialog()}

                {this.renderEditDialog()}


                <h2>Notes list</h2>
                <p>Between <i>{this.state.dateTo.toDateString()}</i> and {this.props.dateStart.toDateString()}</p>
                <hr />
                <button onClick={this.handleEdit}>create</button>
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
