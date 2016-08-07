
var React = require('react'),
    Note = require("./Note").Note,
    dates = require("./dates"),
    client = require("./client");

var NoteList = React.createClass({
    displayName: "NoteList",

    propTypes: {
        dateFrom:     React.PropTypes.object,
        loadDays:     React.PropTypes.number,
        handleEdit:   React.PropTypes.func.isRequired,
        handleRemove: React.PropTypes.func.isRequired
    },

    getDefaultProps: function() {
        return {
            dateFrom: new Date("2016-08-08"),
            loadDays: 1
        };
    },

    getInitialState: function() {
        return {
            dateTo:   this.props.dateFrom,
            loadDays: this.props.loadDays,
            notes:    [],
            loading:  true
        };
    },

    increaseLoadDays: function(loadDays) {
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
        // TODO this.props.handleEdit(id)
    },

    handleRemove: function(index, id) {
        console.log("called NoteList.handleRemove");
        console.log("index:", index);
        console.log("id:", id);

        // TODO if (this.props.handleRemove(id)) {
        // var notes = this.state.notes;
        // notes.splice(index, 1);

        // this.setState({
            // notes: notes
        // });
        // }
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

        console.log(this.state.noteToEdit);

        return (
            <div>
                <h2>Notes list {this.state.page}</h2>
                <p>Between {this.props.dateFrom.toDateString()} and {this.state.dateTo.toDateString()}</p>
                <hr />
                <div>
                    {this.state.loading && this.state.notes.length == 0 ? "loading" : notes}
                </div>
                <hr />
                <p>Between {this.props.dateFrom.toDateString()} and {this.state.dateTo.toDateString()}</p>

                {this.state.loading ? "loading..." : "loaded: "+this.state.notes.length}
                <button
                    onClick={this.handleLoadMore}
                    disabled={this.state.loading}>load more: {this.state.loadDays} days</button>

            </div>
        );
    }
});

exports.NoteList = NoteList;
