
var React = require('react');
var dates = require('./dates');
var client = require('./client');

var Note = React.createClass({
    propTypes: {
        id:             React.PropTypes.number.isRequired,
        name:           React.PropTypes.string.isRequired,
        text:           React.PropTypes.string.isRequired,
        removeCallback: React.PropTypes.func.isRequired,
        editCallback:   React.PropTypes.func.isRequired
    },

    getInitialState: function() {
        return {
            id:   this.props.id,
            name: this.props.name,
            text: this.props.text,
        };
    },

    handleRemove: function(id) {
        console.log("called Note.handleRemove");
        console.log("request to remove note with id", id);
        this.props.removeCallback();
    },

    handleEdit: function() {
        console.log("called Note.handleEdit");
        this.props.editCallback(this.update);
    },

    update: function(time, name, text) {
        this.setState({
            name: name,
            text: text,
            time: time
        })
    },

    handleTest: function() {
        console.log("ok333");
    },

    render: function() {
        return (
            <div>
                <p>{this.state.id}</p>
                <p>{this.state.name}</p>
                <p>{this.state.text}</p>
                <p><button onClick={this.handleRemove.bind(this, this.props.id)}>delete</button></p>
                <p><button onClick={this.handleEdit}>edit</button></p>
                <p><button ref={this.handleTest}>ref</button></p>
            </div>
        );
    }
});

var NoteForm = React.createClass({
    propTypes: {
        id:   React.PropTypes.number,
        time: React.PropTypes.string,
        name: React.PropTypes.string,
        text: React.PropTypes.string
    },

    handleSave: function() {

    },

    render: function() {
        console.log("okok");
        return (
            <form>
                <input value={this.props.note ? this.props.note.id : ""} />
                <input value={this.props.note ? this.props.note.time : ""} />
                <input value={this.props.note ? this.props.note.name : ""} />
                <input value={this.props.note ? this.props.note.text : ""} />
                <button type="button" onClick={this.handleSave}>Save</button>
            </form>
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

    doLoadMore: function(from, to, callback) {
        console.log("called NoteList.doLoadMore");
        console.log("dates", dates.yyyymmdd(from), dates.yyyymmdd(to));

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
        console.log("state", this.state);

        var from = dates.addDays(this.state.from, -this.state.days),
            to   = this.state.from;
        console.log("====>", to, "-", this.state.days, "=", from);
        this.doLoadMore(from, to, function(date, notes) {
            console.log("load more: setting component state");
            this.setState({
                from:  date,
                notes: this.state.notes.concat(notes)
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

        var notes = this.state.notes;
        notes.splice(index, 1);

        this.setState({
            notes: notes
        });
    },

    handleEdit: function(note, handleUpdate) {
        console.log("edit", note, handleUpdate);
        this.setState({
            noteToEdit: note
        });
    },

    render: function() {
        console.log("render note list");
        var notes = this.state.notes.map(function(note, index) {
            return (
                <Note
                    id={note.id}
                    key={index}
                    name={note.name}
                    text={note.text}
                    removeCallback={this.handleRemove.bind(this, index)}
                    editCallback={this.handleEdit.bind(this, note)}
                />
            );
        }.bind(this));

        console.log(this.state.noteToEdit);

        return (
            <div>
                <h1>Notes list {this.state.page}</h1>
                <NoteForm note={this.state.noteToEdit} />
                <div>
                    {notes}
                </div>
                <button onClick={this.handleLoadMore}>load more</button>
            </div>
        );
    }
});

exports.NoteList = NoteList;
